package tools

import (
	"fmt"
	"testing"
)

func Test_getCityLocation(t *testing.T) {
	resp, err := getCityLocation("武侯区")
	if err != nil {
		t.Errorf("getCityLocation() error = %v", err)
		return
	}
	fmt.Println(*resp)
}
