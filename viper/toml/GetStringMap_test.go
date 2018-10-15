package app

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

//TestViper try to use viper
func TestViper(t *testing.T) {
	viper.AddConfigPath("./")
	viper.SetConfigName("test")
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("read viper configure", err)
	}
	//

	var x int
	if viper.IsSet("logger.meta.other.xo") {
		x = viper.GetInt("logger.meta.other.xo")
	} else {
		x = 8
	}

	if viper.IsSet("logger.meta.input") {
		metaInput := viper.GetStringMap("logger.meta.input")
		for k, v := range metaInput {
			// if viper.IsSet("logger.meta.input." + k) {
			// 	kv := viper.GetStringMapString("logger.meta.input." + k)
			// 	for x, y := range kv {
			// 		fmt.Println(x, y)
			// 	}
			// }
			fmt.Println(k, v)

			if value, ok := v.(map[string]string); ok {
				for x, y := range value {
					fmt.Println(x, y)
				}
			} else {
				fmt.Println(ok)
			}
		}
	}
	// for k, v := range logConf["meta"]["input"] {
	// 	fmt.Println(k, v)
	// }
	fmt.Println(x)
}
