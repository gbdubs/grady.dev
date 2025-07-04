package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type HSLColor struct {
	H float64 `json:"h"`
	S float64 `json:"s"`
	L float64 `json:"l"`
}

type BreakpointData struct {
	Breakpoint int      `json:"breakpoint"`
	Color      HSLColor `json:"color"`
	RGB        string   `json:"rgb"`
}

type FrameData struct {
	Frame       int               `json:"frame"`
	Breakpoints []BreakpointData `json:"breakpoints"`
}

type GradientStop struct {
	Offset float64 `json:"offset"`
	Color  string  `json:"color"`
}

type BlockGradient struct {
	BlockIndex       int            `json:"blockIndex"`
	TopFrame         int            `json:"topFrame"`
	BottomFrame      int            `json:"bottomFrame"`
	HorizontalStops  []GradientStop `json:"horizontalStops"`  // Sky to horizon gradient
	VerticalStops    []GradientStop `json:"verticalStops"`    // Overlay fade gradient
}

type BlockConfiguration struct {
	TotalBlocks int             `json:"totalBlocks"`
	Gradients   []BlockGradient `json:"gradients"`
}

type SunriseMatrixData struct {
	Frames            []FrameData          `json:"frames"`
	TotalFrames       int                  `json:"totalFrames"`
	TotalBreakpoints  int                  `json:"totalBreakpoints"`
	BlockConfigs      []BlockConfiguration `json:"blockConfigs"`  // Pre-computed gradients for different numbers of blocks
}

func hslToRgb(h, s, l float64) (int, int, int) {
	// Convert HSL to RGB
	// H is in degrees (0-360), S and L are 0-1
	h = h / 360.0
	
	var r, g, b float64
	
	if s == 0 {
		r = l
		g = l
		b = l
	} else {
		hue2rgb := func(p, q, t float64) float64 {
			if t < 0 {
				t += 1
			}
			if t > 1 {
				t -= 1
			}
			if t < 1.0/6.0 {
				return p + (q-p)*6*t
			}
			if t < 1.0/2.0 {
				return q
			}
			if t < 2.0/3.0 {
				return p + (q-p)*(2.0/3.0-t)*6
			}
			return p
		}
		
		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}
		p := 2*l - q
		
		r = hue2rgb(p, q, h+1.0/3.0)
		g = hue2rgb(p, q, h)
		b = hue2rgb(p, q, h-1.0/3.0)
	}
	
	return int(math.Round(r * 255)), int(math.Round(g * 255)), int(math.Round(b * 255))
}

func main() {
	// Check if the matrix CSV file exists
	matrixPath := "../intermediate_outputs/all_frame_matrix_hbmodeslmean.csv"
	if _, err := os.Stat(matrixPath); os.IsNotExist(err) {
		fmt.Printf("ERROR: Required matrix data file not found: %s\n", matrixPath)
		fmt.Println("Please run the sunrise data pipeline (step 6) to generate this file.")
		os.Exit(1)
	}
	
	// Read the CSV file manually to handle headers
	content, err := os.ReadFile(matrixPath)
	if err != nil {
		fmt.Printf("Error reading matrix file: %v\n", err)
		os.Exit(1)
	}
	
	lines := strings.Split(string(content), "\n")
	var dataRecords [][]string
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "Frame,") {
			continue
		}
		
		fields := strings.Split(line, ",")
		if len(fields) == 5 {
			dataRecords = append(dataRecords, fields)
		}
	}
	
	if len(dataRecords) == 0 {
		fmt.Println("ERROR: No data records found in matrix file")
		os.Exit(1)
	}
	
	// Parse the data
	frameMap := make(map[int][]BreakpointData)
	maxFrame := 0
	maxBreakpoint := 0
	
	for _, record := range dataRecords {
		if len(record) < 5 {
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
		
		// Convert HSL to RGB
		r, g, b := hslToRgb(h, s, l)
		rgb := fmt.Sprintf("rgb(%d,%d,%d)", r, g, b)
		
		breakpointData := BreakpointData{
			Breakpoint: breakpoint,
			Color: HSLColor{H: h, S: s, L: l},
			RGB: rgb,
		}
		
		frameMap[frame] = append(frameMap[frame], breakpointData)
		
		if frame > maxFrame {
			maxFrame = frame
		}
		if breakpoint > maxBreakpoint {
			maxBreakpoint = breakpoint
		}
	}
	
	// Create structured data
	var frames []FrameData
	for i := 1; i <= maxFrame; i++ {
		if breakpoints, exists := frameMap[i]; exists {
			frames = append(frames, FrameData{
				Frame: i,
				Breakpoints: breakpoints,
			})
		}
	}
	
	// Pre-compute gradients for different numbers of blocks (3 to 270)
	var blockConfigs []BlockConfiguration
	for numBlocks := 3; numBlocks <= 270 && numBlocks <= maxFrame; numBlocks++ {
		config := BlockConfiguration{
			TotalBlocks: numBlocks,
			Gradients:   make([]BlockGradient, numBlocks),
		}
		
		// Calculate frame assignments for each block
		for blockIdx := 0; blockIdx < numBlocks; blockIdx++ {
			// For N blocks, we need N+1 gradients (fenceposts)
			// Block B uses gradient B as top (horizontal) and gradient B+1 as bottom (vertical)
			frameForTop := int(float64(blockIdx) / float64(numBlocks) * float64(maxFrame-1))
			frameForBottom := int(float64(blockIdx+1) / float64(numBlocks) * float64(maxFrame-1))
			
			// Clamp to valid range
			if frameForTop >= maxFrame {
				frameForTop = maxFrame - 1
			}
			if frameForBottom >= maxFrame {
				frameForBottom = maxFrame - 1
			}
			
			// Get frame data (+1 because frames are 1-indexed in the data)
			topFrameData := frameMap[frameForTop+1]
			bottomFrameData := frameMap[frameForBottom+1]
			
			// Create horizontal gradient (this block's top - same as previous block's bottom)
			horizStops := make([]GradientStop, len(topFrameData))
			for i, bp := range topFrameData {
				horizStops[i] = GradientStop{
					Offset: float64(i) / float64(len(topFrameData)-1),
					Color:  bp.RGB,
				}
			}
			
			// Create vertical gradient (this block's bottom - will be next block's top)
			vertStops := make([]GradientStop, len(bottomFrameData))
			for i, bp := range bottomFrameData {
				vertStops[i] = GradientStop{
					Offset: float64(i) / float64(len(bottomFrameData)-1),
					Color:  bp.RGB,
				}
			}
			
			config.Gradients[blockIdx] = BlockGradient{
				BlockIndex:      blockIdx,
				TopFrame:        frameForTop + 1,
				BottomFrame:     frameForBottom + 1,
				HorizontalStops: horizStops,
				VerticalStops:   vertStops,
			}
		}
		
		blockConfigs = append(blockConfigs, config)
	}
	
	matrixData := SunriseMatrixData{
		Frames:           frames,
		TotalFrames:      maxFrame,
		TotalBreakpoints: maxBreakpoint + 1,
		BlockConfigs:     blockConfigs,
	}
	
	// Output as JSON for Hugo data
	jsonData, err := json.MarshalIndent(matrixData, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}
	
	// Write to data file
	err = os.WriteFile("exampleSite/data/sunrise_matrix.json", jsonData, 0644)
	if err != nil {
		fmt.Printf("Error writing JSON file: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Printf("Successfully extracted matrix data:\n")
	fmt.Printf("- %d frames\n", len(frames))
	fmt.Printf("- %d breakpoints per frame\n", maxBreakpoint + 1)
	fmt.Printf("- %d block configurations (3-%d blocks)\n", len(blockConfigs), len(blockConfigs)+2)
	fmt.Printf("- %d total data points\n", len(dataRecords))
	fmt.Printf("- Output: exampleSite/data/sunrise_matrix.json\n")
}