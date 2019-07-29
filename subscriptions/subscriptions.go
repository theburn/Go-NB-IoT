package subscriptions

import (
	"encoding/json"
	"fmt"

	"github.com/theburn/Go-NB-IoT/client"
	"github.com/theburn/Go-NB-IoT/configure"
	log "github.com/theburn/Go-NB-IoT/logging"
)

type NodeType struct {
	BindDevice         string
	DeviceAdded        string
	DeviceInfoChanged  string
	DeviceDataChanged  string
	DeviceDatasChanged string
	DeviceDeleted      string
	ServiceInfoChanged string
	RuleEvent          string
	DeviceModelAdded   string
	DeviceModelDeleted string
}

var NodeTypeDict = NodeType{
	BindDevice:         "bindDevice",
	DeviceAdded:        "deviceAdded",
	DeviceInfoChanged:  "deviceInfoChanged",
	DeviceDataChanged:  "deviceDataChanged",
	DeviceDatasChanged: "deviceDatasChanged",
	DeviceDeleted:      "deviceDeleted",
	ServiceInfoChanged: "serviceInfoChanged",
	RuleEvent:          "ruleEvent",
	DeviceModelAdded:   "deviceModelAdded",
	DeviceModelDeleted: "deviceModelDeleted",
}

/*
1.bindDevice（绑定设备，订阅后推送绑定设备通知）
2.deviceAdded（添加新设备，订阅后推送注册设备通知）
3.deviceInfoChanged（设备信息变化，订阅后推送设备信息变化通知）
4.deviceDataChanged（设备数据变化，订阅后推送设备数据变化通知）
5.deviceDatasChanged（设备数据批量变化，订阅后推送批量设备数据变化通知）
6.deviceDeleted（删除设备，订阅后推送删除设备通知）
7.serviceInfoChanged（服务信息变化，订阅后推送设备服务信息变化通知）
8.ruleEvent（规则事件，订阅后推送规则事件通知）
9.deviceModelAdded（添加设备模型，订阅后推送增加设备模型通知）
10.deviceModelDeleted（删除设备模型，订阅后推送删除设备模型通知）
*/

type SubscriptionsBusinessReq struct {
	NotifyType  string `json:"notifyType"`
	CallbackUrl string `json:"callbackUrl"`
	AppId       string `json:"appId"`
}

type SubscriptionsBusinessResp struct {
	SubscriptionId string `json:"subscriptionId"`
	CallbackUrl    string `json:"callbackUrl"`
	NotifyType     string `json:"notifyType"`
}

type SubscriptionsBusinessBatchQueryResp struct {
	TotalCount    int                         `json:"totalCount"`
	PageNo        int                         `json:"pageNo"`
	PageSize      int                         `json:"pageSize"`
	Subscriptions []SubscriptionsBusinessResp `json:"subscriptions"`
}

func (s *SubscriptionsBusinessReq) SubscriptionsBusiness(c *client.NBHttpClient) (*SubscriptionsBusinessResp, error) {

	reqRespParam := client.ReqRespParam{}
	reqRespParam.URL = configure.NBIoTConfig.ReqParam.IoTHost +
		subBusinessURI

	reqRespParam.Method = "POST"
	reqRespParam.ContentType = "application/json"

	var err error

	if reqRespParam.ReqBody, err = json.Marshal(s); err != nil {
		log.Error("json Marshal Failed, ", s, err)
		return nil, err
	}

	if err = c.Request(&reqRespParam); err != nil {
		log.Error("Request error!", err)
		return nil, err
	}

	var subResp SubscriptionsBusinessResp

	if err = json.Unmarshal(reqRespParam.RespBody, &subResp); err != nil {
		log.Error("json SubscriptionsBusinessResp Unmarshal Failed, ", reqRespParam.RespBody, err)
		return nil, err
	}

	log.Debugf("+%v", string(reqRespParam.RespBody))
	return &subResp, err
}

func (s *SubscriptionsBusinessResp) SubscriptionsQuerySingle(c *client.NBHttpClient) (*SubscriptionsBusinessResp, error) {

	reqRespParam := client.ReqRespParam{}
	reqRespParam.URL = fmt.Sprintf(configure.NBIoTConfig.ReqParam.IoTHost+subQuerySingleURI,
		s.SubscriptionId, configure.NBIoTConfig.ReqParam.AppID)

	reqRespParam.Method = "GET"
	reqRespParam.ContentType = "application/json"

	var err error

	if err = c.Request(&reqRespParam); err != nil {
		log.Error("Request error!", err)
		return nil, err
	}

	var subResp SubscriptionsBusinessResp

	if err = json.Unmarshal(reqRespParam.RespBody, &subResp); err != nil {
		log.Error("json SubscriptionsBusinessResp Unmarshal Failed, ", reqRespParam.RespBody, err)
		return nil, err
	}

	log.Debugf("+%v", string(reqRespParam.RespBody))
	return &subResp, err
}

func SubscriptionsQueryBatch(c *client.NBHttpClient, notifyType string) (*SubscriptionsBusinessBatchQueryResp, error) {

	reqRespParam := client.ReqRespParam{}
	reqRespParam.URL = fmt.Sprintf(configure.NBIoTConfig.ReqParam.IoTHost+subQueryBatchURI,
		configure.NBIoTConfig.ReqParam.AppID, notifyType, 0, 20)

	reqRespParam.Method = "GET"
	reqRespParam.ContentType = "application/json"

	var err error

	if err = c.Request(&reqRespParam); err != nil {
		log.Error("Request error!", err)
		return nil, err
	}

	var subResp SubscriptionsBusinessBatchQueryResp

	if err = json.Unmarshal(reqRespParam.RespBody, &subResp); err != nil {
		log.Error("json SubscriptionsBusinessResp Unmarshal Failed, ", reqRespParam.RespBody, err)
		return nil, err
	}

	log.Debugf("+%v", subResp)
	return &subResp, err
}
