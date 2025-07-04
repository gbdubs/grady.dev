package main

import (
	"encoding/csv"
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

const (
	NUM_ROWS_TO_SAMPLE = 26
	MIN_PIXELS_PER_ROW = 5
)

type HSLColor struct {
	H, S, L float64
}

type RowData struct {
	RowIndex int
	Pixels   []HSLColor
}

type RowAnalyzer struct {
	imagePath string
	frameNum  int
}

func NewRowAnalyzer(imagePath string, frameNum int) *RowAnalyzer {
	return &RowAnalyzer{
		imagePath: imagePath,
		frameNum:  frameNum,
	}
}

func (ra *RowAnalyzer) ProcessImage() error {
	img, err := ra.loadPNG()
	if err != nil {
		return fmt.Errorf("failed to load PNG: %w", err)
	}

	validRows := ra.findValidRows(img)
	if len(validRows) == 0 {
		return fmt.Errorf("no valid rows found with more than %d pixels", MIN_PIXELS_PER_ROW)
	}

	selectedRows := ra.selectEvenlyDistributedRows(validRows, NUM_ROWS_TO_SAMPLE)
	
	return ra.processRowsParallel(img, selectedRows)
}

func (ra *RowAnalyzer) loadPNG() (image.Image, error) {
	file, err := os.Open(ra.imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func (ra *RowAnalyzer) findValidRows(img image.Image) []int {
	bounds := img.Bounds()
	var validRows []int

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		pixelCount := 0
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			_, _, _, a := img.At(x, y).RGBA()
			if a > 0 {
				pixelCount++
			}
		}
		if pixelCount > MIN_PIXELS_PER_ROW {
			validRows = append(validRows, y)
		}
	}

	return validRows
}

func (ra *RowAnalyzer) selectEvenlyDistributedRows(validRows []int, numSamples int) []int {
	if len(validRows) <= numSamples {
		return validRows
	}

	selected := make([]int, numSamples)
	step := float64(len(validRows)-1) / float64(numSamples-1)

	for i := 0; i < numSamples; i++ {
		index := int(math.Round(float64(i) * step))
		selected[i] = validRows[index]
	}

	return selected
}

func (ra *RowAnalyzer) processRowsParallel(img image.Image, selectedRows []int) error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(selectedRows))

	for k, rowY := range selectedRows {
		wg.Add(1)
		go func(k, rowY int) {
			defer wg.Done()
			if err := ra.processRow(img, rowY, k); err != nil {
				errChan <- err
			}
		}(k, rowY)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

func (ra *RowAnalyzer) processRow(img image.Image, rowY, k int) error {
	bounds := img.Bounds()
	var pixels []HSLColor

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		r, g, b, a := img.At(x, rowY).RGBA()
		if a > 0 {
			r8 := uint8(r >> 8)
			g8 := uint8(g >> 8)
			b8 := uint8(b >> 8)
			hsl := rgbToHSL(r8, g8, b8)
			pixels = append(pixels, hsl)
		}
	}

	if len(pixels) == 0 {
		return nil
	}

	return ra.writeCSV(pixels, k)
}

func (ra *RowAnalyzer) writeCSV(pixels []HSLColor, k int) error {
	outputDir := fmt.Sprintf("intermediate_outputs/row_data/%d", ra.frameNum)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	filename := filepath.Join(outputDir, fmt.Sprintf("%d.csv", k))
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"H", "S", "L"})

	for _, pixel := range pixels {
		record := []string{
			strconv.FormatFloat(pixel.H, 'f', 6, 64),
			strconv.FormatFloat(pixel.S, 'f', 6, 64),
			strconv.FormatFloat(pixel.L, 'f', 6, 64),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

func rgbToHSL(r, g, b uint8) HSLColor {
	rf := float64(r) / 255.0
	gf := float64(g) / 255.0
	bf := float64(b) / 255.0

	max := math.Max(rf, math.Max(gf, bf))
	min := math.Min(rf, math.Min(gf, bf))
	diff := max - min

	h := 0.0
	s := 0.0
	l := (max + min) / 2.0

	if diff != 0 {
		if l < 0.5 {
			s = diff / (max + min)
		} else {
			s = diff / (2.0 - max - min)
		}

		switch max {
		case rf:
			h = (gf - bf) / diff
			if gf < bf {
				h += 6
			}
		case gf:
			h = (bf-rf)/diff + 2
		case bf:
			h = (rf-gf)/diff + 4
		}
		h /= 6
	}

	return HSLColor{H: h * 360, S: s, L: l}
}

func ProcessMultipleFrames(imagePaths []string, frameNums []int) error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(imagePaths))

	for i, imagePath := range imagePaths {
		wg.Add(1)
		go func(imagePath string, frameNum int) {
			defer wg.Done()
			analyzer := NewRowAnalyzer(imagePath, frameNum)
			if err := analyzer.ProcessImage(); err != nil {
				errChan <- fmt.Errorf("frame %d: %w", frameNum, err)
			}
		}(imagePath, frameNums[i])
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run row_analyzer.go <image_path> <frame_number>")
		os.Exit(1)
	}

	imagePath := os.Args[1]
	frameNum, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid frame number: %v\n", err)
		os.Exit(1)
	}

	analyzer := NewRowAnalyzer(imagePath, frameNum)
	if err := analyzer.ProcessImage(); err != nil {
		fmt.Printf("Error processing image: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully processed frame %d\n", frameNum)
}