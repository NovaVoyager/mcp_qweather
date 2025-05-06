package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/parnurzeal/gorequest"
)

func CityGeoTool() mcp.Tool {
	cityGeoTool := mcp.NewTool("city_geo",
		mcp.WithDescription("城市搜索API提供全球地理位位置、全球城市搜索服务，支持经纬度坐标反查、多语言、模糊搜索等功能，获取到需要查询城市或POI的基本信息，包括查询地区的Location ID（你需要这个ID去查询天气），多语言名称、经纬度、时区、海拔、Rank值、归属上级行政区域、所在行政区域等"),
		mcp.WithString("city_name",
			mcp.Required(),
			mcp.Description("城市名称"),
		),
	)

	return cityGeoTool
}

func HandleCityGeo(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	cityName := request.Params.Arguments["city_name"].(string)
	geoResp, err := getCityLocation(cityName)
	if err != nil {
		return nil, err
	}

	contents := make([]mcp.Content, 0, len(geoResp.Location))
	for _, location := range geoResp.Location {
		content := fmt.Sprintf("城市名称：%s\nLocation ID：%s\n经度：%s\n纬度：%s\n时区：%s\n海拔：%s\nRank值：%s\n归属上级行政区域：%s\n所在行政区域：%s",
			location.Name, location.Id, location.Lat, location.Lon, location.Tz, location.Country, location.Rank, location.Adm1, location.Adm2)
		contents = append(contents, mcp.TextContent{
			Type: "text",
			Text: content,
		})
	}

	return &mcp.CallToolResult{
		Content: contents,
	}, nil
}

func getCityLocation(cityName string) (*GeoResp, error) {
	url := host + lookup + "?location=" + cityName
	request := gorequest.New()
	_, body, errs := request.Get(url).Set(keyHeader, apiKey).End()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	var geoResp GeoResp
	err := json.Unmarshal([]byte(body), &geoResp)
	if err != nil {
		return nil, err
	}

	return &geoResp, nil
}
