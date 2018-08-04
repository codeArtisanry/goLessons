package fs

import (
	"testing"
)

func TestZipFile(t *testing.T) {
	file := "a.txt"
	if err := ZipFile(file, ""); err != nil {
		t.Fatal(err)
	}
}

func TestZip(t *testing.T) {

	Zip("a.zip", "zip.go")
}
