/**
    @author: dongjs
    @date: 2025/2/11
    @description:
**/

package GoodWeOpenApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

// 获取用户电站列表
// 获取当前用户当前组织和下级组织电站列表
func (c *Client) GetUserPowerStation(request GetUserPowerStationRequest) (*GetUserPowerStationResponse, error) {
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("GetUserPowerStation:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doPostRequest("GetUserPowerStation", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetUserPowerStationResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 获取电站信息
// 获取当前用户当前组织或下级电站的信息，包含图片和业主信息
func (c *Client) GetPowerStationByID(request GetPowerStationByIDRequest) (*GetPowerStationByIDResponse, error) {
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("GetPowerStationByID:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doPostRequest("GetPowerStationByID", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetPowerStationByIDResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 电站状态接口
// 查询当前登录者当前组织和下级组织电站,无设备电站将被过滤
func (c *Client) QueryPowerStationMonitor(request QueryPowerStationMonitorRequest) (*QueryPowerStationMonitorResponse, error) {
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("QueryPowerStationMonitor:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doPostRequest("QueryPowerStationMonitor", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := QueryPowerStationMonitorResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 单电站状态详情信息 GetPowerStationMonitorDetail
func (c *Client) GetPowerStationMonitorDetail(request GetPowerStationMonitorDetailRequest) (*GetPowerStationMonitorDetailResponse, error) {
	//requestBytes, err := json.Marshal(&request)
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Println("GetPowerStationMonitorDetail:", string(requestBytes))
	//buf := bytes.NewBuffer(requestBytes)
	params := url.Values{}
	params.Add("id", request.id)
	body, err := c.doGetRequest("GetPowerStationMonitorDetail", params)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetPowerStationMonitorDetailResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 根据 SN 获取电站详情 GetPowerStationMonitorDetailBySn
func (c *Client) GetPowerStationMonitorDetailBySn(request GetPowerStationMonitorDetailBySnRequest) (*GetPowerStationMonitorDetailBySnResponse, error) {
	//requestBytes, err := json.Marshal(&request)
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Println("GetPowerStationMonitorDetail:", string(requestBytes))
	//buf := bytes.NewBuffer(requestBytes)
	params := url.Values{}
	params.Add("sn", request.Sn)
	params.Add("page_index", fmt.Sprintf("%d",request.PageIndex))
	params.Add("page_size", fmt.Sprintf("%d",request.PageSize))
	body, err := c.doGetRequest("GetPowerStationMonitorDetail", params)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetPowerStationMonitorDetailBySnResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

//电站日功率曲线 GetPowerStationPacByDay

// 电站日,月,年发电量 GetPowerStationPower
func (c *Client) GetPowerStationPower(request GetPowerStationPowerRequest) (*GetPowerStationPowerResponse, error) {
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("GetPowerStationPower:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doPostRequest("GetPowerStationPower", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetPowerStationPowerResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 电站设备列表 QueryInventers
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

// 设备实时数据 GetInventersMoreDetailInfoBySn
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

// 设备功率曲线 GetInverterPacByDay
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

// 设备今日发电量/累计发电量/温度/交流电压/曲线 GetInverterDataByColumn
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

// 设备日,月,年发电量 GetInverterPower
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

// 告警列表接口 GetPowerStationWariningInfoByMultiCondition
// 获取当前用户组织自有电站的告警信息，最多查询 60 天内的告警信息
// 注：每条告警信息都有一个 eroorcode,可以根据这个值去查对应错误。
func (c *Client) GetPowerStationWariningInfoByMultiCondition(request GetPowerStationWariningInfoByMultiConditionRequest) (*GetPowerStationWariningInfoByMultiConditionResponse, error) {
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("GetPowerStationWariningInfoByMultiCondition:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doPostRequest("GetPowerStationWariningInfoByMultiCondition", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetPowerStationWariningInfoByMultiConditionResponse{}
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
