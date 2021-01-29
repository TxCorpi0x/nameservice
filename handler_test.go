package nameservicetest

import (
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/vjdmhd/nameservice/x/nameservice/types"
)

var (
	tokenName = "stake"
	ctx       sdk.Context
	handler   sdk.Handler
	buyerAcc  sdk.AccAddress
)

func TestBasicMsgs(t *testing.T) {
	ctx, buyerAcc, handler = CreateTestHandler(t)
	//Unrecognized type
	res, err := handler(ctx, sdk.NewTestMsg())
	require.Error(t, err)
	require.Nil(t, res)
	require.True(t, strings.Contains(err.Error(), "unrecognized nameservice message type: "))
}

func TestMsgBuyName(t *testing.T) {

	bid := sdk.Coins{sdk.NewInt64Coin(tokenName, 0)}
	msg := types.NewMsgBuyName("mehdi", bid, buyerAcc)
	err := msg.ValidateBasic()
	require.Error(t, err)

	bid = sdk.Coins{sdk.NewInt64Coin(tokenName, 2)}
	msg = types.NewMsgBuyName("mehdi", bid, buyerAcc)
	res, err := handler(ctx, msg)
	require.NoError(t, err)
	require.NotNil(t, res)
	t.Log(buyerAcc)

	// for _, event := range res.Events {
	// 	for _, attribute := range event.Attributes {
	// 		value := string(attribute.Value)
	// 		switch key := string(attribute.Key); key {
	// 		case "module":
	// 			require.Equal(t, value, types.ModuleName)
	// 		case "cosmos_receiver":
	// 			require.Equal(t, value, types.TestAddress)
	// 		case "amount":
	// 			require.Equal(t, value, strconv.FormatInt(types.TestCoinsAmount, 10))
	// 		case "symbol":
	// 			require.Equal(t, value, types.TestCoinsSymbol)
	// 		case "token_contract_address":
	// 			require.Equal(t, value, types.TestTokenContractAddress)
	// 		case statusString:
	// 			require.Equal(t, value, oracle.StatusTextToString[oracle.PendingStatusText])
	// 		case "claim_type":
	// 			require.Equal(t, value, types.ClaimTypeToString[types.LockText])
	// 		default:
	// 			require.Fail(t, fmt.Sprintf("unrecognized event %s", key))
	// 		}
	// 	}
	// }

}

func TestMsgSetName(t *testing.T) {

	msg := types.NewMsgSetName(buyerAcc, "mehdiplus", "mehdiname")
	err := msg.ValidateBasic()
	require.NoError(t, err)

	_, err = handler(ctx, msg)
	require.Error(t, err)

	msg = types.NewMsgSetName(buyerAcc, "mehdi", "mehdiname")
	err = msg.ValidateBasic()
	require.NoError(t, err)

	_, err = handler(ctx, msg)
	require.NoError(t, err)
}

func TestMsgDeleteName(t *testing.T) {

	msg := types.NewMsgDeleteName("mehdiplus", buyerAcc)
	err := msg.ValidateBasic()
	require.NoError(t, err)

	_, err = handler(ctx, msg)
	require.Error(t, err)

	msg = types.NewMsgDeleteName("mehdi", buyerAcc)
	err = msg.ValidateBasic()
	require.NoError(t, err)

	_, err = handler(ctx, msg)
	require.NoError(t, err)

}
