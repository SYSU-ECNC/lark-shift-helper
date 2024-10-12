package application

import (
	"fmt"
	"os"

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
	data, err := app.getSubmissions()
	if err != nil {
		fmt.Println("getSubmissions error:", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", data)
}
