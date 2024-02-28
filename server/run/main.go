package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// EventData represents the structure of the JSON data
type EventData struct {
	Adid           string            `json:"adid"`
	AppToken       string            `json:"app_token"`
	CallbackParams map[string]string `json:"callback_params"`
	Currency       string            `json:"currency"`
	Environment    string            `json:"environment"`
	EventToken     string            `json:"event_token"`
	Revenue        string            `json:"revenue"`
	S2S            int               `json:"s2s"`
}

func main() {
	strJson := "[\n{\"adid\":\"7707f2343d1ccb707a6d4544cfa6c17f\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"cd582b59-970c-45ba-a7b0-3c47d252e5de\",\"user_id\":\"3sta\"},\"currency\":\"USD\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"6.99\",\"s2s\":1},\n{\"adid\":\"176479255fa4e5a380c009c573fbb4e9\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"3d5215d2-0f66-4a4f-be99-77e33d75e427\",\"user_id\":\"3ufl\"},\"currency\":\"USD\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"14.99\",\"s2s\":1},\n{\"adid\":\"9ae137fd80cc9413c3cc98b1ec711081\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"07a4fdc4-7523-4a2f-a308-2b12cd8930d2\",\"user_id\":\"3t2f\"},\"currency\":\"CAD\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"14.99\",\"s2s\":1},\n{\"adid\":\"5cc8c7b0c42022047fb5aa4d25fda6dd\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"a3cc26bf-c2e8-4237-9662-bb6de116f47b\",\"user_id\":\"3uoa\"},\"currency\":\"USD\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"2.99\",\"s2s\":1},\n{\"adid\":\"1eb88a39dd75a557d57d6519f80aab7b\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"5d3719fa-d9b4-46ee-88db-f653a7486ef8\",\"user_id\":\"3uu1\"},\"currency\":\"USD\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"6.99\",\"s2s\":1},\n{\"adid\":\"d6a5cb474bf6ea5f402b03422c06ec34\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"7134b7c1-ac0d-4005-ad1d-a9ae46fce48f\",\"user_id\":\"3rsq\"},\"currency\":\"USD\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"6.99\",\"s2s\":1},\n{\"adid\":\"2153fc88b00b9da3b5ca6cdc425a9d4a\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"b757005e-9d5c-4b1c-b3b4-84a9ec4c63e9\",\"user_id\":\"3urx\"},\"currency\":\"USD\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"6.99\",\"s2s\":1},\n{\"adid\":\"2153fc88b00b9da3b5ca6cdc425a9d4a\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"8a7803cb-debd-4158-b524-c516e9272315\",\"user_id\":\"3urx\"},\"currency\":\"USD\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"2.99\",\"s2s\":1},\n{\"adid\":\"c53df3efd02395df78805feeca9bce72\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"f2a3d701-c73c-445c-9657-6e8d5979d556\",\"user_id\":\"3us9\"},\"currency\":\"CAD\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"6.99\",\"s2s\":1},\n{\"adid\":\"5bfa9b0d67253c6d361280d07ecdac52\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"02f81548-2322-460b-b756-d6c9116fb5e1\",\"user_id\":\"3uxq\"},\"currency\":\"USD\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"6.99\",\"s2s\":1},\n{\"adid\":\"f35c1e6c09f88be64c022449039f933d\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"3bf0448b-c4bc-4528-a98e-e9dad1696e0f\",\"user_id\":\"3uzb\"},\"currency\":\"GBP\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"6.99\",\"s2s\":1},\n{\"adid\":\"4654fe725d8e6e81714a0191fd5b6f4b\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"9fa7d1af-2aa0-49c3-82eb-9eb2ed2d1e04\",\"user_id\":\"3v0x\"},\"currency\":\"MYR\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"6.99\",\"s2s\":1},\n{\"adid\":\"72ff6b56e83514471e0392d8a055b36a\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"580e6540-4264-488f-ace7-7b824c3a8d8f\",\"user_id\":\"3v1l\"},\"currency\":\"USD\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"6.99\",\"s2s\":1},\n{\"adid\":\"c82162851e692bdaa7a8d2752f90dfe2\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"25ab55f8-56ee-4e00-91c9-8b6e54379a0f\",\"user_id\":\"3v13\"},\"currency\":\"USD\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"6.99\",\"s2s\":1},\n{\"adid\":\"1fc29a312378777b01afe5ab9a87d937\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"d375a9df-3824-4c4d-a0e8-c7f0a47db684\",\"user_id\":\"3v0c\"},\"currency\":\"USD\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"6.99\",\"s2s\":1},\n{\"adid\":\"32be80955fac5b6d5a0fa1f37c98d182\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"6107ba51-3b7b-4720-be07-15f8af00e507\",\"user_id\":\"3uoc\"},\"currency\":\"USD\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"6.99\",\"s2s\":1},\n{\"adid\":\"0d91c4f016eb7921b01c638f04a13436\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"7c90ee02-db6e-44b8-84f3-5dbeb49b7e5b\",\"user_id\":\"3v69\"},\"currency\":\"CAD\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"2.99\",\"s2s\":1},\n{\"adid\":\"aaeb02325b4322e7b1ae7cfb8ea32f51\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"ae22b6e0-e2e5-46d9-95b5-baaf75067a7c\",\"user_id\":\"3v5r\"},\"currency\":\"USD\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"2.99\",\"s2s\":1},\n{\"adid\":\"aaeb02325b4322e7b1ae7cfb8ea32f51\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"d87ea554-54e5-48d6-b46a-96e93934405f\",\"user_id\":\"3v5r\"},\"currency\":\"USD\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"6.99\",\"s2s\":1},\n{\"adid\":\"cb23e494aeae03f442d60ee50976da26\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"2a84971a-2885-4120-b993-4872985c22ee\",\"user_id\":\"3v6p\"},\"currency\":\"USD\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"6.99\",\"s2s\":1},\n{\"adid\":\"1efaff57183b6e37f9f30b2772926697\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"2fe5f17d-def4-45c1-a8ae-fa7882f5d594\",\"user_id\":\"3v4i\"},\"currency\":\"USD\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"6.99\",\"s2s\":1},\n{\"adid\":\"00cd82b0337de60ebb212ee05e27989d\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"0d193665-6e4e-4591-92ba-5d73de243b5f\",\"user_id\":\"3v8l\"},\"currency\":\"GBP\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"2.99\",\"s2s\":1},\n{\"adid\":\"e725ae994d07634a0faa3d7a75b74b36\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"a2bfccad-3279-4901-ad8f-5da4251a159b\",\"user_id\":\"3v8y\"},\"currency\":\"CAD\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"14.99\",\"s2s\":1},\n{\"adid\":\"f4187e8e90b64fc475c9334939a76e5f\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"508e4051-d963-425e-bd72-b5de85f9d2f4\",\"user_id\":\"3v9i\"},\"currency\":\"USD\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"2.99\",\"s2s\":1},\n{\"adid\":\"db9d57e0e3db3084c2dab2c231d49484\",\"app_token\":\"2k6i15xddsg0\",\"callback_params\":{\"order_no\":\"e75b0455-3ab1-4820-94a4-e7a0f1c0f776\",\"user_id\":\"3ug4\"},\"currency\":\"USD\",\"environment\":\"production\",\"event_token\":\"tkkjbz\",\"revenue\":\"6.99\",\"s2s\":1}]"
	// 解析代码
	var data []EventData
	err := json.Unmarshal([]byte(strJson), &data)
	if err != nil {
		panic(err)
	}
	// Iterate over the JSON data
	for _, entry := range data {
		// Modify app_token and event_token
		entry.AppToken = "qa0o5avjiznk"
		entry.EventToken = "bteuq4"

		// Convert the JSON object to a map to use in the form data
		formData := make(url.Values)
		formData.Set("adid", entry.Adid)
		formData.Set("app_token", entry.AppToken)
		for k, v := range entry.CallbackParams {
			formData.Set(fmt.Sprintf("callback_params[%s]", k), v)
		}
		formData.Set("currency", entry.Currency)
		formData.Set("environment", entry.Environment)
		formData.Set("event_token", entry.EventToken)
		formData.Set("revenue", entry.Revenue)
		formData.Set("s2s", fmt.Sprintf("%d", entry.S2S))

		// Create a new request
		req, err := http.NewRequest("POST", "https://s2s.adjust.com/event", bytes.NewBufferString(formData.Encode()))
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			continue
		}

		// Set headers
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		// Send the request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error sending request for adid %s: %v\n", entry.Adid, err)
			continue
		}
		defer resp.Body.Close()

		// Check the response
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Error response from server for adid %s: %s\n", entry.Adid, resp.Status)
		} else {
			fmt.Printf("Successfully sent data for adid %s\n", entry.Adid)
		}
	}
}
