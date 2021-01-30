package data_service

import "clouddist/pkg/cloud"

// CreatePath 创建路径
func CreatePath(username string) bool {
	return cloud.CreateUSER(username)
}

// GetPath 获取目录下内容
func GetPath(username, path string) []cloud.Item {
	return cloud.ShowDir(username, path)
}
