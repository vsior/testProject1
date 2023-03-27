package configjson

type Config struct {
	IpStartService string `json: ip_start_service`
	IpGetData      string `json: ip_get_data`
	IpSendData     string `json: ip_send_data`
}
