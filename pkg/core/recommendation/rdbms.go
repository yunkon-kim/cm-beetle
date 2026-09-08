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

// Package recommendation provides logic to recommend optimal target cloud resources
package recommendation

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"

	rdbmsmodel "github.com/cloud-barista/cm-beetle/imdl/rdbms-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/rs/zerolog/log"
)

// GetRDBMSSupport retrieves CSP support matrix for managed RDBMS via CB-Tumblebug.
func GetRDBMSSupport(providerName string) (rdbmsmodel.RDBMSSupportResponse, error) {
	log.Debug().Str("providerName", providerName).Msg("Fetching RDBMS support matrix via CB-Tumblebug")
	return tbclient.NewSession().GetRDBMSSupport(providerName)
}

// GetRDBMSCapability retrieves real-time capability information for a connection and optional dbEngine via CB-Tumblebug.
func GetRDBMSCapability(connectionName string, optionalEngine ...string) (rdbmsmodel.RDBMSCapabilityResponse, error) {
	log.Debug().Str("connectionName", connectionName).Msg("Fetching RDBMS capability via CB-Tumblebug")
	return tbclient.NewSession().GetRDBMSCapability(connectionName, optionalEngine...)
}

// ValidateRDBMS performs dry-run validation and default autofill for an RDBMS create request.
// It executes a two-stage validation pipeline:
//  1. Beetle Deep Validation: Referential & cross-connection integrity checks (e.g. verifying VNet/SG connection matches target RDBMS connection)
//  2. Tumblebug Infra Validation: Dry-run check against CSP capability & API constraints via CB-Tumblebug.
func ValidateRDBMS(nsId string, req rdbmsmodel.RDBMSCreateRequest) (rdbmsmodel.RDBMSCreateRequest, error) {
	var emptyRes rdbmsmodel.RDBMSCreateRequest

	nsId = strings.TrimSpace(nsId)
	if nsId == "" {
		return emptyRes, fmt.Errorf("nsId is required")
	}
	targetConn := strings.TrimSpace(req.ConnectionName)
	if targetConn == "" {
		return emptyRes, fmt.Errorf("connectionName is required")
	}

	tbSess := tbclient.NewSession()

	// 1. Referential & Cross-connection check for VNet
	if req.VNetId != "" {
		vNetInfo, err := tbSess.ReadVNet(nsId, req.VNetId)
		if err == nil && vNetInfo.Id != "" {
			if vNetInfo.ConnectionName != "" && !strings.EqualFold(vNetInfo.ConnectionName, targetConn) {
				return emptyRes, fmt.Errorf("cross-connection mismatch: VNet '%s' belongs to connection '%s', but RDBMS target connection is '%s'",
					req.VNetId, vNetInfo.ConnectionName, targetConn)
			}
		}
	}

	// 2. Referential & Cross-connection check for Security Groups
	if len(req.SecurityGroupIds) > 0 {
		for _, sgId := range req.SecurityGroupIds {
			if sgId == "" {
				continue
			}
			sgInfo, err := tbSess.ReadSecurityGroup(nsId, sgId)
			if err == nil && sgInfo.Id != "" {
				if sgInfo.ConnectionName != "" && !strings.EqualFold(sgInfo.ConnectionName, targetConn) {
					return emptyRes, fmt.Errorf("cross-connection mismatch: SecurityGroup '%s' belongs to connection '%s', but RDBMS target connection is '%s'",
						sgId, sgInfo.ConnectionName, targetConn)
				}
			}
		}
	}

	// 3. NHN Cloud specific flag validation
	if req.NHNDBSGToAllowAllInbound {
		if !strings.HasPrefix(strings.ToLower(targetConn), "nhn") {
			return emptyRes, fmt.Errorf("nhnDBSGToAllowAllInbound is only supported for NHN Cloud")
		}
		if !req.PublicAccess {
			return emptyRes, fmt.Errorf("nhnDBSGToAllowAllInbound requires publicAccess=true")
		}
	}

	// 4. Pass to Tumblebug for live CSP capability strict dry-run validation (autoFillDefaults=false)
	req.AutoFillDefaults = false
	log.Info().Str("nsId", nsId).Str("rdbmsName", req.Name).Str("connection", targetConn).Msg("Validating RDBMS create request via CB-Tumblebug (strict mode: autoFillDefaults=false)")
	return tbSess.ValidateRDBMS(nsId, req)
}

// ValidateSourceRDBMS strictly validates the source RDBMS model properties before recommendation.
func ValidateSourceRDBMS(sources []rdbmsmodel.SourceRDBMSProperty) error {
	if len(sources) == 0 {
		err := fmt.Errorf("at least one source RDBMS instance is required")
		log.Warn().Msg(err.Error())
		return err
	}

	for i, src := range sources {
		instName := strings.TrimSpace(src.InstanceName)
		if instName == "" {
			err := fmt.Errorf("instance [%d]: instanceName is required", i)
			log.Warn().Msg(err.Error())
			return err
		}

		engine := strings.ToLower(strings.TrimSpace(src.Engine))
		if engine == "" {
			err := fmt.Errorf("instance '%s': engine is required", instName)
			log.Warn().Msg(err.Error())
			return err
		}

		if strings.TrimSpace(src.EngineVersion) == "" {
			err := fmt.Errorf("instance '%s': engineVersion is required", instName)
			log.Warn().Msg(err.Error())
			return err
		}

		if src.Vcpu <= 0 {
			err := fmt.Errorf("instance '%s': vcpu must be greater than 0 (got %d)", instName, src.Vcpu)
			log.Warn().Msg(err.Error())
			return err
		}

		if src.MemoryMb <= 0 {
			err := fmt.Errorf("instance '%s': memoryMb must be greater than 0 (got %d)", instName, src.MemoryMb)
			log.Warn().Msg(err.Error())
			return err
		}

		if src.StorageSizeGb <= 0 {
			err := fmt.Errorf("instance '%s': storageSizeGb must be greater than 0 (got %d)", instName, src.StorageSizeGb)
			log.Warn().Msg(err.Error())
			return err
		}

		for dbIdx, db := range src.Databases {
			if strings.TrimSpace(db.DatabaseName) == "" {
				err := fmt.Errorf("instance '%s': database [%d] databaseName is required", instName, dbIdx)
				log.Warn().Msg(err.Error())
				return err
			}
		}
	}

	return nil
}

// RecommendRDBMS recommends optimal managed RDBMS instances for target cloud migration.
func RecommendRDBMS(desiredCsp, desiredRegion string, sources []rdbmsmodel.SourceRDBMSProperty) (rdbmsmodel.RecommendedRDBMS, error) {
	var emptyRes rdbmsmodel.RecommendedRDBMS

	desiredCsp = strings.ToLower(strings.TrimSpace(desiredCsp))
	desiredRegion = strings.ToLower(strings.TrimSpace(desiredRegion))

	if desiredCsp == "" || desiredRegion == "" {
		err := fmt.Errorf("desiredCsp and desiredRegion are required")
		log.Warn().Msg(err.Error())
		return emptyRes, err
	}

	// 1. Upfront Source Model Validation
	if err := ValidateSourceRDBMS(sources); err != nil {
		return emptyRes, err
	}

	log.Info().
		Str("csp", desiredCsp).
		Str("region", desiredRegion).
		Int("sourceCount", len(sources)).
		Msg("Starting RDBMS recommendation")

	warnings := make([]string, 0)

	// 2. Fetch CSP support info from Tumblebug & Validate CSP support
	supportResp, err := tbclient.NewSession().GetRDBMSSupport(desiredCsp)
	if err != nil {
		log.Warn().Err(err).Str("csp", desiredCsp).Msg("Failed to fetch CSP RDBMS support info from Tumblebug")
		return emptyRes, fmt.Errorf("failed to fetch RDBMS support info for CSP '%s': %w", desiredCsp, err)
	}

	support, hasSupport := supportResp.Supports[desiredCsp]
	if hasSupport && !support.Supported {
		err := fmt.Errorf("managed RDBMS is not supported on CSP '%s'", desiredCsp)
		log.Warn().Str("csp", desiredCsp).Msg(err.Error())
		return emptyRes, err
	}

	// 3. Validate that each source instance's engine is supported on the target CSP (strictly NO fallback)
	for _, src := range sources {
		targetEngine := strings.ToLower(strings.TrimSpace(src.Engine))
		if !isDBEngineSupported(support, hasSupport, targetEngine) {
			err := fmt.Errorf("dbEngine '%s' is not supported on CSP '%s' (supported engines: %v)", targetEngine, desiredCsp, support.SupportedDBEngines)
			log.Warn().Err(err).Str("csp", desiredCsp).Str("engine", targetEngine).Str("instance", src.InstanceName).Msg("Source engine validation failed against target CSP")
			return emptyRes, err
		}
	}

	// Recommend each target RDBMS instance
	targetInstances := make([]rdbmsmodel.TargetRDBMSInstance, 0, len(sources))

	for i, src := range sources {
		instNum := i + 1
		targetName := fmt.Sprintf("mig-rdbms-%02d", instNum)

		// 1. Engine & Version Recommendation
		targetEngine := strings.ToLower(strings.TrimSpace(src.Engine))

		// Fetch engine-specific live capability from Tumblebug
		connName := fmt.Sprintf("%s-%s", desiredCsp, desiredRegion)
		capaResp, capErr := tbclient.NewSession().GetRDBMSCapability(connName, targetEngine)
		if capErr != nil {
			log.Warn().Err(capErr).Str("connectionName", connName).Str("engine", targetEngine).Msg("Failed to fetch live RDBMS capability from Tumblebug")
			return emptyRes, fmt.Errorf("failed to retrieve live RDBMS capability for connection '%s' and engine '%s': %w", connName, targetEngine, capErr)
		}
		capa := capaResp.Supports

		// Select Engine Version with strict matching (no arbitrary fallback)
		targetVersion, err := selectEngineVersion(targetEngine, src.EngineVersion, capa.SupportedVersions, &warnings, src.InstanceName)
		if err != nil {
			log.Warn().Err(err).Str("instance", src.InstanceName).Str("engine", targetEngine).Str("version", src.EngineVersion).Msg("Engine version selection failed")
			return emptyRes, err
		}

		// 2. DB Spec Recommendation (Strategy: Conservative Capacity Proximity, Target >= Source)
		targetSpec, err := recommendDBInstanceSpec(src.Vcpu, src.MemoryMb, src.StorageSizeGb, targetEngine, capa)
		if err != nil {
			log.Warn().Err(err).Str("instance", src.InstanceName).Int("vcpu", src.Vcpu).Int("memoryMb", src.MemoryMb).Int("storageSizeGb", src.StorageSizeGb).Msg("DBInstanceSpec recommendation failed")
			return emptyRes, err
		}

		// 3. Storage Type Recommendation using live Notes.StorageTypes and StorageTypeOptions
		targetStorageType, selectedNote, err := selectStorageType(src.StorageType, capa, &warnings, src.InstanceName)
		if err != nil {
			log.Warn().Err(err).Str("instance", src.InstanceName).Str("storageType", src.StorageType).Msg("Storage type selection failed")
			return emptyRes, err
		}

		// Determine storage size boundaries using live capability info
		minStorage := capa.StorageSizeRange.Min
		maxStorage := capa.StorageSizeRange.Max
		if selectedNote != nil {
			if selectedNote.MinSize > 0 {
				minStorage = selectedNote.MinSize
			}
			if selectedNote.MaxSize > 0 {
				maxStorage = selectedNote.MaxSize
			}
		}

		targetStorageSize := src.StorageSizeGb
		if minStorage > 0 && targetStorageSize < minStorage {
			warning := fmt.Sprintf("Adjusted storage size for instance '%s' from %dGB to minimum %dGB required by target cloud (%s).",
				src.InstanceName, targetStorageSize, minStorage, targetStorageType)
			warnings = append(warnings, warning)
			targetStorageSize = minStorage
		}
		if maxStorage > 0 && targetStorageSize > maxStorage {
			warning := fmt.Sprintf("Clamped storage size for instance '%s' from %dGB to maximum %dGB supported by target cloud.",
				src.InstanceName, targetStorageSize, maxStorage)
			warnings = append(warnings, warning)
			targetStorageSize = maxStorage
		}

		// IOPS handling: dynamic validation using live StorageTypeNote constraints
		targetIops := ""
		if selectedNote != nil && selectedNote.RequiresIops {
			assignedIops := 3000
			if selectedNote.IopsRange != nil && selectedNote.IopsRange.Min > 0 {
				assignedIops = selectedNote.IopsRange.Min
			}
			if src.Iops > 0 {
				assignedIops = src.Iops
				if selectedNote.IopsRange != nil {
					if selectedNote.IopsRange.Min > 0 && assignedIops < selectedNote.IopsRange.Min {
						assignedIops = selectedNote.IopsRange.Min
					}
					if selectedNote.IopsRange.Max > 0 && assignedIops > selectedNote.IopsRange.Max {
						assignedIops = selectedNote.IopsRange.Max
					}
				}
			}
			targetIops = fmt.Sprintf("%d", assignedIops)
		}

		// 4. Admin Credentials Default from Capability
		targetAdminUser := "cbuser"
		if capa.AdminUserNameRequirement != nil && capa.AdminUserNameRequirement.FixedValue != "" {
			targetAdminUser = capa.AdminUserNameRequirement.FixedValue
		} else if capa.AdminUserNameRequirement != nil && len(capa.AdminUserNameRequirement.ReservedValues) > 0 {
			for _, reserved := range capa.AdminUserNameRequirement.ReservedValues {
				if strings.EqualFold(targetAdminUser, reserved) {
					targetAdminUser = "dbadmin"
					break
				}
			}
		}

		// 5. Network & CSP-specific warnings
		if desiredCsp == "ncp" && src.PublicAccess {
			warning := fmt.Sprintf("NCP Cloud DB does not provide external public IP by default; instance '%s' will be created within private VPC.", src.InstanceName)
			warnings = append(warnings, warning)
		}
		targetNHNDBSG := src.NHNDBSGToAllowAllInbound
		if src.HighAvailability && desiredCsp == "aws" {
			warning := fmt.Sprintf("High availability (Multi-AZ) on AWS requires Subnets in at least two distinct Availability Zones for instance '%s'.", src.InstanceName)
			warnings = append(warnings, warning)
		}

		// 6. Map Inner Databases
		targetDatabases := make([]rdbmsmodel.TargetDatabase, 0, len(src.Databases))
		for _, db := range src.Databases {
			targetDatabases = append(targetDatabases, rdbmsmodel.TargetDatabase{
				DatabaseName: db.DatabaseName,
				CharacterSet: db.CharacterSet,
			})
		}

		targetBackupDays := src.BackupRetentionDays
		if strings.EqualFold(desiredCsp, "ibm") {
			targetBackupDays = 0 // IBM Cloud Databases does not support setting BackupRetentionDays during provisioning
		}

		targetInst := rdbmsmodel.TargetRDBMSInstance{
			SourceInstanceName:       src.InstanceName,
			SourceMachineId:          src.MachineId,
			RDBMSName:                targetName,
			DBEngine:                 targetEngine,
			DBEngineVersion:          targetVersion,
			DBInstanceSpec:           targetSpec,
			StorageType:              targetStorageType,
			StorageSize:              targetStorageSize,
			Iops:                     targetIops,
			AdminUserName:            targetAdminUser,
			HighAvailability:         src.HighAvailability,
			BackupRetentionDays:      targetBackupDays,
			PublicAccess:             src.PublicAccess,
			NHNDBSGToAllowAllInbound: targetNHNDBSG,
			Databases:                targetDatabases,
		}

		targetInstances = append(targetInstances, targetInst)
	}

	result := rdbmsmodel.RecommendedRDBMS{
		Status:               "recommended",
		Description:          fmt.Sprintf("Successfully recommended %d managed RDBMS configuration(s) for %s (%s)", len(targetInstances), desiredCsp, desiredRegion),
		Warnings:             warnings,
		TargetCloud:          rdbmsmodel.CloudProperty{Csp: desiredCsp, Region: desiredRegion},
		TargetRDBMSInstances: targetInstances,
	}

	log.Info().
		Int("targetInstances", len(targetInstances)).
		Int("warnings", len(warnings)).
		Msg("RDBMS recommendation completed")

	return result, nil
}

// selectEngineVersion selects the best-matching target database engine version from capability.
// It handles standard semver ("8.0", "8.0.32") as well as CSP-specific codes (e.g. NHN "MYSQL_V8032", Tencent "8.0").
// If the requested version is not directly supported, it recommends the closest available version with an informative warning.
func selectEngineVersion(engine, requestedVersion string, supported []string, warnings *[]string, instName string) (string, error) {
	requestedVersion = strings.TrimSpace(requestedVersion)
	if len(supported) == 0 {
		if requestedVersion != "" {
			return requestedVersion, nil
		}
		if strings.EqualFold(engine, "mariadb") {
			return "10.6", nil
		}
		return "8.0", nil
	}

	// 1. Exact match (case-insensitive)
	for _, v := range supported {
		if strings.EqualFold(v, requestedVersion) {
			return v, nil
		}
	}

	// 2. Direct substring / prefix / suffix match
	if requestedVersion != "" {
		for _, v := range supported {
			if strings.EqualFold(v, requestedVersion) ||
				strings.HasPrefix(v, requestedVersion) ||
				strings.HasPrefix(requestedVersion, v) ||
				strings.Contains(strings.ToLower(v), strings.ToLower(requestedVersion)) {
				return v, nil
			}
		}
	}

	// 3. CSP Version Code Normalization match (e.g. "8.0" -> "80" matching "MYSQL_V8032", "5.7" -> "57" matching "MYSQL_V5744")
	if requestedVersion != "" {
		reqDigits := extractVersionDigits(requestedVersion)
		if len(reqDigits) >= 2 {
			for _, v := range supported {
				vDigits := extractVersionDigits(v)
				if strings.HasPrefix(vDigits, reqDigits) || strings.Contains(vDigits, reqDigits) {
					return v, nil
				}
			}
		}
	}

	// 4. Proximity fallback: If requestedVersion was specified but not directly supported,
	// find the closest available version (e.g. "8.4" for "8.0", or latest available version)
	// and provide an informative warning instead of failing.
	if requestedVersion != "" {
		chosen := findClosestVersion(requestedVersion, supported)
		warning := fmt.Sprintf("Requested %s version '%s' is not directly supported on target cloud; recommended closest available version '%s' (supported versions: %v).",
			engine, requestedVersion, chosen, supported)
		*warnings = append(*warnings, warning)
		log.Info().
			Str("instance", instName).
			Str("engine", engine).
			Str("requestedVersion", requestedVersion).
			Str("recommendedVersion", chosen).
			Msg("Recommended closest available engine version for target cloud")
		return chosen, nil
	}

	// 5. Preferred default version (e.g., 8.0 for MySQL, 10.6 for MariaDB)
	preferred := "8.0"
	preferredDigits := "80"
	if strings.EqualFold(engine, "mariadb") {
		preferred = "10.6"
		preferredDigits = "106"
	}

	for _, v := range supported {
		if strings.Contains(strings.ToLower(v), preferred) || strings.Contains(extractVersionDigits(v), preferredDigits) {
			return v, nil
		}
	}

	return supported[0], nil
}

// findClosestVersion selects the closest available version from supported versions list.
// It prioritizes matching major version prefix (e.g. "8." or same leading digit),
// falling back to the latest supported version in the supported list.
func findClosestVersion(requested string, supported []string) string {
	if len(supported) == 0 {
		return requested
	}

	reqDigits := extractVersionDigits(requested)

	// 1. Check for shared major version digit (e.g., requested "8.0" matching "8.4" or "8.0.32")
	if len(reqDigits) > 0 {
		reqMajorDigit := reqDigits[0]
		for _, v := range supported {
			vDigits := extractVersionDigits(v)
			if len(vDigits) > 0 && vDigits[0] == reqMajorDigit {
				return v
			}
		}
	}

	// 2. Check for semver major version string (e.g. "8" or "10")
	reqParts := strings.Split(strings.TrimPrefix(strings.ToLower(requested), "v"), ".")
	if len(reqParts) > 0 && reqParts[0] != "" {
		reqMajor := reqParts[0]
		for _, v := range supported {
			vClean := strings.TrimPrefix(strings.ToLower(v), "v")
			vParts := strings.Split(vClean, ".")
			if len(vParts) > 0 && vParts[0] == reqMajor {
				return v
			}
		}
	}

	// 3. Fallback to latest supported option (typically the last or first in list)
	return supported[len(supported)-1]
}

// isPremiumStorageNote checks if a storage type option is an ultra-high performance / expensive enterprise tier.
func isPremiumStorageNote(note rdbmsmodel.StorageTypeNote, opt string) bool {
	if strings.EqualFold(note.RecommendationLevel, "premium") {
		return true
	}
	optLower := strings.ToLower(opt)
	// Common ultra-performance enterprise tiers with high minimums and spec limitations
	if strings.Contains(optLower, "essd2") || strings.Contains(optLower, "essd3") ||
		strings.HasPrefix(optLower, "io") || strings.Contains(optLower, "extreme") {
		return true
	}
	return false
}

// extractVersionDigits extracts consecutive digit sequences for version matching (e.g. "8.0.35" -> "8035", "MYSQL_V8032" -> "8032")
func extractVersionDigits(s string) string {
	var b strings.Builder
	for _, r := range s {
		if unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// selectStorageType selects the optimal storage type and guidance note dynamically from Tumblebug capability.
// If the target cloud does not support storage type selection (e.g. Azure, IBM, NCP), storageType is omitted with an informative log and warning.
// If storage type selection is supported, it finds the exact match or closest equivalent storage type, warning if mapped from generic descriptors.
func selectStorageType(req string, capa rdbmsmodel.RDBMSMetaInfo, warnings *[]string, instName string) (string, *rdbmsmodel.StorageTypeNote, error) {
	// 1. If the CSP does not support storage type selection (e.g. Azure, IBM, NCP), storageType must be omitted
	if !capa.SupportsStorageTypeSelection {
		if req != "" && !strings.EqualFold(req, "NA") && !strings.EqualFold(req, "default") {
			warning := fmt.Sprintf("Storage type selection is not configurable on target cloud (%s); requested storage type '%s' for instance '%s' will be managed automatically by the provider.",
				capa.ProviderName, req, instName)
			*warnings = append(*warnings, warning)
			log.Info().
				Str("instance", instName).
				Str("csp", capa.ProviderName).
				Str("requestedStorageType", req).
				Msg("Target cloud automatically manages storage type (selection not configurable); returning empty storageType")
		}
		return "", nil, nil
	}

	// Filter out non-actionable placeholder options (e.g. "NA")
	usableOptions := make([]string, 0, len(capa.StorageTypeOptions))
	for _, opt := range capa.StorageTypeOptions {
		if !strings.EqualFold(opt, "NA") && strings.TrimSpace(opt) != "" {
			usableOptions = append(usableOptions, opt)
		}
	}
	if len(usableOptions) == 0 {
		return "", nil, nil
	}

	// Collect storage notes map from notes or storageTypeGuidance
	notesMap := make(map[string]rdbmsmodel.StorageTypeNote)
	if capa.Notes != nil && len(capa.Notes.StorageTypes) > 0 {
		for _, note := range capa.Notes.StorageTypes {
			notesMap[strings.ToLower(note.StorageType)] = note
		}
	}
	if len(capa.StorageTypeGuidance) > 0 {
		for k, v := range capa.StorageTypeGuidance {
			notesMap[strings.ToLower(k)] = v
		}
	}

	// 2. Check if user requested a specific or generic storage type
	if req != "" && !strings.EqualFold(req, "NA") && !strings.EqualFold(req, "default") {
		reqLower := strings.ToLower(req)

		// 2.1 Exact / case-insensitive match against usable options
		for _, opt := range usableOptions {
			if strings.EqualFold(opt, req) {
				if note, ok := notesMap[reqLower]; ok {
					return opt, &note, nil
				}
				return opt, nil, nil
			}
		}

		// 2.2 Intent-based matching for generic on-premise hardware types (SSD / HDD / NVMe)
		isSSDIntent := strings.EqualFold(req, "SSD") || strings.Contains(reqLower, "ssd") || strings.Contains(reqLower, "nvme") || strings.Contains(reqLower, "flash")
		isHDDIntent := strings.EqualFold(req, "HDD") || strings.Contains(reqLower, "hdd") || strings.Contains(reqLower, "magnetic")

		if isSSDIntent {
			// A. Priority 1: Pick any option explicitly flagged Recommended in usableOptions (e.g. cloud_auto on Alibaba, gp3 on AWS, PD_SSD on GCP)
			for _, opt := range usableOptions {
				if note, ok := notesMap[strings.ToLower(opt)]; ok && note.Recommended {
					return opt, &note, nil
				}
			}

			// B. Priority 2: Pick defaultStorageType from capability if available and not premium
			if capa.DefaultStorageType != "" && !strings.EqualFold(capa.DefaultStorageType, "NA") {
				for _, opt := range usableOptions {
					if strings.EqualFold(opt, capa.DefaultStorageType) {
						note, hasNote := notesMap[strings.ToLower(opt)]
						if !hasNote || !isPremiumStorageNote(note, opt) {
							if hasNote {
								return opt, &note, nil
							}
							return opt, nil, nil
						}
					}
				}
			}

			// C. Priority 3: Pick standard SSD option from usableOptions (excluding premium enterprise tiers like cloud_essd2/3, io1/2)
			for _, opt := range usableOptions {
				optLower := strings.ToLower(opt)
				note, hasNote := notesMap[optLower]
				if hasNote && isPremiumStorageNote(note, opt) {
					continue
				}
				if strings.Contains(optLower, "ssd") ||
					(hasNote && (strings.Contains(strings.ToLower(note.DisplayName), "ssd") || strings.Contains(strings.ToLower(note.Description), "ssd"))) {
					if hasNote {
						return opt, &note, nil
					}
					return opt, nil, nil
				}
			}

			// D. Priority 4: If no literal "SSD" exists on target cloud (e.g. OpenStack [RBD, __DEFAULT__]):
			// Pick high-performance block backend (e.g. RBD)
			for _, opt := range usableOptions {
				if strings.EqualFold(opt, "RBD") {
					var chosenNote *rdbmsmodel.StorageTypeNote
					if note, ok := notesMap[strings.ToLower(opt)]; ok {
						chosenNote = &note
					}
					warning := fmt.Sprintf("Storage type '%s' requested for instance '%s' is not directly offered on %s; recommended closest available storage type 'RBD' (available: %v).",
						req, instName, capa.ProviderName, usableOptions)
					*warnings = append(*warnings, warning)
					log.Info().
						Str("instance", instName).
						Str("csp", capa.ProviderName).
						Str("requestedStorageType", req).
						Str("recommendedStorageType", "RBD").
						Msg("Recommended closest available storage type for target cloud")
					return opt, chosenNote, nil
				}
			}

			// E. Priority 5: Fallback to first non-premium option in usableOptions with warning
			var chosen string
			var chosenNote *rdbmsmodel.StorageTypeNote
			for _, opt := range usableOptions {
				note, hasNote := notesMap[strings.ToLower(opt)]
				if hasNote && isPremiumStorageNote(note, opt) {
					continue
				}
				chosen = opt
				if hasNote {
					chosenNote = &note
				}
				break
			}
			if chosen == "" {
				chosen = usableOptions[0]
				if note, ok := notesMap[strings.ToLower(chosen)]; ok {
					chosenNote = &note
				}
			}

			warning := fmt.Sprintf("Storage type '%s' requested for instance '%s' is not directly offered on %s; recommended closest available storage type '%s' (available: %v).",
				req, instName, capa.ProviderName, chosen, usableOptions)
			*warnings = append(*warnings, warning)
			log.Info().
				Str("instance", instName).
				Str("csp", capa.ProviderName).
				Str("requestedStorageType", req).
				Str("recommendedStorageType", chosen).
				Msg("Recommended closest available storage type for target cloud")
			return chosen, chosenNote, nil

		} else if isHDDIntent {
			// Check note metadata mentioning HDD or magnetic
			for _, opt := range usableOptions {
				optLower := strings.ToLower(opt)
				note, hasNote := notesMap[optLower]
				if strings.Contains(optLower, "hdd") || strings.Contains(optLower, "magnetic") ||
					(hasNote && (strings.Contains(strings.ToLower(note.DisplayName), "hdd") ||
						strings.Contains(strings.ToLower(note.Description), "hdd") ||
						strings.Contains(strings.ToLower(note.Description), "magnetic"))) {
					if hasNote {
						return opt, &note, nil
					}
					return opt, nil, nil
				}
			}
			chosen := usableOptions[0]
			if capa.DefaultStorageType != "" && !strings.EqualFold(capa.DefaultStorageType, "NA") {
				chosen = capa.DefaultStorageType
			}
			var chosenNote *rdbmsmodel.StorageTypeNote
			if note, ok := notesMap[strings.ToLower(chosen)]; ok {
				chosenNote = &note
			}
			warning := fmt.Sprintf("Storage type '%s' requested for instance '%s' is not directly offered on %s; recommended available storage type '%s' (available: %v).",
				req, instName, capa.ProviderName, chosen, usableOptions)
			*warnings = append(*warnings, warning)
			log.Info().
				Str("instance", instName).
				Str("csp", capa.ProviderName).
				Str("requestedStorageType", req).
				Str("recommendedStorageType", chosen).
				Msg("Recommended closest available storage type for target cloud")
			return chosen, chosenNote, nil
		}

		// 2.3 Requested storage type was not found: recommend closest available with warning instead of failing
		var chosen string
		var chosenNote *rdbmsmodel.StorageTypeNote
		for _, opt := range usableOptions {
			note, hasNote := notesMap[strings.ToLower(opt)]
			if hasNote && note.Recommended {
				chosen = opt
				chosenNote = &note
				break
			}
		}
		if chosen == "" {
			chosen = usableOptions[0]
			if note, ok := notesMap[strings.ToLower(chosen)]; ok {
				chosenNote = &note
			}
		}
		warning := fmt.Sprintf("Requested storageType '%s' for instance '%s' is not supported on target cloud (%s); recommended available storage type '%s' (available: %v).",
			req, instName, capa.ProviderName, chosen, usableOptions)
		*warnings = append(*warnings, warning)
		log.Info().
			Str("instance", instName).
			Str("csp", capa.ProviderName).
			Str("requestedStorageType", req).
			Str("recommendedStorageType", chosen).
			Msg("Recommended available storage type for unsupported request")
		return chosen, chosenNote, nil
	}

	// 3. If user did not request storage type, select recommended storage from notes
	for _, opt := range usableOptions {
		if note, ok := notesMap[strings.ToLower(opt)]; ok && note.Recommended {
			return opt, &note, nil
		}
	}

	// 4. Select defaultStorageType if provided by capability
	if capa.DefaultStorageType != "" && !strings.EqualFold(capa.DefaultStorageType, "NA") {
		if note, ok := notesMap[strings.ToLower(capa.DefaultStorageType)]; ok {
			return capa.DefaultStorageType, &note, nil
		}
		return capa.DefaultStorageType, nil, nil
	}

	// 5. Select first usable storage option from capability
	first := usableOptions[0]
	if note, ok := notesMap[strings.ToLower(first)]; ok {
		return first, &note, nil
	}
	return first, nil, nil
}

// dbSpecCandidate represents a scored DB instance spec candidate during proximity ranking.
type dbSpecCandidate struct {
	spec         rdbmsmodel.RDBMSDBInstanceSpecInfo
	vCpu         int
	memMb        int
	distance     float64
	satisfiesAll bool
}

// recommendDBInstanceSpec recommends the best-fitting DB instance spec dynamically from Tumblebug capability.
func recommendDBInstanceSpec(vcpu, memoryMb, storageSizeGb int, engine string, capa rdbmsmodel.RDBMSMetaInfo) (string, error) {
	engineLower := strings.ToLower(engine)

	// 1. If detailed DBInstanceSpecs list is available, evaluate requirements via proximity ranking
	// * Strategy: Conservative Capacity Proximity (Target >= Source for vCPU, RAM, and Disk)
	if len(capa.DBInstanceSpecs) > 0 {
		candidates := make([]dbSpecCandidate, 0, len(capa.DBInstanceSpecs))

		sourceVcpu := float64(vcpu)
		sourceMemGiB := float64(memoryMb) / 1024.0

		// Workload ratio (memory GiB / vCPU)
		ratio := 2.0 // default
		if sourceVcpu > 0 {
			ratio = sourceMemGiB / sourceVcpu
		}

		for _, spec := range capa.DBInstanceSpecs {
			specCpu, _ := strconv.Atoi(spec.VCpuCount)
			specMemMb, _ := strconv.Atoi(spec.MemSizeMiB)
			if specCpu <= 0 || specMemMb <= 0 {
				continue
			}

			specMemGiB := float64(specMemMb) / 1024.0

			vcpuDiff := math.Abs(float64(specCpu) - sourceVcpu)
			memGiBDiff := math.Abs(specMemGiB - sourceMemGiB)

			// Ratio-weighted Manhattan distance (inspired by resource-node-spec.go)
			var dist float64
			switch {
			case ratio <= 3.0: // compute-intensive (<= 3.0)
				dist = (vcpuDiff * 1.5) + memGiBDiff
			case ratio >= 7.0: // memory-intensive (>= 7.0)
				dist = vcpuDiff + (memGiBDiff * 1.5)
			default: // general-purpose
				dist = vcpuDiff + memGiBDiff
			}

			// Baseline constraint check: Target >= Source across vCPU, RAM, and Disk
			satisfiesCpu := (specCpu >= vcpu)
			satisfiesMem := (specMemMb >= memoryMb)
			satisfiesStorage := (spec.StorageSizeRangeGB.Max == 0 || storageSizeGb <= spec.StorageSizeRangeGB.Max)
			satisfiesAll := satisfiesCpu && satisfiesMem && satisfiesStorage

			candidates = append(candidates, dbSpecCandidate{
				spec:         spec,
				vCpu:         specCpu,
				memMb:        specMemMb,
				distance:     dist,
				satisfiesAll: satisfiesAll,
			})
		}

		if len(candidates) > 0 {
			// Sort candidates:
			// 1. Prioritize specs meeting all requirements (satisfiesAll: true before false)
			// 2. Minimum weighted distance
			// 3. Smaller resource excess (vCPU, then memory)
			// 4. Closer storage range ceiling (if both have Max specified) to avoid excessive overprovisioning
			// 5. Alphabetical spec name for deterministic tie-breaking
			sort.Slice(candidates, func(i, j int) bool {
				if candidates[i].satisfiesAll != candidates[j].satisfiesAll {
					return candidates[i].satisfiesAll
				}
				if math.Abs(candidates[i].distance-candidates[j].distance) > 0.001 {
					return candidates[i].distance < candidates[j].distance
				}
				if candidates[i].vCpu != candidates[j].vCpu {
					return candidates[i].vCpu < candidates[j].vCpu
				}
				if candidates[i].memMb != candidates[j].memMb {
					return candidates[i].memMb < candidates[j].memMb
				}
				if candidates[i].spec.StorageSizeRangeGB.Max > 0 && candidates[j].spec.StorageSizeRangeGB.Max > 0 {
					if candidates[i].spec.StorageSizeRangeGB.Max != candidates[j].spec.StorageSizeRangeGB.Max {
						return candidates[i].spec.StorageSizeRangeGB.Max < candidates[j].spec.StorageSizeRangeGB.Max
					}
				}
				return candidates[i].spec.Name < candidates[j].spec.Name
			})

			return candidates[0].spec.Name, nil
		}
	}

	// 2. Check engine-specific ReferenceDBInstanceSpec in DBMSRequirements from Tumblebug capability
	if req, ok := capa.DBMSRequirements[engineLower]; ok && req.ReferenceDBInstanceSpec != "" {
		return req.ReferenceDBInstanceSpec, nil
	}

	// 3. If string options list is provided from Tumblebug capability, return the first valid option
	for _, opt := range capa.DBInstanceSpecOptions {
		if strings.TrimSpace(opt) != "" && !strings.EqualFold(opt, "NA") {
			return opt, nil
		}
	}

	err := fmt.Errorf("no suitable DB instance spec found for engine '%s' (requested %d vCPU, %d MB RAM, %d GB storage)", engine, vcpu, memoryMb, storageSizeGb)
	log.Warn().Err(err).Str("engine", engine).Int("vcpu", vcpu).Int("memoryMb", memoryMb).Int("storageSizeGb", storageSizeGb).Msg("DBInstanceSpec recommendation failed")
	return "", err
}

// isDBEngineSupported checks if the target CSP supports the specified database engine.
func isDBEngineSupported(support rdbmsmodel.RDBMSCSPSupportInfo, hasSupport bool, engine string) bool {
	if !hasSupport || len(support.SupportedDBEngines) == 0 {
		return false
	}
	for _, eng := range support.SupportedDBEngines {
		if strings.EqualFold(eng, engine) {
			return true
		}
	}
	return false
}
