#!/bin/bash

# Validate frame 250 by roundtripping through golang
# Usage: ./1-validate-frame-number-250.sh

cd "$(dirname "$0")"

echo "Building golang validation tool..."
go build -o tmp_validate_frame 2V_validate_frame.go

if [ $? -ne 0 ]; then
    echo "Error: Failed to build golang validation tool"
    exit 1
fi

echo "Running validation on frame 250..."
./tmp_validate_frame 250

if [ $? -eq 0 ]; then
    echo ""
    echo "Validation complete! Check the following files:"
    echo "- Original: intermediate_outputs/validation/original_250.png"  
    echo "- Recreated: intermediate_outputs/validation/recreated_250.png"
    echo ""
    echo "These files should be visually identical."
else
    echo "Error: Validation failed"
    exit 1
fi