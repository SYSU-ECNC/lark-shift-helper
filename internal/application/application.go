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

	assistantsDetailMap, err := app.getAssistantsDetailMapFromSubmissions(submissions)
	if err != nil {
		fmt.Println("getAssistantsDetailMapFromSubmissions error:", err)
		os.Exit(1)
	}

	for key, value := range assistantsDetailMap {
		fmt.Println(key, value)
	}

	shift := app.generateShift(availableAssistants, assistantsDetailMap)

	// 检查 shift，如果有 shift 有某个时间段的 len 不足 4，那么就补齐 ""
	for i := 0; i < 5; i++ {
		for j := 0; j < 7; j++ {
			for len(shift[i][j]) < 4 {
				shift[i][j] = append(shift[i][j], "")
			}
		}
	}

	if err := app.Write(shift); err != nil {
		fmt.Println("Write error:", err)
		os.Exit(1)
	}
}
