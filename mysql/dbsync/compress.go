// Copyright 2015 Aller Media AS.  All rights reserved.
// License: GPL3
package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

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
