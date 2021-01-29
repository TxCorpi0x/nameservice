package nameservice

import (

	//"github.com/cosmos/cosmos-sdk/std"

	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/tendermint/libs/log"

	//abci "github.com/tendermint/tendermint/abci/types"
	//tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	//tmtypes "github.com/tendermint/tendermint/types"

	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/vjdmhd/nameservice/app"
	"github.com/vjdmhd/nameservice/x/nameservice/keeper"
	"github.com/vjdmhd/nameservice/x/nameservice/types"
)

var (
	testStoreKey = "params"
)

func CreateTestHandler() (sdk.Context, sdk.Handler) {
	// nApp := app.NewInitApp(
	// 	logger, db, traceStore, true, invCheckPeriod,
	// 	baseapp.SetPruning(pruningOpts),
	// 	baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
	// 	baseapp.SetHaltHeight(viper.GetUint64(server.FlagHaltHeight)),
	// 	baseapp.SetHaltTime(viper.GetUint64(server.FlagHaltTime)),
	// 	baseapp.SetInterBlockCache(cache),
	// )
	// ctx := nApp.BaseApp.NewContext(false)

	nApp, cdc := Setup(false) //simapp.Setup(false)
	ctx := nApp.BaseApp.NewContext(false, abci.Header{})

	initCoins := sdk.TokensFromConsensusPower(100)
	buyerAcc, _ := sdk.AccAddressFromBech32(types.TestAddress)
	totalSupply := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, initCoins))

	nApp.BankKeeper.AddCoins(ctx, buyerAcc, totalSupply)

	//app.BankKeeper.SetParams(ctx, types.DefaultParams())
	k := keeper.NewKeeper(nApp.BankKeeper, cdc, nApp.GetKey(testStoreKey))
	k.CoinKeeper.SetCoins(ctx, buyerAcc, totalSupply)
	handler := NewHandler(k)
	return ctx, handler
}

// // DefaultConsensusParams defines the default Tendermint consensus params used in
// // SimApp testing.
// var DefaultConsensusParams = &abci.ConsensusParams{
// 	Block: &abci.BlockParams{
// 		MaxBytes: 200000,
// 		MaxGas:   2000000,
// 	},
// 	Evidence: &tmproto.EvidenceParams{
// 		MaxAgeNumBlocks: 302400,
// 		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
// 		MaxBytes:        10000,
// 	},
// 	Validator: &tmproto.ValidatorParams{
// 		PubKeyTypes: []string{
// 			tmtypes.ABCIPubKeyTypeEd25519,
// 		},
// 	},
// }

func Setup(isCheckTx bool) (*app.NewApp, *codec.Codec) {
	nApp, cdc, genesisState := setup(!isCheckTx, 5)
	if !isCheckTx {
		// init chain must be called to stop deliverState from being nil
		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
		if err != nil {
			panic(err)
		}

		// Initialize the chain
		nApp.InitChain(
			abci.RequestInitChain{
				Validators:      []abci.ValidatorUpdate{},
				ConsensusParams: DefaultConsensusParams,
				AppStateBytes:   stateBytes,
			},
		)
	}

	return nApp, cdc
}

func setup(withGenesis bool, invCheckPeriod uint) (*app.NewApp, *codec.Codec, app.GenesisState) {
	db := dbm.NewMemDB()
	encCdc := app.MakeCodec()
	nApp := app.NewInitApp(log.NewNopLogger(), db, nil, true, invCheckPeriod, nil)
	if withGenesis {
		return nApp, encCdc, app.NewDefaultGenesisState()
	}
	return nApp, encCdc, app.GenesisState{}
}

// // MakeTestEncodingConfig creates an EncodingConfig for testing. This function
// // should be used only in tests or when creating a new app instance (NewApp*()).
// // App user shouldn't create new codecs - use the app.AppCodec instead.
// // [DEPRECATED]
// func MakeTestEncodingConfig() EncodingConfig {
// 	encodingConfig := MakeTestEncodingConfig()
// 	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
// 	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
// 	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
// 	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
// 	return encodingConfig
// }

// // EmptyAppOptions is a stub implementing AppOptions
// type EmptyAppOptions struct{}
