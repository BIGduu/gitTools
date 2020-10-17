package urlinfo

import (
	"fmt"
	"net/url"
	"strings"
)

type RepositoryInfo struct {
	urlString      string
	info           *url.URL
	repositoryName string
}

func (receiver RepositoryInfo) GetInfo() *url.URL {
	return receiver.info
}

func (receiver RepositoryInfo) GetURLString() string {
	return receiver.info.String()
}

func (receiver RepositoryInfo) GetRepositoryName() string {
	return receiver.repositoryName
}

func (receiver RepositoryInfo) SetUserInfo(token string) {
	if token == "" {
		return
	}
	receiver.info.User = url.UserPassword("oauth2", token)
}

func ParseURL(urlString string) *RepositoryInfo {
	parse, err := url.Parse(urlString)
	if err != nil {
		fmt.Println("The url address is incorrect, please copy the https address")
		panic(err)
	}
	path := parse.Path
	index := strings.LastIndex(path, "/")
	gitIndex := strings.LastIndex(path, ".git")
	if gitIndex < 0 {
		fmt.Println("The url address is incorrect, there is no .git ending")
	}
	repositoryName := path[index+1 : gitIndex]
	return &RepositoryInfo{urlString: urlString, info: parse, repositoryName: repositoryName}
}

func ParseURLWithToken(urlString, token string) *RepositoryInfo {
	parse, err := url.Parse(urlString)
	if err != nil {
		fmt.Println("The url address is incorrect, please copy the https address")
		panic(err)
	}
	path := parse.Path
	index := strings.LastIndex(path, "/")
	gitIndex := strings.LastIndex(path, ".git")
	if gitIndex < 0 {
		fmt.Println("The url address is incorrect, there is no .git ending")
	}
	repositoryName := path[index+1 : gitIndex]
	repository := &RepositoryInfo{urlString: urlString, info: parse, repositoryName: repositoryName}
	repository.SetUserInfo(token)
	return repository
}
