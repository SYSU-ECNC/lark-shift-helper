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

	submissions, err := app.convertSubmissionsJSON(data)
	if err != nil {
		fmt.Println("convertSubmissionJSON error:", err)
		os.Exit(1)
	}

	availableAssistants, err := app.getAvailableAssistants(submissions)
	if err != nil {
		fmt.Println("getAvailableAssistants error:", err)
		os.Exit(1)
	}

	_ = availableAssistants

	users, err := app.getUsersFromSubmissions(submissions)
	if err != nil {
		fmt.Println("getUsersFromSubmissions error:", err)
		os.Exit(1)
	}

	for _, user := range users {
		fmt.Println(user)
	}
}
