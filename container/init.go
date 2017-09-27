package container

import (
	"os"
	"syscall"
	"github.com/Sirupsen/logrus"
)

func RunContainerInitProcess(command string, args []string) error {
	logrus.Info("command %s", command)

	//defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	//syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	argv := []string{command}
	logrus.Infof("15, argv = %s", argv)
	if err := syscall.Exec(command, argv, os.Environ()); err != nil {
		logrus.Error(err.Error())
	}
	return nil
}
