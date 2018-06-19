package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/bramvdbogaerde/go-scp"
	"golang.org/x/crypto/ssh"
)

func Genconfig() *ssh.ClientConfig {
	var User string = "root"
	// var connectKey string = "/home/ubuntu/.ssh/id_rsa"
	var connectKey string
	var permBytes []byte
	var connectPass string = "HR2018!!"

	config := &ssh.ClientConfig{}
	if connectKey != "" {
		// Read PublicKey

		buffer, err := ioutil.ReadFile(connectKey)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:%s%n", err)
			return nil
		}
		if len(permBytes) != 0 {
			buffer = permBytes
		}
		key, err := ssh.ParsePrivateKey(buffer)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:%s%n", err)
			return nil
		}

		// Create ssh client config for KeyAuth
		config = &ssh.ClientConfig{
			User: User,
			Auth: []ssh.AuthMethod{
				ssh.PublicKeys(key)},
			Timeout:         60 * time.Second,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
	} else {
		// Create ssh client config for PasswordAuth
		config = &ssh.ClientConfig{
			User: User,
			Auth: []ssh.AuthMethod{
				ssh.Password(connectPass)},
			Timeout:         60 * time.Second,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
	}
	return config
}

func DialSSH(connectHostPort, cmdStr string) int {

	// New Connext create
	conn, err := ssh.Dial("tcp", connectHostPort, Genconfig())
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot connect %v: %v \n", connectHostPort, err)
		return 1
	}

	// New Session
	session, err := conn.NewSession()
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot open new session: %v \n", err)
		return 1
	}

	// go func() {
	// 	time.Sleep(2419200 * time.Second)
	// 	conn.Close()
	// }()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	err = session.Run(cmdStr)
	session.Close()
	conn.Close()
	return 0
}

func SCP(connectHostPort, src, dst string) {

	// Create a new SCP client
	client := scp.NewClient(connectHostPort, Genconfig())

	// Connect to the remote server
	err := client.Connect()
	if err != nil {
		fmt.Println("Couldn't establisch a connection to the remote server ", err)
		return
	}

	// Open a file
	f, _ := os.Open(src)

	// Close session after the file has been copied
	defer client.Session.Close()

	// Close the file after it has been copied
	defer f.Close()

	// Finaly, copy the file over
	// Usage: CopyFile(fileReader, remotePath, permission)

	client.CopyFile(f, dst, "0655")
}

func createDeploy() {

	var data string = `#! /bin/sh
echo "wo shi shui "`

	ioutil.WriteFile("deploy.sh", []byte(data), 0755)
}
func main() {
	createDeploy()
	SCP("118.190.117.250:3009", "deploy.sh", "deploy.sh")
	DialSSH("118.190.117.250:3009", `sh -c ~/deploy.sh`)
	// commandSSH("d2", "echo 1 ;hostname；sleep 3")
}
