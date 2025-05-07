/**
    @author: dongjs
    @date: 2025/2/10
    @description:
**/

package GoodWeOpenApi

import (
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Username string
	Password string
	Token string
}

func InitClient(username string, password string) *Client {
	return &Client{
		Username: username,
		Password: password,
	}
}

// 发起请求
func (c *Client) doPostRequest(url string, buf io.Reader) ([]byte, error) {
	req, _ := http.NewRequest("POST", ApiDomain+url, buf)
	req.Header.Set("Content-Type", "application/json")
	if url != "GetToken"{
		req.Header.Set("Token", c.Token)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body, nil
}

func (c *Client) doGetRequest(url string, buf io.Reader) ([]byte, error) {
	req, _ := http.NewRequest("GET", ApiDomain+url, buf)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body, nil
}
