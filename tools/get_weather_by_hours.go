package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/parnurzeal/gorequest"
)

func GetWeatherByHourlyTool() mcp.Tool {
	weatherTool := mcp.NewTool("get_weather_by_hours",
		mcp.WithDescription("提供城市未来24-168小时逐小时天气预报，包括：温度、天气状况、风力、风速、风向、相对湿度、大气压强、降水概率、露点温度、云量"),
		mcp.WithString("hours",
			mcp.Required(),
			mcp.Description("(必选)需要查询的小时数，支持 24h、72h、168h最少24d，最多168d，(h指的小时)"),
		),
		mcp.WithString("location",
			mcp.Required(),
			mcp.Description("需要查询地区的LocationID或以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位），LocationID可通过GeoAPI获取。例如 location=101010100 或 location=116.41,39.92"),
		),
	)

	return weatherTool
}

func GetWeatherByHourly(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	hours := request.Params.Arguments["hours"].(string)
	location := request.Params.Arguments["location"].(string)
	URI := getHourlyURI(hours)
	url := weatherHost + URI
	resp, err := getHourlyWeather(url, location)
	if err != nil {
		return nil, err
	}

	content := make([]mcp.Content, 0)
	for _, item := range resp.Hourly {
		str := fmt.Sprintf(hourlyText,
			item.FxTime, item.Temp, item.Text, item.WindDir, item.WindScale, item.WindSpeed, item.Humidity, item.Precip,
			item.Pop, item.Pressure, item.Cloud, item.Dew,
		)

		content = append(content, mcp.TextContent{
			Text: str,
			Type: "text",
		})
	}

	return &mcp.CallToolResult{
		Content: content,
	}, nil
}

func getHourlyWeather(url, location string) (*HourlyResp, error) {
	request := gorequest.New()
	_, body, errs := request.Get(url).Set(keyHeader, apiKey).Param("location", location).End()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	var hoursResp HourlyResp
	err := json.Unmarshal([]byte(body), &hoursResp)
	if err != nil {
		return nil, err
	}

	return &hoursResp, nil
}

func getHourlyURI(hours string) string {
	uri := ""
	switch hours {
	case "24h":
		uri = weather24h
	case "72h":
		uri = weather72h
	case "168h":
		uri = weather168h
	default:
		uri = weather24h
	}

	return uri
}

var hourlyText = `
预报时间：%s
温度：%s
天气状况：%s
风向：%s
风力等级：%s
风速：%s
相对湿度：%s
当前小时累计降水量：%s
逐小时预报降水概率：%s
大气压强：%s
云量：%s
露点温度：%s
`
