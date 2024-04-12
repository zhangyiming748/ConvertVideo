package util

import (
	"testing"
)

func TestGetAllFiles(t *testing.T) {
	dir := "/Users/zen/Downloads"
	files := GetAllFiles(dir)
	for _, file := range files {
		t.Log(file)
	}
}
