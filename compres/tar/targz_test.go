package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestTarAndGz(t *testing.T) {
	z, err := os.OpenFile("a.tar.gz", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0744)
	if err != nil {
		return
	}
	defer z.Close()
	if err := TarAndGz("../zip", z); err != nil {
		// t.Logf("TarAndGz() error = %v, wantErr %v", err)
		return
	}
}

func TarAndGz(dir string, w io.Writer) error {
	zipf := gzip.NewWriter(w)
	defer zipf.Close()
	tarf := tar.NewWriter(zipf)
	defer func() {
		if err := tarf.Close(); err != nil {
			panic(err)
		}
	}()

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	fmt.Println(files)
	for _, f := range files {
		of, err := os.Open(filepath.Join(dir, string(filepath.Separator), f.Name()))
		if err != nil {
			return err
		}

		th := &tar.Header{
			Name: f.Name(),
			Size: f.Size(),
			Mode: int64(f.Mode()),
		}

		if err := tarf.WriteHeader(th); err != nil {
			of.Close()
			return err
		}

		io.Copy(tarf, of)

		of.Close()
	}

	return nil
}
