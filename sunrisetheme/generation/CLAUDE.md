# BOBOLINK Project

## Objective
Clean implementation of video processing pipeline, starting from scratch.

## Input Files (DO NOT MODIFY)
- `data/video.mp4` - stabilized video input
- `data/mask.png` - mask for processing (white=keep, black=transparent)
- `data/mountains-only.png` - mountain reference image

## Processing Pipeline
1. Extract video frames to PNG format in `intermediate_outputs/frames/`
2. Apply mask to each frame (black areas become transparent)
3. Process first 270 frames
4. Validate frame processing with golang tool

## Masking Process
The masking uses ImageMagick's CopyOpacity composite operation:
- Mask: Grayscale image where white (255) = opaque, black (0) = transparent
- Command: `magick input.png mask.png -alpha off -compose CopyOpacity -composite output.png`
- Result: 8-bit/color RGBA PNG with true alpha channel
- The mask's grayscale values directly become the alpha channel values

### Color Encoding Details
- Input frames: RGB 8-bit per channel (24-bit color)
- Output frames: RGBA 8-bit per channel (32-bit color with alpha)
- Color space: sRGB
- Alpha channel: 8-bit (0-255), where 0=fully transparent, 255=fully opaque
- File format: PNG with RGBA color type, non-interlaced

## Numbered Scripts Workflow

### Step 1: Extract Frames
Run `./1_extract_frames.sh` to:
- Clean frames directory
- Extract frames 1-270 from video.mp4 to PNG format
- Apply mask.png to each frame (areas outside mask become transparent)
- Save processed frames to intermediate_outputs/frames/

Run `./1_extract_frames.sh [frame_number]` to apply mask to single frame for testing

### Step 2: Validate Processing
Run `./1-validate-frame-number-250.sh` to:
- Build golang validation tool
- Read frame 250 pixel by pixel
- Recreate PNG from extracted pixel data
- Save both original and recreated versions for comparison
- Verify golang can properly ingest the PNG format

### Step 3: Analyze Frames for HSL Data
Run `./3_analyze_frames.sh` to:
- Build row_analyzer.go tool
- Process all 270 frames in parallel batches
- Extract HSL data from evenly distributed rows
- Save CSV data to `intermediate_outputs/row_data/[frame_number]/[K].csv`

The analysis process:
1. Identifies rows with >5 non-transparent pixels
2. Selects 21 evenly distributed rows from valid rows
3. Converts all pixels in each row from RGB to HSL
4. Saves HSL values as CSV files (H, S, L columns)
5. Processes frames in parallel for performance

### Step 4: Visualization System
Use visualization tools to compare averaging strategies:

**Single Frame Visualization:**
```bash
go run 4_averaging_lib.go 4_visualizer_main.go <frame_number> <strategy1> [strategy2] ...
```

**Grid Visualization (Multiple Frames):**
```bash
go run 4_averaging_lib.go 5_visualizer_grid.go <strategy1> [strategy2] ...
```

## Averaging Strategies

### Available Strategies
1. **mean** - Standard mean averaging of all channels
2. **mode** - Most frequent value per channel (with bucketing)
3. **median** - Middle value per channel
4. **t3lmean** - Top Third Lightness Mean (brightest 1/3 of pixels)
5. **m3lmean** - Middle Third Lightness Mean (moderate brightness pixels)
6. **hbmodeslmean** - Hue Bucketized Mode + Saturation/Lightness Mean ⭐ **RECOMMENDED**

### HBModeSLMean Strategy (Recommended)
This sophisticated strategy treats color channels differently:

**Hue (H)**: Uses bucketized mode algorithm
- Starts with 1° buckets, increases to find 15% threshold
- Tests all offsets to eliminate bias
- Returns median of largest qualifying bucket
- Handles circular hue nature and wraparound

**Saturation (S) & Lightness (L)**: Uses simple mean averaging

**Why it works well:**
- Identifies dominant hue ranges in each row
- Resists hue outliers and noise
- Provides stable, perceptually meaningful colors
- Balances statistical rigor with practical results

## Visualization Features

### Test Strips and Gradients
- **4 test strip positions**: Evenly spaced across image width
- **Thin strips**: Minimal visual interference with original image
- **Gradient strips**: Full 150px width reference gradients on right
- **RGB interpolation**: Standard linear interpolation between row averages

### Labeling System
- **Strategy numbers**: Large, readable numbers below each strip
- **Strategy key**: Centered at bottom with complete mapping
- **6x scaled fonts**: Highly visible text for presentations

### Grid Layout
- **Multiple frames**: 30, 60, 90, 120, 150, 180, 210, 240, 270
- **3x3 grid**: Comprehensive view across sunrise sequence
- **Consistent strategies**: Same averaging applied to all frames

## Color Processing Details

### HSL Color Space
- **H (Hue)**: 0-360° circular color representation
- **S (Saturation)**: 0-1 color intensity
- **L (Lightness)**: 0-1 brightness level
- **Conversion**: RGB ↔ HSL for analysis and display

### Row Selection Logic
- **Valid rows**: More than 5 non-transparent pixels
- **Even distribution**: 21 rows selected across valid range
- **Gradient height**: From first to last valid row

## Tools
- `row_analyzer.go` - Extracts HSL data from PNG frames to CSV
- `4_averaging_lib.go` - Core averaging strategy implementations
- `4_visualizer_main.go` - Single frame visualization with test strips
- `5_visualizer_grid.go` - Multi-frame grid comparison
- `validate_frame.go` - PNG validation tool (legacy)

## Output Structure
```
intermediate_outputs/
├── frames/                    # Masked PNG frames (001.png - 270.png)
├── row_data/                  # HSL analysis data
│   └── [frame_number]/        # Per-frame directories
│       └── [0-20].csv         # Row HSL data (H,S,L columns)
├── validation/                # Frame validation outputs
└── grid_visualization_hbmodeslmean_mean.png  # Final comparison
```

## Rules
- Never modify input files
- Always derive outputs from inputs
- Keep all processing in BOBOLINK directory
- Use PNG format for golang compatibility
- Numbered scripts indicate processing order
- Use HBModeSLMean strategy for final color analysis