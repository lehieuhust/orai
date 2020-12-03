package websocket

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	aiRequest "github.com/oraichain/orai/x/airequest/types"
	provider "github.com/oraichain/orai/x/provider/types"
	"github.com/oraichain/orai/x/websocket/exported"
	"github.com/oraichain/orai/x/websocket/types"
)

// GetEventAIRequest returns the event AI request in the given log.
func GetEventAIRequest(log sdk.ABCIMessageLog) (AIRequest, error) {
	ev := aiRequest.EventTypeRequestWithData
	requestID, err := GetEventValue(log, ev, aiRequest.AttributeRequestID)

	req := AIRequest{}
	if err != nil {
		return req, err
	}
	oscriptName, err := GetEventValue(log, ev, aiRequest.AttributeOracleScriptName)
	if err != nil {
		return req, err
	}
	creatorStr, err := GetEventValue(log, ev, aiRequest.AttributeRequestCreator)
	if err != nil {
		return req, err
	}

	valCountStr, err := GetEventValue(log, ev, aiRequest.AttributeRequestValidatorCount)
	if err != nil {
		return req, err
	}

	inputStr, err := GetEventValue(log, ev, aiRequest.AttributeRequestInput)

	expectedOutputStr, err := GetEventValue(log, ev, aiRequest.AttributeRequestExpectedOutput)

	fmt.Println("expected output str: ", expectedOutputStr)

	creator, err := sdk.AccAddressFromBech32(creatorStr)
	if err != nil {
		return req, err
	}

	valCount, _ := strconv.ParseInt(valCountStr, 10, 64)

	req = NewAIRequest(requestID, oscriptName, creator, valCount, inputStr, expectedOutputStr)

	fmt.Println("AI request: ", req)

	return req, nil
}

func handleAIRequestLog(c *Context, l *Logger, log sdk.ABCIMessageLog) {
	l.Info(":delivery_truck: Processing incoming request event before checking validators")

	// Skip if not related to this validator
	validators := GetEventValues(log, aiRequest.EventTypeSetPriceRequest, aiRequest.AttributeRequestValidator)
	hasMe := false
	for _, validator := range validators {
		l.Info(":delivery_truck: validator: %s", validator)
		if validator == c.validator.String() {
			hasMe = true
			break
		}
	}

	if !hasMe {
		l.Debug(":next_track_button: Skip request not related to this validator")
		return
	}

	l.Info(":delivery_truck: Processing incoming request event")

	req, err := GetEventAIRequest(log)
	if err != nil {
		l.Error(":skull: Failed to parse raw requests with error: %s", err.Error())
	}

	key := <-c.keys
	defer func() {
		c.keys <- key
	}()
	go func(l *Logger, req AIRequest) {
		// collect data source name from the oScript script
		oscriptPath := getOScriptPath(req.OracleScriptName)

		// collect ai data sources and test cases from the ai request event.
		aiDataSources, testCases, err := getPaths(log)
		if err != nil {
			l.Error(":skull: Failed to parse ai data sources and test cases with error: %s", err.Error())
		}

		var finalResultStr string
		// create data source results to store in the report
		var dataSourceResultsTest []exported.DataSourceResultI
		var dataSourceResults []exported.DataSourceResultI
		var testCaseResults []exported.TestCaseResultI

		// we have different test cases, so we need to loop through them
		for i := range testCases {
			//put the results from the data sources into the test case to verify if they are good enough
			for j := range aiDataSources {
				// Aggregate the required fees for an AI request
				// run the test case script
				fmt.Println("test case path: ", getTCasePath(testCases[i])+provider.DataSourceStoreKeyString(aiDataSources[j]))
				cmdTestCase := exec.Command("python", getTCasePath(testCases[i]), provider.DataSourceStoreKeyString(aiDataSources[j]), req.Input, req.ExpectedOutput)
				var outTestCase bytes.Buffer
				cmdTestCase.Stdout = &outTestCase
				err = cmdTestCase.Run()
				if err != nil {
					l.Error(":skull: failed to execute test case 1st loop: %s", err.Error())
				}

				// collect test case result from the script
				result := strings.TrimSuffix(outTestCase.String(), "\n")

				fmt.Println("result after running test case: ", result)

				dataSourceResult := types.NewDataSourceResult(aiDataSources[j], []byte(result), types.ResultSuccess)

				// By default, we consider returning null as failure. If any datasource does not follow this rule then it should not be used by any oracle scripts.
				if result == types.FailedResult || len(result) == 0 {
					// change status to fail so the datasource cannot be rewarded afterwards
					dataSourceResult.Status = types.ResultFailure
					dataSourceResult.Result = []byte(types.FailedResponseTc)
				}
				// append an data source result into the list
				dataSourceResultsTest = append(dataSourceResultsTest, dataSourceResult)
			}

			// add test case result
			testCaseResults = append(testCaseResults, types.NewTestCaseResult(testCases[i], dataSourceResultsTest))
		}

		// after passing the test cases, we run the actual data sources to collect their results
		// create data source results to store in the report
		// we use dataSourceResultsTest since this list is the complete list of data sources that have passed the test cases
		for i := range dataSourceResultsTest {
			// run the data source script
			var outTestCase bytes.Buffer
			var dataSourceResult types.DataSourceResult
			if dataSourceResultsTest[i].GetStatus() == types.ResultSuccess {
				cmdTestCase := exec.Command("python", getDSourcePath(dataSourceResultsTest[i].GetName()))
				cmdTestCase.Stdout = &outTestCase
				err = cmdTestCase.Run()
				if err != nil {
					l.Error(":skull: failed to execute data source script: %s", err.Error())
				}
				// collect test case result from the script
				result := strings.TrimSuffix(outTestCase.String(), "\n")
				fmt.Println("result from data sources: ", result)
				// By default, we consider returning null as failure. If any datasource does not follow this rule then it should not be used by any oracle scripts.
				dataSourceResult = types.NewDataSourceResult(dataSourceResultsTest[i].GetName(), []byte(result), types.ResultSuccess)
				if result == types.FailedResult || len(result) == 0 {
					// change status to fail so the datasource cannot be rewarded afterwards
					dataSourceResult.Status = types.ResultFailure
					dataSourceResult.Result = []byte(types.FailedResponseDs)
				} else {
					finalResultStr = finalResultStr + result + delimiter
				}
			} else {
				dataSourceResult = types.NewDataSourceResult(dataSourceResultsTest[i].GetName(), []byte(dataSourceResultsTest[i].GetResult()), types.ResultFailure)
			}
			// append an data source result into the list
			dataSourceResults = append(dataSourceResults, dataSourceResult)
		}

		fmt.Println("final result string: ", finalResultStr)
		fmt.Println("final result after trimming: ", strings.TrimSuffix(finalResultStr, "-"))
		msgReport := NewReport(req.RequestID, c.validator, dataSourceResults, testCaseResults, key.GetAddress(), sdk.NewCoins(sdk.NewCoin("orai", sdk.NewInt(int64(5000)))), []byte(finalResultStr))
		if len(finalResultStr) == 0 {
			msgReport.AggregatedResult = []byte(types.FailedResponseOs)
			// Create a new MsgCreateReport to the Oraichain
		} else {
			// "2" here is the expected output that the user wants to get
			cmd := exec.Command("python", oscriptPath, "aggregation", finalResultStr)
			var res bytes.Buffer
			cmd.Stdout = &res
			err = cmd.Run()
			if err != nil {
				l.Error(":skull: failed to aggregate results: %s", err.Error())
			}

			// collect data source result from the script
			ress := strings.TrimSuffix(res.String(), "\n")
			fmt.Printf("final result from oScript: %s\n", ress)
			msgReport.AggregatedResult = []byte(ress)
		}
		// Create a new MsgCreateReport to the Oraichain
		SubmitReport(c, l, key, msgReport)
	}(l.With("reqid", req.RequestID, "oscriptname", req.OracleScriptName), req)
}