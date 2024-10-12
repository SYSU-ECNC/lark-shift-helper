package main

import (
	"github.com/SYSU-ECNC/lark-shift-helper/internal/application"
	"github.com/SYSU-ECNC/lark-shift-helper/internal/config"
)

func main() {
	cfg := config.NewConfig()
	cfg.Parse()

	app := application.NewApplication(cfg)
	app.Run()
}
