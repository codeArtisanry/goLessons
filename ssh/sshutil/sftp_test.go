package sshutil

import (
	"testing"
)

func Test_Dobackup(t *testing.T) {

	DoBackup("192.168.5.150", "22", "root", "HR2018!!", "", "/home/ubuntu/backup", "/docker/update")

}
