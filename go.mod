module github.com/sanguohot/medichain

require (
	github.com/BurntSushi/toml v0.3.0 // indirect
	github.com/aristanetworks/goarista v0.0.0-20180907105523-ff33da284e76 // indirect
	github.com/btcsuite/btcd v0.0.0-20180903232927-cff30e1d23fc // indirect
	github.com/cespare/cp v1.0.0 // indirect
	github.com/deckarep/golang-set v0.0.0-20180831180637-cbaa98ba5575 // indirect
	github.com/edsrzf/mmap-go v0.0.0-20170320065105-0bce6a688712 // indirect
	github.com/ethereum/go-ethereum v1.8.15
	github.com/fjl/memsize v0.0.0-20180427083637-f6d5545993d6 // indirect
	github.com/gin-contrib/sse v0.0.0-20170109093832-22d885f9ecc7 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/protobuf v1.2.0 // indirect
	github.com/golang/snappy v0.0.0-20180518054509-2e65f85255db // indirect
	github.com/google/uuid v1.0.0
	github.com/hashicorp/golang-lru v0.5.0 // indirect
	github.com/hpcloud/tail v1.0.0 // indirect
	github.com/huin/goupnp v0.0.0-20180415215157-1395d1447324 // indirect
	github.com/jackpal/go-nat-pmp v1.0.1 // indirect
	github.com/json-iterator/go v1.1.5 // indirect
	github.com/karalabe/hid v0.0.0-20180420081245-2b4488a37358 // indirect
	github.com/luopotaotao/IdValidator v0.0.0-20170417064513-0c93c35562c6
	github.com/mattn/go-colorable v0.0.9 // indirect
	github.com/mattn/go-isatty v0.0.4 // indirect
	github.com/mattn/go-sqlite3 v1.9.0
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/onsi/ginkgo v1.6.0 // indirect
	github.com/onsi/gomega v1.4.1 // indirect
	github.com/pborman/uuid v0.0.0-20180909234540-25cd46ecac86 // indirect
	github.com/pkg/errors v0.8.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rjeczalik/notify v0.9.1 // indirect
	github.com/rs/cors v1.5.0 // indirect
	github.com/sanguohot/chardet v0.0.0-20180903090850-20e22483e848
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24
	github.com/spacemonkeygo/openssl v0.0.0-20180905133406-3b86b4289698
	github.com/spf13/viper v1.2.0
	github.com/stretchr/testify v1.2.2 // indirect
	github.com/syndtr/goleveldb v0.0.0-20180815032940-ae2bd5eed72d // indirect
	github.com/ugorji/go/codec v0.0.0-20180831062425-e253f1f20942 // indirect
	github.com/urfave/cli v1.20.0
	github.com/vedranvuk/magnet v0.0.0-20130912231909-9a8503cc3543
	go.uber.org/atomic v1.3.2 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.9.1
	golang.org/x/crypto v0.0.0-20180910181607-0e37d006457b // indirect
	golang.org/x/net v0.0.0-20180906233101-161cd47e91fd // indirect
	golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f // indirect
	golang.org/x/text v0.3.0
	golang.org/x/tools v0.0.0-20180910180008-18207bb12d3a // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/karalabe/cookiejar.v2 v2.0.0-20150724131613-8dcd6a7f4951 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/urfave/cli.v1 v1.20.0 // indirect
)

replace (
	github.com/ethereum/go-ethereum v1.8.15 => github.com/sanguohot/go-ethereum v1.8.14-0.20180905013405-eb10ed570864
	golang.org/x/crypto v0.0.0-20180910181607-0e37d006457b => github.com/golang/crypto v0.0.0-20180910181607-0e37d006457b // indirect
	golang.org/x/net v0.0.0-20180906233101-161cd47e91fd => github.com/golang/net v0.0.0-20180906233101-161cd47e91fd // indirect
	golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f => github.com/golang/sync v0.0.0-20180314180146-1d60e4601c6f // indirect
	golang.org/x/sys v0.0.0-20180906133057-8cf3aee42992 => github.com/golang/sys v0.0.0-20180906133057-8cf3aee42992
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
	golang.org/x/tools v0.0.0-20180910180008-18207bb12d3a => github.com/golang/tools v0.0.0-20180910180008-18207bb12d3a // indirect

)
