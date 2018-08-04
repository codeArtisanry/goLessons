package fs

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//Zipa um arquivo e o salva na pasta dir.
//O nome do arquivo gerado é o mesmo que path, exceto pela extensão, que é zip.
//Se dir for vazio, salva no mesmo diretório que file.
func ZipFile(path, dir string) error {
	if dir == "" {
		dir = filepath.Dir(path)
	}
	nome := filepath.Base(path)
	pathZip := filepath.Join(dir, nome[:strings.LastIndex(nome, ".")]+".zip")
	z, err := os.OpenFile(pathZip, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0744)
	if err != nil {
		return err
	}
	defer z.Close()

	w := zip.NewWriter(z)
	defer w.Close()
	f, err := w.Create(nome)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	f.Write(data)
	return nil
}

func Zip(zipname, srcname string) error {
	z, err := os.OpenFile(zipname, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0744)
	if err != nil {
		return nil
	}
	defer z.Close()
	ziper := zip.NewWriter(z)
	defer ziper.Close()

	f, err := ziper.Create(srcname)
	if err != nil {
		return nil
	}
	data, err := ioutil.ReadFile(srcname)
	if err != nil {
		return nil
	}
	f.Write(data)
	return nil
}

func unzip(filename string, password string) bool {
	// open a zip archive for reading
	r, err := zip.OpenReader(filename)
	if err != nil {
		return false
	}
	defer r.Close()
	// create a buffer to write archive to
	buffer := new(bytes.Buffer)
	// iterate through the files in the archive
	for _, f := range r.File {
		f.SetPassword(password)
		// open a file for reading
		r, err := f.Open()
		if err != nil {
			return false
		}
		defer r.Close()
		// try to copy file data to buffer
		n, err := io.Copy(buffer, r)
		if n == 0 || err != nil {
			return false
		}
		break
	}
	return true
}

// func NetCon() {
// 	/*
// 		在这个时候，文件已经不再局限于io，可以是一个内存buffer，
// 		也可以是一个计算hash的对象，甚至是一个计数器，流量限速器。
// 		golang灵活的接口机制为我们提供了无限可能。
// 	*/
// 	conn, _ := net.Dial("tcp", "127.0.0.1:8000")
// 	//加上一个zip压缩，还可以利用加上crypto/aes来个AES加密...
// 	zip := zlib.NewWriter(conn)
// 	//对socket加上一个buffer来增加吞吐量
// 	bufconn := bufio.NewWriter(zip)
// 	ioutils.EncodePacket3(bufconn, []byte("hello,client \r\n"))
// }
