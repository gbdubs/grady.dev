#!/bin/bash

# Extract frames from video to PNG format with mask applied
# Usage: ./1_extract_frames.sh [frame_number]
# If frame_number is provided, only process that specific frame

cd "$(dirname "$0")"

if [ $# -eq 1 ]; then
    # Single frame mode
    FRAME_NUM=$1
    FRAME_FILE=$(printf "%03d.png" $FRAME_NUM)
    
    echo "Single frame mode: processing frame $FRAME_NUM"
    
    if [ ! -f "intermediate_outputs/frames/$FRAME_FILE" ]; then
        echo "Error: Frame $FRAME_FILE does not exist. Run without parameters to extract all frames first."
        exit 1
    fi
    
    echo "Applying mask to frame $FRAME_NUM..."
    magick "intermediate_outputs/frames/$FRAME_FILE" data/mask.png -alpha off -compose CopyOpacity -composite "intermediate_outputs/frames/$FRAME_FILE"
    echo "Masking complete for frame $FRAME_NUM"
    
else
    # Full extraction mode
    echo "Cleaning frames directory..."
    rm -rf intermediate_outputs/frames/*

    echo "Extracting frames from video.mp4 to PNG format (up to frame 270)..."

    # Use ffmpeg to extract frames as PNG, limiting to 270 frames
    ffmpeg -i data/video.mp4 -frames:v 270 -f image2 intermediate_outputs/frames/%03d.png

    echo "Applying mask to extracted frames..."

    # Apply mask to each frame using ImageMagick
    for frame in intermediate_outputs/frames/*.png; do
        if [[ -f "$frame" ]]; then
            echo "Processing $(basename "$frame")..."
            # Apply mask to create transparency where mask is black (mountains)
            magick "$frame" data/mask.png -alpha off -compose CopyOpacity -composite "$frame"
        fi
    done

    echo "Frame extraction and masking complete."
    echo "270 PNG frames saved to intermediate_outputs/frames/"
fi