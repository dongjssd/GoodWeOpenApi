/**
    @author: dongjs
    @date: 2025/2/10
    @description:
**/

package GoodWeOpenApi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetToken(request GetTokenRequest) (*GetTokenResponse, error) {
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doPostRequest("GetToken", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetTokenResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

type GetTokenRequest struct {
	Account string `json:"account"`
	Pwd     string `json:"pwd"`
}

type GetTokenResponse struct {
	Data struct {
		Token   string `json:"token"`
		Expired int `json:"expired"`
	} `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
