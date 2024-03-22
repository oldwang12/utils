package common

import (
	"fmt"
	"os/exec"

	"k8s.io/klog/v2"
)

const (
	color_red     = uint8(iota + 91)
	color_green   //	绿
	color_yellow  //	黄
	color_blue    // 	蓝
	color_magenta //	洋红
	clour_blue2   // 	浅蓝
)

func Red(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_red, s)
}

func Green(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_green, s)
}

func Yellow(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_yellow, s)
}

func Blue(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_blue, s)
}

// 洋红
func Magenta(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_magenta, s)
}

func LocalExec(command string) (string, error) {
	c := exec.Command("/bin/sh", "-c", command)
	result, err := c.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%v, err: %v", Red("[本地执行]: "+command), c.Stderr)
	}
	klog.Info(Green("[本地执行]: " + command))

	if string(result) != "" {
		klog.Info(Yellow(string(result)))
	}
	return string(result), nil
}

func LocalExecNoLog(command string) (string, error) {
	c := exec.Command("/bin/sh", "-c", command)
	result, err := c.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%v, err: %v", Red("[本地执行]: "+command), c.Stderr)
	}
	return string(result), nil
}
