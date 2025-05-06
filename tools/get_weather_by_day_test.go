package tools

import (
	"testing"
)

func Test_getWeather(t *testing.T) {
	res, err := getDayWeather("https://nb3wry4k6f.re.qweatherapi.com/v7/weather/3d", "104.05,30.63")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
