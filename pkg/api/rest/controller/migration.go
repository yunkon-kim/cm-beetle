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

// Package controller has handlers and their request/response bodies for migration APIs
package controller

import (
	"fmt"
	"net/http"

	model "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/beetle"
	// cloudmodel "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/cloud/infra"
	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"

	"github.com/cloud-barista/cm-beetle/pkg/core/migration"
	"github.com/labstack/echo/v4"

	"github.com/rs/zerolog/log"
)

type MigrateInfraRequest struct {
	// [NOTE] Failed to embed the struct in CB-Tumblebug as follows:
	// mci.TbMciDynamicReq

	tbmodel.TbMciDynamicReq
}

type MigrateInfraResponse struct {
	tbmodel.TbMciDynamicReq
}

// MigrateInfra godoc
// @ID MigrateInfra
// @Summary Migrate an infrastructure to the multi-cloud infrastructure (MCI)
// @Description Migrate an infrastructure to the multi-cloud infrastructure (MCI)
// @Tags [Migration] Infrastructure
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param mciInfo body MigrateInfraRequest true "Specify the information for the targeted mulci-cloud infrastructure (MCI)"
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 {object} MigrateInfraResponse "Successfully migrated to the multi-cloud infrastructure"
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /migration/ns/{nsId}/mci [post]
func MigrateInfra(c echo.Context) error {

	// [Input]
	nsId := c.Param("nsId")
	if nsId == "" {
		err := fmt.Errorf("invalid request, namespace ID (nsId: %s) is required", nsId)
		log.Warn().Msg(err.Error())
		res := model.Response{
			Success: false,
			Text:    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, res)
	}
	// nsId := common.DefaulNamespaceId

	req := new(MigrateInfraRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	log.Debug().Msgf("req: %v\n", req)
	log.Debug().Msgf("req.TbMciDynamicReq: %v\n", req.TbMciDynamicReq)

	// [Process]
	// Create the VM infrastructure for migration
	mciInfo, err := migration.CreateVMInfra(nsId, &req.TbMciDynamicReq)

	log.Debug().Msgf("mciInfo: %v\n", mciInfo)

	// [Output]
	if err != nil {
		log.Error().Err(err).Msg("failed to create VM infrastructure")

		res := model.Response{
			Success: false,
			Text:    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := mciInfo
	return c.JSON(http.StatusOK, res)

}

// ListInfra godoc
// @ID ListInfra
// @Summary Get the migrated multi-cloud infrastructure (MCI)
// @Description Get the migrated multi-cloud infrastructure (MCI)
// @Tags [Migration] Infrastructure
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param option query string false "Option for getting the migrated multi-cloud infrastructure" Enums(status,id) default(status)
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 {object} migration.IdList "The ID list of The migrated multi-cloud infrastructure (MCI)"
// @Success 200 {object} migration.MciInfoList "The info list of the migrated multi-cloud infrastructure (MCI)"
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /migration/ns/{nsId}/mci [get]
func ListInfra(c echo.Context) error {

	// [Input]
	nsId := c.Param("nsId")
	if nsId == "" {
		err := fmt.Errorf("invalid request, the nanespace ID (nsId: %s) is required", nsId)
		log.Warn().Msg(err.Error())
		res := model.Response{
			Success: false,
			Text:    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, res)
	}
	// nsId := common.DefaulNamespaceId

	option := c.QueryParam("option")
	if option != "" && option != "status" && option != "id" {
		err := fmt.Errorf("invalid request, the option (option: %s) is invalid", option)
		log.Warn().Msg(err.Error())
		res := model.Response{
			Success: false,
			Text:    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, res)
	}

	// [Process] List the migrated multi-cloud infrastructures as the option
	switch option {
	case "status":
		infraInfoList, err := migration.ListAllVMInfraInfo(nsId)
		if err != nil {
			log.Error().Err(err).Msg("failed to get the migrated multi-cloud infrastructures")
			res := model.Response{
				Success: false,
				Text:    err.Error(),
			}
			return c.JSON(http.StatusInternalServerError, res)
		}
		return c.JSON(http.StatusOK, infraInfoList)

	case "id":
		idList, err := migration.ListVMInfraIDs(nsId, option)
		if err != nil {
			log.Error().Err(err).Msg("failed to get the migrated multi-cloud infrastructure IDs")
			res := model.Response{
				Success: false,
				Text:    err.Error(),
			}
			return c.JSON(http.StatusInternalServerError, res)
		}

		return c.JSON(http.StatusOK, idList)
	}

	return c.JSON(http.StatusInternalServerError, nil)
}

// GetInfra godoc
// @ID GetInfra
// @Summary Get the migrated multi-cloud infrastructure (MCI)
// @Description Get the migrated multi-cloud infrastructure (MCI)
// @Tags [Migration] Infrastructure
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param mciId path string true "Migrated Multi-Cloud Infrastructure (MCI) ID" default(mmci01)
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 {object} MigrateInfraResponse "The migrated multi-cloud infrastructure (MCI) information"
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /migration/ns/{nsId}/mci/{mciId} [get]
func GetInfra(c echo.Context) error {

	// [Input]
	nsId := c.Param("nsId")
	if nsId == "" {
		err := fmt.Errorf("invalid request, the nanespace ID (nsId: %s) is required", nsId)
		log.Warn().Msg(err.Error())
		res := model.Response{
			Success: false,
			Text:    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, res)
	}
	// nsId := common.DefaulNamespaceId

	mciId := c.Param("mciId")
	if mciId == "" {
		err := fmt.Errorf("invalid request, the multi-cloud infrastructure ID (mciId: %s) is required", mciId)
		log.Warn().Msg(err.Error())
		res := model.Response{
			Success: false,
			Text:    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, res)
	}

	// [Process]
	vmInfraInfo, err := migration.GetVMInfra(nsId, mciId)
	if err != nil {
		log.Error().Err(err).Msg("failed to get the migrated multi-cloud infrastructure")
		res := model.Response{
			Success: false,
			Text:    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}

	// [Ouput]
	return c.JSON(http.StatusOK, vmInfraInfo)
}

// DeleteInfra godoc
// @ID DeleteInfra
// @Summary Delete the migrated mult-cloud infrastructure (MCI)
// @Description Delete the migrated mult-cloud infrastructure (MCI)
// @Tags [Migration] Infrastructure
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param mciId path string true "Migrated Multi-Cloud Infrastructure (MCI) ID" default(mmci01)
// @Param action query string false "Action for deletion" Enums(terminate,force) default(terminate)
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 {object} model.Response "The result of deleting the migrated multi-cloud infrastructure (MCI)"
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /migration/ns/{nsId}/mci/{mciId} [delete]
func DeleteInfra(c echo.Context) error {

	// [Input]
	nsId := c.Param("nsId")
	if nsId == "" {
		err := fmt.Errorf("invalid request, the namespace ID (nsId: %s) is required", nsId)
		log.Warn().Msg(err.Error())
		res := model.Response{
			Success: false,
			Text:    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, res)
	}
	// nsId := common.DefaulNamespaceId

	mciId := c.Param("mciId")
	if mciId == "" {
		err := fmt.Errorf("invalid request, the multi-cloud infrastructure ID (mciId: %s) is required", mciId)
		log.Warn().Msg(err.Error())
		res := model.Response{
			Success: false,
			Text:    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, res)
	}

	action := c.QueryParam("action")
	if action != "" && action != "terminate" && action != "force" {
		err := fmt.Errorf("invalid request, the action (action: %s) is invalid", action)
		log.Warn().Msg(err.Error())
		res := model.Response{
			Success: false,
			Text:    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, res)
	}

	// [Process]
	retMsg, err := migration.DeleteVMInfra(nsId, mciId, action)

	if err != nil {
		log.Error().Err(err).Msg("failed to delete the migrated multi-cloud infrastructure")
		res := model.Response{
			Success: false,
			Text:    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}

	// [Ouput]
	res := model.Response{
		Success: true,
		Text:    retMsg.Message,
	}

	return c.JSON(http.StatusOK, res)
}

////////////////////////

// type MigrateNetworkRequest struct {
// 	cloudmodel.DummyNetwork
// }

// type MigrateNetworkResponse struct {
// 	cloudmodel.DummyNetwork
// }

// // MigrateNetwork godoc
// // @Summary (Skeleton) Migrate network on a cloud platform
// // @Description It migrates network on a cloud platform. Network includes name, ID, IPv4 CIDR block, IPv6 CIDR block, and so on.
// // @Tags [Migration] Infrastructure
// // @Accept  json
// // @Produce  json
// // @Param NetworkInfo body MigrateNetworkRequest true "Specify name, IPv4 CIDR block, etc."
// // @Success 200 {object} MigrateNetworkResponse "Successfully migrated network on a cloud platform"
// // @Failure 404 {object} common.SimpleMsg
// // @Failure 500 {object} common.SimpleMsg
// // @Router /migration/infra/network [post]
// func MigrateNetwork(c echo.Context) error {

// 	// [Input]
// 	req := &MigrateNetworkRequest{}
// 	if err := c.Bind(req); err != nil {
// 		return err
// 	}

// 	log.Trace().Msgf("req: %v\n", req)
// 	log.Trace().Msgf("req.DummyNetwork: %v\n", req.DummyNetwork)

// 	// [Process]
// 	// Something to process here like,
// 	// Perform some functions,
// 	// Calls external APIs and so on

// 	res := &MigrateNetworkResponse{}
// 	log.Trace().Msgf("res: %v\n", res)
// 	log.Trace().Msgf("res.DummyNetwork: %v\n", res.DummyNetwork)

// 	// This is an intentionally created variable.
// 	// You will have to delete this later.
// 	var err error = nil

// 	// [Ouput]
// 	if err != nil {
// 		log.Error().Err(err).Msg("Failed to migrate network on a cloud platform")
// 		mapA := map[string]string{"message": err.Error()}
// 		return c.JSON(http.StatusInternalServerError, &mapA)
// 	}

// 	return c.JSON(http.StatusOK, res)

// }

// ////////////////////////

// ////////////////////////

// type MigrateStorageRequest struct {
// 	cloudmodel.DummyStorage
// }

// type MigrateStorageResponse struct {
// 	cloudmodel.DummyStorage
// }

// // MigrateStorage godoc
// // @Summary (Skeleton) Migrate storage on a cloud platform
// // @Description It migrates storage on a cloud platform. Storage includes name, ID, type, size, and so on.
// // @Tags [Migration] Infrastructure
// // @Accept  json
// // @Produce  json
// // @Param StorageInfo body MigrateStorageRequest true "Specify name, type, size, affiliated Network ID, and so on."
// // @Success 200 {object} MigrateStorageResponse "Successfully migrated storage on a cloud platform"
// // @Failure 404 {object} common.SimpleMsg
// // @Failure 500 {object} common.SimpleMsg
// // @Router /migration/infra/storage [post]
// func MigrateStorage(c echo.Context) error {

// 	// [Input]
// 	req := &MigrateStorageRequest{}
// 	if err := c.Bind(req); err != nil {
// 		return err
// 	}

// 	log.Trace().Msgf("req: %v\n", req)
// 	log.Trace().Msgf("req.DummyStorage: %v\n", req.DummyStorage)

// 	// [Process]
// 	// Something to process here like,
// 	// Perform some functions,
// 	// Calls external APIs and so on

// 	res := &MigrateStorageResponse{}
// 	log.Trace().Msgf("res: %v\n", res)
// 	log.Trace().Msgf("res.DummyStorage: %v\n", res.DummyStorage)

// 	// This is an intentionally created variable.
// 	// You will have to delete this later.
// 	var err error = nil

// 	// [Ouput]
// 	if err != nil {
// 		log.Error().Err(err).Msg("Failed to migrate storage on a cloud platform")
// 		mapA := map[string]string{"message": err.Error()}
// 		return c.JSON(http.StatusInternalServerError, &mapA)
// 	}

// 	return c.JSON(http.StatusOK, res)

// }

// ////////////////////////

// ////////////////////////

// type MigrateInstanceRequest struct {
// 	cloudmodel.DummyInstance
// }

// type MigrateInstanceResponse struct {
// 	cloudmodel.DummyInstance
// }

// // MigrateInstance godoc
// // @Summary (Skeleton) Migrate instance on a cloud platform
// // @Description It migrates instance on a cloud platform. Storage includes name, spec, OS, and so on.
// // @Tags [Migration] Infrastructure
// // @Accept  json
// // @Produce  json
// // @Param InstanceInfo body MigrateInstanceRequest true "Specify name, spec, OS, and so on."
// // @Success 200 {object} MigrateInstanceResponse "Successfully migrated storage on a cloud platform"
// // @Failure 404 {object} common.SimpleMsg
// // @Failure 500 {object} common.SimpleMsg
// // @Router /migration/infra/instance [post]
// func MigrateInstance(c echo.Context) error {

// 	// [Input]
// 	req := &MigrateInstanceRequest{}
// 	if err := c.Bind(req); err != nil {
// 		return err
// 	}

// 	log.Trace().Msgf("req: %v\n", req)
// 	log.Trace().Msgf("req.DummyInstance: %v\n", req.DummyInstance)

// 	// [Process]
// 	// Something to process here like,
// 	// Perform some functions,
// 	// Calls external APIs and so on

// 	res := &MigrateInstanceResponse{}
// 	log.Trace().Msgf("res: %v\n", res)
// 	log.Trace().Msgf("res.DummyInstance: %v\n", res.DummyInstance)

// 	// This is an intentionally created variable.
// 	// You will have to delete this later.
// 	var err error = nil

// 	// [Ouput]
// 	if err != nil {
// 		log.Error().Err(err).Msg("Failed to migrate instance on a cloud platform")
// 		mapA := map[string]string{"message": err.Error()}
// 		return c.JSON(http.StatusInternalServerError, &mapA)
// 	}

// 	return c.JSON(http.StatusOK, res)

// }

// ////////////////////////
