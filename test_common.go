// test_common has functions needed by test functions as a factory
// we should create it in the root of project because of recursive import happening while initializing the newApp instance inside x/namespace

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

// CreateTestHandler is the main Test Handler instance generator
func CreateTestHandler(t *testing.T) (sdk.Context, sdk.AccAddress, sdk.AccAddress, sdk.Handler, keeper.Keeper) {

	nApp, cdc := SetupApp(false)
	ctx := nApp.BaseApp.NewContext(false, abci.Header{Height: nApp.LastBlockHeight()})

	//get auto-generated accounts
	allAcc := nApp.AccountKeeper.GetAllAccounts(ctx)

	// set first as buyer 1
	buyerAcc := allAcc[0].GetAddress()
	// make buyer 1 rich!
	buyerAccCoins := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.TokensFromConsensusPower(100)))

	// set second as buyer 2
	buyer2Acc := allAcc[1].GetAddress()
	// make buyer 2 poor!
	buyer2AccCoins := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.TokensFromConsensusPower(0)))

	// initialize keeper
	k := keeper.NewKeeper(nApp.BankKeeper, cdc, nApp.Keys[bam.MainStoreKey])

	// set account coins
	k.CoinKeeper.SetCoins(ctx, buyerAcc, buyerAccCoins)
	k.CoinKeeper.SetCoins(ctx, buyer2Acc, buyer2AccCoins)

	// initialize handler
	handler := nameservice.NewHandler(k)
	return ctx, buyerAcc, buyer2Acc, handler, k
}

// SetupApp creates application instance that simulate real initiation
func SetupApp(isCheckTx bool) (*app.NewApp, *codec.Codec) {
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
				ConsensusParams: defaultConsensusParams,
				AppStateBytes:   stateBytes,
			},
		)
	}

	return nApp, cdc
}

// copied form simapp
var defaultConsensusParams = &abci.ConsensusParams{
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
	// initialize tendermint DB
	db := dbm.NewMemDB()
	// make app codec
	encCdc := app.MakeCodec()
	// init newApp instance ready for testing
	nApp := app.NewInitApp(log.NewNopLogger(), db, nil, true, invCheckPeriod, func(*bam.BaseApp) {})
	if withGenesis {
		return nApp, encCdc, app.NewDefaultGenesisState()
	}
	return nApp, encCdc, app.GenesisState{}
}
