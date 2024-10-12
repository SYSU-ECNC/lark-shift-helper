package application

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (app *Application) getSubmissions() (map[string]interface{}, error) {
	client := &http.Client{}
	jsonData := `{}`
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"https://open.feishu.cn/open-apis/bitable/v1/apps/%s/tables/%s/records/search",
			app.cfg.APPToken,
			app.cfg.TableID,
		),
		bytes.NewBuffer([]byte(jsonData)),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", app.cfg.UserAccessToken))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}
