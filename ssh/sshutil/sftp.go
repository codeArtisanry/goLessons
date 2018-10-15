package sshutil

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"github.com/pkg/sftp"
)

func NewSFTP(host, port, username, password, keyfile string) (*sftp.Client, error) {
	sshClient, err := NewSSHClient(host, port, username, password, keyfile)
	// create sftp client
	var sftpClient *sftp.Client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}
	return sftpClient, nil
}

// 2、上传文件

func uploadFile(sftpClient *sftp.Client, localFilePath string, remotePath string) {
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		fmt.Println("os.Open error : ", localFilePath)
		log.Fatal(err)

	}
	defer srcFile.Close()

	var remoteFileName = path.Base(localFilePath)

	dstFile, err := sftpClient.Create(path.Join(remotePath, remoteFileName))
	if err != nil {
		fmt.Println("sftpClient.Create error : ", path.Join(remotePath, remoteFileName))
		log.Fatal(err)

	}
	defer dstFile.Close()

	ff, err := ioutil.ReadAll(srcFile)
	if err != nil {
		fmt.Println("ReadAll error : ", localFilePath)
		log.Fatal(err)

	}
	dstFile.Write(ff)
	fmt.Println(localFilePath + "  copy file to remote server finished!")
}

// 3、上传文件夹

func uploadDirectory(sftpClient *sftp.Client, localPath string, remotePath string) {
	localFiles, err := ioutil.ReadDir(localPath)
	if err != nil {
		log.Fatal("read dir list fail ", err)
	}

	for _, backupDir := range localFiles {
		localFilePath := path.Join(localPath, backupDir.Name())
		remoteFilePath := path.Join(remotePath, backupDir.Name())
		if backupDir.IsDir() {
			sftpClient.Mkdir(remoteFilePath)
			uploadDirectory(sftpClient, localFilePath, remoteFilePath)
		} else {
			uploadFile(sftpClient, path.Join(localPath, backupDir.Name()), remotePath)
		}
	}

	fmt.Println(localPath + "  copy directory to remote server finished!")
}

// 4、上传测试

func DoBackup(host, port, username, password, keyfile, localPath string, remotePath string) {
	var (
		err        error
		sftpClient *sftp.Client
	)
	start := time.Now()
	sftpClient, err = NewSFTP(host, port, username, password, keyfile)
	if err != nil {
		log.Fatal(err)
	}
	defer sftpClient.Close()

	_, errStat := sftpClient.Stat(remotePath)
	if errStat != nil {
		log.Fatal(remotePath + " remote path not exists!")
	}

	backupDirs, err := ioutil.ReadDir(localPath)
	if err != nil {
		log.Fatal(localPath + " local path not exists!")
	}

	fmt.Println(backupDirs)
	uploadDirectory(sftpClient, localPath, remotePath)

	elapsed := time.Since(start)

	fmt.Println("elapsed time : ", elapsed)

}

// //TODO nice to have, a progress bar of download
// func (s *sftpClient) save(file *sftp.File, dest string, product string) error {
// 	_, fileName := path.Split(file.Name())
// 	downFile, err := os.Create(path.Join(dest, fileName))
// 	if err != nil {
// 		log.WithError(err).WithFields(log.Fields{"fs_product": product}).Errorf("Could not create file %s/%s", dest, fileName)
// 		return err
// 	}
// 	defer downFile.Close()

// 	fileStat, err := file.Stat()
// 	if err != nil {
// 		log.WithError(err).WithFields(log.Fields{"fs_product": product}).Errorf("Could not get file stats for file %s/%s", dest, fileName)
// 		return err
// 	}
// 	size := fileStat.Size()

// 	log.WithFields(log.Fields{"fs_product": product}).Infof("Downloading %s from sftp server", fileName)
// 	n, err := io.Copy(downFile, io.LimitReader(file, size))
// 	if n != size || err != nil {
// 		log.WithError(err).WithFields(log.Fields{"fs_product": product}).Errorf("Download stopped at [%d] when copying sftp file to %s/%s", n, dest, fileName)
// 		return err
// 	}

// 	return nil
// }
// func (s *sftpClient) Download(path string, dest string, product string) error {
// 	file, err := s.sftp.Open(path)
// 	if err != nil {
// 		log.WithError(err).WithFields(log.Fields{"fs_product": product}).Errorf("Could not open %s on sftp server", path)
// 		return err
// 	}
// 	defer file.Close()
// 	return s.save(file, dest, product)
// }
