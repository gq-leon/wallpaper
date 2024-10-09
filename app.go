package main

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"net/http"
	"wallpaper/utils"
)

var (
	imageSource = "https://picsum.photos/"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet() string {
	width, height := utils.ScreenSize()

	url := fmt.Sprintf("%s%d/%d", imageSource, width, height)
	// 创建一个自定义的 HTTP 客户端，不跟随重定向
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 返回 http.ErrUseLastResponse 来阻止跟随重定向
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Get(url)
	if err != nil {
		return fmt.Sprintf("未获取到图片地址%s", err)
	}
	defer resp.Body.Close()

	// 如果是 302 重定向，获取重定向地址
	if resp.StatusCode == http.StatusFound {
		location := resp.Header.Get("Location")
		return location
	}
	return url
}

func (a *App) SetWallpaper(src string) {
	file, err := utils.Download(a.ctx, src, utils.WallpaperDir())
	if err != nil {
		runtime.LogErrorf(a.ctx, "download failed: %s", err)
		return
	}

	if err = utils.SetFromFile(file); err != nil {
		runtime.LogErrorf(a.ctx, "set wallpaper failed: %s", err)
	}
}
