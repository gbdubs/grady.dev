package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

const (
	DEFAULT_AVERAGING_METHOD = "hbmodeslmean"
	NUM_FRAMES              = 270
	NUM_BREAKPOINTS         = 26
)

type ColorMatrix struct {
	Frames      int
	Breakpoints int
	Data        [][]HSLColorAveraging // [frame][breakpoint]
	Strategy    AveragingStrategy
}

func main() {
	var strategy AveragingStrategy = HBMODESLMEAN // Default

	// Parse command line arguments
	if len(os.Args) > 1 {
		strategies, err := ParseStrategies([]string{os.Args[1]})
		if err != nil {
			fmt.Printf("Error parsing strategy: %v\n", err)
			fmt.Printf("Using default strategy: %s\n", DEFAULT_AVERAGING_METHOD)
		} else {
			strategy = strategies[0]
		}
	}

	fmt.Printf("Processing all frame data with strategy: %s\n", strategy)
	fmt.Printf("Building %dx%d matrix (frames x breakpoints)\n", NUM_FRAMES, NUM_BREAKPOINTS)

	matrix, err := ProcessAllFrameData(strategy)
	if err != nil {
		fmt.Printf("Error processing frame data: %v\n", err)
		os.Exit(1)
	}

	outputPath := fmt.Sprintf("intermediate_outputs/all_frame_matrix_%s.csv", strategy)
	if err := SaveMatrixToCSV(matrix, outputPath); err != nil {
		fmt.Printf("Error saving matrix: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Matrix saved to: %s\n", outputPath)
	fmt.Printf("Successfully processed %d frames with %d breakpoints each\n", NUM_FRAMES, NUM_BREAKPOINTS)
}

func ProcessAllFrameData(strategy AveragingStrategy) (*ColorMatrix, error) {
	matrix := &ColorMatrix{
		Frames:      NUM_FRAMES,
		Breakpoints: NUM_BREAKPOINTS,
		Data:        make([][]HSLColorAveraging, NUM_FRAMES),
		Strategy:    strategy,
	}

	// Initialize matrix
	for f := 0; f < NUM_FRAMES; f++ {
		matrix.Data[f] = make([]HSLColorAveraging, NUM_BREAKPOINTS)
	}

	// Create work items for parallel processing
	type WorkItem struct {
		Frame      int
		Breakpoint int
	}

	// Channel for work items
	workChan := make(chan WorkItem, NUM_FRAMES*NUM_BREAKPOINTS)
	resultChan := make(chan struct {
		Frame      int
		Breakpoint int
		Color      HSLColorAveraging
		Error      error
	}, NUM_FRAMES*NUM_BREAKPOINTS)

	// Add all work items
	for f := 1; f <= NUM_FRAMES; f++ { // Frames are 1-indexed
		for k := 0; k < NUM_BREAKPOINTS; k++ {
			workChan <- WorkItem{Frame: f, Breakpoint: k}
		}
	}
	close(workChan)

	// Start worker goroutines
	numWorkers := 16 // Adjust based on system capabilities
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for work := range workChan {
				color, err := ProcessSingleCell(work.Frame, work.Breakpoint, strategy)
				resultChan <- struct {
					Frame      int
					Breakpoint int
					Color      HSLColorAveraging
					Error      error
				}{work.Frame, work.Breakpoint, color, err}
			}
		}()
	}

	// Close result channel when all workers are done
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect results
	processed := 0
	for result := range resultChan {
		if result.Error != nil {
			fmt.Printf("Warning: Error processing frame %d, breakpoint %d: %v\n", 
				result.Frame, result.Breakpoint, result.Error)
			// Use zero color for errors
			result.Color = HSLColorAveraging{}
		}
		
		// Store in matrix (convert to 0-indexed)
		matrix.Data[result.Frame-1][result.Breakpoint] = result.Color
		
		processed++
		if processed%500 == 0 {
			fmt.Printf("Processed %d/%d cells (%.1f%%)\n", 
				processed, NUM_FRAMES*NUM_BREAKPOINTS, 
				float64(processed)/float64(NUM_FRAMES*NUM_BREAKPOINTS)*100)
		}
	}

	fmt.Printf("Completed processing all %d cells\n", NUM_FRAMES*NUM_BREAKPOINTS)
	return matrix, nil
}

func ProcessSingleCell(frameNumber, breakpoint int, strategy AveragingStrategy) (HSLColorAveraging, error) {
	// Load CSV data for this specific frame and breakpoint
	csvPath := fmt.Sprintf("intermediate_outputs/row_data/%d/%d.csv", frameNumber, breakpoint)
	
	colors, err := loadCSVColors(csvPath)
	if err != nil {
		return HSLColorAveraging{}, fmt.Errorf("failed to load CSV %s: %w", csvPath, err)
	}

	if len(colors) == 0 {
		return HSLColorAveraging{}, nil // Return zero color for empty data
	}

	// Apply the averaging strategy
	switch strategy {
	case MEAN:
		return calculateMean(colors), nil
	case MODE:
		return calculateMode(colors), nil
	case MEDIAN:
		return calculateMedian(colors), nil
	case T3LMEAN:
		return calculateT3LMean(colors), nil
	case M3LMEAN:
		return calculateM3LMean(colors), nil
	case HBMODESLMEAN:
		return calculateHBModeSLMean(colors), nil
	default:
		return calculateHBModeSLMean(colors), nil // Default fallback
	}
}

func SaveMatrixToCSV(matrix *ColorMatrix, outputPath string) error {
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return err
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header with metadata
	writer.Write([]string{"# Matrix:", fmt.Sprintf("%dx%d", matrix.Frames, matrix.Breakpoints)})
	writer.Write([]string{"# Strategy:", string(matrix.Strategy)})
	writer.Write([]string{"# Format:", "Frame,Breakpoint,H,S,L"})
	writer.Write([]string{"Frame", "Breakpoint", "H", "S", "L"})

	// Write data
	for f := 0; f < matrix.Frames; f++ {
		for k := 0; k < matrix.Breakpoints; k++ {
			color := matrix.Data[f][k]
			record := []string{
				strconv.Itoa(f + 1), // Convert back to 1-indexed
				strconv.Itoa(k),
				strconv.FormatFloat(color.H, 'f', 6, 64),
				strconv.FormatFloat(color.S, 'f', 6, 64),
				strconv.FormatFloat(color.L, 'f', 6, 64),
			}
			writer.Write(record)
		}
	}

	return nil
}