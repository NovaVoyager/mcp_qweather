package main

import (
	"fmt"
	"github.com/mark3labs/mcp-go/server"
	"os"
	"weather/tools"
)

func main() {
	setEnv()

	svr := server.NewMCPServer("weather", "1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	svr.AddTool(tools.CityGeoTool(), tools.HandleCityGeo)
	svr.AddTool(tools.WeatherTool(), tools.GetWeather)
	svr.AddTool(tools.GetWeatherByDayTool(), tools.GetWeatherByDay)
	svr.AddTool(tools.GetWeatherByHourlyTool(), tools.GetWeatherByHourly)
	svr.AddTool(tools.WeatherWarningTool(), tools.GetWeatherWarning)

	if err := server.ServeStdio(svr); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func setEnv() {
	key, exists := os.LookupEnv("WEATHER_API_KEY")
	if !exists {
		// 处理环境变量未设置的逻辑
		fmt.Println("WEATHER_API_KEY environment variable is not set")
		os.Exit(1)
	}
	tools.SetApiKey(key)

	domain, exists := os.LookupEnv("WEATHER_HOST")
	if !exists {
		// 处理环境变量未设置的逻辑
		fmt.Println("WEATHER_HOST environment variable is not set")
		os.Exit(1)
	}
	tools.SetHost(domain)
}
