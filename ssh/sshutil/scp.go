package sshutil

import (
	"fmt"
	"os"

	scp "github.com/bramvdbogaerde/go-scp"
)

func Copy(hostport, src, dst string) (err error) {
	var username = "root"
	var password = ""
	var keyfile = ""
	sshClient, err := NewClientConfig(username, password, keyfile)
	// Create a new SCP client
	client := scp.NewClient(hostport, sshClient)

	// Connect to the remote server
	err = client.Connect()
	if err != nil {
		fmt.Println("Couldn't establisch a connection to the remote server ", err)
		return err
	}

	// Open a file
	f, _ := os.Open(src)

	// Close session after the file has been copied
	defer client.Session.Close()

	// Close the file after it has been copied
	defer f.Close()

	// Finaly, copy the file over
	// Usage: CopyFile(fileReader, remotePath, permission)

	client.CopyFile(f, dst, "0644")
	return nil
}
