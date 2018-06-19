package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/bramvdbogaerde/go-scp"
	"golang.org/x/crypto/ssh"
)

//go:generate goversioninfo -icon=icon.ico

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

var nginx_host = "15.14.12.150:22"
var tomcat_host = "15.14.12.151:22"

// var nginx_host = "118.190.117.250:3009"
// var tomcat_host = "118.190.117.250:3009"
func deployFrontend() {
	// 1. check frontend
	if Exists("lotus.tar.gz") {
		DialSSH(nginx_host, `mkdir -p /docker/update/ /docker/bianban/ /docker/rollback/`)
		SCP(nginx_host, "lotus.tar.gz", "/docker/update/lotus.tar.gz")
		DialSSH(nginx_host, `rm -fr /docker/rollback/lotus; mv /docker/bianban/lotus /docker/rollback/;tar -zxf /docker/update/lotus.tar.gz -C /docker/bianban/`)
		os.Mkdir("del", 0755)
		os.Rename("lotus.tar.gz", "del/lotus.tar.gz")
	}
}

func deployTomcat() {
	// 1. check frontend
	if Exists("ROOT.war") {
		DialSSH(nginx_host, `mkdir -p /docker/bianban/lzkpv4/ /docker/rollback/`)
		SCP(nginx_host, "ROOT.war", "/docker/bianban/lzkpv4/ROOT.war")
		DialSSH(nginx_host, `sh -c ~/lzkpv4/deploy.sh tomcat;rm -f /docker/rollback/ROOT.war; mv /docker/bianban/lzkpv4/ROOT.war /docker/rollback/ROOT.war `)
		DialSSH(tomcat_host, `systemctl restart tomcat8`)
		os.Mkdir("del", 0755)
		os.Rename("ROOT.war", "del/ROOT.war")
	}
}
func deployBackend() {
	// 1. check frontend
	if Exists("backendv4.jar") {
		DialSSH(nginx_host, `mkdir -p /docker/bianban/backendv4/ /docker/rollback/`)
		SCP(nginx_host, "backendv4.jar", "/docker/bianban/backendv4/backendv4.jar; rm -f /docker/rollback/backendv4.jar; mv /docker/bianban/backendv4/backendv4.jar /docker/rollback/backendv4.jar")
		DialSSH(nginx_host, `sh -c ~/lzkpv4/deploy.sh backend`)
		os.Mkdir("del", 0755)
		os.Rename("backendv4.jar", "del/backendv4.jar")
	}
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
func main() {
	deployFrontend()
	deployTomcat()
	deployBackend()
}
