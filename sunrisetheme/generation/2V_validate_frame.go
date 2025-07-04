package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <frame_number>\n", os.Args[0])
		os.Exit(1)
	}

	frameNum, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: frame number must be an integer: %v\n", err)
		os.Exit(1)
	}

	// Format frame number with leading zeros
	frameFile := fmt.Sprintf("%03d.png", frameNum)
	inputPath := filepath.Join("intermediate_outputs", "frames", frameFile)
	
	// Create validation output directory
	validationDir := "intermediate_outputs/validation"
	err = os.MkdirAll(validationDir, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating validation directory: %v\n", err)
		os.Exit(1)
	}

	// Read the input PNG file
	fmt.Printf("Reading frame: %s\n", inputPath)
	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", inputPath, err)
		os.Exit(1)
	}
	defer file.Close()

	// Decode the PNG
	img, err := png.Decode(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding PNG: %v\n", err)
		os.Exit(1)
	}

	bounds := img.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y
	
	fmt.Printf("Image dimensions: %dx%d\n", width, height)

	// Create a new RGBA image
	newImg := image.NewRGBA(bounds)

	// Copy pixel by pixel
	pixelCount := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.At(x, y)
			r, g, b, a := c.RGBA()
			
			// Convert from 16-bit to 8-bit
			r8 := uint8(r >> 8)
			g8 := uint8(g >> 8)
			b8 := uint8(b >> 8)
			a8 := uint8(a >> 8)
			
			newImg.Set(x, y, color.RGBA{r8, g8, b8, a8})
			pixelCount++
		}
	}

	fmt.Printf("Processed %d pixels\n", pixelCount)

	// Save the recreated image
	outputPath := filepath.Join(validationDir, fmt.Sprintf("recreated_%s", frameFile))
	outputFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, newImg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding PNG: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Recreated image saved to: %s\n", outputPath)
	
	// Copy original for comparison
	originalCopyPath := filepath.Join(validationDir, fmt.Sprintf("original_%s", frameFile))
	copyFile(inputPath, originalCopyPath)
	fmt.Printf("Original image copied to: %s\n", originalCopyPath)
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	img, err := png.Decode(sourceFile)
	if err != nil {
		return err
	}

	return png.Encode(destFile, img)
}