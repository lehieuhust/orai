package provider

import (
	"github.com/oraichain/orai/x/ai-request/types"
	"github.com/oraichain/orai/x/provider/keeper"
)

const (
	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	QuerierRoute      = types.QuerierRoute
	IPFSUrl           = types.IPFSUrl
	IPFSAdd           = types.IPFSAdd
	IPFSCat           = types.IPFSCat
	DefaultParamspace = types.DefaultParamspace
)

var (
	NewKeeper                = keeper.NewKeeper
	NewQuerier               = keeper.NewQuerier
	NewMsgCreateOracleScript = types.NewMsgCreateOracleScript
	NewMsgCreateAIDataSource = types.NewMsgCreateAIDataSource
	NewMsgSetAIRequest       = types.NewMsgSetAIRequest
	ModuleCdc                = types.ModuleCdc
	RegisterCodec            = types.RegisterCodec
	NewGenesisState          = types.NewGenesisState
	RequestKeyPrefix         = types.RequestKeyPrefix
	ResultKeyPrefix          = types.ResultKeyPrefix
	ReportKeyPrefix          = types.ReportKeyPrefix
	ReporterKeyPrefix        = types.ReporterKeyPrefix
	RewardKeyPrefix          = types.RewardKeyPrefix
	StrategyKeyPrefix        = types.StrategyKeyPrefix
)

type (
	Keeper                = keeper.Keeper
	MsgCreateOracleScript = types.MsgCreateOracleScript
	MsgCreateAIDataSource = types.MsgCreateAIDataSource
	MsgSetAIRequest       = types.MsgSetAIRequest
	AIRequest             = types.AIRequest
	GenesisState          = types.GenesisState
)