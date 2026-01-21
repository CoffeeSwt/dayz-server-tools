package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"dayz-server-tools/app"
	"dayz-server-tools/logger"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	logger.SetOutput(logger.GetFileWriter())

	logger.Info("正在启动Dayz Server Tools")
	// Create an instance of the app structure
	a := app.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "Dayz Server Tools",
		Width:     1056,
		MinWidth:  1056,
		Height:    752,
		MinHeight: 752,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: a.Startup,
		Bind: append([]interface{}{
			a,
		}, app.GetBind()...),
		Frameless:        true,
		WindowStartState: options.Minimised,
	})

	if err != nil {
		logger.Error("Wails 启动失败:", "error", err.Error())
	}
}
