package localfilesystem

import "errors"

var (
	ErrFileExists = errors.New("file already exists") // for when user attempts to save a file that already exists (prevents overwriting)
)
