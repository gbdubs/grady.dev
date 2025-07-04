# Sunrise Theme - Gradient Generation Documentation

## Overview
The Sunrise theme implements a sophisticated gradient background system for project pages that uses actual sunrise animation data from the BOBOLINK project. Each content block (paragraph, header, etc.) gets two gradients derived from sunrise frame data, creating a seamless visual progression down the page.

## Key Concepts

### Fencepost Gradient Logic
The gradients work like "fenceposts" where N content blocks require N+1 gradients:
- Block 0 uses gradient 0 (top) and gradient 1 (bottom)
- Block 1 uses gradient 1 (top) and gradient 2 (bottom)
- Block 2 uses gradient 2 (top) and gradient 3 (bottom)
- etc.

This creates seamless transitions where each block's top gradient matches the previous block's bottom gradient.

### Two-Layer Gradient System
Each content block has two gradient layers:
1. **Background layer** (horizontal gradient): The actual sunrise colors from sky to horizon
2. **Foreground layer** (vertical gradient with opacity mask): Creates a fade effect from transparent to the next gradient

## Data Flow

### 1. Source Data
- **Input**: `../intermediate_outputs/all_frame_matrix_hbmodeslmean.csv`
- **Format**: CSV with columns: Frame, Breakpoint, H, S, L
- **Content**: 270 frames × 26 breakpoints = 7,020 HSL color values
- **Strategy**: HBModeSLMean (Hue Bucketized Mode + Saturation/Lightness Mean)

### 2. Data Extraction (`extract_matrix_data.go`)
This Go program:
1. Reads the CSV matrix file
2. Converts HSL values to RGB for web display
3. Pre-computes gradient configurations for 3-270 blocks
4. Implements the fencepost logic for gradient assignment
5. Outputs JSON to `exampleSite/data/sunrise_matrix.json`

Key logic:
```go
// For N blocks, we need N+1 gradients (fenceposts)
// Block B uses gradient B as top (horizontal) and gradient B+1 as bottom (vertical)
frameForTop := int(float64(blockIdx) / float64(numBlocks) * float64(maxFrame-1))
frameForBottom := int(float64(blockIdx+1) / float64(numBlocks) * float64(maxFrame-1))
```

### 3. Hugo Template Processing

#### Content Parsing (`layouts/partials/content-parser.html`)
- Splits HTML content into individual blocks by closing tags
- Handles: `</h1>`, `</h2>`, `</h3>`, `</h4>`, `</h5>`, `</h6>`, `</p>`, `</ul>`, `</ol>`, `</blockquote>`, `</pre>`, `</div>`, `</figure>`
- Returns array of HTML content blocks

#### Project Template (`layouts/projects/single.html`)
1. Loads sunrise matrix data from JSON (fails loudly if missing)
2. Parses content into blocks using content-parser partial
3. Adds navigation and title blocks at the beginning
4. Finds matching gradient configuration for the number of blocks
5. Renders each block with its assigned gradients

#### Gradient Rendering
For each content block:
```html
<div class="content-block">
    <div class="block-background" style="background: {{ $horizGradient | safeCSS }};"></div>
    <div class="block-foreground" style="background: {{ $vertGradient | safeCSS }};"></div>
    <div class="block-content">
        {{ $blockContent | safeHTML }}
    </div>
</div>
```

The CSS mask creates the opacity fade:
```css
.block-foreground {
    -webkit-mask-image: linear-gradient(to bottom, transparent 0%, black 100%);
    mask-image: linear-gradient(to bottom, transparent 0%, black 100%);
}
```

## Gradient Details

### Color Stops
- Each gradient has 26 color stops (from the 26 breakpoints)
- Colors are evenly distributed across the gradient (0% to 100%)
- Format: `rgb(r,g,b) offset%`

### Direction
- Both gradients go "to right" (horizontal)
- The vertical fade effect is achieved through CSS mask, not gradient direction

### Frame Selection
- Frames are distributed evenly across the total number of blocks
- For 10 blocks: frames 1, 30, 60, 90, 120, 150, 180, 210, 240, 270
- For 50 blocks: frames 1, 6, 11, 16, 21, 26, 31, 36, 41, 46...

## Building and Deployment

### Prerequisites
1. Complete BOBOLINK sunrise data pipeline (steps 1-6)
2. Generated CSV matrix file with HBModeSLMean averaging

### Build Process
1. Run data extraction:
   ```bash
   cd sunrise-theme
   go run extract_matrix_data.go
   ```
   This creates `exampleSite/data/sunrise_matrix.json`

2. Build Hugo site:
   ```bash
   cd exampleSite
   hugo
   ```

### Validation
- Template fails with clear error if sunrise_matrix.json is missing
- No fallback gradients - system requires real sunrise data
- Check `public/projects/*/index.html` for generated gradient CSS

## Technical Notes

### Performance
- Gradients are pre-computed in Go, not calculated in Hugo templates
- JSON includes configurations for 3-270 blocks to avoid runtime computation
- All 26 color stops are included for visual accuracy

### Browser Compatibility
- Uses standard CSS linear-gradient (wide support)
- Includes -webkit-mask-image prefix for Safari
- RGB color format for maximum compatibility

### Hugo Security
- All dynamic CSS values use `safeCSS` filter to prevent XSS
- HTML content uses `safeHTML` for proper rendering

## Example Output

For a page with 5 content blocks:
- Block 0: Frame 1 (top) to Frame 54 (bottom)
- Block 1: Frame 54 (top) to Frame 108 (bottom)
- Block 2: Frame 108 (top) to Frame 162 (bottom)
- Block 3: Frame 162 (top) to Frame 216 (bottom)
- Block 4: Frame 216 (top) to Frame 270 (bottom)

Each frame provides 26 color stops creating smooth sky-to-horizon gradients that progress through the sunrise sequence as you scroll down the page.