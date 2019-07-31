package callback

import "encoding/json"

// 2.6.4. deviceDataChanged
type DeviceServiceData struct {
	ServiceId   string          `json:"serviceId"`
	ServiceType string          `json:"serviceType"`
	Data        json.RawMessage `json:"data"` // custom by yourself
	EventTime   string          `json:"eventTime"`
}

type ConnectivityData struct {
	SignalPower int `json:"SignalPower "`
	CellID      int `json:"CellID"`
	SNR         int `json:"SNR"`
	Battery     int `json:"battery"`
}

type CallbackDeviceDataChanged struct {
	NotifyType string            `json:"notifyType"`
	DeviceId   string            `json:"deviceId"`
	GatewayId  string            `json:"gatewayId"`
	RequestId  string            `json:"requestId"`
	Service    DeviceServiceData `json:"service"`
}

type DeviceInfoData struct {
	DeviceType        string `json:"deviceType"`
	SupportedSecurity string `json:"supportedSecurity"`
	IsSecurity        string `json:"isSecurity"`
	SwVersion         string `json:"swVersion"`
	SerialNumber      string `json:"serialNumber"`
	ManufacturerName  string `json:"manufacturerName"`
	SignalStrength    string `json:"signalStrength"`
	ManufacturerId    string `json:"manufacturerId"`
	Description       string `json:"description"`
	StatusDetail      string `json:"statusDetail"`
	Mute              string `json:"mute"`
	ProtocolType      string `json:"protocolType"`
	Mac               string `json:"mac"`
	HwVersion         string `json:"hwVersion"`
	SigVersion        string `json:"sigVersion"`
	BridgeId          string `json:"bridgeId"`
	Name              string `json:"name"`
	Location          string `json:"location"`
	Model             string `json:"model"`
	FwVersion         string `json:"fwVersion"`
	NodeId            string `json:"nodeId"`
	Status            string `json:"status"`
	BatteryLevel      string `json:"batteryLevel"`
}

// 2.6.1 deviceAdded
type CallbackDeviceAdded struct {
	NotifyType string         `json:"notifyType"`
	DeviceId   string         `json:"deviceId"`
	GatewayId  string         `json:"gatewayId"`
	NodeType   string         `json:"nodeType"`
	DeviceInfo DeviceInfoData `json:"deviceInfo"`
}

// 2.6.2 bindDevice
type CallbackBindDevice struct {
	NotifyType string         `json:"notifyType"`
	DeviceId   string         `json:"deviceId"`
	ResultCode string         `json:"resultCode"`
	DeviceInfo DeviceInfoData `json:"deviceInfo"`
}

// 2.6.3 deviceInfoChanged
type CallbackDeviceInfoChanged struct {
	NotifyType string         `json:"notifyType"`
	DeviceId   string         `json:"deviceId"`
	GatewayId  string         `json:"gatewayId"`
	DeviceInfo DeviceInfoData `json:"deviceInfo"`
}
