package aioracle

import (
	"github.com/oraichain/orai/x/aioracle/keeper"
	"github.com/oraichain/orai/x/aioracle/types"
)

const (
	ModuleName               = types.ModuleName
	RouterKey                = types.RouterKey
	StoreKey                 = types.StoreKey
	QuerierRoute             = types.QuerierRoute
	DefaultParamspace        = types.DefaultParamspace
	EventTypeSetAIOracle     = types.EventTypeSetAIOracle
	EventTypeRequestWithData = types.EventTypeRequestWithData
	AttributeRequest         = types.AttributeRequest
	Denom                    = types.Denom
)

var (
	NewKeeper            = keeper.NewKeeper
	NewQuerier           = keeper.NewQuerier
	NewMsgSetAIOracleReq = types.NewMsgSetAIOracleReq
	ModuleCdc            = types.ModuleCdc
	RegisterCodec        = types.RegisterCodec
	NewGenesisState      = types.NewGenesisState
	RequestKeyPrefix     = types.RequestKeyPrefix
	ErrRequestNotFound   = types.ErrRequestNotFound
	NewAIOracle          = types.NewAIOracle
	ResultKeyPrefix      = types.ResultKeyPrefix
	RewardKeyPrefix      = types.RewardKeyPrefix
	ReportKeyPrefix      = types.ReportKeyPrefix
)

type (
	Keeper                  = keeper.Keeper
	MsgSetAIOracleReq       = types.MsgSetAIOracleReq
	MsgCreateReport         = types.MsgCreateReport
	MsgCreateTestCaseReport = types.MsgCreateTestCaseReport
	AIOracle                = types.AIOracle
	GenesisState            = types.GenesisState
	Report                  = types.Report
)
