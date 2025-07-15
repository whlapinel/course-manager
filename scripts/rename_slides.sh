# Add prefix to filenames
find . -name "slides.html" -execdir mv {} ".slides.html" \;

find . -name "slides.md" -execdir mv {} ".slides.md" \;
