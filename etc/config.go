package etc

import (
	"github.com/spf13/viper"
	"fmt"
	"path"
)

var (
	defaultFilePath  = "./etc/config.json"
	Config *viper.Viper
	ContractUsersData = "UsersData"
	ContractFilesData = "FilesData"
	ContractOrgsData = "OrgsData"
	ContractController = "Controller"
	ContractEasyCns = "EasyCns"
	ContractMap = map[string]bool{ContractUsersData:true, ContractFilesData:true, ContractOrgsData:true, ContractController:true, ContractEasyCns:true}
)

func init()  {
	InitConfig(defaultFilePath)
}
func InitConfig(filePath string) {
	Config = viper.New()
	if filePath == "" {
		Config.SetConfigFile(defaultFilePath)
	}else {
		Config.SetConfigFile(filePath)
	}

	err := Config.ReadInConfig()
	if err != nil {
		// do not exit
		//log.Fatal(err)
		fmt.Println("error ===>", err.Error())
	}
}
func GetConfig() *viper.Viper {
	return Config
}
func GetBcosOwnerAddress() string {
	return GetConfig().GetString("bcos.owner.address")
}
func GetBcosOwnerPassword() string {
	return GetConfig().GetString("bcos.owner.password")
}
func GetBcosHostAddress() string {
	return GetConfig().GetString("bcos.host.address")
}
func GetBcosHostRpcPort() int {
	return GetConfig().GetInt("bcos.host.rpc_port")
}
func GetBcosKeystore() string {
	return path.Join(GetServerDir(), GetConfig().GetString("bcos.keystore"))
}
func GetBcosEasyCnsAddress() string {
	return GetConfig().GetString("bcos.easy_cns.address")
}
func GetServerDir() string {
	return GetConfig().GetString("server.dir")
}