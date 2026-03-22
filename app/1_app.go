package app

import (
	"context"
	"dayz-server-tools/logger"
	"os"
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
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) CloseApp() {
	a.ctx.Done()
	logger.Info("应用程序关闭")

	os.Exit(0)
}
