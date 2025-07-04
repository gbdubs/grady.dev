package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	ANIMATION_DURATION = 20.0 // seconds
	DEFAULT_KEYFRAMES  = 30
	DEFAULT_SMOOTH     = 0
	SVG_WIDTH         = 800
	SVG_HEIGHT        = 600
)

type AnimationConfig struct {
	SmoothLevel     int
	SpatialSmooth   bool
	NumKeyframes    int
	InputFile       string
	OutputFile      string
	AnimationData   [][]HSLColorAveraging // [frame][breakpoint]
	TotalFrames     int
	NumBreakpoints  int
}

func main() {
	config := parseArguments()
	
	fmt.Printf("Loading matrix data from: %s\n", config.InputFile)
	if err := loadMatrixData(config); err != nil {
		fmt.Printf("Error loading matrix data: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Printf("Loaded %d frames with %d breakpoints each\n", config.TotalFrames, config.NumBreakpoints)
	fmt.Printf("Configuration: smooth-level=%d, spatial-smooth=%t, keyframes=%d\n", config.SmoothLevel, config.SpatialSmooth, config.NumKeyframes)
	
	// Apply temporal smoothing if requested
	if config.SmoothLevel > 0 {
		fmt.Printf("Applying temporal smoothing with level %d...\n", config.SmoothLevel)
		applySmoothingLS(config)
	}
	
	// Apply spatial smoothing if requested
	if config.SpatialSmooth {
		fmt.Printf("Applying spatial smoothing...\n")
		applySpatialSmoothing(config)
	}
	
	// Sample keyframes
	fmt.Printf("Sampling %d keyframes from %d total frames...\n", config.NumKeyframes, config.TotalFrames)
	keyframes := sampleKeyframes(config)
	
	// Generate animated SVG
	fmt.Printf("Generating animated SVG...\n")
	if err := generateAnimatedSVG(config, keyframes); err != nil {
		fmt.Printf("Error generating SVG: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Printf("Animated SVG saved to: %s\n", config.OutputFile)
}

func parseArguments() *AnimationConfig {
	config := &AnimationConfig{
		SmoothLevel:   DEFAULT_SMOOTH,
		SpatialSmooth: false,
		NumKeyframes:  DEFAULT_KEYFRAMES,
		InputFile:     "intermediate_outputs/all_frame_matrix_hbmodeslmean.csv",
	}
	
	// Parse command line arguments
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--smooth-level":
			if i+1 < len(args) {
				if val, err := strconv.Atoi(args[i+1]); err == nil {
					config.SmoothLevel = val
				}
				i++
			}
		case "--spatial-smooth":
			config.SpatialSmooth = true
		case "--number-of-keyframes":
			if i+1 < len(args) {
				if val, err := strconv.Atoi(args[i+1]); err == nil {
					config.NumKeyframes = val
				}
				i++
			}
		case "--input":
			if i+1 < len(args) {
				config.InputFile = args[i+1]
				i++
			}
		}
	}
	
	// Generate output filename based on parameters
	spatialSuffix := ""
	if config.SpatialSmooth {
		spatialSuffix = "_spatial"
	}
	config.OutputFile = fmt.Sprintf("intermediate_outputs/animation_smooth%d%s_keyframes%d.svg", 
		config.SmoothLevel, spatialSuffix, config.NumKeyframes)
	
	return config
}

func loadMatrixData(config *AnimationConfig) error {
	file, err := os.Open(config.InputFile)
	if err != nil {
		return err
	}
	defer file.Close()
	
	reader := csv.NewReader(file)
	// Allow variable number of fields to handle the quoted format description
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
	
	// Parse dimensions and find data start
	dataStartRow := 0
	for i, record := range records {
		if len(record) >= 2 && record[0] == "# Matrix:" {
			// Parse dimensions like "270x21"
			dims := strings.Split(record[1], "x")
			if len(dims) == 2 {
				config.TotalFrames, _ = strconv.Atoi(dims[0])
				config.NumBreakpoints, _ = strconv.Atoi(dims[1])
			}
		}
		if len(record) >= 1 && record[0] == "Frame" {
			dataStartRow = i + 1
			break
		}
	}
	
	// Initialize data matrix
	config.AnimationData = make([][]HSLColorAveraging, config.TotalFrames)
	for i := range config.AnimationData {
		config.AnimationData[i] = make([]HSLColorAveraging, config.NumBreakpoints)
	}
	
	// Parse CSV data
	for i := dataStartRow; i < len(records); i++ {
		record := records[i]
		if len(record) != 5 {
			continue
		}
		
		frame, err := strconv.Atoi(record[0])
		if err != nil {
			continue
		}
		breakpoint, err := strconv.Atoi(record[1])
		if err != nil {
			continue
		}
		h, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			continue
		}
		s, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			continue
		}
		l, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			continue
		}
		
		// Convert to 0-indexed
		frameIdx := frame - 1
		if frameIdx >= 0 && frameIdx < config.TotalFrames && 
		   breakpoint >= 0 && breakpoint < config.NumBreakpoints {
			config.AnimationData[frameIdx][breakpoint] = HSLColorAveraging{H: h, S: s, L: l}
		}
	}
	
	return nil
}

func applySmoothingLS(config *AnimationConfig) {
	if config.SmoothLevel <= 0 {
		return
	}
	
	// Create a copy for smoothing calculations
	smoothed := make([][]HSLColorAveraging, config.TotalFrames)
	for i := range smoothed {
		smoothed[i] = make([]HSLColorAveraging, config.NumBreakpoints)
		copy(smoothed[i], config.AnimationData[i])
	}
	
	// Apply smoothing to L and S values only (preserve H)
	for frame := 0; frame < config.TotalFrames; frame++ {
		for breakpoint := 0; breakpoint < config.NumBreakpoints; breakpoint++ {
			var sSum, lSum float64
			count := 0
			
			// Average with adjacent frames within smoothing window
			for offset := -config.SmoothLevel; offset <= config.SmoothLevel; offset++ {
				targetFrame := frame + offset
				if targetFrame >= 0 && targetFrame < config.TotalFrames {
					sSum += config.AnimationData[targetFrame][breakpoint].S
					lSum += config.AnimationData[targetFrame][breakpoint].L
					count++
				}
			}
			
			if count > 0 {
				smoothed[frame][breakpoint].S = sSum / float64(count)
				smoothed[frame][breakpoint].L = lSum / float64(count)
				// Keep original hue
				smoothed[frame][breakpoint].H = config.AnimationData[frame][breakpoint].H
			}
		}
	}
	
	config.AnimationData = smoothed
}

func applySpatialSmoothing(config *AnimationConfig) {
	// Create a copy for smoothing calculations
	smoothed := make([][]HSLColorAveraging, config.TotalFrames)
	for i := range smoothed {
		smoothed[i] = make([]HSLColorAveraging, config.NumBreakpoints)
		copy(smoothed[i], config.AnimationData[i])
	}
	
	// Apply spatial smoothing with weighted averaging:
	// 6x center + 4x 1-time neighbors + 2x 2-time neighbors + 1x 3-time neighbors
	// Only smooth S and L, preserve original H
	for frame := 0; frame < config.TotalFrames; frame++ {
		for breakpoint := 0; breakpoint < config.NumBreakpoints; breakpoint++ {
			var sSum, lSum float64
			totalWeight := 0.0
			
			// Define the sampling pattern with weights based on Manhattan distance
			// Format: [timeOffset, locationOffset, weight]
			samples := [][]int{
				// Center (6x weight)
				{0, 0, 6},
				
				// 1-time neighbors (4x weight each) - Manhattan distance 1
				{-1, 0, 4},  // Time neighbors
				{1, 0, 4},
				{0, -1, 4},  // Location neighbors
				{0, 1, 4},
				
				// 2-time neighbors (2x weight each) - Manhattan distance 2
				{-2, 0, 2},  // Time neighbors at distance 2
				{2, 0, 2},
				{0, -2, 2},  // Location neighbors at distance 2
				{0, 2, 2},
				{-1, -1, 2}, // Diagonal neighbors (also distance 2)
				{-1, 1, 2},
				{1, -1, 2},
				{1, 1, 2},
				
				// 3-time neighbors (1x weight each) - Manhattan distance 3
				{-3, 0, 1},  // Time neighbors at distance 3
				{3, 0, 1},
				{0, -3, 1},  // Location neighbors at distance 3
				{0, 3, 1},
				{-2, -1, 1}, // Mixed distance 3 neighbors
				{-2, 1, 1},
				{2, -1, 1},
				{2, 1, 1},
				{-1, -2, 1},
				{-1, 2, 1},
				{1, -2, 1},
				{1, 2, 1},
			}
			
			for _, sample := range samples {
				timeOffset := sample[0]
				locOffset := sample[1]
				weight := float64(sample[2])
				
				targetFrame := frame + timeOffset
				targetBreakpoint := breakpoint + locOffset
				
				// Check bounds
				if targetFrame >= 0 && targetFrame < config.TotalFrames &&
				   targetBreakpoint >= 0 && targetBreakpoint < config.NumBreakpoints {
					
					color := config.AnimationData[targetFrame][targetBreakpoint]
					// Only accumulate S and L, not H
					sSum += color.S * weight
					lSum += color.L * weight
					totalWeight += weight
				}
			}
			
			if totalWeight > 0 {
				// Preserve original hue, smooth only saturation and lightness
				smoothed[frame][breakpoint] = HSLColorAveraging{
					H: config.AnimationData[frame][breakpoint].H, // Keep original hue
					S: sSum / totalWeight,
					L: lSum / totalWeight,
				}
			}
		}
	}
	
	config.AnimationData = smoothed
}

func sampleKeyframes(config *AnimationConfig) []int {
	if config.NumKeyframes >= config.TotalFrames {
		// Use all frames
		keyframes := make([]int, config.TotalFrames)
		for i := range keyframes {
			keyframes[i] = i
		}
		return keyframes
	}
	
	// Sample evenly across the total frames
	keyframes := make([]int, config.NumKeyframes)
	for i := 0; i < config.NumKeyframes; i++ {
		// Distribute evenly across the frame range
		frameIndex := int(float64(i) * float64(config.TotalFrames-1) / float64(config.NumKeyframes-1))
		keyframes[i] = frameIndex
	}
	
	return keyframes
}

func generateAnimatedSVG(config *AnimationConfig, keyframes []int) error {
	if err := os.MkdirAll(filepath.Dir(config.OutputFile), 0755); err != nil {
		return err
	}
	
	file, err := os.Create(config.OutputFile)
	if err != nil {
		return err
	}
	defer file.Close()
	
	// Write SVG header
	fmt.Fprintf(file, `<?xml version="1.0" encoding="UTF-8"?>
<svg width="%d" height="%d" xmlns="http://www.w3.org/2000/svg">
<defs>
`, SVG_WIDTH, SVG_HEIGHT)
	
	// Create a single vertical linear gradient with animated stops
	fmt.Fprintf(file, `<linearGradient id="sunriseGradient" x1="0%%" y1="0%%" x2="0%%" y2="100%%">
`)
	
	// Create animated stops for each breakpoint (vertical position)
	for breakpoint := 0; breakpoint < config.NumBreakpoints; breakpoint++ {
		// Calculate the offset percentage for this breakpoint
		offset := float64(breakpoint) / float64(config.NumBreakpoints-1) * 100
		
		fmt.Fprintf(file, `<stop offset="%.1f%%" stop-color="rgb(0,0,0)">
<animate attributeName="stop-color" dur="%.1fs" repeatCount="indefinite" values="`, offset, ANIMATION_DURATION)
		
		// Generate color values for animation for this breakpoint
		colorValues := make([]string, len(keyframes))
		for i, frameIdx := range keyframes {
			color := config.AnimationData[frameIdx][breakpoint]
			r, g, b := HSLToRGB(color.H, color.S, color.L)
			colorValues[i] = fmt.Sprintf("rgb(%d,%d,%d)", r, g, b)
		}
		
		// Add the first color at the end to create seamless loop
		colorValues = append(colorValues, colorValues[0])
		
		fmt.Fprintf(file, "%s\" />\n</stop>\n", strings.Join(colorValues, ";"))
	}
	
	fmt.Fprintf(file, "</linearGradient>\n")
	fmt.Fprintf(file, "</defs>\n")
	
	// Create a single rectangle that fills the entire SVG with the animated gradient
	fmt.Fprintf(file, `<rect x="0" y="0" width="%d" height="%d" fill="url(#sunriseGradient)" />
`, SVG_WIDTH, SVG_HEIGHT)
	
	// Add metadata text overlay
	spatialText := ""
	if config.SpatialSmooth {
		spatialText = " | Spatial: ON"
	}
	fmt.Fprintf(file, `<text x="10" y="30" fill="white" font-family="Arial" font-size="14" opacity="0.7">
Smooth Level: %d%s | Keyframes: %d | Duration: %.0fs
</text>
`, config.SmoothLevel, spatialText, config.NumKeyframes, ANIMATION_DURATION)
	
	fmt.Fprintf(file, "</svg>\n")
	
	return nil
}