package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run 4_averaging_lib.go 4_visualizer_main.go <frame_number> <strategy1> [strategy2] [strategy3] ...")
		fmt.Println("Available strategies: mean, mode, median")
		os.Exit(1)
	}

	frameNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Invalid frame number: %v\n", err)
		os.Exit(1)
	}

	strategyCodes := os.Args[2:]
	strategies, err := ParseStrategies(strategyCodes)
	if err != nil {
		fmt.Printf("Error parsing strategies: %v\n", err)
		os.Exit(1)
	}

	if err := CreateVisualization(frameNumber, strategies); err != nil {
		fmt.Printf("Error creating visualization: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Visualization created for frame %d with strategies: %v\n", frameNumber, strategyCodes)
}

func CreateVisualization(frameNumber int, strategies []AveragingStrategy) error {
	frameData, err := LoadFrameData(frameNumber)
	if err != nil {
		return fmt.Errorf("failed to load frame data: %w", err)
	}

	originalImg, err := loadOriginalFrame(frameNumber)
	if err != nil {
		return fmt.Errorf("failed to load original frame: %w", err)
	}

	minRow, maxRow, err := FindValidRowBounds(frameNumber)
	if err != nil {
		return fmt.Errorf("failed to find valid row bounds: %w", err)
	}

	originalBounds := originalImg.Bounds()
	gradientHeight := maxRow - minRow + 1
	gradientWidth := 150

	// Calculate test strip sizing with spacing
	imageWidth := originalBounds.Dx()
	stripGroupSpacing := imageWidth / 5         // Distance between the 4 groups
	testStripSpacing := stripGroupSpacing / 3   // Test strip width is 1/3 of spacing
	testStripWidth := int(testStripSpacing) / 5 // Decrease width by 5x
	if testStripWidth < 4 {                     // Minimum width for visibility (reduced from 8)
		testStripWidth = 4
	}

	totalWidth := originalBounds.Dx() + len(strategies)*gradientWidth
	keyHeight := 120 // Space for strategy key and gradient labels at bottom (increased for 6x scaled font)
	totalHeight := originalBounds.Dy() + keyHeight

	resultImg := image.NewRGBA(image.Rect(0, 0, totalWidth, totalHeight))

	draw.Draw(resultImg, originalBounds, originalImg, image.Point{0, 0}, draw.Src)

	// Add 4 test strips evenly spaced across the original image
	stripPositions := []int{
		stripGroupSpacing * 1,
		stripGroupSpacing * 2,
		stripGroupSpacing * 3,
		stripGroupSpacing * 4,
	}

	for strategyIndex, strategy := range strategies {
		result := ApplyAveragingStrategy(frameData, strategy)

		testStripImg, err := createGradientImage(result.Colors, testStripWidth, gradientHeight)
		if err != nil {
			return fmt.Errorf("failed to create test strip for %s: %w", strategy, err)
		}

		// Add 4 test strips at evenly spaced positions across the original image
		// Offset each strategy's strips so they don't overlap
		stripGap := testStripWidth * 2 // Gap between strips is 2x strip width
		strategyOffset := strategyIndex * (testStripWidth + stripGap)

		for _, baseStripX := range stripPositions {
			totalGroupWidth := len(strategies)*testStripWidth + (len(strategies)-1)*stripGap
			stripX := baseStripX - totalGroupWidth/2 + strategyOffset

			if stripX < 0 {
				stripX = 0
			}
			if stripX+testStripWidth > imageWidth {
				stripX = imageWidth - testStripWidth
			}

			stripRect := image.Rect(stripX, minRow, stripX+testStripWidth, minRow+gradientHeight)
			draw.Draw(resultImg, stripRect, testStripImg, image.Point{0, 0}, draw.Over)

			// Add strategy number label below each test strip
			labelY := minRow + gradientHeight + 15
			labelText := fmt.Sprintf("%d", strategyIndex+1)
			addTextLabel(resultImg, stripX+testStripWidth/2, labelY, labelText, color.RGBA{255, 255, 255, 255})
		}
	}

	// Add the main gradient strips on the right side
	currentX := originalBounds.Dx()
	for strategyIndex, strategy := range strategies {
		result := ApplyAveragingStrategy(frameData, strategy)

		gradientImg, err := createGradientImage(result.Colors, gradientWidth, gradientHeight)
		if err != nil {
			return fmt.Errorf("failed to create gradient for %s: %w", strategy, err)
		}

		gradientRect := image.Rect(currentX, minRow, currentX+gradientWidth, minRow+gradientHeight)
		draw.Draw(resultImg, gradientRect, gradientImg, image.Point{0, 0}, draw.Src)

		// Add strategy number label below each gradient strip
		labelY := minRow + gradientHeight + 25
		labelText := fmt.Sprintf("%d", strategyIndex+1)
		addTextLabel(resultImg, currentX+gradientWidth/2, labelY, labelText, color.RGBA{255, 255, 255, 255})

		currentX += gradientWidth
	}

	outputPath := fmt.Sprintf("intermediate_outputs/visualization_%d_%s.png", frameNumber, strings.Join(strategyCodesToStrings(strategies), "_"))

	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	// Add strategy key at the bottom of the image
	addStrategyKey(resultImg, strategies, totalWidth, originalBounds.Dy()+5)

	if err := png.Encode(outputFile, resultImg); err != nil {
		return fmt.Errorf("failed to encode PNG: %w", err)
	}

	return nil
}

func loadOriginalFrame(frameNumber int) (image.Image, error) {
	framePath := fmt.Sprintf("intermediate_outputs/frames/%03d.png", frameNumber)

	file, err := os.Open(framePath)
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

func createGradientImage(colors []HSLColorAveraging, width, height int) (image.Image, error) {
	if len(colors) == 0 {
		return image.NewRGBA(image.Rect(0, 0, width, height)), nil
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		progress := float64(y) / float64(height-1)

		colorIndex := progress * float64(len(colors)-1)
		index1 := int(colorIndex)
		index2 := index1 + 1

		if index2 >= len(colors) {
			index2 = len(colors) - 1
			index1 = index2
		}

		if index1 < 0 {
			index1 = 0
		}

		var pixelColor color.RGBA
		if index1 == index2 {
			// Single color, convert HSL to RGB
			r, g, b := HSLToRGB(colors[index1].H, colors[index1].S, colors[index1].L)
			pixelColor = color.RGBA{r, g, b, 255}
		} else {
			// Interpolate in RGB space
			t := colorIndex - float64(index1)
			
			// Convert both HSL colors to RGB
			r1, g1, b1 := HSLToRGB(colors[index1].H, colors[index1].S, colors[index1].L)
			r2, g2, b2 := HSLToRGB(colors[index2].H, colors[index2].S, colors[index2].L)
			
			// Linear interpolation in RGB space
			r := uint8(float64(r1)*(1-t) + float64(r2)*t)
			g := uint8(float64(g1)*(1-t) + float64(g2)*t)
			b := uint8(float64(b1)*(1-t) + float64(b2)*t)
			
			pixelColor = color.RGBA{r, g, b, 255}
		}

		for x := 0; x < width; x++ {
			img.Set(x, y, pixelColor)
		}
	}

	return img, nil
}

func interpolateHSL(color1, color2 HSLColorAveraging, t float64) HSLColorAveraging {
	h1, h2 := color1.H, color2.H

	diff := h2 - h1
	if diff > 180 {
		h1 += 360
	} else if diff < -180 {
		h2 += 360
	}

	h := h1 + t*(h2-h1)
	if h >= 360 {
		h -= 360
	} else if h < 0 {
		h += 360
	}

	return HSLColorAveraging{
		H: h,
		S: color1.S + t*(color2.S-color1.S),
		L: color1.L + t*(color2.L-color1.L),
	}
}

func strategyCodesToStrings(strategies []AveragingStrategy) []string {
	var codes []string
	for _, strategy := range strategies {
		codes = append(codes, string(strategy))
	}
	return codes
}

func addTextLabel(img *image.RGBA, x, y int, text string, textColor color.RGBA) {
	// Create larger text by drawing each pixel of the font as a larger block
	scale := 6 // Make each font pixel 6x6
	face := basicfont.Face7x13

	// Create a temporary small image to render the text first
	tempImg := image.NewRGBA(image.Rect(0, 0, 200, 50))
	drawer := &font.Drawer{
		Dst:  tempImg,
		Src:  image.NewUniform(textColor),
		Face: face,
	}

	// Draw text on temporary image
	textWidth := drawer.MeasureString(text).Round()
	drawer.Dot = fixed.Point26_6{
		X: fixed.I(10),
		Y: fixed.I(20),
	}
	drawer.DrawString(text)

	// Now scale up the temporary image and draw it onto the main image
	tempBounds := tempImg.Bounds()
	for ty := tempBounds.Min.Y; ty < tempBounds.Max.Y; ty++ {
		for tx := tempBounds.Min.X; tx < tempBounds.Max.X; tx++ {
			r, g, b, a := tempImg.At(tx, ty).RGBA()
			if a > 0 { // If there's a pixel here
				// Draw a scale x scale block for each pixel
				for sy := 0; sy < scale; sy++ {
					for sx := 0; sx < scale; sx++ {
						px := x - (textWidth*scale)/2 + (tx-10)*scale + sx
						py := y - 20*scale + (ty-tempBounds.Min.Y)*scale + sy
						if px >= 0 && py >= 0 && px < img.Bounds().Max.X && py < img.Bounds().Max.Y {
							img.Set(px, py, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 255})
						}
					}
				}
			}
		}
	}
}

func addStrategyKey(img *image.RGBA, strategies []AveragingStrategy, totalWidth, startY int) {
	keyText := ""
	for i, strategy := range strategies {
		if i > 0 {
			keyText += ", "
		}
		keyText += fmt.Sprintf("%d=%s", i+1, string(strategy))
	}

	// Add the key text centered at the bottom using the larger font function
	addTextLabel(img, totalWidth/2, startY+60, keyText, color.RGBA{255, 255, 255, 255})
}
