#!/bin/bash
# This script processes images in the original_images directory and saves them to the static/images directory.
# It uses ImageMagick to convert the images to a more web-friendly format and applies some optimizations.

for file in original_images/*.jpg; do
  filename=$(basename "$file") # Extract the filename from the path
  magick "$file" -strip -interlace Plane -gaussian-blur 0.05 -compress JPEG -quality 85 "static/images/$filename"
done

rm -rf original_images