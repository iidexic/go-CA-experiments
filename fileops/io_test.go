package fileops

import (
	"path/filepath"
	"testing"
)

func TestDetermineFileType(t *testing.T) {
	tests := []struct {
		name string
		file string
		want string
	}{
		{
			name: "image",
			file: "test.png",
			want: "image",
		},
		{
			name: "audio",
			file: "test.mp3",
			want: "audio",
		},
		{
			name: "text",
			file: "test.txt",
			want: "text",
		},
		{
			name: "data",
			file: "test.toml",
			want: "data",
		},
		{
			name: "other",
			file: "test.exe",
			want: "other",
		},
		{
			name: "dir",
			file: "test",
			want: "dir",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetermineFileType(tt.file); got != tt.want {
				t.Errorf("DetermineFileType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCatchDir(t *testing.T) {
	testDir := "../Resources/"
	t.Logf("path = %s, cleanpath = %s\n", testDir, filepath.Clean(testDir))
	cont, e := GetDirContents(testDir)
	if e != nil {
		t.Errorf("GetDirContents() error = %v, want nil", e)
	}
	for k, v := range cont {
		cat := DetermineFileType(k)
		t.Logf("%s [name %s, isdir %t, cat %s]", k, v.name, v.isdir, cat)
	}
}
