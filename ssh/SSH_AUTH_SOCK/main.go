package main

import (
	"net"
	"os"
	"strconv"

	"github.com/Sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

func main() {
	WithAuthSock("root", "zybb", 443)
}

// 你的私钥是不是有密码，试着在Linux 客户端执行
// ssh-agent
// ssh-add ~/.ssh/id_dsa (你的私钥)
func WithAuthSock(username, hostport string, port int) (*ssh.Client, error) {
	sock, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		logrus.Infof("error login,details: %s", err.Error())
		return nil, err
	}

	agent := agent.NewClient(sock)

	signers, err := agent.Signers()
	if err != nil {
		logrus.Infof("error login,details: %s", err.Error())
		return nil, err
	}

	auths := []ssh.AuthMethod{ssh.PublicKeys(signers...)}

	ClientConfig := &ssh.ClientConfig{
		User: username,
		Auth: auths,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	ClientConfig.SetDefaults()

	logrus.Infof("tcp dial to %s", hostport+":"+strconv.Itoa(port))
	client, err := ssh.Dial("tcp", hostport+":"+strconv.Itoa(port), ClientConfig)
	if err != nil {
		logrus.Infof("error login,details: %s", err.Error())
		return nil, err
	}
	return client, nil
}
