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
func TestGetCrf(t *testing.T) {
	code := "h265"
	width := 1080
	height := 1920
	crf := GetCrf(code, width, height)
	t.Log(crf)
}
