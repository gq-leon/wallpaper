//go:build windows

package utils

import (
	"os/user"
	"path/filepath"
	"syscall"
	"unsafe"
)

const (
	spiGetDeskWallpaper = 0x0073
	spiSetDeskWallpaper = 0x0014
	uiParam             = 0x0000
	spiFUpdateINIFile   = 0x01
	spiFSendChange      = 0x02
)

var (
	picturesDir               string
	screenWidth, screenHeight uintptr

	user32               = syscall.NewLazyDLL("user32.dll")
	systemParametersInfo = user32.NewProc("SystemParametersInfoW")
)

func ScreenSize() (uintptr, uintptr) {
	if screenWidth == 0 || screenHeight == 0 {
		screenWidth, _, _ = user32.NewProc(`GetSystemMetrics`).Call(0)
		screenHeight, _, _ = user32.NewProc(`GetSystemMetrics`).Call(1)
	}
	return screenWidth, screenHeight
}

func SetFromFile(filename string) error {
	filenameUTF16, err := syscall.UTF16PtrFromString(filename)
	if err != nil {
		return err
	}

	systemParametersInfo.Call(
		uintptr(spiSetDeskWallpaper),
		uintptr(uiParam),
		uintptr(unsafe.Pointer(filenameUTF16)),
		uintptr(spiFUpdateINIFile|spiFSendChange),
	)
	return nil
}

func WallpaperDir() string {
	if picturesDir == "" {
		current, _ := user.Current()
		picturesDir = filepath.Join(current.HomeDir, filepath.Join("Pictures", "Wallpapers"))
	}
	return picturesDir
}
