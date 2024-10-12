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
	data, err := app.getSubmissionsJSON()
	if err != nil {
		fmt.Println("getSubmissions error:", err)
		os.Exit(1)
	}

	submissions, err := convertSubmissionsJSON(data)
	if err != nil {
		fmt.Println("convertSubmissionJSON error:", err)
		os.Exit(1)
	}

	for _, submission := range submissions {
		fmt.Println(submission)
	}
}
