package config

import "github.com/liam-z/bitgate/constants"

const (
	BaseUrl = "https://api.bitget.com"
	WsUrl   = "wss://ws.bitget.com/mix/v1/stream"

	ApiKey        = ""
	SecretKey     = ""
	PASSPHRASE    = ""
	TimeoutSecond = 30
	SignType      = constants.SHA256
)
