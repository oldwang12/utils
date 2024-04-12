package common

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

// 记得 defer client.Close()
func ClientSSH(serverAddr, sshPort, sshUser, privateKeyPath string) (*ssh.Client, error) {
	// 读取私钥文件
	key, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("读取私钥文件失败:%v", err)
	}

	// 使用私钥创建认证对象
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, fmt.Errorf("使用私钥创建认证对象:%v", err)
	}

	// 配置SSH客户端
	config := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 忽略主机密钥检查（不建议在生产环境中使用）
	}

	// 建立SSH连接
	return ssh.Dial("tcp", serverAddr+":"+sshPort, config)
}

func ExecSSHCommand(client *ssh.Client, command string) (string, error) {
	// 执行远程命令
	session, err := client.NewSession()
	if err != nil {
		return "", fmt.Errorf("无法创建SSH会话: %v", err)
	}
	defer session.Close()
	// 运行命令
	output, err := session.CombinedOutput(command)
	return string(output), err
}
