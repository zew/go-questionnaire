#!/bin/bash

# Usage:
#   ./hash_sha512.sh [ext1 ext2 ...]
#   If no extensions are given, only *.pdf and *.xlsx files are hashed (excluding *.sha512)


extensions=("$@")

# Default to "pdf" and "xlsx" if no extensions provided
if [ ${#extensions[@]} -eq 0 ]; then
  extensions=("pdf" "xlsx")
fi


echo "effective extensions ${extensions[@]}"


/usr/bin/find ./ -type f -name "*.sha512" -exec rm -f {} +
echo "deletions complete"

/usr/bin/find  ./  -type f | while read file; do
  case "$file" in *.sha512)
      continue
      ;;
  esac

  matched=0
  for ext in "${extensions[@]}"; do
    case "$file" in *."$ext")
        matched=1
        break
        ;;
    esac
  done

  if [ "$matched" -eq 1 ]; then
    hashFile="${file}.sha512"
    if [ -f "$hashFile" ]; then
      echo "  exists $hashFile"
      continue
    fi
    echo "  writing $hashFile"
    sha512sum "$file" > "$hashFile"
  fi
done


echo "end"
