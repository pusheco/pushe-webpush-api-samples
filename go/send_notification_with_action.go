package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	// Obtain token -> https://docs.pushe.co/docs/web-api/authentication
	const token = "YOUR_TOKEN"

	// Webpush doc -> https://docs.pushe.co/docs/web-api/notification-actions
	reqData := map[string]interface{}{
		"app_ids":  []string{"YOUR_APP_ID"},
		"data": map[string]interface{}{
			"title":   "Title",
			"content": "Content",

			// Actions -> https://docs.pushe.co/docs/web-api/notification-actions
			"action": map[string]interface{}{
				"action_type": "U",
				"url":         "https://pushe.co",
			},

			"buttons": []map[string]interface{}{
				{
					"btn_content": "YOUR_CONTENT",
					"btn_action": map[string]interface{}{
						"action_type":         "U",
						"url":                 "https://pushe.co",
					},
					"btn_order": 0,
				},
                {
					"btn_content": "YOUR_CONTENT",
					"btn_action": map[string]interface{}{
						"action_type":         "U",
						"url":                 "https://pushe.co",
					},
					"btn_order": 1,
				},
			},
		},
		// additional keywords -> https://docs.pushe.co/docs/web-api/notification-keys
	}

	// Marshal returns the JSON encoding of reqData.
	reqJSON, err := json.Marshal(reqData)

	// check encoded json
	if err != nil {
		fmt.Println("json:", err)
		return
	}

	// create request obj
	request, err := http.NewRequest(
		http.MethodPost,
		"https://api.pushe.co/v2/messaging/notifications/web/",
		bytes.NewBuffer(reqJSON),
	)

	// check request
	if err != nil {
		fmt.Println("Req error:", err)
		return
	}

	// set header
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Token "+token)

	// send request and get response
	client := http.Client{}
	response, err := client.Do(request)

	// check response
	if err != nil {
		fmt.Println("Resp error:", err)
		return
	}

	defer response.Body.Close()

	// check status_code and response
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(response.Body)
	respContent := buf.String()

	fmt.Println("status code =>", response.StatusCode)
	fmt.Println("response =>", respContent)
	fmt.Println("==========")

	if response.StatusCode == http.StatusCreated {
		fmt.Println("success!")

		var respData map[string]interface{}
		_ = json.Unmarshal(buf.Bytes(), &respData)

		var reportURL string

		// hashed_id just generated for Non-Free plan
		if respData["hashed_id"] != nil {
			reportURL = "https://pushe.co/report?id=" + respData["hashed_id"].(string)
		} else {
			reportURL = "no report url for your plan"
		}

		fmt.Println("report_url:", reportURL)
		fmt.Println("notification id:", int(respData["wrapper_id"].(float64)))
	} else {
		fmt.Println("failed!")
	}
}
