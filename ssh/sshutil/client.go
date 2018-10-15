package sshutil

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/Sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

// TemporaryKey creates a new temporary public and private key
func TemporaryKey() (string, string, error) {
	priv, err := rsa.GenerateKey(rand.Reader, 2014)
	if err != nil {
		return "", "", err
	}

	// ASN.1 DER encoded form
	priv_der := x509.MarshalPKCS1PrivateKey(priv)
	priv_blk := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   priv_der,
	}

	privateKey := string(pem.EncodeToMemory(&priv_blk))

	// Marshal the public key into SSH compatible format
	// TODO properly handle the public key error
	pub, _ := ssh.NewPublicKey(&priv.PublicKey)
	pub_sshformat := string(ssh.MarshalAuthorizedKey(pub))

	return privateKey, pub_sshformat, nil
}

func getHostPublicKey(host string) (ssh.PublicKey, error) {
	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				return nil, errors.New(fmt.Sprintf("error parsing %q: %v", fields[2], err))
			}
			break
		}
	}

	if hostKey == nil {
		return nil, errors.New(fmt.Sprintf("no hostkey for %s", host))
	}
	return hostKey, nil
	// HostKeyCallback: ssh.FixedHostKey(hostKey),
}

// Generate a password-auth'd ssh ClientConfig
func WithPassword(username, password string) (*ssh.ClientConfig, error) {
	// password := state.Get("Password").(string)
	return &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
			ssh.KeyboardInteractive(passwordKeyboardInteractive(password)),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}, nil
}

// WithPrivateKey Generate a PKI- Public Key Infrastructure ssh ClientConfig
func WithPrivateKey(username, passphrase, keyfile string) (*ssh.ClientConfig, error) {
	pemBytes, err := ioutil.ReadFile(keyfile)
	if err != nil {
		return nil, fmt.Errorf("failed reading key: %s", err.Error())
	}

	var signer ssh.Signer
	signer, err = ssh.ParsePrivateKey(pemBytes)
	if passphrase == "" && err != nil {
		return nil, err
	}

	if strings.Contains(err.Error(), "cannot decode encrypted private keys") {
		signer, err = ssh.ParsePrivateKeyWithPassphrase(pemBytes, []byte(passphrase))
	}
	if err != nil {
		return nil, err
	}

	var hostKey ssh.PublicKey
	return &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}, nil
}

// WithAthSock 你的私钥是不是有密码，试着在Linux 客户端执行
// ssh-agent
// ssh-add ~/.ssh/id_dsa (你的私钥)
func WithAthSock(username string) (*ssh.ClientConfig, error) {
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
	// ClientConfig.SetDefaults()

	return ClientConfig, nil
}

func NewClientConfig(username, password, keyfile string) (*ssh.ClientConfig, error) {
	// var authMethod ssh.AuthMethod
	var ClientConfig *ssh.ClientConfig
	var err error
	sock := os.Getenv("SSH_AUTH_SOCK")
	if sock != "" {
		log.Println("Creating ssh client with ssh agent")
		ClientConfig, err = WithAthSock(username)
	} else if keyfile != "" {
		ClientConfig, err = WithPrivateKey(username, password, keyfile)
	} else {
		ClientConfig, err = WithPassword(username, password)
		// return nil, fmt.Errorf("No ssh connection authentication provided")
		if err != nil {
			return nil, err
		}

	}

	ClientConfig.SetDefaults()
	return ClientConfig, nil
}

const defaultSSHPort = "22"

func NewSSHClient(host, port, username, password, keyfile string) (*ssh.Client, error) {
	ClientConfig, err := NewClientConfig(username, password, keyfile)
	if port == "" {
		port = defaultSSHPort
	}
	return ssh.Dial("tcp", host+":"+port, ClientConfig)

}

type SSH struct {
	host    string
	user    string
	passwd  string
	keyfile string
}

func NewClient(config *SSHConnConfig) (*ssh.Client, error) {
	sshConfig := &ssh.ClientConfig{
		User: config.User,
		Auth: []ssh.AuthMethod{SSHAgent()},
	}

	if config.Password != "" {
		sshConfig.Auth = append(sshConfig.Auth, ssh.Password(config.Password))
	}

	return ssh.Dial("tcp", net.JoinHostPort(config.Host, config.Port), sshConfig)
}

func SSHAgent() ssh.AuthMethod {
	if sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		return ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers)
	}
	return nil
}
