package application

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Submission struct {
	name                    string
	isAvailableForFieldWork bool
	isEligibleAsLeader      bool
	isNewAssistant          bool
	availableTime           [7][]string
}

func (app *Application) getSubmissionsJSON() (map[string]interface{}, error) {
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

func convertSubmissionsJSON(v map[string]interface{}) ([]*Submission, error) {
	data, ok := v["data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("data field not found")
	}

	hasMore, ok := data["has_more"].(bool)
	if !ok {
		return nil, fmt.Errorf("has_more field not found")
	}
	if hasMore {
		fmt.Println("Warning: has_more is true, the result may be incomplete")
	}

	items, ok := data["items"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("items field not found")
	}

	submissions := make([]*Submission, 0)
	for _, item := range items {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("item is not a map[string]interface{}")
		}

		fields, ok := itemMap["fields"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("fields field not found")
		}

		submission, err := convertFields2Submission(fields)
		if err != nil {
			return nil, err
		}

		submissions = append(submissions, submission)
	}

	return submissions, nil
}

func convertFields2Submission(fields map[string]interface{}) (*Submission, error) {
	// Get name
	nameArraySlice, ok := fields["姓名"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("姓名 field not found")
	}
	nameField, ok := nameArraySlice[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("姓名 field is not a map[string]interface{}")
	}

	name, ok := nameField["text"].(string)
	if !ok {
		return nil, fmt.Errorf("can not get name from name field")
	}

	// Get isAvailableForFieldWork
	isAvailableForFieldWorkField, ok := fields["是否可以安排外勤"].(string)
	if !ok {
		return nil, fmt.Errorf("isAvailableForFieldWork field not found")
	}
	isAvailableForFieldWork := isAvailableForFieldWorkField == "是"

	// Get isEligibleAsLeader
	isEligibleAsLeaderField, ok := fields["是否可以担任负责人"].(string)
	if !ok {
		return nil, fmt.Errorf("isEligibleAsLeader field not found")
	}
	isEligibleAsLeader := isEligibleAsLeaderField == "是"

	// Get isNewAssistant
	isNewAssistantField, ok := fields["是否是新助理"].(string)
	if !ok {
		return nil, fmt.Errorf("isNewAssistant field not found")
	}
	isNewAssistant := isNewAssistantField == "是"

	// Get availableTime
	sundayAvailableTimeSlice, _ := fields["周日空闲时间"].([]interface{})
	mondayAvailableTimeSlice, _ := fields["周一空闲时间"].([]interface{})
	tuesdayAvailableTimeSlice, _ := fields["周二空闲时间"].([]interface{})
	wednesdayAvailableTimeSlice, _ := fields["周三空闲时间"].([]interface{})
	thursdayAvailableTimeSlice, _ := fields["周四空闲时间"].([]interface{})
	fridayAvailableTimeSlice, _ := fields["周五空闲时间"].([]interface{})
	saturdayAvailableTimeSlice, _ := fields["周六空闲时间"].([]interface{})

	sundayAvailableTime := convertToStringSlice(sundayAvailableTimeSlice)
	mondayAvailableTime := convertToStringSlice(mondayAvailableTimeSlice)
	tuesdayAvailableTime := convertToStringSlice(tuesdayAvailableTimeSlice)
	wednesdayAvailableTime := convertToStringSlice(wednesdayAvailableTimeSlice)
	thursdayAvailableTime := convertToStringSlice(thursdayAvailableTimeSlice)
	fridayAvailableTime := convertToStringSlice(fridayAvailableTimeSlice)
	saturdayAvailableTime := convertToStringSlice(saturdayAvailableTimeSlice)

	availableTime := [7][]string{
		sundayAvailableTime,
		mondayAvailableTime,
		tuesdayAvailableTime,
		wednesdayAvailableTime,
		thursdayAvailableTime,
		fridayAvailableTime,
		saturdayAvailableTime,
	}

	return &Submission{
		name:                    name,
		isAvailableForFieldWork: isAvailableForFieldWork,
		isEligibleAsLeader:      isEligibleAsLeader,
		isNewAssistant:          isNewAssistant,
		availableTime:           availableTime,
	}, nil
}

func convertToStringSlice(slice []interface{}) []string {
	strSlice := make([]string, len(slice))
	for i, v := range slice {
		strSlice[i], _ = v.(string)
	}
	return strSlice
}
