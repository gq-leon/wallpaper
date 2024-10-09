package utils

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Download(ctx context.Context, src, dst string) (string, error) {
	if err := ensureDirExists(ctx, dst); err != nil {
		return "", err
	}

	resp, err := http.Get(src)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	fileName := fileNameFromContentDisposition(resp.Header["Content-Disposition"])
	dst = filepath.Join(dst, fileName)
	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	if _, err = io.Copy(out, resp.Body); err != nil {
		return "", err
	}
	return dst, nil
}

func ensureDirExists(ctx context.Context, dirPath string) error {
	// 检查目录是否存在
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// 如果目录不存在，则创建
		err := os.MkdirAll(dirPath, os.ModePerm) // os.ModePerm 设置为 0777
		if err != nil {
			return err
		}
		runtime.LogInfof(ctx, "目录已创建:%s", dirPath)
	}
	return nil
}

func fileNameFromContentDisposition(contentDisposition []string) string {
	for _, part := range contentDisposition {
		keyValue := strings.SplitN(part, "=", 2)
		if len(keyValue) == 2 && strings.Contains(keyValue[0], "filename") {
			return strings.Trim(keyValue[1], `"`)
		}
	}
	return ""
}
