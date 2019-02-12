// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"log"
// 	"net/http"
// )

// // func main() {
// // 	message := map[string]interface{}{
// // 		"name":        "Test Project",
// // 		"description": "Test Description",
// // 		"capabilities": {
// // 			"versioncontrol": {
// // 				"sourceControlType": "Git"},
// // 			"processTemplate": {
// // 				"templateTypeId": "adcc42ab-9882-485e-a3ed-7678f01f66bc"
// // 			}
// // 		}
// // 	}

// 	bytesRepresentation, err := json.Marshal(message)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(bytesRepresentation))
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	var result map[string]interface{}

// 	json.NewDecoder(resp.Body).Decode(&result)

// 	log.Println(result)
// 	log.Println(result["data"])
// }
