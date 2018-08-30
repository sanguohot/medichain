package etc

import (
	"github.com/spf13/viper"
	"log"
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
	TEST_OK = "ok"
	TEST_FAIL = "fail"
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
		if filePath != defaultFilePath {
			log.Fatal(err)
		}
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
func GetServerHostAddress() string {
	return GetConfig().GetString("server.host.address")
}
func GetServerHostPort() int {
	return GetConfig().GetInt("server.host.port")
}
func GetServerPkiKey() string {
	return path.Join(GetServerDir(), GetConfig().GetString("server.pki.key"))
}
func GetServerPkiCert() string {
	return path.Join(GetServerDir(), GetConfig().GetString("server.pki.cert"))
}
func GetServerTlsVerifyPeer() bool {
	return GetConfig().GetBool("server.tls.verify_peer")
}
func GetServerTlsEnable() bool {
	return GetConfig().GetBool("server.tls.enable")
}
func GetServerPkiCa() string {
	return path.Join(GetServerDir(), GetConfig().GetString("server.pki.ca"))
}
func GetBigDataHostAddress() string {
	return GetConfig().GetString("big_data.host.address")
}
func GetBigDataHostPort() int {
	return GetConfig().GetInt("big_data.host.port")
}