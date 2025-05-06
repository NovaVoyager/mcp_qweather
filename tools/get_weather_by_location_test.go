package tools

import (
	"testing"
)

func Test_getWeatherByLocation(t *testing.T) {
	resp, err := getWeatherByLocation("104.05,30.63")
	if err != nil {
		t.Errorf("getWeatherByLocation() error = %v", err)
		return
	}
	fmt.Println(*resp)
}
