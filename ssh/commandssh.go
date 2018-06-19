package main

import (
	"log"
	"os/exec"
)

func commandSSH(host, mdst string) ([]byte, error) {

	log.Println("[SSH：] Runing")

	args := []string{
		"-A",
		"-o", "StrictHostKeyChecking=no",
		"-o", "UserKnownHostsFile=/dev/null",
		"-o", "LogLevel=quiet",
	}
	args = append(args, "-p", "3009")

	args = append(args, "-l", "root")

	args = append(args, host)
	args = append(args, mdst)

	cmd := exec.Command("ssh", args...)
	log.Println(args)
	b, err := cmd.Output()
	if err != nil {
		log.Println(err)
	}
	log.Println("[SSH:] output-> ", string(b))
	return b, err
}
