package sshutil

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	scp "github.com/bramvdbogaerde/go-scp"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

const (
	sshConnectRetryInterval = 5 * time.Second
	sshConnectMaxWait       = 1 * time.Minute
)

type SSHClient struct {
	client     *ssh.Client  `json:"client,omitempty"`
	scpClinet  *scp.Client  `json:"scp_clinet,omitempty"`
	sftpClient *sftp.Client `json:"sftp_client,omitempty"`
}

func (s *SSHClient) Command(command string, out io.Writer, errW io.Writer) error {
	if s.client == nil {
		return ClientError{msg: "session not created"}
	}
	session, err := s.NewSession()
	if err == nil {
		session.Stdout = out
		session.Stderr = errW
		err = session.Run(command)
		if err == nil {
			session.Close()
		}
	}
	return err
}
func (s *SSHClient) Command(command string) (string, error) {
	session, err := s.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	combinedOutput := new(bytes.Buffer)
	session.Stdout = combinedOutput
	session.Stderr = combinedOutput

	if err := session.Start(command); err != nil {
		return "", err
	}

	// Wait for the SCP connection to close, meaning it has consumed all
	// our data and has completed. Or has errored.
	err = session.Wait()
	if err != nil {
		if exitErr, ok := err.(*ssh.ExitError); ok {
			// Otherwise, we have an ExitErorr, meaning we can just read
			// the exit status
			log.Printf("non-zero exit status: %d", exitErr.ExitStatus())

			// If we exited with status 127, it means SCP isn't available.
			// Return a more descriptive error for that.
			if exitErr.ExitStatus() == 127 {
				return "", errors.New(
					"SCP failed to start. This usually means that SCP is not\n" +
						"properly installed on the remote system.")
			}
		}

		return combinedOutput.String(), err
	}

	return combinedOutput.String(), nil
}

func GetSSH() (*SSHClient, error) {
	client, err := ssh.Dial("tcp", ip, config)
	if err != nil {
		return nil, err
	}
	return &SSHClient{Client: client}, nil

}

// ConnectSSH tries to connect to the given IP and returns a new client.
func ConnectSSH(ip string, config *ssh.ClientConfig) (*SSHClient, error) {
	// dialFunc := func() (*SSHClient, error) {
	// 	client, err := ssh.Dial("tcp", ip, config)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return &SSHClient{Client: client}, nil
	// }

	// run it before we pass it to the ticker for re-dials
	client, err := func() (*SSHClient, error) {
		client, err := ssh.Dial("tcp", ip, config)
		if err != nil {
			return nil, err
		}
		return &SSHClient{Client: client}, nil
	}()
	if err == nil {
		return client, nil
	}

	timeout := time.After(sshConnectMaxWait)

	var dialError error
	for {
		select {
		case <-time.Tick(sshConnectRetryInterval):
			client, err := dialFunc()
			if err != nil {
				dialError = err
				continue
			}
			return client, nil
		case <-timeout:
			if dialError != nil {
				return nil, fmt.Errorf("cannot connect with ssh: %s", dialError)
			}

			return nil, errors.New("cannot connect with ssh. timeout")
		}
	}
}

// SshConfig returns a new clientConfig based on the given privateKey
func SshConfig(username, privateKey string) (*ssh.ClientConfig, error) {
	signer, err := ssh.ParsePrivateKey([]byte(privateKey))
	if err != nil {
		return nil, fmt.Errorf("Error setting up SSH config: %s", err)
	}

	// fallback to root if the username is empty
	if username == "" {
		username = "root"
	}

	return &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
	}, nil
}
