package main

import "fmt"

func main() {

	inputJson := `{
			"count": 3,
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
		}`

	fmt.Printf(inputJson)
}
