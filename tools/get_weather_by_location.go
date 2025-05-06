package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/parnurzeal/gorequest"
)

func WeatherTool() mcp.Tool {
	weatherTool := mcp.NewTool("city_weather",
		mcp.WithDescription("获取中国3000+市县区的实时天气数据，包括实时温度、体感温度、风力风向、相对湿度、大气压强、降水量、能见度、露点温度、云量等"),
		mcp.WithString("location",
			mcp.Required(),
			mcp.Description("需要查询地区的LocationID或以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位），LocationID可通过GeoAPI获取。例如 location=101010100 或 location=116.41,39.92"),
		),
	)

	return weatherTool
}

func GetWeather(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	location := request.Params.Arguments["location"].(string)
	weatherResp, err := getWeatherByLocation(location)
	if err != nil {
		return nil, err
	}

	content := mcp.TextContent{
		Type: "text",
		Text: fmt.Sprintf("当前时间：%s\n天气更新时间：%s\n天气链接：%s\n天气状况:%s\n当前温度：%s\n体感温度：%s\n风力风向：%s\n相对湿度：%s\n大气压强：%s\n降水量：%s\n能见度：%s\n露点温度：%s\n云量：%s",
			weatherResp.Now.ObsTime, weatherResp.UpdateTime, weatherResp.FxLink, weatherResp.Now.Text, weatherResp.Now.Temp, weatherResp.Now.FeelsLike, weatherResp.Now.WindDir, weatherResp.Now.Humidity, weatherResp.Now.Pressure, weatherResp.Now.Precip, weatherResp.Now.Vis, weatherResp.Now.Dew, weatherResp.Now.Cloud,
		),
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			content,
		},
	}, nil
}

func getWeatherByLocation(location string) (*WeatherResp, error) {
	url := host + weatherNow + "?location=" + location
	request := gorequest.New()
	_, body, errs := request.Get(url).Set(keyHeader, apiKey).End()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	var weatherResp WeatherResp
	err := json.Unmarshal([]byte(body), &weatherResp)
	if err != nil {
		return nil, err
	}

	return &weatherResp, nil
}
