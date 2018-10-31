package md5

import (
	"bufio"
	"crypto"
	"encoding/hex"
	"io"
	"os"
)

// 计算md5
func Md5sum(filename string) (string, error) {
	md5 := crypto.MD5.New()
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	r := bufio.NewReader(f)

	_, err = io.Copy(md5, r)
	if err != nil {
		return "", err
	}
	temp := hex.EncodeToString(md5.Sum(nil))
	return temp, nil
}
