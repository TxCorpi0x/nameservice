module github.com/vjdmhd/nameservice

go 1.15

require (
	github.com/cosmos/cosmos-sdk v0.39.1
	github.com/cosmos/sdk-tutorials/nameservice/nameservice v0.0.0-20210126204308-89adb1fe6914
	github.com/gorilla/mux v1.8.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.33.9
	github.com/tendermint/tm-db v0.5.2
)

//replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

//replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
