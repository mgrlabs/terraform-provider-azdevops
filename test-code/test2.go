package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type RestAPI struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Capabilities struct {
		Test string `json:"test"`
	} `json:"capabilities"`
}

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

func main() {
	name := "potatoes"
	description := "test description"
	test := "hello there"

	project := RestAPI{
		Name:        name,
		Description: description,
		Test:        test,
	}
	var jsonData []byte
	jsonData, err := json.Marshal(project)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))
}
