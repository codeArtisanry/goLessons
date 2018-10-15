package sshutil

// func (s *SSH) ExecCmd(cmd string) error {
// 	session, err := s.createSession()
// 	if err != nil {
// 		logrus.Errorf("create %s session err", s.Ip, err)
// 		return err
// 	}
// 	defer session.Close()

// 	session.Stdout = s.stdout
// 	session.Stderr = s.stdout
// 	session.Run(cmd)
// 	logrus.Debugf("[%s] run [%s] out ---> %s", s.Ip, cmd)
// 	return nil
// }

// func (s *SSH) ExecMulti(cmds ...string) error {
// 	cmd := strings.Join(cmds, ";")
// 	return s.ExecCmd(cmd)
// }
