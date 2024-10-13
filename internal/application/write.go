package application

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Fields struct {
	TimeSlot  string `json:"时间段"`
	Monday    string `json:"周一"`
	Tuesday   string `json:"周二"`
	Wednesday string `json:"周三"`
	Thursday  string `json:"周四"`
	Friday    string `json:"周五"`
	Saturday  string `json:"周六"`
	Sunday    string `json:"周日"`
}

type Resquest struct {
	Fields Fields `json:"fields"`
}

func (app *Application) Write(shift [5][7][]string) error {
	number2TimeSlot := func(number int) string {
		switch number {
		case 0:
			return "09:00-10:00"
		case 1:
			return "10:00-12:00"
		case 2:
			return "13:30-16:10"
		case 3:
			return "16:10-18:00"
		case 4:
			return "19:00-21:00"
		default:
			return ""
		}
	}

	// timeslot 1~4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if err := app.WriteHelper(number2TimeSlot(i), shift[i][1][j], shift[i][2][j], shift[i][3][j], shift[i][4][j], shift[i][5][j], shift[i][6][j], shift[i][0][j]); err != nil {
				return err
			}
		}
		if err := app.WriteHelper("", "", "", "", "", "", "", ""); err != nil {
			return err
		}
	}

	// timeslot 5
	i := 4
	for j := 0; j < 3; j++ {
		if err := app.WriteHelper(number2TimeSlot(i), shift[i][1][j], shift[i][2][j], shift[i][3][j], shift[i][4][j], shift[i][5][j], shift[i][6][j], shift[i][0][j]); err != nil {
			return err
		}
	}

	return nil
}

func (app *Application) WriteHelper(timeSlot string, mon string, tue string, wed string, thu string, fri string, sat string, sun string) error {
	client := &http.Client{}

	jsonData := Resquest{
		Fields: Fields{
			TimeSlot:  timeSlot,
			Monday:    mon,
			Tuesday:   tue,
			Wednesday: wed,
			Thursday:  thu,
			Friday:    fri,
			Saturday:  sat,
			Sunday:    sun,
		},
	}

	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"https://open.feishu.cn/open-apis/bitable/v1/apps/%s/tables/%s/records",
			app.cfg.APPToken,
			app.cfg.OutputTableID,
		),
		bytes.NewBuffer(jsonBytes),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", app.cfg.UserAccessToken))

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}
