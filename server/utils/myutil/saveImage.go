package myutil

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"os"
	"path/filepath"
)

// SaveBase64AsPNG 将 Base64 编码的图像数据保存为本地 PNG 文件
// SaveBase64AsPNG 将 Base64 编码的图像数据保存为本地 PNG 文件
func SaveBase64AsPNG(base64Data string, filename string) (string, error) {
	// 创建 avatars 目录（如果不存在）
	err := os.MkdirAll("avatars", 0755)
	if err != nil {
		return "", err
	}

	// 拼接完整的文件路径，并加上 .png 扩展名
	filePath := filepath.Join("avatars", filename+".png")

	// 解码 Base64 数据
	decodedData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return "", err
	}

	// 将解码后的数据写入文件
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 将解码后的数据解析为图像
	img, _, err := image.Decode(bytes.NewReader(decodedData))
	if err != nil {
		return "", err
	}

	// 将图像保存为 PNG 文件
	err = png.Encode(file, img)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
