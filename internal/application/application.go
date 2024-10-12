package application

import (
	"fmt"

	"github.com/SYSU-ECNC/lark-shift-helper/internal/config"
)

type Application struct {
	cfg *config.Config
}

func NewApplication(cfg *config.Config) *Application {
	return &Application{
		cfg: cfg,
	}
}

func (app *Application) Run() {
	fmt.Println("APP ID:", app.cfg.APPID)
	fmt.Println("APP Secret:", app.cfg.APPSecret)
}
