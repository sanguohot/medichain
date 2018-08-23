package util

import (
	"github.com/spf13/viper"
	"log"
	"fmt"
)
var (
	filePath  = "./etc/config.json"
	Config *viper.Viper
)
func InitConfig() {
	Config = viper.New()
	Config.SetConfigFile(filePath)
	err := Config.ReadInConfig()
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println(Config.GetString("bcos.host.address"))
	fmt.Println(Config.GetInt32("server.host.port"))
	fmt.Println(Config.GetString("bcos.pki.ca"))
	fmt.Println(Config.GetString("bcos.pki.cert"))
}
