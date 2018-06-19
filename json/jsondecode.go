package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type ServerCache struct {
	Host string
	Port int
}

type GrayLog struct {
	Host string
	Port int
}

type Config struct {
	Cpu    int    `json:"cpu,string"`
	Web    string `json:"web"`
	Rpc    string `json:"-"`
	Ns     string
	Set    string
	Cache  []ServerCache `json:"caches,omitempty"`
	Logger GrayLog
}

// загрузка конфига
func New() *Config {
	c := new(Config)

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	file, _ := os.Open("conf.json")
	fmt.Println("Config loaded from: " + dir + "/conf.json")
	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		fmt.Println("Configure failed: ", err)
	}

	return c
}
func main() {
	conf := New()
	log.Println(conf.Web)
	log.Println(conf.Cpu)
	log.Println(conf.Rpc)
}
