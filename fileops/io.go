package fileops

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
)

var ErrorNilStat = errors.New("nil Stat but threw no error")
var ErrNotExist = os.ErrNotExist

// MakeOpenFileF will open the given fpath as a file. It will make the file if it does not exist,
// and it will make any missing directories necessary.
func MakeOpenFileF(fpath string) (*os.File, error) {
	// makedir->create+open|if exists->open
	e := os.MkdirAll(filepath.Dir(fpath), 0o755)
	if e != nil {
		return nil, e
	}
	file, e := os.OpenFile(fpath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0o666)
	if os.IsExist(e) {
		file, e = os.OpenFile(fpath, os.O_RDWR, 0)
		return file, e
	}
	return file, e
}

// ReadFile will read contents of file into a ReadResult object and return a ptr
// result contains file and/or operation outcome/error if e!=nil
func ReadFile(pathElements ...string) ([]byte, error) {
	fpath := filepath.Join(pathElements...)
	file, e := os.ReadFile(fpath)
	if e != nil {
		return nil, e
	}
	return file, nil
}

type QuickFile struct {
	path  string
	isdir bool
	cat   string
	name  string
}

type DirMap map[string]QuickFile

func (d DirMap) Add(path string, isdir bool, name string) bool {
	if _, ok := d[path]; ok {
		return false
	}
	d[path] = QuickFile{
		path:  path,
		isdir: isdir,
		name:  name,
	}
	return true
}

type filecat []string

var (
	extensions = map[string]filecat{
		"image": {".png", ".jpg", ".jpeg", ".gif", ".bmp", ".tiff", ".webp"},
		"audio": {".mp3", ".wav", ".ogg", ".flac", ".aac", ".m4a", ".wma", ".aiff", ".ape"},
		"text":  {".txt", ".md"},
		"data":  {".toml", ".json", ".yaml", ".yml", ".csv", ".tsv"},
	}
)

func DetermineFileType(file string) string {
	if ext := filepath.Ext(file); ext == "" || ext == "." {
		return "dir"
	} else {
		for k, v := range extensions {
			if slices.Contains(v, ext) {
				return k
			}
		}
	}
	return "other"
}

func (d DirMap) Ext(file string) string {
	if qf, ok := d[file]; ok {
		return filepath.Ext(qf.path)
	}
	return ""
}

func GetDirContents(path string) (DirMap, error) {
	path = filepath.Clean(path)
	dc := make(DirMap)
	dir := os.DirFS(path)
	err := fs.WalkDir(dir, ".", func(path string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		name := filepath.Base(path)
		dc.Add(path, d.IsDir(), name)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return dc, nil
}
