# 和风天气 MCP Server
## 简介
和风天气的 MCP server，目前支持国内的天气查询，包含有地理位置坐标查询，城市天气实时、每日天气、逐小时天气预报以及天气预警。

## 工具列表
- 获取城市位置信息 `city_geo`
    - 把城市名称转为天气查询使用的地理坐标信息
    - 输入：city_name
    - 输出：返回城市地理位置信息

- 每日天气预报 `get_weather_by_day`
    - 获取每日的天气预报
    - 输入：day (需要查询的天数，支持 3d、7d、10d、15d、30d，最少3d，最多30d)
    - 输入：location

- 24-168小时逐小时天气预报 `get_weather_by_hours`
    - 输入：hours （需要查询的小时数，支持 24h、72h、168h最少24d，最多168d）
    - 输入：location

- 获取城市天气预报 `city_weather`
    - 输入：location

- 获取天气预警信息 `weather_warning`
    - 输入：location

## 快速开始
#### 注册账号
1. 在和风天气创建账号并登录
2. 进入开发者控制台
3. 创建项目，获取 api key和api domain

### 使用服务
服务是由 golang 开发，直接打包为二进制进行使用

进入到项目根目录执行 `go build`
根据需求打包为不同的平台

### 在 vscode 中使用

```
{
  "mcpServers": {
	  "weather": {
      "command": "C:/project/MCP/weather/weather.exe",
      "args": [],
	  "env":{
		  "WEATHER_API_KEY":和风天气API key,
		  "WEATHER_HOST":和风天气请求的domain，每个用户的不一样
	  }
    }
  }
}
```

### 更新
| 版本     | 内容     | 时间        |
|--------|--------|-----------|
| v1.0.0 | 天气查询工具 | 2025年5月6日 |
