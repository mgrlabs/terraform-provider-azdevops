package main

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"server"`
	Postgres struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		DB       string `json:"db"`
	} `json:"database"`
}

func main() {
	jsonConfig := []byte(`{
			"server":{
					"host":"localhost",
					"port":"8080"},
			"database":{
					"host":"localhost",
					"user":"db_user",
					"password":"supersecret",
					"db":"my_db"}}`)
	var config Config
	err := json.Unmarshal(jsonConfig, &config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Config: %+v\n", config)
}
