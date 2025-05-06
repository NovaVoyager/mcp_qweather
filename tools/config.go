package tools

const (
	keyHeader      = "X-QW-Api-Key"
	lookup         = "/geo/v2/city/lookup"
	weatherNow     = "/v7/weather/now"
	weather3d      = "/v7/weather/3d"
	weather7d      = "/v7/weather/7d"
	weather10d     = "/v7/weather/10d"
	weather15d     = "/v7/weather/15d"
	weather30d     = "/v7/weather/30d"
	weather24h     = "/v7/weather/24h"
	weather72h     = "/v7/weather/72h"
	weather168h    = "/v7/weather/168h"
	weatherWarning = "/v7/warning/now"
)

var (
	apiKey string
	host   string
)

func SetApiKey(key string) {
	apiKey = key
}

func SetHost(domain string) {
	host = domain
}
