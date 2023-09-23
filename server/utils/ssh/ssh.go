package ssh

import (
	"io/ioutil"
	"log"
	"strconv"

	"golang.org/x/crypto/ssh"
)

type SSHConfig struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	User       string `json:"user"`
	Password   string `json:"password"`
	PrivateKey string `json:"privateKey"`
}

/** 创建客户端后 需要defer client.close()*/
/** 优先私钥访问模式*/
func NewClient(config SSHConfig) (*ssh.Client, error) {
	var sshConfig *ssh.ClientConfig
	// ssh 私钥访问方式
	if config.PrivateKey != "" {
		// 读取私钥文件
		key, err := ioutil.ReadFile(config.PrivateKey)
		if err != nil {
			return nil, err
		}
		// 解析私钥
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			return nil, err
		}
		// 解析私钥
		sshConfig = &ssh.ClientConfig{
			User: config.User,
			Auth: []ssh.AuthMethod{
				// 使用私钥访问
				ssh.PublicKeys(signer),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
	}
	// user password 访问方式
	if config.Password != "" {
		sshConfig = &ssh.ClientConfig{
			User: config.User,
			Auth: []ssh.AuthMethod{
				// 使用password访问
				ssh.Password(config.Password),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
	}
	// 连接ssh服务器
	sshClient, err := ssh.Dial("tcp", config.Host+":"+strconv.Itoa(config.Port), sshConfig)
	if err != nil {
		log.Fatal("无法连接到SSH服务器", err)
		return nil, err
	}
	return sshClient, nil
}

/**远程服务器执行命令*/
func ExecuteRemoteCommand(client *ssh.Client, command string) (string, error) {
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("创建ssh会话失败", err)
	}
	defer session.Close()
	// 执行远程命令
	output, err := session.CombinedOutput(command)
	if err != nil {
		log.Fatalf("执行SSH命令失败: %v", err)
	}
	return string(output), nil
}
