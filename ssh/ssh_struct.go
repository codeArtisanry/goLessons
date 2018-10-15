package sshutil

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	scp "github.com/bramvdbogaerde/go-scp"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

var (
	cipherList = []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"}
)

type SSHCommand struct {
	Path   string
	Env    []string
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

// type SSHClient struct {
// 	session *ssh.Session
// 	Config  *ssh.ClientConfig
// 	Host    string
// 	Port    int
// }
type SSH struct {
	Ip           string            `json:"ip,omitempty"`
	Port         int               `json:"port,omitempty"`
	Username     string            `json:"username,omitempty"`
	Password     string            `json:"password,omitempty"`
	Key          string            `json:"key,omitempty"`
	clientConfig *ssh.ClientConfig `json:"client_config,omitempty"`
	client       *ssh.Client       `json:"client,omitempty"`
	scpClinet    *scp.Client       `json:"scp_clinet,omitempty"`
	sftpClient   *sftp.Client      `json:"sftp_client,omitempty"`
	stdout       io.Writer         `json:"stdout,omitempty"`
}

func Heart() {
	ticker := time.NewTicker(time.Second * 10)
	go func() {
		for f := range ticker.C {
			tcpConn.SendRequest(fmt.Sprintf("keepalive%s@ft.com", f), true, nil)
		}
	}()

}
func NewSSH(ip, username, password, key string, port int, out io.Writer) (*SSH, error) {
	var ssh *SSH
	ssh = &SSH{
		Ip:       ip,
		Username: username,
		Password: password,
		Key:      key,
		Port:     port,
		stdout:   out,
	}

	err := ssh.Connect()
	if err != nil {
		logrus.Errorf("connect ssh %s error!", ip, err)
		return ssh, err
	}
	return ssh, err
}

func (s *SSH) Connect() error {
	var (
		auth       []ssh.AuthMethod
		addr       string
		config     ssh.Config
		err        error
		cipherList []string
	)

	auth = make([]ssh.AuthMethod, 0)

	if s.Key == "" {
		auth = append(auth, ssh.Password(s.Password))
	} else {
		signer, err := readKeyString(s.Key, s.Password)
		if err != nil {
			return err
		}
		auth = append(auth, ssh.PublicKeys(signer))
	}

	config = ssh.Config{
		Ciphers: cipherList,
	}

	s.clientConfig = &ssh.ClientConfig{
		User:            s.Username,
		Auth:            auth,
		Timeout:         30 * time.Second,
		Config:          config,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	addr = fmt.Sprintf("%s:%d", s.Ip, s.Port)

	if s.client, err = ssh.Dial("tcp", addr, s.clientConfig); err != nil {
		return err
	}

	if s.sftpClient, err = sftp.NewClient(s.client); err != nil {
		return err
	}
	return nil
}

func (s *SSH) createSession() (session *ssh.Session, err error) {
	if session, err = s.client.NewSession(); err != nil {
		return nil, err
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		return nil, err
	}
	return session, nil
}

func (s *SSH) ExecCmd(cmd string) error {
	session, err := s.createSession()
	if err != nil {
		logrus.Errorf("create %s session err", s.Ip, err)
		return err
	}
	defer session.Close()

	session.Stdout = s.stdout
	session.Stderr = s.stdout
	session.Run(cmd)
	logrus.Debugf("[%s] run [%s] out ---> %s", s.Ip, cmd)
	return nil
}

func (s *SSH) ExecMulti(cmds ...string) error {
	cmd := strings.Join(cmds, ";")
	return s.ExecCmd(cmd)
}

func readKeyString(key string, password string) (ssh.Signer, error) {
	var signer ssh.Signer
	var err error
	if password == "" {
		signer, err = ssh.ParsePrivateKey([]byte(key))
	} else {
		signer, err = ssh.ParsePrivateKeyWithPassphrase([]byte(key), []byte(password))
	}
	if err != nil {
		return nil, err
	}
	return signer, nil
}

func readKeyFile(key string, password string) (ssh.Signer, error) {
	pemBytes, err := ioutil.ReadFile(key)
	if err != nil {
		return nil, err
	}

	var signer ssh.Signer
	if password == "" {
		signer, err = ssh.ParsePrivateKey(pemBytes)
	} else {
		signer, err = ssh.ParsePrivateKeyWithPassphrase(pemBytes, []byte(password))
	}
	if err != nil {
		return nil, err
	}
	return signer, nil
}

func (s *SSH) Put(localFile string, remotePath string) error {
	src, err := os.Open(localFile)
	if err != nil {
		fmt.Println("os.Open error : ", localFile)
	}
	defer src.Close()

	var remoteFileName = path.Base(localFile)
	dst, err := s.sftpClient.Create(path.Join(remotePath, remoteFileName))
	if err != nil {
		fmt.Println("sftpClient.Create error : ", path.Join(remotePath, remoteFileName))
	}
	defer dst.Close()

	c, err := ioutil.ReadAll(src)
	if err != nil {
		fmt.Println("ReadAll error : ", localFile)
	}
	dst.Write(c)

	fmt.Println(localFile + "  copy file to remote server finished!")
	return nil
}

func (s *SSH) Close() {
	defer s.sftpClient.Close()
	defer s.scpClient.Close()
	defer s.client.Close()
}
