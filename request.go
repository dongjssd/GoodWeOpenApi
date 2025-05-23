/**
    @author: dongjs
    @date: 2025/2/10
    @description:
**/

package GoodWeOpenApi

import "time"

type GetUserPowerStationRequest struct {
	Key string `json:"key"` // 电站名称或设备 sn
	Type int `json:"type"` // 0:电站 1:sn 默认 0
	PageIndex          int    `json:"page_index"`          //当前页 是
	PageSize           int    `json:"page_size"`           //每页记录数 默认 20 是
}

type GetPowerStationByIDRequest struct {
	Id string `json:"id"`
}

type QueryPowerStationMonitorRequest struct {
	Key                string `json:"key"`                 //50 电站名称或 sn 编码 否
	Adcode             string `json:"adcode"`              //区域编码 否
	Orderby            string `json:"orderby"`             //排序规则入升序：xxxx 降序：xxx desc；capacity：  容量pac_kw：单 kw 功率to_hour：等效小时 eday：今日发电量 etotal：累计发电量 目前只提供这 5 类排序 不填默认 按照电站名称升序
	Condition          *int   `json:"condition,omitempty"`           //0:根据电站查询 1：根据 sn 查询 否,
	PowerstationStatus *int   `json:"powerstation_status,omitempty"` //电站状态(“” or null:全部,-1：离线 //0：待机 1：发电中 2：停机 -2： //无设备),不填代表全部
	PageIndex          int    `json:"page_index"`          //当前页 是
	PageSize           int    `json:"page_size"`           //每页记录数 默认 20 是
}

type GetPowerStationMonitorDetailRequest struct {
	id string `json:"id"` //50 电站 id 是
}

type GetPowerStationMonitorDetailBySnRequest struct {
	Sn        string `json:"sn"`         //50 设备序列号 是
	PageIndex int    `json:"page_index"` //页码起始 1 是
	PageSize  int    `json:"page_size"`  //每页显示记录数 20 是
}

type GetPowerStationPowerRequest struct {
	Id    string    `json:"id"`    //50 电站 Id 是
	Date  time.Time `json:"date"`  //日期 是
	Count int       `json:"count"` //数量(包含当前)最小为 1，如果为 0 则取默认 值 日发电量：最大 31 天，默认 31 月发电量：最大 12 月，默认 12 年发电量：当年默认为 0,最大 10 年 是
	QType int       `json:"type"`  //0:日发电量 1：月发电量 2：年发电量 默认是 0
}

type QueryInventersRequest struct {
	PwId      string `json:"pw_id"`      //50 电站 id 是
	PageIndex int    `json:"page_index"` //页码 是
	PageSize  int    `json:"page_size"`  //每页记录 是
}

type GetInventersMoreDetailInfoBySnRequest struct {
	PwId string `json:"pwId"` //50 电站 id 否
	Sn   string `json:"sn"`   //设备 Sn 是
}

type GetInverterPacByDayRequest struct {
	Id   string `json:"id"`   //50 设备 sn 是
	Date string `json:"date"` //选择时间，日期格式 2018-07-11 是
}

type GetInverterDataByColumnRequest struct {
	Id     string `json:"id"`     //50 设备 sn 是
	Date   string `json:"date"`   //选择时间，日期格式 2018-07-11 是
	Column int    `json:"column"` //0：eday:今日发电量； 是 1：etotal:累计发电量 2：tempperature：温度 3：vac：交流电压 默认 0
}

type GetInverterPowerRequest struct {
	Sn    string `json:"sn"`    //50 设备 Sn 是
	Date  string `json:"date"`  //日期，日期格式 2018-07-11 是
	Count int    `json:"count"` //数量值最小为 1，为 0 则取默认值 包含当前所 选时间 日发电量：最大 31 天，默认 31 月发电量：最大 12 月，默认 12 年发电量：当年默认为 0,最大 10 如果设置值超过规定范围则取默认值 是
	QType int    `json:"type"`  //0:日发电量 1：月发电量 2：年发电量 默认为 0
}

type GetPowerStationWariningInfoByMultiConditionRequest struct {
	Stationid string `json:"stationid"`  //电站 id 否
	Sn        string `json:"sn"`         //设备 Sn 备注：电站和 Sn 都不为空时,按照 Sn 查询 否
	Starttime string `json:"starttime"`  //查询开始时间，开始时间和结束时间 间隔最大 30 天。 是
	Endtime   string `json:"endtime"`    //查询结束时间，如果不赋值默认为当 前时间 否
	PageSize  int    `json:"page_size"`  //默认 20 条 是
	PageIndex int    `json:"page_index"` //页码默认 1 是
	Status    int    `json:"status"`     //1:恢复=恢复未关注+恢复已关注 2:发生=发生未关注+发生已关注 3:全部=恢复+发生 默认显示发生和关注的
}

type GetInverterMoreDataRequest struct {
	Sn    string `json:"sn"`    //设备 Sn 是
	Start string `json:"start"` //开始时间 是
	End   string `json:"end"`   //结束时间 开始和结束时间跨度不允 许超过 24 小时 是
	PwId  string `json:"pwId"`  //电站 id 否
}

type GetInventersDatasRequest struct {
	Sns []string `json:"sns"` //设备序列号集合
}

type GetInvertDatasByPwIdsRequest struct {
	Ids []string `json:"ids"` //电站 id 集合
}
