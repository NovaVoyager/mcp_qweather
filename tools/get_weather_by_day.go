package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/parnurzeal/gorequest"
)

func GetWeatherByDayTool() mcp.Tool {
	weatherTool := mcp.NewTool("get_weather_by_day",
		mcp.WithDescription("每日天气预报API，提供未来3-30天天气预报，包括：日出日落、月升月落、最高最低温度、天气白天和夜间状况、风力、风速、风向、相对湿度、大气压强、降水量、露点温度、紫外线强度、能见度等"),
		mcp.WithString("day",
			mcp.Required(),
			mcp.Description("(必选)需要查询的天数，支持 3d、7d、10d、15d、30d，最少3d，最多30d，(d指的天)"),
		),
		mcp.WithString("location",
			mcp.Required(),
			mcp.Description("需要查询地区的LocationID或以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位），LocationID可通过GeoAPI获取。例如 location=101010100 或 location=116.41,39.92"),
		),
	)

	return weatherTool
}

func GetWeatherByDay(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	day := request.Params.Arguments["day"].(string)
	location := request.Params.Arguments["location"].(string)
	URI := getDayURI(day)
	url := host + URI
	resp, err := getDayWeather(url, location)
	if err != nil {
		return nil, err
	}

	content := make([]mcp.Content, 0)
	for _, item := range resp.Daily {
		str := fmt.Sprintf(dayText,
			item.FxDate, item.Sunrise, item.Sunset, item.TempMax, item.TempMin, item.TextDay, item.TextNight, item.WindDirDay,
			item.WindScaleDay, item.WindSpeedDay, item.WindDirNight, item.WindScaleNight, item.WindSpeedNight, item.Precip,
			item.UvIndex, item.Humidity, item.Pressure, item.Vis, item.Cloud,
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

func getDayWeather(url, location string) (*DailyResp, error) {
	request := gorequest.New()
	_, body, errs := request.Get(url).Set(keyHeader, apiKey).Param("location", location).End()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	var dailyResp DailyResp
	err := json.Unmarshal([]byte(body), &dailyResp)
	if err != nil {
		return nil, err
	}

	return &dailyResp, nil
}

func getDayURI(day string) string {
	uri := ""
	switch day {
	case "3d":
		uri = weather3d
	case "7d":
		uri = weather7d
	case "10d":
		uri = weather10d
	case "15d":
		uri = weather15d
	case "30d":
		uri = weather30d
	default:
		uri = weather3d
	}

	return uri
}

var dayText = `
预报日期：%s
日出时间：%s
日落时间：%s
预报当天最高温度：%s
预报当天最低温度：%s
预报白天天气状况：%s
预报晚间天气状况：%s
预报白天风向：%s
预报白天风力等级：%s
预报白天风速：%s
预报夜间当天风向：%s
预报夜间风力等级：%s
预报夜间风速：%s
预报当天总降水量：%s
紫外线强度指数：%s
相对湿度：%s
大气压强：%s
能见度：%s
云量：%s
`
