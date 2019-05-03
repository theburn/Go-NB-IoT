package callback

const (
	ServiceTypeTransmission = "Transmission"
	ServiceTypeConnectivity = "Connectivity"
)

// 2.6.4. deviceDataChanged
type DeviceServiceData struct {
	ServiceId   string `json:"serviceId"`
	ServiceType string `json:"serviceType"`
	Data        interface{}
	EventTime   string `json:"eventTime"`
}

type ConnectivityData struct {
	SignalPower int `json:"SignalPower "`
	CellID      int `json:"CellID"`
	SNR         int `json:"SNR"`
	Battery     int `json:"battery"`
}

type LiquidTransmissionData struct {
	Tank_ID_sign   int `json:"Tank_ID_sign"`
	Liquid_precent int `json:"liquid_precent"`
}

type CallbackDeviceDataChanged struct {
	NotifyType string            `json:"notifyType"`
	DeviceId   string            `json:"deviceId"`
	GatewayId  string            `json:"gatewayId"`
	RequestId  string            `json:"requestId"`
	Service    DeviceServiceData `json:"service"`
}
