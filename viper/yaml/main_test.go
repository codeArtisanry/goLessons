package main

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

// Server example struct for config
type Server struct {
	Host string
	Port int
	List int
}
type Config struct {
	Name string
	Host string
	Port int
}
type Project struct {
	Config []string
}

func TestYaml(t *testing.T) {
	viper.SetConfigName("config")
	// viper.SetConfigType("toml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	fmt.Println(viper.GetStringSlice("hobbies"))
	fmt.Println(viper.GetStringMap("clothing"))
	fmt.Println(viper.GetStringMapString("GetStringMapString"))
	fmt.Println(viper.GetStringMapStringSlice("GetStringMapStringSlice"))

	fmt.Println(viper.Get(""))
	// var s []Server
	// viper.Unmarshal(&s)
	// _ = viper.UnmarshalKey("server", &s)
	// fmt.Printf("%#+v", s)
	var c interface{}
	viper.UnmarshalKey("project", &c)
	fmt.Println(c)
	// // proj := viper.GetStringMapStringSlice("project")
	// proj := viper.GetStringMapStringSlice("project")
	// // proj := viper.GetStringMapStringSlice("scenarios")
	// for k, v := range proj {
	// 	fmt.Println(k, v)
	// }
	// fmt.Println(proj)
	// workers:
	// - period: 10s
	//   job:
	// 	name: hello-world
	// - period: 5s
	//   job:
	// 	name: http-requestor
	// - schedule:
	// 	hour: 17
	// 	minute: 29
	//   job:
	// 	name: http-requestor

	var specs []map[string]interface{}
	err := viper.UnmarshalKey("workers", &specs)
	if err != nil {
		return
	}
	for k, v := range specs {
		fmt.Println(k, v)
		for kk, vv := range v {
			fmt.Println(kk, vv)
			switch vv.(type) {
			case map[string]interface{}:
				for kkk, vvv := range vv.(map[string]interface{}) {
					println()
					fmt.Println(kkk, vvv)
					println()

				}
			}
		}
	}
}
