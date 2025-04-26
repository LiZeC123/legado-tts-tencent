package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config 结构体用于映射JSON配置
type Config struct {
	SecretId  string `json:"secretId"`
	SecretKey string `json:"secretKey"`
	Region    string `json:"region"`
}

// LoadConfig 从当前目录读取config.json文件并解析为Config结构体
func LoadConfig() (*Config, error) {
	// 1. 读取配置文件
	configFile := "config.json"
	fileContent, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("无法读取配置文件: %v", err)
	}

	// 2. 初始化配置结构体
	var config Config

	// 3. 反序列化JSON内容
	if err := json.Unmarshal(fileContent, &config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}

	return &config, nil
}
