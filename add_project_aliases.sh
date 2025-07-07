#!/bin/bash

# Add aliases to all project markdown files
for file in hugodata/content/projects/*.md; do
    # Skip _index.md
    if [[ "$file" == *"_index.md" ]]; then
        continue
    fi
    
    # Get the filename without path and extension
    filename=$(basename "$file" .md)
    
    # Check if aliases already exists in the file
    if grep -q "^aliases:" "$file"; then
        echo "Aliases already exists in $file, skipping..."
        continue
    fi
    
    # Add the alias after the front matter opening
    # Find the line number of the first ---
    line_num=$(grep -n "^---$" "$file" | head -1 | cut -d: -f1)
    
    if [ -n "$line_num" ]; then
        # Insert aliases after the first ---
        sed -i '' "${line_num}a\\
aliases: [\"/project/${filename}/\"]\\
" "$file"
        echo "Added alias /project/${filename}/ to $file"
    else
        echo "Could not find front matter in $file"
    fi
done

echo "Done adding aliases!"