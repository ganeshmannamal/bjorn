package util

import "path/filepath"

func ResolvePath(parent string, path string) string {
	if filepath.IsAbs(path) {
		return path
	}

	pathBase := filepath.Dir(path)

	if pathBase == ".." {
		return filepath.Join(filepath.Dir(parent), filepath.Base(path))
	}

	if pathBase == "." {
		return filepath.Join(parent, filepath.Base(path))
	}

	return filepath.Join(parent, path)
}
