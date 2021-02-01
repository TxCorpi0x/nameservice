package nameservicetest

import (
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/vjdmhd/nameservice/x/nameservice/keeper"
	"github.com/vjdmhd/nameservice/x/nameservice/types"
)

var (
	tokenName = "stake"
	ctx       sdk.Context
	handler   sdk.Handler
	k         keeper.Keeper
	buyerAcc  sdk.AccAddress
	buyer2Acc sdk.AccAddress
)

// TestBasicMsgs tests general validation of test message to ensure the message send/receive functionality
func TestBasicMsgs(t *testing.T) {
	ctx, buyerAcc, buyer2Acc, handler, k = CreateTestHandler(t)
	//Unrecognized type
	res, err := handler(ctx, sdk.NewTestMsg())
	require.Error(t, err)
	require.Nil(t, res)
	require.True(t, strings.Contains(err.Error(), "unrecognized nameservice message type: "))
}

// TestMsgBuyName tests all possible validation and logical possibilities MsgBuyName
func TestMsgBuyName(t *testing.T) {

	t.Log("Test TestMsgBuyName for invalid data prevention")
	//*** Validate Basic Test ***\\
	// Test if the empty name impossible
	bid := sdk.Coins{sdk.NewInt64Coin(tokenName, 1)}
	msg := types.NewMsgBuyName("", bid, buyerAcc)
	err := msg.ValidateBasic()
	require.Error(t, err)

	// Test if the empty buyer be impossible
	bid = sdk.Coins{sdk.NewInt64Coin(tokenName, 1)}
	msg = types.NewMsgBuyName("Mehdi", bid, nil)
	err = msg.ValidateBasic()
	require.Error(t, err)

	// Test if nil bid be impossible
	msg = types.NewMsgBuyName("Mehdi", nil, buyer2Acc)
	err = msg.ValidateBasic()
	require.Error(t, err)

	// Test if insufficient funds be impossible
	msg = types.NewMsgBuyName("Mehdi", sdk.Coins{}, buyer2Acc)
	err = msg.ValidateBasic()
	require.Error(t, err)

	t.Log("Test TestMsgBuyName for invalid data preventio passed!")

	t.Log("Test TestMsgBuyName for logical data")
	//*** Handler logics Test ***\\
	// Test if the lower bid be impossible. The initial price of a name is 1 so the zero should raise error.
	bid = sdk.Coins{sdk.NewInt64Coin(tokenName, 0)}
	msg = types.NewMsgBuyName("mehdi", bid, buyerAcc)
	err = msg.ValidateBasic()
	require.Error(t, err)

	// Test if the correct info be possible
	bid = sdk.Coins{sdk.NewInt64Coin(tokenName, 2)}
	msg = types.NewMsgBuyName("mehdi", bid, buyerAcc)
	res, err := handler(ctx, msg)
	require.NoError(t, err)
	require.NotNil(t, res)

	// Test if the name belongs to new owner
	nameOnwer := k.GetOwner(ctx, "mehdi")
	require.Equal(t, nameOnwer, buyerAcc)
	require.NotEqual(t, nameOnwer, buyer2Acc)

	// Test if Price is set according to bid
	namePrice := k.GetPrice(ctx, "mehdi")
	require.True(t, namePrice.IsEqual(bid))

	t.Log("Test TestMsgBuyName for logical data passed!")

}

// TestMsgSetName tests all possible validation and logical possibilities MsgSetName
func TestMsgSetName(t *testing.T) {

	t.Log("Test TestMsgSetName for invalid data prevention")
	//*** Validate Basic Test ***\\
	// Test if the empty buyer be impossible
	t.Log("Test if the empty buyer be impossible")
	msg := types.NewMsgSetName(nil, "mehdi", "mehdiname")
	err := msg.ValidateBasic()
	require.Error(t, err)

	// Test if the empty name be impossible
	t.Log("Test if the empty name be impossible")
	msg = types.NewMsgSetName(buyerAcc, "", "mehdiname")
	err = msg.ValidateBasic()
	require.Error(t, err)

	// Test if the empty value be impossible
	t.Log("Test if the empty value be impossible")
	msg = types.NewMsgSetName(buyerAcc, "mehdi", "")
	err = msg.ValidateBasic()
	require.Error(t, err)

	t.Log("Test TestMsgSetName for invalid data prevention passed!")

	t.Log("Test TestMsgSetName for logical data")
	//*** Handler logics Test ***\\
	// Test if the wrong owner of name be impossible
	t.Log("Test if the wrong owner of name be impossible")
	msg = types.NewMsgSetName(buyerAcc, "mehdiplus", "mehdiname")
	err = msg.ValidateBasic()
	require.NoError(t, err)

	_, err = handler(ctx, msg)
	require.Error(t, err)

	// Test if correct parameters be possible to set
	t.Log("Test if correct parameters be possible to set")
	msg = types.NewMsgSetName(buyerAcc, "mehdi", "mehdiname")
	err = msg.ValidateBasic()
	require.NoError(t, err)

	_, err = handler(ctx, msg)
	require.NoError(t, err)

	// Test if the value of the name is set corrcetly
	t.Log("Test if the value of the name is set corrcetly")
	nameToResolve := k.ResolveName(ctx, "mehdi")
	require.Equal(t, nameToResolve, "mehdiname")

	t.Log("Test TestMsgSetName for logical data passed")
}

// TestMsgSetName tests all possible validation and logical possibilities of the MsgDeleteName
func TestMsgDeleteName(t *testing.T) {

	t.Log("Test MsgDeleteName for invalid data prevention")

	//*** Validate Basic Test ***\\
	// Test if the empty buyer be impossible
	t.Log("Test if the empty buyer be impossible")
	msg := types.NewMsgDeleteName("mehdi", nil)
	err := msg.ValidateBasic()
	require.Error(t, err)

	t.Log("Test MsgDeleteName for invalid data prevention passed!")

	t.Log("Test MsgDeleteName for logical data")
	//*** Handler logics Test ***\\
	// Test if unavailable name can not be deleted.
	t.Log("Test if unavailable name can not be deleted.")
	msg = types.NewMsgDeleteName("mehdiplus", buyerAcc)
	err = msg.ValidateBasic()
	require.NoError(t, err)

	_, err = handler(ctx, msg)
	require.Error(t, err)

	// Test if deleting the be impossible by incorrect owner
	t.Log("Test if deleting the be impossible by incorrect owner")
	msg = types.NewMsgDeleteName("mehdi", buyer2Acc)
	err = msg.ValidateBasic()
	require.NoError(t, err)

	_, err = handler(ctx, msg)
	require.Error(t, err)

	// Test if corrcet parameters works
	t.Log("Test if corrcet parameters works")
	msg = types.NewMsgDeleteName("mehdi", buyerAcc)
	err = msg.ValidateBasic()
	require.NoError(t, err)

	_, err = handler(ctx, msg)
	require.NoError(t, err)

	// Test if the name is deleted successfully
	t.Log("Test if the name is deleted successfully")
	require.False(t, k.Exists(ctx, "mehdi"))

	t.Log("Test MsgDeleteName for logical data passed!")
}
