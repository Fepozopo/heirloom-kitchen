#!/bin/bash
# This script processes images in the original_images directory and saves them to the static/images directory.
# It uses ImageMagick to convert the images to a more web-friendly format and applies some optimizations.

for file in new_images/*.jpg; do
  filename=$(basename "$file") # Extract the filename from the path
  magick "$file" -strip -interlace Plane -gaussian-blur 0.05 -compress JPEG -quality 85 "static/images/$filename"
done

# Move all the new images to the backup directory
mv new_images/* backup/original_images/
# Remove the new_images directory
rm -rf new_images/

echo "Image processing complete. New images have been backed up."