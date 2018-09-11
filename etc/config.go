package etc

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
)
// auto generate struct
// https://mholt.github.io/json-to-go/
// use mapstructure to replace json for '_' key words, e.g. rpc_port,big_data
type ConfigStruct struct {
	Server struct {
		Host struct {
			Address string `json:"address"`
			Port    int    `json:"port"`
		} `mapstructure:"host"`
		TLS struct {
			Enable     bool `json:"enable"`
			VerifyPeer bool `json:"verify_peer"`
		} `mapstructure:"tls"`
		Pki struct {
			Key  string `json:"key"`
			Cert string `json:"cert"`
			Ca   string `json:"ca"`
		} `mapstructure:"pki"`
	} `mapstructure:"server"`
	BigData struct {
		Host struct {
			Address string `json:"address"`
			Port    int    `json:"port"`
		} `json:"host"`
	} `mapstructure:"big_data"`
	Bcos struct {
		Host struct {
			Address     string `json:"address"`
			RPCPort     int    `mapstructure:"rpc_port"`
			P2PPort     int    `mapstructure:"p2p_port"`
			ChannelPort int    `mapstructure:"channel_port"`
		} `mapstructure:"host"`
		Pki struct {
			Key  string `json:"key"`
			Cert string `json:"cert"`
			Ca   string `json:"ca"`
		} `mapstructure:"pki"`
		Owner struct {
			PrivateKey string `json:"private_key"`
			PublicKey  string `json:"public_key"`
			Address    string `json:"address"`
			Password   string `json:"password"`
		} `mapstructure:"owner"`
		Keystore string `json:"keystore"`
		EasyCns  struct {
			Address string `json:"address"`
			Tx      string `json:"tx"`
		} `mapstructure:"easy_cns"`
	} `mapstructure:"bcos"`
}

var (
	defaultFilePath  = "./etc/config.json"
	ViperConfig *viper.Viper
	Config *ConfigStruct
	ContractUsersData = "UsersData"
	ContractFilesData = "FilesData"
	ContractOrgsData = "OrgsData"
	ContractController = "Controller"
	ContractEasyCns = "EasyCns"
	ContractAuth = "Auth"
	ContractMap = map[string]bool{ContractUsersData:true, ContractFilesData:true, ContractOrgsData:true, ContractController:true, ContractEasyCns:true, ContractAuth:true}
	TEST_OK = "ok"
	TEST_FAIL = "fail"
	FileTypeList = []string{"预约", "接诊", "影像采集", "治疗方案", "诊断报告"}
	FileTypeMap = map[common.Hash]string{}
	serverPath = os.Getenv("MEDICHAIN_PATH")
)

func initFileTypeMap()  {
	for _, item := range FileTypeList {
		hash := crypto.Keccak256Hash([]byte(item))
		FileTypeMap[hash] = item
	}
}
func init()  {
	if serverPath == "" {
		log.Fatal("MEDICHAIN_PATH env required")
	}
	InitConfig(defaultFilePath)
	initFileTypeMap()
}
func InitConfig(filePath string) {
	ViperConfig = viper.New()
	if filePath == "" {
		ViperConfig.SetConfigFile(defaultFilePath)
	}else {
		ViperConfig.SetConfigFile(filePath)
	}

	err := ViperConfig.ReadInConfig()
	if err != nil {
		if filePath != defaultFilePath {
			log.Fatal(err)
		}
	}
	err = ViperConfig.Unmarshal(&Config)
	if err != nil {
		log.Fatal(err)
	}
}
func GetConfig() *ConfigStruct {
	return Config
}
func GetViperConfig() *viper.Viper {
	return ViperConfig
}
func GetBcosOwnerAddress() string {
	return GetViperConfig().GetString("bcos.owner.address")
}
func GetBcosOwnerPassword() string {
	return GetViperConfig().GetString("bcos.owner.password")
}
func GetBcosHostAddress() string {
	return GetViperConfig().GetString("bcos.host.address")
}
func GetBcosHostRpcPort() int {
	return GetViperConfig().GetInt("bcos.host.rpc_port")
}
func GetBcosKeystore() string {
	return path.Join(GetServerDir(), GetViperConfig().GetString("bcos.keystore"))
}
func GetBcosEasyCnsAddress() string {
	return GetViperConfig().GetString("bcos.easy_cns.address")
}
func GetServerDir() string {
	//return GetViperConfig().GetString("server.dir")
	return serverPath
}
func GetServerHostAddress() string {
	return GetViperConfig().GetString("server.host.address")
}
func GetServerHostPort() int {
	return GetViperConfig().GetInt("server.host.port")
}
func GetServerPkiKey() string {
	return path.Join(GetServerDir(), GetViperConfig().GetString("server.pki.key"))
}
func GetServerPkiCert() string {
	return path.Join(GetServerDir(), GetViperConfig().GetString("server.pki.cert"))
}
func GetServerTlsVerifyPeer() bool {
	return GetViperConfig().GetBool("server.tls.verify_peer")
}
func GetServerTlsEnable() bool {
	return GetViperConfig().GetBool("server.tls.enable")
}
func GetServerPkiCa() string {
	return path.Join(GetServerDir(), GetViperConfig().GetString("server.pki.ca"))
}
func GetBigDataHostAddress() string {
	return GetViperConfig().GetString("big_data.host.address")
}
func GetBigDataHostPort() int {
	return GetViperConfig().GetInt("big_data.host.port")
}
func GetBigDataRootFolderCode() string {
	return GetViperConfig().GetString("big_data.root_folder.code")
}
func GetBigDataRootFolderName() string {
	return GetViperConfig().GetString("big_data.root_folder.name")
}
func GetSqliteFilePath() string {
	return path.Join(GetServerDir(), "databases", GetViperConfig().GetString("sqlite.file_name"))
}