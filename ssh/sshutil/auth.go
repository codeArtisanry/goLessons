package sshutil

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"golang.org/x/crypto/ssh/terminal"
)

// An implementation of ssh.KeyboardInteractiveChallenge that simply sends
// back the password for all questions. The questions are logged.
func passwordKeyboardInteractive(password string) ssh.KeyboardInteractiveChallenge {
	return func(user, instruction string, questions []string, echos []bool) ([]string, error) {
		// log.Printf("Keyboard interactive challenge: ")
		// log.Printf("-- User: %s", user)
		// log.Printf("-- Instructions: %s", instruction)
		// for i, question := range questions {
		// 	log.Printf("-- Question %d: %s", i+1, question)
		// }

		// Just send the password back for all questions
		answers := make([]string, len(questions))
		for i := range answers {
			answers[i] = password
		}

		return answers, nil
	}
}

// Generate a password-auth'd ssh ClientConfig
func WithKeyboardPasswordAuth(password string) ssh.AuthMethod {
	return ssh.KeyboardInteractive(passwordKeyboardInteractive(password))
}

// Generate a password-auth'd ssh ClientConfig
func WithPasswordAuth(password string) ssh.AuthMethod {
	return ssh.Password(password)
}

// WithAgentAuth use already authed user
func WithAgentAuth() ssh.AuthMethod {
	sock := os.Getenv("SSH_AUTH_SOCK")
	if sock != "" {
		fmt.Println(errors.New("Agent Disabled"))
		return nil
	}
	socks, err := net.Dial("unix", sock)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	signer := agent.NewClient(socks).Signers
	return ssh.PublicKeysCallback(signer)

	// 简写方式
	// if sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
	// 	return ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers)
	// }
	// return nil
}

func WithKeyString(key string, password string) ssh.AuthMethod{
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
}
// PublicKeyAuth 自动监测是否带有密码
func WithPublicKeyAuth(keyfile string, password string) ssh.AuthMethod {
	pemBytes, err := ioutil.ReadFile(keyfile)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var signer ssh.Signer
	signer, err = ssh.ParsePrivateKey(pemBytes)
	if password == "" && err != nil {
		fmt.Println(err)
		return nil
	}

	if strings.Contains(err.Error(), "cannot decode encrypted private keys") {
		if signer, err = ssh.ParsePrivateKeyWithPassphrase(pemBytes, []byte(password)); err == nil {
			return ssh.PublicKeys(signer)
		}
	}
	return nil
}

// PublicKeyTerminalAuth 通过终端读取带密码的 PublicKey
func WithPublicKeyTerminalAuth(keyfile string) ssh.AuthMethod {
	// fmt.Fprintf(os.Stderr, "This SSH key is encrypted. Please enter passphrase for key '%s':", priv.path)
	passphrase, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Fprintln(os.Stderr)
	pemBytes, err := ioutil.ReadFile(keyfile)
	if err != nil {

		fmt.Println(err)
		return nil
	}
	signer, err := ssh.ParsePrivateKeyWithPassphrase(pemBytes, passphrase)
	if err != nil {

		fmt.Println(err)
		return nil
	}

	return ssh.PublicKeys(signer)
}
