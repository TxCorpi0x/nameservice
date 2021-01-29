package nameservicetest

import (
	"encoding/json"
	"testing"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/vjdmhd/nameservice/app"
	"github.com/vjdmhd/nameservice/x/nameservice"

	"github.com/vjdmhd/nameservice/x/nameservice/keeper"
)

var (
	testStoreKey = "params"
)

func CreateTestHandler(t *testing.T) (sdk.Context, sdk.AccAddress, sdk.Handler) {

	nApp, cdc := Setup(false) //simapp.Setup(false)
	ctx := nApp.BaseApp.NewContext(false, abci.Header{Height: nApp.LastBlockHeight()})

	initCoins := sdk.TokensFromConsensusPower(100)

	allAcc := nApp.AccountKeeper.GetAllAccounts(ctx)
	buyerAcc := allAcc[0].GetAddress()
	totalSupply := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, initCoins))

	nApp.BankKeeper.AddCoins(ctx, buyerAcc, totalSupply)

	//app.BankKeeper.SetParams(ctx, types.DefaultParams())
	k := keeper.NewKeeper(nApp.BankKeeper, cdc, nApp.Keys[bam.MainStoreKey])
	k.CoinKeeper.SetCoins(ctx, buyerAcc, totalSupply)
	handler := nameservice.NewHandler(k)
	return ctx, buyerAcc, handler
}

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

var DefaultConsensusParams = &abci.ConsensusParams{
	Block: &abci.BlockParams{
		MaxBytes: 200000,
		MaxGas:   2000000,
	},
	// Evidence: &tmproto.EvidenceParams{
	// 	MaxAgeNumBlocks: 302400,
	// 	MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
	// 	MaxBytes:        10000,
	// },
	// Validator: &tmproto.ValidatorParams{
	// 	PubKeyTypes: []string{
	// 		tmtypes.ABCIPubKeyTypeEd25519,
	// 	},
	// },
}

func setup(withGenesis bool, invCheckPeriod uint) (*app.NewApp, *codec.Codec, app.GenesisState) {
	db := dbm.NewMemDB()
	encCdc := app.MakeCodec()
	nApp := app.NewInitApp(log.NewNopLogger(), db, nil, true, invCheckPeriod, func(*bam.BaseApp) {})
	if withGenesis {
		return nApp, encCdc, app.NewDefaultGenesisState()
	}
	return nApp, encCdc, app.GenesisState{}
}

// EmptyAppOptions is a stub implementing AppOptions
type EmptyAppOptions struct{}

// Get implements AppOptions
func (ao EmptyAppOptions) Get(o string) interface{} {
	return nil
}
