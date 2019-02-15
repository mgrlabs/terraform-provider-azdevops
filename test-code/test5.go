package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type App struct {
		Count int `json:"count"`
		Value []struct {
			TypeID string `json:"typeId"`
			Name   string `json:"name"`
			// ReferenceName       interface{} `json:"referenceName"`
			// Description         string      `json:"description"`
			// ParentProcessTypeID string      `json:"parentProcessTypeId"`
			// IsEnabled           bool        `json:"isEnabled"`
			// IsDefault           bool        `json:"isDefault"`
			// CustomizationType   string      `json:"customizationType"`
		} `json:"value"`
	}

	data := []byte(`{
		"count": 4,
		"value": [
			{
				"typeId": "27450541-8e31-4150-9947-dc59f998fc01",
				"name": "CMMI",
				"referenceName": null,
				"description": "This template is for more formal projects requiring a framework for process improvement and an auditable record of decisions.",
				"parentProcessTypeId": "00000000-0000-0000-0000-000000000000",
				"isEnabled": true,
				"isDefault": false,
				"customizationType": "system"
			},
			{
				"typeId": "6b724908-ef14-45cf-84f8-768b5384da45",
				"name": "Scrum",
				"referenceName": null,
				"description": "This template is for teams who follow the Scrum framework.",
				"parentProcessTypeId": "00000000-0000-0000-0000-000000000000",
				"isEnabled": true,
				"isDefault": false,
				"customizationType": "system"
			},
			{
				"typeId": "b8a3a935-7e91-48b8-a94c-606d37c3e9f2",
				"name": "Basic",
				"referenceName": null,
				"description": "This template is flexible for any process and great for teams getting started with Azure DevOps.",
				"parentProcessTypeId": "00000000-0000-0000-0000-000000000000",
				"isEnabled": true,
				"isDefault": false,
				"customizationType": "system"
			},
			{
				"typeId": "adcc42ab-9882-485e-a3ed-7678f01f66bc",
				"name": "Agile",
				"referenceName": null,
				"description": "This template is flexible and will work great for most teams using Agile planning methods, including those practicing Scrum.",
				"parentProcessTypeId": "00000000-0000-0000-0000-000000000000",
				"isEnabled": true,
				"isDefault": true,
				"customizationType": "system"
			}
		]
	}`)

	var app App
	err := json.Unmarshal(data, &app)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", app)
	// fmt.Println(app.Value)

}
