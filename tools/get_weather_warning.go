package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/parnurzeal/gorequest"
)

func WeatherWarningTool() mcp.Tool {
	tool := mcp.NewTool("weather_warning",
		mcp.WithDescription("天气灾害预警API可以获取中国及全球多个国家或地区官方发布的实时天气灾害预警数据"),
		mcp.WithString("location",
			mcp.Required(),
			mcp.Description("需要查询地区的LocationID或以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位），LocationID可通过GeoAPI获取。例如 location=101010100 或 location=116.41,39.92"),
		),
	)

	return tool
}

func GetWeatherWarning(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	location := request.Params.Arguments["location"].(string)
	url := host + weatherWarning
	resp, err := getWeatherWarning(url, location)
	if err != nil {
		return nil, err
	}

	if resp.Warning == nil || len(resp.Warning) == 0 {
		return mcp.NewToolResultText("当前区域没有预警信息"), nil
	}

	content := make([]mcp.Content, 0)
	for _, item := range resp.Warning {
		str := fmt.Sprintf(weatherWarningText, item.Sender, item.PubTime, item.Title, item.StartTime, item.EndTime,
			item.Status, item.Severity, item.SeverityColor, item.Text,
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

func getWeatherWarning(url, location string) (*WeatherWarningResp, error) {
	request := gorequest.New()
	_, body, errs := request.Get(url).Set(keyHeader, apiKey).Param("location", location).End()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	var resp WeatherWarningResp
	err := json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

var weatherWarningText = `
预警发布单位:%s
发布时间：%s
预警信息标题:%s
预警开始时间:%s
预警结束时间:%s
预警信息的发布状态:%s
预警严重等级:%s
预警严重等级颜色:%s
预警详细文字描述:%s
`
