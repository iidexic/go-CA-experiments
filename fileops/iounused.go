package fileops

import (
	"errors"
	"os"
	"path/filepath"
)

var ErrorNilStat = errors.New("nil Stat but threw no error")

// PathExists attempts to check if a path exists
// returns true unless received a NotExist error from os.Stat
// This should work in most cases, but probably not all
func PathExists(path string) (bool, error) {

	path = filepath.Clean(path)
	s, e := os.Stat(path)
	if e != nil && errors.Is(e, os.ErrNotExist) {
		return false, nil
	} else if e != nil {
		return true, e
	}
	if s != nil {
		return true, nil
	}
	return true, ErrorNilStat
}

// PathExistsUsable simplifies PathExists; returns true if path exists and no error on os.Stat
func PathExistsUsable(path string) bool {
	exists, e := PathExists(path)
	if e != nil {
		return false
	}
	return exists
}
