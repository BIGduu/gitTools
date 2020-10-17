package gittools

import (
	"fmt"
	"gitTools/urlinfo"
	"os"
	"os/exec"
)

type GitTools struct {
	repositoryUrl string
	operation     OperationType
}

func NewGitTools(info urlinfo.RepositoryInfo, operationType OperationType) *GitTools {
	getInfo := info.GetInfo()
	URL := getInfo.String()
	return &GitTools{URL, operationType}
}

func (receiver GitTools) Execute(sign chan int) {
	command := exec.Command("git", receiver.operation.String(), "--progress", receiver.repositoryUrl)
	//显示运行的命令
	fmt.Println(command.Args)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Run()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
	sign <- 1
}

type OperationType int32

const (
	OperationClone  OperationType = 0
	OperationUpdate OperationType = 1
)

func (receiver OperationType) String() string {
	switch receiver {
	case 0:
		return "clone"
	case 1:
		return "pull"
	}
	panic("wrong operation")
}
