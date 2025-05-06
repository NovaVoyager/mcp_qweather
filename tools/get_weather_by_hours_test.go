package tools

import (
	"testing"
)

func Test_getHourlyWeather(t *testing.T) {
	res, err := getHourlyWeather("https://nb3wry4k6f.re.qweatherapi.com/v7/weather/24h", "104.05,30.63")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
