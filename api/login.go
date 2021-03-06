package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (cli *Client) LoginWithGithub(endpoint string, githubToken string) (*User, error) {
	url := endpoint + pathPrefix + "/login"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-GitHub-Token", githubToken)
	b, err := cli.rawRequest(req)
	if err != nil {
		return nil, err
	}

	var userResp UserResponse
	err = json.Unmarshal(b, &userResp)
	if err != nil {
		return nil, err
	}
	return userResp.User, nil
}

func (cli *Client) LoginWithVault(vault_url string, githubToken string) (*User, error) {
	req, err := http.NewRequest("POST",
		vault_url+"/v1/auth/github/login",
		strings.NewReader("{\"token\":\""+githubToken+"\"}"))
	if err != nil {
		return nil, err
	}
	rawResponse, err := cli.rawRequest(req)
	if err != nil {
		return nil, err
	}

	var resp VaultAuthResponse
	json.Unmarshal(rawResponse, &resp)

	user := &User{Token: resp.Auth.ClientToken}

	return user, nil
}
