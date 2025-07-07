#!/bin/bash

# Script to minify HTML, CSS, and JS files after Hugo build
# Install minify first: brew install tdewolff/tap/minify

cd hugodata

# Build Hugo site
echo "Building Hugo site..."
hugo

# Check if minify is installed
if ! command -v minify &> /dev/null; then
    echo "minify could not be found. Install it with: brew install tdewolff/tap/minify"
    exit 1
fi

# Minify HTML files
echo "Minifying HTML files..."
find public -name "*.html" -exec minify -o {} {} \;

# Minify CSS files
echo "Minifying CSS files..."
find public -name "*.css" -exec minify -o {} {} \;

# Minify JS files
echo "Minifying JS files..."
find public -name "*.js" -exec minify -o {} {} \;

echo "Minification complete!"

# Show size reduction
echo "Size comparison:"
echo "Before minification:"
du -sh public
echo "After minification:"
du -sh public