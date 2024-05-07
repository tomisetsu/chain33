// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/33cn/chain33/common/crypto"

	"github.com/33cn/chain33/common"
	"github.com/33cn/chain33/util"
	"github.com/stretchr/testify/require"

	"github.com/33cn/chain33/account"
	"github.com/33cn/chain33/client/mocks"
	"github.com/33cn/chain33/common/address"
	slog "github.com/33cn/chain33/common/log"
	"github.com/33cn/chain33/pluginmgr"
	qmock "github.com/33cn/chain33/queue/mocks"
	cty "github.com/33cn/chain33/system/dapp/coins/types"
	"github.com/33cn/chain33/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Init(cfg *types.Chain33Config) {
	slog.SetLogLevel("error")
	pluginmgr.InitExec(cfg)
}

func newTestChannelClient() *ChannelClient {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := &mocks.QueueProtocolAPI{}
	api.On("GetConfig", mock.Anything).Return(cfg)
	return &ChannelClient{
		QueueProtocolAPI: api,
	}
}

// TODO
func TestInit(t *testing.T) {
	client := newTestChannelClient()
	qm := &qmock.Client{}
	cfg := client.GetConfig()
	qm.On("GetConfig", mock.Anything).Return(cfg)
	client.Init(qm, nil)
}

func testCreateRawTransactionNil(t *testing.T) {
	client := newTestChannelClient()
	_, err := client.CreateRawTransaction(nil)
	assert.Equal(t, types.ErrInvalidParam, err)
}

func testCreateRawTransactionExecNameErr(t *testing.T) {
	tx := types.CreateTx{ExecName: "aaa", To: "1MY4pMgjpS2vWiaSDZasRhN47pcwEire32"}

	client := newTestChannelClient()
	cfg := client.GetConfig()
	Init(cfg)
	_, err := client.CreateRawTransaction(&tx)
	assert.Equal(t, types.ErrExecNameNotMatch, err)
}

func testCreateRawTransactionAmoutErr(t *testing.T) {
	client := newTestChannelClient()
	cfg := client.GetConfig()
	tx := types.CreateTx{ExecName: cfg.ExecName(cty.CoinsX), Amount: -1, To: "1MY4pMgjpS2vWiaSDZasRhN47pcwEire32"}

	_, err := client.CreateRawTransaction(&tx)
	assert.Equal(t, types.ErrAmount, err)
}

func testCreateRawTransactionTo(t *testing.T) {
	client := newTestChannelClient()
	cfg := client.GetConfig()
	name := cfg.ExecName(cty.CoinsX)
	tx := types.CreateTx{ExecName: name, Amount: 1, To: "1MY4pMgjpS2vWiaSDZasRhN47pcwEire32", Fee: 1}

	rawtx, err := client.CreateRawTransaction(&tx)
	assert.NoError(t, err)

	reqDecode := &types.ReqDecodeRawTransaction{TxHex: hex.EncodeToString(rawtx)}
	_, err = client.DecodeRawTransaction(reqDecode)
	assert.NoError(t, err)

	assert.Nil(t, err)
	var mytx types.Transaction
	err = types.Decode(rawtx, &mytx)
	assert.Nil(t, err)
	if cfg.IsPara() {
		assert.Equal(t, address.ExecAddress(name), mytx.To)
	} else {
		assert.Equal(t, tx.To, mytx.To)
	}
}

func testCreateRawTransactionCoinTransfer(t *testing.T) {
	ctx := types.CreateTx{
		ExecName:   "",
		Amount:     10,
		IsToken:    false,
		IsWithdraw: false,
		To:         "1JkbMq5yNMZHtokjg5XxkC3RZbqjoPJm84",
		Note:       []byte("note"),
		Fee:        1,
	}

	client := newTestChannelClient()
	cfg := client.GetConfig()
	txHex, err := client.CreateRawTransaction(&ctx)
	assert.Nil(t, err)
	var tx types.Transaction
	types.Decode(txHex, &tx)
	assert.Equal(t, []byte(cfg.ExecName(cty.CoinsX)), tx.Execer)

	var transfer cty.CoinsAction
	types.Decode(tx.Payload, &transfer)
	assert.Equal(t, int32(cty.CoinsActionTransfer), transfer.Ty)
}

func testCreateRawTransactionCoinTransferExec(t *testing.T) {
	client := newTestChannelClient()
	cfg := client.GetConfig()
	name := cfg.ExecName("coins")
	ctx := types.CreateTx{
		ExecName:   name,
		Amount:     10,
		IsToken:    false,
		IsWithdraw: false,
		To:         "1JkbMq5yNMZHtokjg5XxkC3RZbqjoPJm84",
		Note:       []byte("note"),
		Fee:        1,
	}
	txHex, err := client.CreateRawTransaction(&ctx)
	assert.Nil(t, err)
	var tx types.Transaction
	types.Decode(txHex, &tx)
	assert.Equal(t, []byte(cfg.ExecName(cty.CoinsX)), tx.Execer)

	var transfer cty.CoinsAction
	types.Decode(tx.Payload, &transfer)
	assert.Equal(t, int32(cty.CoinsActionTransferToExec), transfer.Ty)
	if cfg.IsPara() {
		assert.Equal(t, address.ExecAddress(cfg.ExecName(cty.CoinsX)), tx.To)
	} else {
		assert.Equal(t, ctx.To, tx.To)
	}
}

func testCreateRawTransactionCoinWithdraw(t *testing.T) {
	client := newTestChannelClient()
	cfg := client.GetConfig()
	ctx := types.CreateTx{
		ExecName:   cfg.ExecName("coins"),
		Amount:     10,
		IsToken:    false,
		IsWithdraw: true,
		To:         "1JkbMq5yNMZHtokjg5XxkC3RZbqjoPJm84",
		Note:       []byte("note"),
		Fee:        1,
	}

	txHex, err := client.CreateRawTransaction(&ctx)
	assert.Nil(t, err)
	var tx types.Transaction
	types.Decode(txHex, &tx)
	assert.Equal(t, []byte(cfg.ExecName(cty.CoinsX)), tx.Execer)

	var transfer cty.CoinsAction
	types.Decode(tx.Payload, &transfer)
	assert.Equal(t, int32(cty.CoinsActionWithdraw), transfer.Ty)

	if cfg.IsPara() {
		assert.Equal(t, address.ExecAddress(cfg.ExecName(cty.CoinsX)), tx.To)
	} else {
		assert.Equal(t, ctx.To, tx.To)
	}
}

func TestChannelClient_CreateRawTransaction(t *testing.T) {
	testCreateRawTransactionNil(t)
	testCreateRawTransactionExecNameErr(t)
	testCreateRawTransactionAmoutErr(t)
	testCreateRawTransactionTo(t)
	testCreateRawTransactionCoinTransfer(t)
	testCreateRawTransactionCoinTransferExec(t)
	testCreateRawTransactionCoinWithdraw(t)
}

func testChannelClientGetAddrOverviewNil(t *testing.T) {
	parm := &types.ReqAddr{
		Addr: "abcde",
	}
	client := newTestChannelClient()
	_, err := client.GetAddrOverview(parm)
	assert.Equal(t, types.ErrInvalidAddress, err)
}

func testChannelClientGetAddrOverviewErr(t *testing.T) {
	parm := &types.ReqAddr{
		Addr: "1Jn2qu84Z1SUUosWjySggBS9pKWdAP3tZt",
	}
	api := new(mocks.QueueProtocolAPI)
	client := &ChannelClient{
		QueueProtocolAPI: api,
	}

	api.On("GetAddrOverview", mock.Anything).Return(nil, fmt.Errorf("error"))
	_, err := client.GetAddrOverview(parm)
	assert.EqualError(t, err, "error")

}

func testChannelClientGetAddrOverviewOK(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	db := new(account.DB)
	client := &ChannelClient{
		QueueProtocolAPI: api,
		accountdb:        db,
	}

	addr := &types.AddrOverview{}
	api.On("GetAddrOverview", mock.Anything).Return(addr, nil)

	head := &types.Header{StateHash: []byte("sdfadasds")}
	api.On("GetLastHeader").Return(head, nil)

	var acc = &types.Account{Addr: "1Jn2qu84Z1SUUosWjySggBS9pKWdAP3tZt", Balance: 100}
	accv := types.Encode(acc)
	storevalue := &types.StoreReplyValue{}
	storevalue.Values = append(storevalue.Values, accv)
	api.On("StoreGet", mock.Anything).Return(storevalue, nil)

	//reply := types.AddrOverview{}
	parm := &types.ReqAddr{
		Addr: "1Jn2qu84Z1SUUosWjySggBS9pKWdAP3tZt",
	}
	data, err := client.GetAddrOverview(parm)
	assert.Nil(t, err, "error")
	assert.Equal(t, acc.Balance, data.Balance)

}

func TestChannelClient_GetAddrOverview(t *testing.T) {
	testChannelClientGetAddrOverviewNil(t)
	testChannelClientGetAddrOverviewOK(t)
	testChannelClientGetAddrOverviewErr(t)
}

func testChannelClientGetBalanceCoin(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := &mocks.QueueProtocolAPI{}
	api.On("GetConfig", mock.Anything).Return(cfg)
	db := new(account.DB)
	client := &ChannelClient{
		QueueProtocolAPI: api,
		accountdb:        db,
	}

	//addr := &types.AddrOverview{}
	//api.On("GetAddrOverview", mock.Anything).Return(addr, nil)

	head := &types.Header{StateHash: []byte("sdfadasds")}
	api.On("GetLastHeader").Return(head, nil)

	var acc = &types.Account{Addr: "1Jn2qu84Z1SUUosWjySggBS9pKWdAP3tZt", Balance: 100}
	accv := types.Encode(acc)
	storevalue := &types.StoreReplyValue{}
	storevalue.Values = append(storevalue.Values, accv)
	api.On("StoreGet", mock.Anything).Return(storevalue, nil)

	var addrs = make([]string, 1)
	addrs = append(addrs, "1Jn2qu84Z1SUUosWjySggBS9pKWdAP3tZt")
	var in = &types.ReqBalance{
		Execer:    "coins",
		Addresses: addrs,
	}
	data, err := client.GetBalance(in)
	assert.Nil(t, err)
	assert.Equal(t, acc.Addr, data[0].Addr)

}

func testChannelClientGetBalanceOther(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := &mocks.QueueProtocolAPI{}
	api.On("GetConfig", mock.Anything).Return(cfg)
	db := new(account.DB)
	client := &ChannelClient{
		QueueProtocolAPI: api,
		accountdb:        db,
	}

	//addr := &types.AddrOverview{}
	//api.On("GetAddrOverview", mock.Anything).Return(addr, nil)

	head := &types.Header{StateHash: []byte("sdfadasds")}
	api.On("GetLastHeader").Return(head, nil)

	var acc = &types.Account{Addr: "1Jn2qu84Z1SUUosWjySggBS9pKWdAP3tZt", Balance: 100}
	accv := types.Encode(acc)
	storevalue := &types.StoreReplyValue{}
	storevalue.Values = append(storevalue.Values, accv)
	api.On("StoreGet", mock.Anything).Return(storevalue, nil)

	var addrs = make([]string, 1)
	addrs = append(addrs, "1Jn2qu84Z1SUUosWjySggBS9pKWdAP3tZt")
	var in = &types.ReqBalance{
		Execer:    cfg.ExecName("ticket"),
		Addresses: addrs,
	}
	data, err := client.GetBalance(in)
	assert.Nil(t, err)
	assert.Equal(t, acc.Addr, data[0].Addr)

}

func TestChannelClient_GetBalance(t *testing.T) {
	testChannelClientGetBalanceCoin(t)
	testChannelClientGetBalanceOther(t)
}

func TestChannelClient_GetTotalCoins(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	client := new(ChannelClient)
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	qm := &qmock.Client{}
	qm.On("GetConfig", mock.Anything).Return(cfg)
	client.Init(qm, api)
	api.On("StoreGetTotalCoins", mock.Anything).Return(&types.ReplyGetTotalCoins{}, nil)
	_, err := client.GetTotalCoins(&types.ReqGetTotalCoins{})
	assert.NoError(t, err)

	// accountdb =
	//token := &types.ReqGetTotalCoins{
	//	Symbol:    "CNY",
	//	StateHash: []byte("1234"),
	//	StartKey:  []byte("sad"),
	//	Count:     1,
	//	Execer:    "coin",
	//}
	//data, err = client.GetTotalCoins(token)
	//assert.NotNil(t, data)
	//assert.Nil(t, err)
}

func TestChannelClient_CreateNoBalanceTransaction(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	client := new(ChannelClient)
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	qm := &qmock.Client{}
	qm.On("GetConfig", mock.Anything).Return(cfg)
	client.Init(qm, api)
	fee := cfg.GetMinTxFeeRate() * 2
	api.On("GetProperFee", mock.Anything).Return(&types.ReplyProperFee{ProperFee: fee}, nil)
	in := &types.NoBalanceTx{}
	params := &types.NoBalanceTxs{
		TxHexs:  []string{in.GetTxHex()},
		PayAddr: in.GetPayAddr(),
		Privkey: in.GetPrivkey(),
		Expire:  in.GetExpire(),
	}
	tx, err := client.CreateNoBalanceTxs(params)
	assert.NoError(t, err)
	gtx, _ := tx.GetTxGroup()
	assert.NoError(t, gtx.Check(cfg, 0, fee, cfg.GetMaxTxFee(0)))
	assert.NoError(t, err)
	params.Expire = "300s"
	tx, err = client.CreateNoBalanceTxs(params)
	assert.NoError(t, err)
	assert.Equal(t, true, (tx.GetExpire() > types.Now().Unix()) && (tx.GetExpire() <= types.Now().Unix()+300))
	params.TxHexs[0] = hex.EncodeToString(types.Encode(&types.Transaction{Execer: []byte("user.p.para.coins")}))
	params.Expire = "100"
	tx, err = client.CreateNoBalanceTxs(params)
	assert.Equal(t, types.ErrInvalidExpire, err)
	params.Expire = "0"
	_, err = client.CreateNoBalanceTxs(params)
	assert.NotEqual(t, types.ErrInvalidExpire, err)
	params.Expire = fmt.Sprintf("%d", types.ExpireBound+1)
	_, err = client.CreateNoBalanceTxs(params)
	assert.NotEqual(t, types.ErrInvalidExpire, err)
}

func TestClientReWriteRawTx(t *testing.T) {
	//交易组原始交易的修改测试
	txHex1 := "0a0a757365722e7772697465121d236d642368616b6468676f7177656a6872676f716a676f6a71776c6a6720c0843d30aab4d59684b5cce7143a2231444e615344524739524431397335396d65416f654e34613246365248393766536f400a4ab50c0aa3010a0a757365722e7772697465121d236d642368616b6468676f7177656a6872676f716a676f6a71776c6a6720c0843d30aab4d59684b5cce7143a2231444e615344524739524431397335396d65416f654e34613246365248393766536f400a4a201f533ac07c3fc4c716f65cdb0f1f02e7f5371b5164277210dafb1dbdd4a5f4f5522008217c413b035fddd8f34a303e90a29e661746ed9b23a97768c1f25817c2c3450a9f010a0a757365722e7772697465121d236d642368616b6468676f7177656a6872676f716a676f6a71776c6a673094fbcabe96c99ea7163a2231444e615344524739524431397335396d65416f654e34613246365248393766536f400a4a201f533ac07c3fc4c716f65cdb0f1f02e7f5371b5164277210dafb1dbdd4a5f4f552203c6a2b11cce466891f084b49450472b1d4c39213f63117d3d4ce2a3851304ebc0a9f010a0a757365722e7772697465121d236d642368616b6468676f7177656a6872676f716a676f6a71776c6a6730c187fb80fe88ce9e3c3a2231444e615344524739524431397335396d65416f654e34613246365248393766536f400a4a201f533ac07c3fc4c716f65cdb0f1f02e7f5371b5164277210dafb1dbdd4a5f4f5522066419d70492f757d7285fd226dff62da8d803c8121ded95242d222dbb10f2d9b0a9f010a0a757365722e7772697465121d236d642368616b6468676f7177656a6872676f716a676f6a71776c6a673098aa929ab292b3f0023a2231444e615344524739524431397335396d65416f654e34613246365248393766536f400a4a201f533ac07c3fc4c716f65cdb0f1f02e7f5371b5164277210dafb1dbdd4a5f4f552202bab08051d24fe923f66c8aeea4ce3f425d47a72f7c5c230a2b1427e04e2eb510a9f010a0a757365722e7772697465121d236d642368616b6468676f7177656a6872676f716a676f6a71776c6a6730bfe9abb3edc6d9cb163a2231444e615344524739524431397335396d65416f654e34613246365248393766536f400a4a201f533ac07c3fc4c716f65cdb0f1f02e7f5371b5164277210dafb1dbdd4a5f4f55220e1ba0493aa431ea3071026bd8dfa8280efab53ce86441fc474a1c19550a554ba0a9f010a0a757365722e7772697465121d236d642368616b6468676f7177656a6872676f716a676f6a71776c6a6730d2e196a8ecada9d53e3a2231444e615344524739524431397335396d65416f654e34613246365248393766536f400a4a201f533ac07c3fc4c716f65cdb0f1f02e7f5371b5164277210dafb1dbdd4a5f4f5522016600fbfa23b3f0e8f9a14b716ce8f4064c091fbf6fa94489bc9d14b5b6049a60a9f010a0a757365722e7772697465121d236d642368616b6468676f7177656a6872676f716a676f6a71776c6a6730a0b7b1b1dda2f4c5743a2231444e615344524739524431397335396d65416f654e34613246365248393766536f400a4a201f533ac07c3fc4c716f65cdb0f1f02e7f5371b5164277210dafb1dbdd4a5f4f5522089d0442d76713369022499d054db65ccacbf5c627a525bd5454e0a30d23fa2990a9f010a0a757365722e7772697465121d236d642368616b6468676f7177656a6872676f716a676f6a71776c6a6730c5838f94e2f49acb4b3a2231444e615344524739524431397335396d65416f654e34613246365248393766536f400a4a201f533ac07c3fc4c716f65cdb0f1f02e7f5371b5164277210dafb1dbdd4a5f4f5522018f208938606b390d752898332a84a9fbb900c2ed55ec33cd54d09b1970043b90a9f010a0a757365722e7772697465121d236d642368616b6468676f7177656a6872676f716a676f6a71776c6a67308dfddb82faf7dfc4113a2231444e615344524739524431397335396d65416f654e34613246365248393766536f400a4a201f533ac07c3fc4c716f65cdb0f1f02e7f5371b5164277210dafb1dbdd4a5f4f5522013002bab7a9c65881bd937a6fded4c3959bb631fa84434572970c1ec3e6fccf90a7d0a0a757365722e7772697465121d236d642368616b6468676f7177656a6872676f716a676f6a71776c6a6730b8b082d799a4ddc93a3a2231444e615344524739524431397335396d65416f654e34613246365248393766536f400a4a201f533ac07c3fc4c716f65cdb0f1f02e7f5371b5164277210dafb1dbdd4a5f4f5522008217c413b035fddd8f34a303e90a29e661746ed9b23a97768c1f25817c2c345"
	//修改交易组的所有交易
	ctx := types.ReWriteRawTx{
		Tx:     txHex1,
		Fee:    29977777777,
		Expire: "130s",
		To:     "14KEKbYtKKQm4wMthSK9J4La4nAiidGozt",
		Index:  0,
	}

	client := newTestChannelClient()

	txHex, err := client.ReWriteRawTx(&ctx)
	assert.Nil(t, err)
	rtTx := hex.EncodeToString(txHex)
	txData, err := hex.DecodeString(rtTx)
	assert.Nil(t, err)
	tx := &types.Transaction{}
	err = types.Decode(txData, tx)
	assert.Nil(t, err)
	assert.Equal(t, ctx.Fee, tx.Fee)

	//只修改交易组中指定的交易
	ctx2 := types.ReWriteRawTx{
		Tx:     txHex1,
		Fee:    29977777777,
		Expire: "130s",
		To:     "14KEKbYtKKQm4wMthSK9J4La4nAiidGozt",
		Index:  2,
	}
	txHex22, err := client.ReWriteRawTx(&ctx2)
	assert.Nil(t, err)
	rtTx22 := hex.EncodeToString(txHex22)
	txData22, err := hex.DecodeString(rtTx22)
	assert.Nil(t, err)
	tx22 := &types.Transaction{}
	err = types.Decode(txData22, tx22)
	assert.Nil(t, err)
	group22, err := tx22.GetTxGroup()
	assert.Nil(t, err)

	for index, tmptx := range group22.GetTxs() {
		if tmptx.GetExpire() != 0 && index != 1 {
			t.Error("TestClientReWriteRawTx Expire !=0 index != 1 ")
		}
		if tmptx.GetFee() != 0 && index != 0 {
			t.Error("TestClientReWriteRawTx Fee !=0")
		}
	}
}

func TestChannelClient_GetWalletRecoverScript(t *testing.T) {
	cli := &ChannelClient{}
	req := &types.ReqGetWalletRecoverAddr{}
	_, err := cli.GetWalletRecoverAddr(req)
	require.Equal(t, types.ErrInvalidParam, err)

	req.RelativeDelayTime = 10

	addr1, priv1 := util.Genaddress()
	_, priv2 := util.Genaddress()
	req.CtrPubKey = common.ToHex(priv1.PubKey().Bytes())
	req.RecoverPubKeys = []string{common.ToHex(priv2.PubKey().Bytes())}
	addr, err := cli.GetWalletRecoverAddr(req)
	require.Nil(t, err)

	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	_, priv3 := util.Genaddress()
	tx := util.CreateNoneTx(cfg, priv3)
	privKeyHex := hex.EncodeToString(priv1.Bytes())
	mockAPI := new(mocks.QueueProtocolAPI)
	cli.QueueProtocolAPI = mockAPI
	mockAPI.On("ExecWalletFunc", "wallet", "DumpPrivkey",
		&types.ReqString{Data: addr1}).Return(&types.ReplyString{Data: privKeyHex}, nil)
	req2 := &types.ReqSignWalletRecoverTx{
		RawTx:              hex.EncodeToString(types.Encode(tx)),
		WalletRecoverParam: req,
		SignAddr:           addr1,
	}

	reply, err := cli.SignWalletRecoverTx(req2)
	require.Nil(t, err)
	txByte, err := common.FromHex(reply.TxHex)
	require.Nil(t, err)
	err = types.Decode(txByte, tx)
	require.Nil(t, err)
	require.True(t, tx.CheckSign(0))
	require.Equal(t, addr.Data, tx.From())
}

func TestChannelClient_SignWalletRecoverTx(t *testing.T) {

	cli := &ChannelClient{}
	req := &types.ReqGetWalletRecoverAddr{}
	signReq := &types.ReqSignWalletRecoverTx{}

	_, err := cli.SignWalletRecoverTx(signReq)
	require.Equal(t, types.ErrInvalidParam, err)
	signReq.WalletRecoverParam = req
	_, err = cli.SignWalletRecoverTx(signReq)
	require.Equal(t, types.ErrInvalidParam, err)
}

func TestQueueProtocol_GetCryptoList(t *testing.T) {
	q := &ChannelClient{}
	list := q.GetCryptoList()
	for _, driver := range list.Cryptos {
		id := int(driver.TypeID)
		require.Equal(t, crypto.GetType(driver.Name), id)
		require.Equal(t, crypto.GetName(id), driver.Name)
	}
}

func TestQueueProtocol_GetAddressDrivers(t *testing.T) {
	q := &ChannelClient{}
	list := q.GetAddressDrivers()
	for _, driver := range list.Drivers {
		d, err := address.LoadDriver(driver.TypeID, -1)
		require.Nil(t, err)
		require.Equal(t, d.GetName(), driver.Name)
	}
}
