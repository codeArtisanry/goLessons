package sshutil

import (
	"bytes"
)

func Command(command string) (string, error) {
	var host, port, username, password, keyfile string
	sshClient, err := NewSSHClient(host, port, username, password, keyfile)

	session, err := sshClient.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	buf := new(bytes.Buffer)
	session.Stdout = buf
	session.Stderr = buf

	if err := session.Run(command); err != nil {
		return "", err
	}
	session.Close()
	return buf.String(), err
}
