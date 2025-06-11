/*
Copyright 2019 The Cloud-Barista Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package migration is to privision targat multi-cloud infra for migration
package migration

import (
	"fmt"
	"time"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	// cloudmodel "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/cloud/infra"
	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/core/recommendation"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
)

//"log"

//csv file handling

// REST API (echo)

// "github.com/cloud-barista/cm-beetle/pkg/core/mcir"

const (
	// ActionCreate is const for Create
	ActionCreate string = "Create"

	// ActionTerminate is const for Terminate
	ActionTerminate string = "Terminate"

	// ActionSuspend is const for Suspend
	ActionSuspend string = "Suspend"

	// ActionResume is const for Resume
	ActionResume string = "Resume"

	// ActionReboot is const for Reboot
	ActionReboot string = "Reboot"

	// ActionRefine is const for Refine
	ActionRefine string = "Refine"

	// ActionComplete is const for Complete
	ActionComplete string = "None"
)
const (
	// StatusRunning is const for Running
	StatusRunning string = "Running"

	// StatusSuspended is const for Suspended
	StatusSuspended string = "Suspended"

	// StatusFailed is const for Failed
	StatusFailed string = "Failed"

	// StatusTerminated is const for Terminated
	StatusTerminated string = "Terminated"

	// StatusCreating is const for Creating
	StatusCreating string = "Creating"

	// StatusSuspending is const for Suspending
	StatusSuspending string = "Suspending"

	// StatusResuming is const for Resuming
	StatusResuming string = "Resuming"

	// StatusRebooting is const for Rebooting
	StatusRebooting string = "Rebooting"

	// StatusTerminating is const for Terminating
	StatusTerminating string = "Terminating"

	// StatusUndefined is const for Undefined
	StatusUndefined string = "Undefined"

	// StatusComplete is const for Complete
	StatusComplete string = "None"
)

// const labelAutoGen string = "AutoGen"

// DefaultSystemLabel is const for string to specify the Default System Label
const DefaultSystemLabel string = "Managed by CM-Beetle"

// CreateVMInfraWithDefaults Create a VM infrastructure with defaults for the computing infra migration
func CreateVMInfraWithDefaults(nsId string, infraModel *tbmodel.TbMciDynamicReq) (tbmodel.TbMciInfo, error) {

	// Set timeout duration
	timeoutDuration := 40 * time.Minute

	client := resty.New()
	apiUser := config.Tumblebug.API.Username
	apiPass := config.Tumblebug.API.Password
	client.SetBasicAuth(apiUser, apiPass)

	// set Tumblebug rest url
	epTumblebug := config.Tumblebug.RestUrl

	method := "POST"
	url := epTumblebug + fmt.Sprintf("/ns/%s/mciDynamic", nsId)
	// url := fmt.Sprintf("%s/ns/{nsId}/mciDynamic%s", cbTumblebugApiEndpoint, idDetails.IdInSp)

	// Set request body
	requestBody := *infraModel

	// Set response body
	responseBody := tbmodel.TbMciInfo{}

	client.SetTimeout(timeoutDuration)

	err := common.ExecuteHttpRequest(
		client,
		method,
		url,
		nil,
		common.SetUseBody(requestBody),
		&requestBody,
		&responseBody,
		common.MediumDuration,
	)

	if err != nil {
		log.Error().Err(err).Msgf("failed to create/migrate the infrastructure (nsId: %s)", nsId)
		return tbmodel.TbMciInfo{}, err
	}

	return responseBody, nil
}

// CreateVMInfra creates a VM infrastructure for the computing infra migration
func CreateVMInfra(nsId string, infraModel *recommendation.RecommendedVmInfra) (tbmodel.TbMciInfo, error) {

	// Set timeout duration
	timeoutDuration := 40 * time.Minute

	client := resty.New()
	apiUser := config.Tumblebug.API.Username
	apiPass := config.Tumblebug.API.Password
	client.SetBasicAuth(apiUser, apiPass)

	// set Tumblebug rest url
	epTumblebug := config.Tumblebug.RestUrl

	// 1. Create a vpc/subnet
	method := "POST"
	url := epTumblebug + fmt.Sprintf("/ns/%s/resources/vNet", nsId)

	// Set request body
	requestBody := infraModel.TargetVNet

	// Set response body
	responseBody := tbmodel.TbMciInfo{}

	client.SetTimeout(timeoutDuration)

	err := common.ExecuteHttpRequest(
		client,
		method,
		url,
		nil,
		common.SetUseBody(requestBody),
		&requestBody,
		&responseBody,
		common.MediumDuration,
	)

	if err != nil {
		log.Error().Err(err).Msgf("failed to create/migrate the infrastructure (nsId: %s)", nsId)
		return tbmodel.TbMciInfo{}, err
	}

	return responseBody, nil
}

// List all migrated VM infrastructures
func ListAllVMInfraInfo(nsId string) (MciInfoList, error) {

	// Set timeout duration
	timeoutDuration := 5 * time.Minute

	var emptyRet MciInfoList
	var mciInfoList MciInfoList

	// Initialize resty client with basic auth
	client := resty.New()
	apiUser := config.Tumblebug.API.Username
	apiPass := config.Tumblebug.API.Password
	client.SetBasicAuth(apiUser, apiPass)

	// set Tumblebug rest url
	epTumblebug := config.Tumblebug.RestUrl

	// Set qeury parameters
	queryParams := "?option=status"

	// check readyz
	method := "GET"
	url := fmt.Sprintf("%s/ns/%s/mci", epTumblebug, nsId)
	if queryParams != "" {
		url += queryParams
	}

	// Set request body
	requestBody := common.NoBody

	// Set response body
	client.SetTimeout(timeoutDuration)

	err := common.ExecuteHttpRequest(
		client,
		method,
		url,
		nil,
		common.SetUseBody(requestBody),
		&requestBody,
		&mciInfoList,
		common.MediumDuration,
	)

	if err != nil {
		log.Error().Err(err).Msgf("failed to get the infrastructure info list (nsId: %s)", nsId)
		return emptyRet, err
	}

	return mciInfoList, nil
}

// Get all migrated VM infrastructures
func ListVMInfraIDs(nsId string, option string) (IdList, error) {

	// Set timeout duration
	timeoutDuration := 5 * time.Minute

	var emptyRet IdList
	var idList IdList
	idList.IdList = make([]string, 0)

	/*
	 * Validate the input
	 */

	var queryParams string
	if option != "id" {
		log.Error().Msgf("invalid option: %s", option)
		return emptyRet, fmt.Errorf("invalid option: %s", option)
	}

	// Set qeury parameters
	queryParams = "?option=id"

	// Initialize resty client with basic auth
	client := resty.New()
	apiUser := config.Tumblebug.API.Username
	apiPass := config.Tumblebug.API.Password
	client.SetBasicAuth(apiUser, apiPass)

	// set Tumblebug rest url
	epTumblebug := config.Tumblebug.RestUrl

	// check readyz
	method := "GET"
	url := fmt.Sprintf("%s/ns/%s/mci", epTumblebug, nsId)
	if queryParams != "" {
		url += queryParams
	}

	// Set request body
	requestBody := common.NoBody

	// Set response body
	tbResp := new(tbmodel.IdList)

	client.SetTimeout(timeoutDuration)
	err := common.ExecuteHttpRequest(
		client,
		method,
		url,
		nil,
		common.SetUseBody(requestBody),
		&requestBody,
		tbResp,
		common.MediumDuration,
	)

	if err != nil {
		log.Error().Err(err).Msgf("failed to get the infrastructure IDs (nsId: %s)", nsId)
		return emptyRet, err
	}

	// Return the result
	idList.IdList = append(idList.IdList, tbResp.IdList...)

	return idList, nil
}

// Get the migrated VM infrastructure
func GetVMInfra(nsId, infraId string) (tbmodel.TbMciInfo, error) {

	// Set timeout duration
	timeoutDuration := 5 * time.Minute

	// Initialize resty client with basic auth
	client := resty.New()
	apiUser := config.Tumblebug.API.Username
	apiPass := config.Tumblebug.API.Password
	client.SetBasicAuth(apiUser, apiPass)

	// set Tumblebug rest url
	epTumblebug := config.Tumblebug.RestUrl

	// check readyz
	method := "GET"
	url := fmt.Sprintf("%s/ns/%s/mci/%s", epTumblebug, nsId, infraId)

	// Set request body
	requestBody := common.NoBody

	// Set response body
	responseBody := new(tbmodel.TbMciInfo)

	client.SetTimeout(timeoutDuration)

	err := common.ExecuteHttpRequest(
		client,
		method,
		url,
		nil,
		common.SetUseBody(requestBody),
		&requestBody,
		responseBody,
		common.MediumDuration,
	)

	if err != nil {
		log.Error().Err(err).Msgf("failed to get the infrastructure info (nsId: %s, infraId: %s)", nsId, infraId)
		return tbmodel.TbMciInfo{}, err
	}

	return *responseBody, nil
}

// Delete the migrated VM infrastructure
func DeleteVMInfra(nsId, infraId, action string) (common.SimpleMsg, error) {

	// Set timeout duration
	timeoutDuration := 40 * time.Minute

	// Initialize resty client with basic auth
	client := resty.New()
	apiUser := config.Tumblebug.API.Username
	apiPass := config.Tumblebug.API.Password
	client.SetBasicAuth(apiUser, apiPass)

	// set Tumblebug rest url
	epTumblebug := config.Tumblebug.RestUrl

	// delete the infrastructure with terminate option
	method := "DELETE"
	url := fmt.Sprintf("%s/ns/%s/mci/%s", epTumblebug, nsId, infraId)
	if action != "" {
		url += "?option=" + action
	}

	// Set request body
	requestBody := common.NoBody

	// Set response body
	responseBody := new(common.SimpleMsg)

	client.SetTimeout(timeoutDuration)

	err := common.ExecuteHttpRequest(
		client,
		method,
		url,
		nil,
		common.SetUseBody(requestBody),
		&requestBody,
		responseBody,
		common.MediumDuration,
	)

	if err != nil {
		log.Error().Err(err).Msgf("failed to delete the infrastructure (nsId: %s, infraId: %s)", nsId, infraId)
		return common.SimpleMsg{}, err
	}

	time.Sleep(15 * time.Second)
	// delete the infrastructure with terminate option
	method = "DELETE"
	url = fmt.Sprintf("%s/ns/%s/sharedResources", epTumblebug, nsId)

	// Set request body
	requestBody = common.NoBody

	// Set response body
	resDeleteDefaultResources := new(common.IdList)

	client.SetTimeout(timeoutDuration)

	err = common.ExecuteHttpRequest(
		client,
		method,
		url,
		nil,
		common.SetUseBody(requestBody),
		&requestBody,
		resDeleteDefaultResources,
		common.MediumDuration,
	)

	if err != nil {
		log.Error().Err(err).Msgf("failed to delete the infrastructure (nsId: %s, infraId: %s)", nsId, infraId)
		return common.SimpleMsg{}, err
	}

	return *responseBody, nil
}
