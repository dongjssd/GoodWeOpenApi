/**
    @author: dongjs
    @date: 2025/7/14
    @description: 设备相关接口
**/

package GoodWeOpenApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

// QueryInventers 电站设备列表 QueryInventers
func (c *Client) QueryInventers(request QueryInventersRequest) (*QueryInventersResponse, error) {
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("QueryInventers:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doPostRequest("QueryInventers", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := QueryInventersResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetInventersMoreDetailInfoBySn 设备实时数据 GetInventersMoreDetailInfoBySn
func (c *Client) GetInventersMoreDetailInfoBySn(request GetInventersMoreDetailInfoBySnRequest) (*GetInventersMoreDetailInfoBySnResponse, error) {
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("GetInventersMoreDetailInfoBySn:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doPostRequest("GetInventersMoreDetailInfoBySn", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetInventersMoreDetailInfoBySnResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetInverterPacByDay 设备功率曲线 GetInverterPacByDay
func (c *Client) GetInverterPacByDay(request GetInverterPacByDayRequest) (*GetInverterPacByDayResponse, error) {
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("GetInverterPacByDay:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doPostRequest("GetInverterPacByDay", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetInverterPacByDayResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetInverterDataByColumn 设备今日发电量/累计发电量/温度/交流电压/曲线 GetInverterDataByColumn
func (c *Client) GetInverterDataByColumn(request GetInverterDataByColumnRequest) (*GetInverterDataByColumnResponse, error) {
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("GetInverterDataByColumn:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doPostRequest("GetInverterDataByColumn", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetInverterDataByColumnResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetInverterPower 设备日,月,年发电量 GetInverterPower
func (c *Client) GetInverterPower(request GetInverterPowerRequest) (*GetInverterPowerResponse, error) {
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("GetInverterPower:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doPostRequest("GetInverterPower", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetInverterPowerResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 获取设备历史发电数据 GetInverterMoreData
// 根据设备序列号和给定日期范围获取设备指定时间内的发电数据
func (c *Client) GetInverterMoreData(request GetInverterMoreDataRequest) (*GetInverterMoreDataResponse, error) {
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("GetInverterMoreData:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doPostRequest("GetInverterMoreData", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetInverterMoreDataResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 批量获取逆变器实时发电数据 GetInventersDatas Get
func (c *Client) GetInventersDatas(request GetInventersDatasRequest) (*GetInventersDatasResponse, error) {
	//requestBytes, err := json.Marshal(&request)
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Println("GetInventersDatas:", string(requestBytes))
	//buf := bytes.NewBuffer(requestBytes)
	params := url.Values{}
	for _, sn := range request.Sns {
		params.Add("sns", sn)
	}
	body, err := c.doGetRequest("GetInventersDatas", params)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetInventersDatasResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 批量获取电站逆变器实时发电数据 GetInvertDatasByPwIds
func (c *Client) GetInvertDatasByPwIds(request GetInvertDatasByPwIdsRequest) (*GetInvertDatasByPwIdsResponse, error) {
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("GetInvertDatasByPwIds:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doPostRequest("GetInvertDatasByPwIds", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetInvertDatasByPwIdsResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
