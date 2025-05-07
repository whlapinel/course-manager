package localfilesystem

import "errors"

var (
	ErrIsDirectory = errors.New("path is a directory") // for when user is attempting to download a directory
	ErrFileExists  = errors.New("file already exists") // for when user attempts to save a file that already exists (prevents overwriting)
)
