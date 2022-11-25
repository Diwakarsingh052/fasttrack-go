package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type repoRequest struct {
	RepoName        string `json:"name"`
	RepoDescription string `json:"description"`
}

type RepoResponse struct {
	RepoId    int    `json:"id"`
	RepoName  string `json:"name"`
	RepoOwner `json:"owner"`
}

type RepoOwner struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
	Url   string `json:"url"`
}

type ArrayErrGithub struct {
	Resource string `json:"resource"`
	Message  string `json:"message"`
	Code     string `json:"code"`
}

type ErrorsGithub struct {
	Message string           `json:"message"`
	Err     []ArrayErrGithub `json:"errors"`
}

func main() {

	repo := repoRequest{
		RepoName:        "test404",
		RepoDescription: "test-repo",
	}
	fmt.Println(createRepo(repo))
}

func createRepo(repo repoRequest) error {
	jsonData, err := json.Marshal(repo)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, "https://api.github.com/user/repos", bytes.NewReader(jsonData))
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", "Bearer "+"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")

	//making the request to the github
	resp, err := http.DefaultClient.Do(request) // this is not going to tell whether error happened at github or not
	if err != nil {
		return err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		var githubError ErrorsGithub
		err = json.Unmarshal(data, &githubError)
		if err != nil {
			return err
		}

		fmt.Printf("%+v\n", githubError)
		return errors.New("repo creation failed")
	}

	var result RepoResponse

	err = json.Unmarshal(data, &result)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", result)
	return nil
}
