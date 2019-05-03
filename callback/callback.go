package callback

import "encoding/json"

const (
	ServiceTypeTransmission = "Transmission"
	ServiceTypeConnectivity = "Connectivity"
)

// 2.6.4. deviceDataChanged
type DeviceServiceData struct {
	ServiceId   string          `json:"serviceId"`
	ServiceType string          `json:"serviceType"`
	Data        json.RawMessage `json:"data"`
	EventTime   string          `json:"eventTime"`
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

// 2.6.1 deviceAdded
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

type CallbackDeviceAdded struct {
	NotifyType string         `json:"notifyType"`
	DeviceId   string         `json:"deviceId"`
	GatewayId  string         `json:"gatewayId"`
	NodeType   string         `json:"nodeType"`
	DeviceInfo DeviceInfoData `json:"deviceInfo"`
}
