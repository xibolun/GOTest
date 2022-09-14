package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
)

func main() {
	stdout, stderr, err := remoteRun("root", "10.200.20.219", "NiuLink2022++", "22", "uname -a")
	if err != nil {
		panic(err)
	}

	fmt.Println(stdout)
	fmt.Println(stderr)
}

func remoteRun(user, addr, pwd, port, cmd string) (stdoutstr, stderrstr string, err error) {
	// Authentication
	config := &ssh.ClientConfig{
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(pwd),
		},
	}
	// Connect
	client, err := ssh.Dial("tcp", net.JoinHostPort(addr, port), config)
	if err != nil {
		return
	}

	session, err := client.NewSession()
	if err != nil {
		return
	}
	defer session.Close()
	var stdout, stderr bytes.Buffer

	session.Stdout = &stdout
	session.Stderr = &stderr

	err = session.Run(cmd)

	return stdout.String(), stderr.String(), err
}
