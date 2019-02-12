package main

import (
	"encoding/json"
)

// type Config struct {
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// 	Server      struct {
// 		Capabilities struct {
// 			sourceControlType string `json:"sourceControlType"`
// 		} `json:"versioncontrol"`
// 		Postgres struct {
// 			TemplateID string `json:"templateTypeId"`
// 		} `json:"processTemplate"`
// 	} `json:"capabilities"`
// }

// {
//   "name": "${managedBy}",
//   "description": "Department: ${department} \nPortfolio Manager: ${portfolioManager} \nProduct Team: ${managedBy} \nProject created: ${dateStamp}",
//   "capabilities": {
//     "versioncontrol": {
//       "sourceControlType": "Git"
//     },
//     "processTemplate": {
//       "templateTypeId": "adcc42ab-9882-485e-a3ed-7678f01f66bc"
//     }
//   }
// }

type App struct {
	Id string `json:"id"`
}

type Org struct {
	Name string `json:"name"`
}

type AppWithOrg struct {
	App
	Org
}

func main() {
	var app AppWithOrg
	test := json.Marshal(app)
}
