#!/bin/bash

# Step 3: Analyze frames using row_analyzer.go
# Processes all frames in intermediate_outputs/frames/ to extract HSL data

echo "Starting frame analysis..."

# Build the row analyzer
echo "Building row analyzer..."
go build -o tmp_row_analyzer 3_row_analyzer.go

if [ $? -ne 0 ]; then
    echo "Failed to build row analyzer"
    exit 1
fi

# Get the frame directory
FRAME_DIR="intermediate_outputs/frames"

if [ ! -d "$FRAME_DIR" ]; then
    echo "Frame directory not found: $FRAME_DIR"
    echo "Please run 1_extract_frames.sh first"
    exit 1
fi

# Count total frames
TOTAL_FRAMES=$(ls -1 "$FRAME_DIR"/*.png | wc -l | tr -d ' ')
echo "Found $TOTAL_FRAMES frames to process"

# Create array of frame paths and numbers for parallel processing
FRAME_PATHS=()
FRAME_NUMBERS=()

for frame_file in "$FRAME_DIR"/*.png; do
    if [ -f "$frame_file" ]; then
        # Extract frame number from filename (remove leading zeros)
        frame_basename=$(basename "$frame_file" .png)
        frame_number=$((10#$frame_basename))
        
        FRAME_PATHS+=("$frame_file")
        FRAME_NUMBERS+=("$frame_number")
    fi
done

# Process frames in parallel batches
BATCH_SIZE=10
CURRENT_BATCH=0

echo "Processing frames in batches of $BATCH_SIZE..."

for ((i=0; i<${#FRAME_PATHS[@]}; i+=BATCH_SIZE)); do
    CURRENT_BATCH=$((CURRENT_BATCH + 1))
    BATCH_END=$((i + BATCH_SIZE - 1))
    
    if [ $BATCH_END -ge ${#FRAME_PATHS[@]} ]; then
        BATCH_END=$((${#FRAME_PATHS[@]} - 1))
    fi
    
    echo "Processing batch $CURRENT_BATCH: frames $((i+1)) to $((BATCH_END+1))"
    
    # Start background processes for this batch
    for ((j=i; j<=BATCH_END; j++)); do
        if [ $j -lt ${#FRAME_PATHS[@]} ]; then
            frame_path="${FRAME_PATHS[$j]}"
            frame_number="${FRAME_NUMBERS[$j]}"
            
            # Run analyzer in background
            ./tmp_row_analyzer "$frame_path" "$frame_number" &
        fi
    done
    
    # Wait for all processes in this batch to complete
    wait
    
    echo "Batch $CURRENT_BATCH completed"
done

echo "All frames processed successfully!"
echo "HSL data saved to intermediate_outputs/row_data/"

# Clean up
rm -f row_analyzer

echo "Analysis complete!"