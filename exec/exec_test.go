package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestExec(t *testing.T) {
	cmd := exec.Command("/bin/sh", "-c", `ss|grep "ESTAB"|grep -v "grep"|awk '{print $1}'`)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func TestExecPipe(t *testing.T) {
	cmd := exec.Command("/bin/sh", "-c", `ss|grep "ESTAB"|grep -v "grep"|awk '{print $1}'`)
	cmd2 := exec.Command("wc", "-l")
	cmd.Env = os.Environ()
	// cmd.Stdout = os.Stdout
	cmd2.Stdin, _ = cmd.StdoutPipe()
	cmd2.Stdout = os.Stdout
	cmd2.Start()
	cmd.Run()
	cmd2.Wait()
}
