package subscribe

import (
	"context"
	"fmt"
	"strings"

	"time"

	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	aiRequest "github.com/oraichain/orai/x/airequest"

	"github.com/oraichain/orai/x/websocket/types"
)

const (
	QuoteString = `"'`
	JoinString  = `-`
)

// SubmitReport creates a new MsgCreateReport and submits it to the Oraichain to create a new report
func (subscriber *Subscriber) submitReport(msgReport *types.MsgCreateReport) (err error) {

	if err = msgReport.ValidateBasic(); err != nil {
		subscriber.log.Error(":exploding_head: Failed to validate basic with error: %s", err.Error())
		return err
	}

	txf := subscriber.newTxFactory("websocket")
	for try := uint64(1); try <= subscriber.config.MaxTry; try++ {
		subscriber.log.Info(":e-mail: Try to broadcast report transaction(%d/%d)", try, subscriber.config.MaxTry)
		err = tx.BroadcastTx(*subscriber.cliCtx, txf, msgReport)
		if err == nil {
			break
		}
		subscriber.log.Error(":warning: Failed to broadcast tx with error: %s", err.Error())
		time.Sleep(subscriber.config.RPCPollInterval)
	}

	// the last error
	return err
}

func (subscriber *Subscriber) handleAIRequestLog(queryClient types.QueryClient, ev *sdk.StringEvent) (*types.MsgCreateReport, error) {
	subscriber.log.Info(":delivery_truck: Processing incoming request event before checking validators")

	attrMap := GetAttributeMap(ev.GetAttributes())

	// Skip if not related to this validator
	validators := attrMap[aiRequest.AttributeRequestValidator]
	subscriber.log.Info(":delivery_truck: Validator lists: ", validators)
	currentValidator := sdk.ValAddress(subscriber.cliCtx.GetFromAddress()).String()
	hasMe := false
	for _, validator := range validators {
		subscriber.log.Debug(":delivery_truck: validator: %s", validator)
		if validator == currentValidator {
			hasMe = true
			break
		}
	}

	if !hasMe {
		subscriber.log.Info(":next_track_button: Skip request not related to this validator")
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Skip request not related to this validator")
	}

	subscriber.log.Info(":delivery_truck: Processing incoming request event")

	var requestID, oscriptName, inputStr, expectedOutputStr string

	if val, ok := attrMap[aiRequest.AttributeRequestID]; ok {
		requestID = val[0]
	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, aiRequest.AttributeRequestID)
	}
	if val, ok := attrMap[aiRequest.AttributeOracleScriptName]; ok {
		oscriptName = val[0]
	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, aiRequest.AttributeOracleScriptName)
	}
	if val, ok := attrMap[aiRequest.AttributeRequestInput]; ok {
		inputStr = val[0]
	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, aiRequest.AttributeRequestInput)
	}
	if val, ok := attrMap[aiRequest.AttributeRequestExpectedOutput]; ok {
		expectedOutputStr = val[0]
	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, aiRequest.AttributeRequestExpectedOutput)
	}

	// collect ai data sources and test cases from the ai request event.
	aiDataSources := attrMap[aiRequest.AttributeRequestDSources]
	testCases := attrMap[aiRequest.AttributeRequestTCases]

	// create data source results to store in the report
	var dataSourceResultsTest []*types.DataSourceResult
	var dataSourceResults []*types.DataSourceResult
	var testCaseResults []*types.TestCaseResult
	var resultArr = []string{}

	// we have different test cases, so we need to loop through them
	ctx := context.Background()
	for _, testCase := range testCases {
		//put the results from the data sources into the test case to verify if they are good enough
		for _, aiDataSource := range aiDataSources {
			// collect test case result from the script
			outTestCase, err := queryClient.TestCaseContract(ctx, &types.QueryTestCaseContract{
				Name:           testCase,
				DataSourceName: aiDataSource,
				Request: &types.RequestTestCase{
					Input:  inputStr,
					Output: expectedOutputStr,
				},
			})

			if err != nil {
				subscriber.log.Info(":skull: failed to execute test case 1st loop: %s", err.Error())
				return nil, err
			}
			result := string(outTestCase.Data)

			subscriber.log.Info("result after running test case: ", result)

			dataSourceResult := types.NewDataSourceResult(aiDataSource, outTestCase.Data, types.ResultSuccess)

			// By default, we consider returning null as failure. If any datasource does not follow this rule then it should not be used by any oracle scripts.
			if len(outTestCase.Data) == 0 || result == "null" {
				// change status to fail so the datasource cannot be rewarded afterwards
				dataSourceResult.Status = types.ResultFailure
				dataSourceResult.Result = []byte(types.FailedResponseTc)
			}
			// append an data source result into the list
			dataSourceResultsTest = append(dataSourceResultsTest, dataSourceResult)
		}

		// add test case result
		testCaseResults = append(testCaseResults, types.NewTestCaseResult(testCase, dataSourceResultsTest))
	}

	// after passing the test cases, we run the actual data sources to collect their results
	// create data source results to store in the report
	// we use dataSourceResultsTest since this list is the complete list of data sources that have passed the test cases
	for _, dataSourceResultTest := range dataSourceResultsTest {
		// run the data source script

		var dataSourceResult *types.DataSourceResult
		if dataSourceResultTest.GetStatus() == types.ResultSuccess {

			outDataSource, err := queryClient.DataSourceContract(ctx, &types.QueryDataSourceContract{
				Name: dataSourceResultTest.GetName(),
				Request: &types.RequestDataSource{
					Input: inputStr,
				},
			})
			if err != nil {
				subscriber.log.Error(":skull: failed to execute data source script: %s", err.Error())
				return nil, err
			}
			// remove all quote at start and begine
			result := strings.Trim(string(outDataSource.Data), QuoteString)

			subscriber.log.Info("star: result from data sources: ", result)
			// By default, we consider returning null as failure. If any datasource does not follow this rule then it should not be used by any oracle scripts.
			dataSourceResult = types.NewDataSourceResult(dataSourceResultTest.GetName(), []byte(result), types.ResultSuccess)
			if len(result) == 0 {
				// change status to fail so the datasource cannot be rewarded afterwards
				dataSourceResult.Status = types.ResultFailure
				dataSourceResult.Result = []byte(types.FailedResponseDs)
			} else {
				resultArr = append(resultArr, result)
			}
		} else {
			dataSourceResult = types.NewDataSourceResult(dataSourceResultTest.GetName(), []byte(dataSourceResultTest.GetResult()), types.ResultFailure)
		}
		// append an data source result into the list
		dataSourceResults = append(dataSourceResults, dataSourceResult)
	}
	subscriber.log.Info("star: final result: ", resultArr)
	// Create a new MsgCreateReport with a new reporter to the Oraichain
	reporter := types.NewReporter(
		subscriber.cliCtx.GetFromAddress(), subscriber.cliCtx.GetFromName(),
		sdk.ValAddress(subscriber.cliCtx.GetFromAddress()),
	)
	finalResult := []byte(strings.Join(resultArr, JoinString))
	msgReport := types.NewMsgCreateReport(
		requestID, dataSourceResults,
		testCaseResults, reporter,
		subscriber.config.Fees, finalResult,
		types.ResultSuccess,
	)

	if len(resultArr) == 0 {
		msgReport.AggregatedResult = []byte(types.FailedResponseOs)
		msgReport.ResultStatus = types.ResultFailure
	} else {

		query := &types.QueryOracleScriptContract{
			Name: oscriptName,
			Request: &types.RequestOracleScript{
				Results: resultArr,
			},
		}
		fmt.Printf("query :%v\n", query)
		outOScript, err := queryClient.OracleScriptContract(ctx, query)

		if err != nil {
			subscriber.log.Error(":skull: failed to aggregate results: %s", err.Error())
			return nil, err
		}
		subscriber.log.Info("final result from oScript: %s", string(outOScript.Data))
		msgReport.AggregatedResult = outOScript.Data
	}

	// TODO: check aggregated result value of the oracle script to verify if it fails or success
	return msgReport, nil

}
