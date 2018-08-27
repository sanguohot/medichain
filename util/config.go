package util

import (
	"github.com/spf13/viper"
	"log"
)

var (
	defaultFilePath  = "./etc/config.json"
	Config *viper.Viper
)
func InitConfig(filePath string) {
	Config = viper.New()
	if filePath == "" {
		Config.SetConfigFile(defaultFilePath)
	}else {
		Config.SetConfigFile(filePath)
	}

	err := Config.ReadInConfig()
	if err!=nil {
		log.Fatal(err)
	}
	//fmt.Println(Config.GetString("bcos.host.address"))
	//fmt.Println(Config.GetInt32("server.host.port"))
	//fmt.Println(Config.GetString("bcos.pki.ca"))
	//fmt.Println(Config.GetString("bcos.pki.cert"))
}
