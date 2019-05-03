package callback

var ()

// 2.6.4. deviceDataChanged
type DeviceServiceData struct {
	ServiceId   string `json:"serviceId"`
	ServiceType string `json:"serviceType"`
	Data        map[string]interface{}
	EventTime   string `json:"eventTime"`
}

type CallbackDeviceDataChanged struct {
	NotifyType string            `json:"notifyType"`
	DeviceId   string            `json:"deviceId"`
	GatewayId  string            `json:"gatewayId"`
	RequestId  string            `json:"requestId"`
	service    DeviceServiceData `json:"service"`
}
