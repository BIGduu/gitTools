package main

import (
	"fmt"
	"gitTools/banner"
	"gitTools/gittools"
	"gitTools/urlinfo"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var sum int
	sign := make(chan int)
	hashmap := make(map[string]*urlinfo.RepositoryInfo)
	banner.GetBanner().PrintBanner()
	repositoryString, tokenString := readFileIfNotFileCreate()
	nextLineSign := getNextLineSign(repositoryString)
	split := strings.Split(repositoryString, nextLineSign)
	for singleRepositoryIndex := range split {
		url := urlinfo.ParseURLWithToken(split[singleRepositoryIndex], tokenString)
		fmt.Println(url.GetURLString())
		fmt.Println(url.GetRepositoryName())
		hashmap[url.GetRepositoryName()] = url
	}

	sum = len(hashmap)
	for _, v := range hashmap {
		if Exists(v.GetRepositoryName()) {
			fmt.Println("skip")
			go func() {
				sign <- 1
			}()
			continue
		}
		tools := gittools.NewGitTools(*v, gittools.OperationClone)
		go tools.Execute(sign)

	}

	count := 0
	for i := range sign {
		count = count + i
		if sum == count {
			break
		}
	}
	fmt.Println("end")
}

func getNextLineSign(repositoryString string) string {
	index := strings.Index(repositoryString, "\r\n")
	var nextLineSign string
	if index == -1 {
		i := strings.Index(repositoryString, "\n")
		if i == -1 {
			nextLineSign = "\r"
		} else {
			nextLineSign = "\n"
		}
	} else {
		nextLineSign = "\r\n"
	}
	return nextLineSign
}

func readFileIfNotFileCreate() (string, string) {
	repositoryExists := Exists("repository.txt")
	if !repositoryExists {
		err := ioutil.WriteFile("repository.txt", []byte(""), os.ModeAppend)
		if err != nil {
			fmt.Println("Failed to create repository.txt, you can try to manually create the file under this directory")
		}
	}
	tokenExists := Exists("token.txt")
	if !tokenExists {
		err := ioutil.WriteFile("token.txt", []byte(""), os.ModeAppend)
		if err != nil {
			fmt.Println("Failed to create token.txt, you can try to create the file manually under this directory")
		}
	}
	if !repositoryExists || !tokenExists {
		var scanln string
		for {
			fmt.Println("Please enter Y after the creation and fill in the information")
			fmt.Scanln(&scanln)
			if scanln == "Y" || scanln == "y" {
				break
			} else {
				fmt.Println("The Input Information is Illegal")
			}
		}
	}
	repositoryAll, _ := ioutil.ReadFile("repository.txt")
	token, _ := ioutil.ReadFile("token.txt")
	repositoryAllString := string(repositoryAll)
	tokenString := string(token)
	if tokenString == "" {
		fmt.Println("The token has not read the information, if it is a public warehouse, no token is needed")
	}
	if repositoryAllString == "" {
		fmt.Println("The repository did not read the information")
	}
	return repositoryAllString, tokenString
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}
