package httpx

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(url string, outPath string) error {
	// 创建一个 HTTP 客户端
	client := &http.Client{}

	// 创建一个 GET 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error downloading file:", err)
		return err
	}
	defer resp.Body.Close()

	// 创建一个文件用于保存下载的内容
	out, err := os.Create(outPath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer out.Close()

	// 将响应体内容复制到文件中
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return err
	}
	return nil
}
