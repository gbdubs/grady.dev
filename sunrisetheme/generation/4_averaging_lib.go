package main

import (
	"encoding/csv"
	"fmt"
	"image/png"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type HSLColorAveraging struct {
	H, S, L float64
}

type AveragingStrategy string

const (
	MEAN      AveragingStrategy = "mean"
	MODE      AveragingStrategy = "mode"
	MEDIAN    AveragingStrategy = "median"
	T3LMEAN   AveragingStrategy = "t3lmean"
	M3LMEAN   AveragingStrategy = "m3lmean"
	HBMODESLMEAN AveragingStrategy = "hbmodeslmean"
)

type FrameData struct {
	FrameNumber int
	RowData     [][]HSLColorAveraging
}

type AveragingResult struct {
	Strategy AveragingStrategy
	Colors   []HSLColorAveraging
}

func LoadFrameData(frameNumber int) (*FrameData, error) {
	frameDir := fmt.Sprintf("intermediate_outputs/row_data/%d", frameNumber)
	
	if _, err := os.Stat(frameDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("frame data not found for frame %d", frameNumber)
	}

	var rowData [][]HSLColorAveraging
	
	for k := 0; k < 21; k++ {
		csvPath := filepath.Join(frameDir, fmt.Sprintf("%d.csv", k))
		colors, err := loadCSVColors(csvPath)
		if err != nil {
			return nil, fmt.Errorf("failed to load CSV %s: %w", csvPath, err)
		}
		rowData = append(rowData, colors)
	}

	return &FrameData{
		FrameNumber: frameNumber,
		RowData:     rowData,
	}, nil
}

func loadCSVColors(csvPath string) ([]HSLColorAveraging, error) {
	file, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var colors []HSLColorAveraging
	for i, record := range records {
		if i == 0 {
			continue
		}
		
		if len(record) != 3 {
			continue
		}

		h, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			continue
		}
		s, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			continue
		}
		l, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			continue
		}

		colors = append(colors, HSLColorAveraging{H: h, S: s, L: l})
	}

	return colors, nil
}

func ApplyAveragingStrategy(frameData *FrameData, strategy AveragingStrategy) AveragingResult {
	var colors []HSLColorAveraging
	
	for _, rowColors := range frameData.RowData {
		if len(rowColors) == 0 {
			continue
		}
		
		var avgColor HSLColorAveraging
		switch strategy {
		case MEAN:
			avgColor = calculateMean(rowColors)
		case MODE:
			avgColor = calculateMode(rowColors)
		case MEDIAN:
			avgColor = calculateMedian(rowColors)
		case T3LMEAN:
			avgColor = calculateT3LMean(rowColors)
		case M3LMEAN:
			avgColor = calculateM3LMean(rowColors)
		case HBMODESLMEAN:
			avgColor = calculateHBModeSLMean(rowColors)
		default:
			avgColor = calculateMean(rowColors)
		}
		
		colors = append(colors, avgColor)
	}

	return AveragingResult{
		Strategy: strategy,
		Colors:   colors,
	}
}

func calculateMean(colors []HSLColorAveraging) HSLColorAveraging {
	if len(colors) == 0 {
		return HSLColorAveraging{}
	}

	var hSum, sSum, lSum float64
	for _, color := range colors {
		hSum += color.H
		sSum += color.S
		lSum += color.L
	}

	count := float64(len(colors))
	return HSLColorAveraging{
		H: hSum / count,
		S: sSum / count,
		L: lSum / count,
	}
}

func calculateMode(colors []HSLColorAveraging) HSLColorAveraging {
	if len(colors) == 0 {
		return HSLColorAveraging{}
	}

	hMode := calculateChannelMode(colors, func(c HSLColorAveraging) float64 { return c.H })
	sMode := calculateChannelMode(colors, func(c HSLColorAveraging) float64 { return c.S })
	lMode := calculateChannelMode(colors, func(c HSLColorAveraging) float64 { return c.L })

	return HSLColorAveraging{H: hMode, S: sMode, L: lMode}
}

func calculateChannelMode(colors []HSLColorAveraging, extractor func(HSLColorAveraging) float64) float64 {
	if len(colors) == 0 {
		return 0
	}

	buckets := make(map[int]int)
	for _, color := range colors {
		value := extractor(color)
		bucket := int(math.Round(value * 100))
		buckets[bucket]++
	}

	maxCount := 0
	modeBucket := 0
	for bucket, count := range buckets {
		if count > maxCount {
			maxCount = count
			modeBucket = bucket
		}
	}

	return float64(modeBucket) / 100.0
}

func calculateMedian(colors []HSLColorAveraging) HSLColorAveraging {
	if len(colors) == 0 {
		return HSLColorAveraging{}
	}

	hMedian := calculateChannelMedian(colors, func(c HSLColorAveraging) float64 { return c.H })
	sMedian := calculateChannelMedian(colors, func(c HSLColorAveraging) float64 { return c.S })
	lMedian := calculateChannelMedian(colors, func(c HSLColorAveraging) float64 { return c.L })

	return HSLColorAveraging{H: hMedian, S: sMedian, L: lMedian}
}

func calculateT3LMean(colors []HSLColorAveraging) HSLColorAveraging {
	if len(colors) == 0 {
		return HSLColorAveraging{}
	}

	// Sort colors by lightness (L value)
	sortedColors := make([]HSLColorAveraging, len(colors))
	copy(sortedColors, colors)
	
	sort.Slice(sortedColors, func(i, j int) bool {
		return sortedColors[i].L > sortedColors[j].L // Sort descending by lightness
	})

	// Take the top third (brightest colors)
	topThirdCount := len(sortedColors) / 3
	if topThirdCount == 0 {
		topThirdCount = 1 // At least one color
	}
	
	topThirdColors := sortedColors[:topThirdCount]
	
	// Calculate mean of the top third
	return calculateMean(topThirdColors)
}

func calculateM3LMean(colors []HSLColorAveraging) HSLColorAveraging {
	if len(colors) == 0 {
		return HSLColorAveraging{}
	}

	// Sort colors by lightness (L value)
	sortedColors := make([]HSLColorAveraging, len(colors))
	copy(sortedColors, colors)
	
	sort.Slice(sortedColors, func(i, j int) bool {
		return sortedColors[i].L > sortedColors[j].L // Sort descending by lightness
	})

	// Calculate middle third range
	totalCount := len(sortedColors)
	thirdSize := totalCount / 3
	if thirdSize == 0 {
		thirdSize = 1 // At least one color
	}
	
	// Take the middle third
	startIndex := thirdSize
	endIndex := startIndex + thirdSize
	
	// Handle edge cases for small arrays
	if endIndex > totalCount {
		endIndex = totalCount
	}
	if startIndex >= totalCount {
		startIndex = totalCount - 1
		endIndex = totalCount
	}
	
	middleThirdColors := sortedColors[startIndex:endIndex]
	
	// Calculate mean of the middle third
	return calculateMean(middleThirdColors)
}

func calculateHBModeSLMean(colors []HSLColorAveraging) HSLColorAveraging {
	if len(colors) == 0 {
		return HSLColorAveraging{}
	}

	// Use bucketized mode for hue
	hue := bucketizedModeHue(colors)
	
	// Use simple mean for saturation and lightness
	var sSum, lSum float64
	for _, color := range colors {
		sSum += color.S
		lSum += color.L
	}
	
	count := float64(len(colors))
	return HSLColorAveraging{
		H: hue,
		S: sSum / count,
		L: lSum / count,
	}
}

func bucketizedModeHue(colors []HSLColorAveraging) float64 {
	if len(colors) == 0 {
		return 0
	}
	
	if len(colors) == 1 {
		return colors[0].H
	}

	threshold := float64(len(colors)) * 0.15 // 15% threshold
	maxBucketSize := 180 // Don't go beyond 180 degrees bucket size
	
	// Try increasing bucket sizes
	for bucketSize := 1; bucketSize <= maxBucketSize; bucketSize++ {
		bestBucket, found := findBestBucketForSize(colors, bucketSize, threshold)
		if found {
			return calculateMedianInBucket(colors, bestBucket, bucketSize)
		}
	}
	
	// Fallback to simple mean if no bucket found
	var hSum float64
	for _, color := range colors {
		hSum += color.H
	}
	return hSum / float64(len(colors))
}

func findBestBucketForSize(colors []HSLColorAveraging, bucketSize int, threshold float64) (int, bool) {
	bestBucketStart := -1
	maxBucketCount := 0
	
	// Try all possible offsets
	for offset := 0; offset < bucketSize; offset++ {
		bucketCounts := make(map[int]int)
		
		// Count hues in buckets with this offset
		for _, color := range colors {
			adjustedHue := math.Mod(color.H + float64(offset), 360)
			bucketIndex := int(adjustedHue) / bucketSize
			bucketCounts[bucketIndex]++
		}
		
		// Find the largest bucket with this offset
		for bucketIndex, count := range bucketCounts {
			if float64(count) >= threshold && count > maxBucketCount {
				maxBucketCount = count
				// Convert back to original coordinate system
				bestBucketStart = (bucketIndex*bucketSize - offset + 360) % 360
			}
		}
	}
	
	return bestBucketStart, maxBucketCount > 0
}

func calculateMedianInBucket(colors []HSLColorAveraging, bucketStart, bucketSize int) float64 {
	var bucketHues []float64
	
	bucketEnd := (bucketStart + bucketSize) % 360
	
	for _, color := range colors {
		hue := color.H
		
		// Handle wraparound case
		if bucketStart > bucketEnd {
			// Bucket wraps around 0/360
			if hue >= float64(bucketStart) || hue < float64(bucketEnd) {
				bucketHues = append(bucketHues, hue)
			}
		} else {
			// Normal bucket
			if hue >= float64(bucketStart) && hue < float64(bucketEnd) {
				bucketHues = append(bucketHues, hue)
			}
		}
	}
	
	if len(bucketHues) == 0 {
		return float64(bucketStart + bucketSize/2) // Fallback to bucket center
	}
	
	// Calculate median, handling wraparound
	if bucketStart > bucketEnd {
		// For wraparound buckets, normalize to 0-360 range for median calculation
		var normalizedHues []float64
		for _, hue := range bucketHues {
			if hue >= float64(bucketStart) {
				normalizedHues = append(normalizedHues, hue - float64(bucketStart))
			} else {
				normalizedHues = append(normalizedHues, hue + 360 - float64(bucketStart))
			}
		}
		sort.Float64s(normalizedHues)
		median := normalizedHues[len(normalizedHues)/2]
		return math.Mod(median + float64(bucketStart), 360)
	} else {
		// Normal case
		sort.Float64s(bucketHues)
		return bucketHues[len(bucketHues)/2]
	}
}

func calculateChannelMedian(colors []HSLColorAveraging, extractor func(HSLColorAveraging) float64) float64 {
	if len(colors) == 0 {
		return 0
	}

	values := make([]float64, len(colors))
	for i, color := range colors {
		values[i] = extractor(color)
	}

	sort.Float64s(values)

	n := len(values)
	if n%2 == 0 {
		return (values[n/2-1] + values[n/2]) / 2
	}
	return values[n/2]
}

func ParseStrategies(strategyCodes []string) ([]AveragingStrategy, error) {
	var strategies []AveragingStrategy
	
	for _, code := range strategyCodes {
		switch strings.ToLower(code) {
		case "mean":
			strategies = append(strategies, MEAN)
		case "mode":
			strategies = append(strategies, MODE)
		case "median":
			strategies = append(strategies, MEDIAN)
		case "t3lmean":
			strategies = append(strategies, T3LMEAN)
		case "m3lmean":
			strategies = append(strategies, M3LMEAN)
		case "hbmodeslmean":
			strategies = append(strategies, HBMODESLMEAN)
		default:
			return nil, fmt.Errorf("unknown strategy code: %s", code)
		}
	}
	
	return strategies, nil
}

func FindValidRowBounds(frameNumber int) (int, int, error) {
	framePath := fmt.Sprintf("intermediate_outputs/frames/%03d.png", frameNumber)
	
	file, err := os.Open(framePath)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return 0, 0, err
	}

	bounds := img.Bounds()
	minRow := bounds.Max.Y
	maxRow := bounds.Min.Y

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		pixelCount := 0
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			_, _, _, a := img.At(x, y).RGBA()
			if a > 0 {
				pixelCount++
			}
		}
		if pixelCount > 5 {
			if y < minRow {
				minRow = y
			}
			if y > maxRow {
				maxRow = y
			}
		}
	}

	if minRow > maxRow {
		return 0, 0, fmt.Errorf("no valid rows found")
	}

	return minRow, maxRow, nil
}

func HSLToRGB(h, s, l float64) (uint8, uint8, uint8) {
	h = h / 360.0
	
	if s == 0 {
		c := uint8(l * 255)
		return c, c, c
	}

	var q float64
	if l < 0.5 {
		q = l * (1 + s)
	} else {
		q = l + s - l*s
	}
	
	p := 2*l - q
	
	r := hueToRGB(p, q, h+1.0/3.0)
	g := hueToRGB(p, q, h)
	b := hueToRGB(p, q, h-1.0/3.0)

	return uint8(r * 255), uint8(g * 255), uint8(b * 255)
}

func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t++
	}
	if t > 1 {
		t--
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