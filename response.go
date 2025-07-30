/**
    @author: dongjs
    @date: 2025/2/10
    @description:
**/

package GoodWeOpenApi

type GetUserPowerStationResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Record int `json:"record"`
		List   []struct {
			ID             string `json:"id"`              // 电站ID
			PwName         string `json:"pw_name"`         // 电站名称
			ConnDate       string `json:"conn_date"`       // 接入时间
			PwCapacity     string `json:"pw_capacity"`     // 电站容量
			PwType         string `json:"pw_type"`         // 电站类型
			PwInventersnum string `json:"pw_inventersnum"` // 组件数量
			Notes          string `json:"notes"`           // 电站备注
			OrgID          string `json:"org_id"`          // 组织ID
		} `json:"list"`
	} `json:"data"`
}

type GetPowerStationByIDResponse struct {
	Code int    `json:"code"` //0,
	Msg  string `json:"msg"`
	Data struct {
		Id             string `json:"id"`
		PwName         string `json:"pw_name"`
		ConnDate       string `json:"conn_date"`       //接入时间
		PwYieldrate    string `json:"pw_yieldrate"`    // 收益率
		PwInventersnum string `json:"pw_inventersnum"` //组件数量
		PwCapacity     string `json:"pw_capacity"`     //电站容量
		PwAddress      string `json:"pw_address"`      //所在地区
		PwAreacode     string `json:"pw_areacode"`     //区域编码
		PwType         string `json:"pw_type"`         //电站类型
		PwLongitude    string `json:"pw_longitude"`    //经度
		PwLatitude     string `json:"pw_latitude"`     //纬度
		PwSynopsishide bool   `json:"pw_synopsishide"` //否启用简介 false：隐藏 true： 显示
		PwSynopsis     string `json:"pw_synopsis"`     //电站简介
		ListPwowner    []struct {
			OwnerId      string `json:"owner_id"`       //主键
			OwnerPhone   string `json:"owner_phone"`    //业主电话
			OwnerName    string `json:"owner_name"`     //业主备注姓名
			IsFirstOwner string `json:"is_first_owner"` // 是否为第一业主
		} `json:"list_pwowner"`
		ListAttachment []struct {
			Id        string `json:"id"`
			File_name string `json:"file_Name"` //文件名称
			Base64    string `json:"base_64"`   //图片 base64 位编码 是
		} `json:"list_attachment"`
	} `json:"data"`
}

type QueryPowerStationMonitorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Record int `json:"record"` // 记录数
		List   []struct {
			PowerstationId   string  `json:"powerstation_id"`   //电站 id
			Stationname      string  `json:"stationname"`       //电站名称
			Location         string  `json:"location"`          //电站位置
			Adcode           string  `json:"adcode"`            //区域代码
			Status           int     `json:"status"`            //电站状态(-1：离线 0：待机 1： 正常 2：停机 -2：无设备)
			Pac              float64 `json:"pac"`               // 功率(w)
			Capacity         float64 `json:"capacity"`          //容量(kw)
			Longitude        string  `json:"longitude"`         // 经度
			Latitude         string  `json:"latitude"`          //纬度
			Eday             float64 `json:"eday"`              // 今日发电量(kWh)
			Etotal           float64 `json:"etotal"`            //累计发电量(kWh)
			PowerstationType string  `json:"powerstation_type"` //电站类型
			PacKw            float64 `json:"pac_kw"`            //单 kw 功率
			ToHour           float64 `json:"to_hour"`           //日等效小时数
		} `json:"list"`
	} `json:"data"`
}

type GetPowerStationMonitorDetailResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		PowerstationId   string   `json:"powerstation_id"`   // 电站 id
		Stationname      string   `json:"stationname"`       //电站名称
		Address          string   `json:"address"`           //电站地址
		OwnerName        string   `json:"owner_name"`        //联系人
		OwnerPhone       string   `json:"owner_phone"`       //联系电话
		TurnonTime       string   `json:"turnon_time"`       //接入时间
		Capacity         float64  `json:"capacity"`          //电站容量 kw
		Longitude        float64  `json:"longitude"`         //经度
		Latitude         float64  `json:"latitude"`          //纬度
		PowerstationType string   `json:"powerstation_type"` //电站类型
		Status           float64  `json:"status"`            //状态
		Pac              float64  `json:"pac"`               //功率
		Eday             float64  `json:"eday"`              //今日发电量
		Etotal           float64  `json:"etotal"`            //累计发电量
		Images           []string `json:"images"`
	} `json:"data"`
}

type GetPowerStationMonitorDetailBySnResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Record string `json:"record"` //记录数
		List   []struct {
			Powerstation_id   string   `json:"powerstation_id"`   //电站 id
			Stationname       string   `json:"stationname"`       //电站名称
			Address           string   `json:"address"`           //电站地址
			Owner_name        string   `json:"owner_name"`        //联系人
			Owner_phone       string   `json:"owner_phone"`       //联系电话
			Turnon_time       string   `json:"turnon_time"`       //接入时间
			Capacity          float64  `json:"capacity"`          //电站容量 kw
			Longitude         float64  `json:"longitude"`         //经度
			Latitude          float64  `json:"latitude"`          //纬度
			Powerstation_type string   `json:"powerstation_type"` //电站类型
			Status            float64  `json:"status"`            //状态
			Pac               float64  `json:"pac"`               //功率(千瓦)
			Eday              float64  `json:"eday"`              //今日发电量(度)
			Etotal            float64  `json:"etotal"`            //累积发电量(度)
			Images            []string `json:"images"`
		} `json:"list"` // 数据集合
	} `json:"data"`
}

type GetPowerStationPowerResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Date  any     `json:"date"`  //日期
		Power float64 `json:"power"` //发电量
	} `json:"data"`
}

type QueryInventersResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Record int `json:"record"`
		List   []struct {
			Id           string `json:"id"`             //设备 id
			PwId         string `json:"pw_id"`          //电站 id
			ItName       string `json:"it_name"`        //设备名称
			ItSn         string `json:"it_sn"`          //SN
			ItType       string `json:"it_type"`        //设备类型
			ItCapacity   string `json:"it_capacity"`    //容量 kw
			ItCheckcode  string `json:"it_checkcode"`   //校验码 默认不提供
			ConnDate     string `json:"conn_date"`      //接入日期
			ItChangeFlag bool   `json:"it_change_flag"` //是否更换 0:未更换 1：已更换 更换后设备将认 为已经不再发电
		} `json:"list"`
	} `json:"data"`
}

type GetInventersMoreDetailInfoBySnResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Sn           string  `json:"sn"`   //设备序列号
		Name         string  `json:"name"` //设备名称
		InPac        float64 `json:"in_pac"`
		OutPac       float64 `json:"out_pac"`        //发电功率(千瓦)
		Eday         float64 `json:"eday"`           //今日发电量（度）
		Etotal       float64 `json:"etotal"`         //累计发电量（度）
		Status       string  `json:"status"`         //状态
		Turnon_time  string  `json:"turnon_time"`    // 接入时间
		Releation_id string  `json:"releation_id"`   //设备与电站关系 id
		DType        string  `json:"type"`           //设备类型
		Capacity     float64 `json:"capacity"`       //容量 kw
		ItChangeFlag bool    `json:"it_change_flag"` //是否变更
		Tempperature float64 `json:"tempperature"`   //温度
		CheckCode    string  `json:"check_code"`     //验证码
		D            struct {
			CreationDate string  `json:"creationDate"` //数据最新时间
			EDay         float64 `json:"eDay"`         //今日发电量
			ETotal       float64 `json:"eTotal"`       //累计发电量
			Pac          float64 `json:"pac"`          //发电功率
			HTotal       float64 `json:"hTotal"`       //累计发电小时
			Vpv1         float64 `json:"vpv1"`         //直流电压 1
			Vpv2         float64 `json:"vpv2"`         //直流电压 2
			Vpv3         float64 `json:"vpv3"`         //直流电压 3
			Vpv4         float64 `json:"vpv4"`         //直流电压 4
			Ipv1         float64 `json:"ipv1"`         //直流电流 1
			Ipv2         float64 `json:"ipv2"`         //直流电流 2
			Ipv3         float64 `json:"ipv3"`         //直流电流 3
			Ipv4         float64 `json:"ipv4"`         //直流电流 4
			Vac1         float64 `json:"vac1"`         //交流电压 1
			Vac2         float64 `json:"vac2"`         //交流电压 2
			Vac3         float64 `json:"vac3"`         //交流电压 3
			Iac1         float64 `json:"iac1"`         //交流电流 1
			Iac2         float64 `json:"iac2"`         //交流电流 2
			Iac3         float64 `json:"iac3"`         //交流电流 3
			Fac1         float64 `json:"fac1"`         //交流频率 1
			Fac2         float64 `json:"fac2"`         //交流频率 2
			Fac3         float64 `json:"fac3"`         //交流频率 3
			Istr1        float64 `json:"istr1"`        //组串电流 1
			Istr2        float64 `json:"istr2"`        //组串电流 2
			Istr3        float64 `json:"istr3"`        //组串电流 3
			Istr4        float64 `json:"istr4"`        //组串电流 4
			Istr5        float64 `json:"istr5"`        //组串电流 5
			Istr6        float64 `json:"istr6"`        //组串电流 6
			Istr7        float64 `json:"istr7"`        //组串电流 7
			Istr8        float64 `json:"istr8"`        //组串电流 8
			Istr9        float64 `json:"istr9"`        //组串电流 9
			Istr10       float64 `json:"istr10"`       //组串电流 10
			Istr11       float64 `json:"istr11"`       //组串电流 11
			Istr12       float64 `json:"istr12"`       //组串电流 12
			Istr13       float64 `json:"istr13"`       //组串电流 13
			Istr14       float64 `json:"istr14"`       //组串电流 14
			Istr15       float64 `json:"istr15"`       //组串电流 15
			Istr16       float64 `json:"istr16"`       //组串电流 16
		} `json:"d"`
	} `json:"data"`
}

type GetInverterPacByDayResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Date string  `json:"date"` //日期
		Pac  float64 `json:"pac"`  //数值
	} `json:"data"`
}

type GetInverterDataByColumnResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Column1 []struct {
			Date   string  `json:"date"`   //日期时间点
			Column float64 `json:"column"` //数值
		} `json:"column1"`
	} `json:"data"`
}

type GetInverterPowerResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Date   string  `json:"date"`   //日期时间点
		Column float64 `json:"column"` //数值
	} `json:"data"`
}

type GetPowerStationWariningInfoByMultiConditionResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Record int `json:"record"`
		List   []struct {
			StationId    string `json:"stationId"`    //电站 ID
			Adcode       string `json:"adcode"`       //区域编码
			Stationname  string `json:"stationname"`  //电站名称
			Devicesn     string `json:"devicesn"`     //设备 SN
			Devicename   string `json:"devicename"`   //设备名称
			Warningid    string `json:"warningid"`    //告警 id
			Warningname  string `json:"warningname"`  //告警名称
			Status       int    `json:"status"`       //状态 1：已处理 0：未处理
			Happentime   string `json:"happentime"`   //发生时间
			Recoverytime string `json:"recoverytime"` //恢复时间
			IsAddTask    int    `json:"is_add_task"`  //是否派送 1：已派送 2：未派送
			ErrorCode    string `json:"error_code"`   //错误编码-详细参考 3.3 告警错误
		} `json:"list"`
	}
}

type GetInverterMoreDataResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Sn          string  `json:"sn"`   //设备序列号
		Name        string  `json:"name"` //设备名称
		InPac       float64 `json:"in_pac"`
		OutPac      float64 `json:"out_pac"`      //发电功率(千瓦)
		Eday        float64 `json:"eday"`         //今日发电量（度）
		Etotal      float64 `json:"etotal"`       //累计发电量（度）
		Status      int     `json:"status"`       //状态
		TurnonTime  string  `json:"turnon_time"`  //接入时间
		ReleationId string  `json:"releation_id"` //设备与电站关系 id-已取消默认返回空
		Qtype       string  `json:"type"`         //设备类型
		Capacity    float64 `json:"capacity"`     //容量 kw
		D           struct {
			CreationDate string  `json:"creationDate"` //数据最新时间
			EDay         float64 `json:"eDay"`         //今日发电量
			ETotal       float64 `json:"eTotal"`       //累计发电量
			Pac          float64 `json:"pac"`          //发电功率
			HTotal       float64 `json:"hTotal"`       //累计发电小时
			Vpv1         float64 `json:"vpv1"`         //直流电压 1
			Vpv2         float64 `json:"vpv2"`         //直流电压 2
			Vpv3         float64 `json:"vpv3"`         //直流电压 3
			Vpv4         float64 `json:"vpv4"`         //直流电压 4
			Ipv1         float64 `json:"ipv1"`         //直流电流 1
			Ipv2         float64 `json:"ipv2"`         //直流电流 1
			Ipv3         float64 `json:"ipv3"`         //直流电流 1
			Ipv4         float64 `json:"ipv4"`         //直流电流 1
			Vac1         float64 `json:"vac1"`         //交流电压 1
			Vac2         float64 `json:"vac2"`         //交流电压 2
			Vac3         float64 `json:"vac3"`         //交流电压 3
			Iac1         float64 `json:"iac1"`         //交流电流 1
			Iac2         float64 `json:"iac2"`         //交流电流 2
			Iac3         float64 `json:"iac3"`         //交流电流 3
			Fac1         float64 `json:"fac1"`         //交流频率 1
			Fac2         float64 `json:"fac2"`         //交流频率 2
			Fac3         float64 `json:"fac3"`         //交流频率 3
			Istr1        float64 `json:"istr1"`        //组串电流 1
			Istr2        float64 `json:"istr2"`        //组串电流 2
			Istr3        float64 `json:"istr3"`        //组串电流 3
			Istr4        float64 `json:"istr4"`        //组串电流 4
			Istr5        float64 `json:"istr5"`        //组串电流 5
			Istr6        float64 `json:"istr6"`        //组串电流 6
			Istr7        float64 `json:"istr7"`        //组串电流 7
			Istr8        float64 `json:"istr8"`        //组串电流 8
			Istr9        float64 `json:"istr9"`        //组串电流 9
			Istr10       float64 `json:"istr10"`       //组串电流 10
			Istr11       float64 `json:"istr11"`       //组串电流 11
			Istr12       float64 `json:"istr12"`       //组串电流 12
			Istr13       float64 `json:"istr13"`       //组串电流 13
			Istr14       float64 `json:"istr14"`       //组串电流 14
			Istr15       float64 `json:"istr15"`       //组串电流 15
			Istr16       float64 `json:"istr16"`       //组串电流 16
		} `json:"d"`
		ItChangeFlag bool    `json:"it_change_flag"` //是否变更
		Tempperature float64 `json:"tempperature"`   //温度
		CheckCode    string  `json:"check_code"`     //验证码
	} `json:"data"`
}

type GetInventersDatasResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Total     int           `json:"total"`
		Count     int           `json:"count"`
		UnAuthSns []interface{} `json:"unAuthSns"`
		List      []struct {
			Sn          string  `json:"sn"`   //设备序列号
			Name        string  `json:"name"` //设备名称
			InPac       float64 `json:"in_pac"`
			OutPac      float64 `json:"out_pac"`      //发电功率(千瓦)
			Eday        float64 `json:"eday"`         //今日发电量（度）
			Etotal      float64 `json:"etotal"`       //累计发电量（度）
			Status      int     `json:"status"`       //状态
			TurnonTime  string  `json:"turnon_time"`  //接入时间
			ReleationId string  `json:"releation_id"` //设备与电站关系 id-已取消默认返回空
			Qtype       string  `json:"type"`         //设备类型
			Capacity    float64 `json:"capacity"`     //容量 kw
			D           struct {
				CreationDate string  `json:"creationDate"` //数据最新时间
				EDay         float64 `json:"eDay"`         //今日发电量
				ETotal       float64 `json:"eTotal"`       //累计发电量
				Pac          float64 `json:"pac"`          //发电功率
				HTotal       float64 `json:"hTotal"`       //累计发电小时
				Vpv1         float64 `json:"vpv1"`         //直流电压 1
				Vpv2         float64 `json:"vpv2"`         //直流电压 2
				Vpv3         float64 `json:"vpv3"`         //直流电压 3
				Vpv4         float64 `json:"vpv4"`         //直流电压 4
				Ipv1         float64 `json:"ipv1"`         //直流电流 1
				Ipv2         float64 `json:"ipv2"`         //直流电流 1
				Ipv3         float64 `json:"ipv3"`         //直流电流 1
				Ipv4         float64 `json:"ipv4"`         //直流电流 1
				Vac1         float64 `json:"vac1"`         //交流电压 1
				Vac2         float64 `json:"vac2"`         //交流电压 2
				Vac3         float64 `json:"vac3"`         //交流电压 3
				Iac1         float64 `json:"iac1"`         //交流电流 1
				Iac2         float64 `json:"iac2"`         //交流电流 2
				Iac3         float64 `json:"iac3"`         //交流电流 3
				Fac1         float64 `json:"fac1"`         //交流频率 1
				Fac2         float64 `json:"fac2"`         //交流频率 2
				Fac3         float64 `json:"fac3"`         //交流频率 3
				Istr1        float64 `json:"istr1"`        //组串电流 1
				Istr2        float64 `json:"istr2"`        //组串电流 2
				Istr3        float64 `json:"istr3"`        //组串电流 3
				Istr4        float64 `json:"istr4"`        //组串电流 4
				Istr5        float64 `json:"istr5"`        //组串电流 5
				Istr6        float64 `json:"istr6"`        //组串电流 6
				Istr7        float64 `json:"istr7"`        //组串电流 7
				Istr8        float64 `json:"istr8"`        //组串电流 8
				Istr9        float64 `json:"istr9"`        //组串电流 9
				Istr10       float64 `json:"istr10"`       //组串电流 10
				Istr11       float64 `json:"istr11"`       //组串电流 11
				Istr12       float64 `json:"istr12"`       //组串电流 12
				Istr13       float64 `json:"istr13"`       //组串电流 13
				Istr14       float64 `json:"istr14"`       //组串电流 14
				Istr15       float64 `json:"istr15"`       //组串电流 15
				Istr16       float64 `json:"istr16"`       //组串电流 16
			} `json:"d"`
			ItChangeFlag bool    `json:"it_change_flag"` //是否变更
			Tempperature float64 `json:"tempperature"`   //温度
			CheckCode    string  `json:"check_code"`     //验证码
		} `json:"list"`
	} `json:"data"`
}

type GetInvertDatasByPwIdsResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Total     int           `json:"total"`
		Count     int           `json:"count"`
		UnAuthIds []interface{} `json:"unAuthIds"`
		List      []struct {
			Id            string `json:"id"`
			InverterDatas []struct {
				Sn           string  `json:"sn"`
				Name         string  `json:"name"`
				InPac        float64 `json:"in_pac"`
				OutPac       float64 `json:"out_pac"`
				Eday         float64 `json:"eday"`
				Etotal       float64 `json:"etotal"`
				Status       int     `json:"status"`
				TurnonTime   string  `json:"turnon_time"`
				ReleationId  string  `json:"releation_id"`
				Type         string  `json:"type"`
				Capacity     float64 `json:"capacity"`
				ItChangeFlag bool    `json:"it_change_flag"`
				Tempperature float64 `json:"tempperature"`
				CheckCode    string  `json:"check_code"`
				D            struct {
					CreationDate string  `json:"creationDate"` //数据最新时间
					EDay         float64 `json:"eDay"`         //今日发电量
					ETotal       float64 `json:"eTotal"`       //累计发电量
					Pac          float64 `json:"pac"`          //发电功率
					HTotal       float64 `json:"hTotal"`       //累计发电小时
					Vpv1         float64 `json:"vpv1"`         //直流电压 1
					Vpv2         float64 `json:"vpv2"`         //直流电压 2
					Vpv3         float64 `json:"vpv3"`         //直流电压 3
					Vpv4         float64 `json:"vpv4"`         //直流电压 4
					Ipv1         float64 `json:"ipv1"`         //直流电流 1
					Ipv2         float64 `json:"ipv2"`         //直流电流 1
					Ipv3         float64 `json:"ipv3"`         //直流电流 1
					Ipv4         float64 `json:"ipv4"`         //直流电流 1
					Vac1         float64 `json:"vac1"`         //交流电压 1
					Vac2         float64 `json:"vac2"`         //交流电压 2
					Vac3         float64 `json:"vac3"`         //交流电压 3
					Iac1         float64 `json:"iac1"`         //交流电流 1
					Iac2         float64 `json:"iac2"`         //交流电流 2
					Iac3         float64 `json:"iac3"`         //交流电流 3
					Fac1         float64 `json:"fac1"`         //交流频率 1
					Fac2         float64 `json:"fac2"`         //交流频率 2
					Fac3         float64 `json:"fac3"`         //交流频率 3
					Istr1        float64 `json:"istr1"`        //组串电流 1
					Istr2        float64 `json:"istr2"`        //组串电流 2
					Istr3        float64 `json:"istr3"`        //组串电流 3
					Istr4        float64 `json:"istr4"`        //组串电流 4
					Istr5        float64 `json:"istr5"`        //组串电流 5
					Istr6        float64 `json:"istr6"`        //组串电流 6
					Istr7        float64 `json:"istr7"`        //组串电流 7
					Istr8        float64 `json:"istr8"`        //组串电流 8
					Istr9        float64 `json:"istr9"`        //组串电流 9
					Istr10       float64 `json:"istr10"`       //组串电流 10
					Istr11       float64 `json:"istr11"`       //组串电流 11
					Istr12       float64 `json:"istr12"`       //组串电流 12
					Istr13       float64 `json:"istr13"`       //组串电流 13
					Istr14       float64 `json:"istr14"`       //组串电流 14
					Istr15       float64 `json:"istr15"`       //组串电流 15
					Istr16       float64 `json:"istr16"`       //组串电流 16
				} `json:"d"`
			} `json:"inverter_datas"`
		} `json:"list"`
	} `json:"data"`
}
