package common

import (
	"fmt"
	"os/exec"

	"k8s.io/klog/v2"
)

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
