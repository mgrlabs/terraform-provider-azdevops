package main

import (
	"encoding/json"
	"fmt"
)

func dumpMap(space string, m map[string]interface{}) {
	for k, v := range m {
		if mv, ok := v.(map[string]interface{}); ok {
			fmt.Printf("{ \"%v\": \n", k)
			dumpMap(space+"\t", mv)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v %v : %v\n", space, k, v)
		}
	}
}

var jsonStr = `
{
	"count": 3,
	"value": [
		{
			"typeId": "27450541-8e31-4150-99x47-dc59f998fc01",
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
}
`

func main() {
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		panic(err)
	}
	dumpMap("", jsonMap)
}
