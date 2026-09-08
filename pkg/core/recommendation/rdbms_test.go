package recommendation

import (
	"strings"
	"testing"
	"time"

	rdbmsmodel "github.com/cloud-barista/cm-beetle/imdl/rdbms-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
)

func init() {
	tbclient.Init(tbclient.ApiConfig{
		RestUrl:  "http://127.0.0.1:1323/tumblebug",
		Username: "default",
		Password: "default",
		Timeout:  100 * time.Millisecond,
	})
}

func TestValidateSourceRDBMS(t *testing.T) {
	tests := []struct {
		name        string
		sources     []rdbmsmodel.SourceRDBMSProperty
		expectErr   bool
		errContains string
	}{
		{
			name:        "Empty sources list",
			sources:     []rdbmsmodel.SourceRDBMSProperty{},
			expectErr:   true,
			errContains: "at least one source RDBMS instance is required",
		},
		{
			name: "Missing InstanceName",
			sources: []rdbmsmodel.SourceRDBMSProperty{
				{
					InstanceName:  "  ",
					Engine:        "mysql",
					EngineVersion: "8.0",
					Vcpu:          2,
					MemoryMb:      4096,
					StorageSizeGb: 50,
				},
			},
			expectErr:   true,
			errContains: "instanceName is required",
		},
		{
			name: "Missing Engine",
			sources: []rdbmsmodel.SourceRDBMSProperty{
				{
					InstanceName:  "db-01",
					Engine:        "",
					EngineVersion: "8.0",
					Vcpu:          2,
					MemoryMb:      4096,
					StorageSizeGb: 50,
				},
			},
			expectErr:   true,
			errContains: "engine is required",
		},
		{
			name: "Missing EngineVersion",
			sources: []rdbmsmodel.SourceRDBMSProperty{
				{
					InstanceName:  "db-01",
					Engine:        "mysql",
					EngineVersion: "  ",
					Vcpu:          2,
					MemoryMb:      4096,
					StorageSizeGb: 50,
				},
			},
			expectErr:   true,
			errContains: "engineVersion is required",
		},
		{
			name: "Non-positive vCPU",
			sources: []rdbmsmodel.SourceRDBMSProperty{
				{
					InstanceName:  "db-01",
					Engine:        "mysql",
					EngineVersion: "8.0",
					Vcpu:          0,
					MemoryMb:      4096,
					StorageSizeGb: 50,
				},
			},
			expectErr:   true,
			errContains: "vcpu must be greater than 0",
		},
		{
			name: "Non-positive MemoryMb",
			sources: []rdbmsmodel.SourceRDBMSProperty{
				{
					InstanceName:  "db-01",
					Engine:        "mysql",
					EngineVersion: "8.0",
					Vcpu:          2,
					MemoryMb:      -512,
					StorageSizeGb: 50,
				},
			},
			expectErr:   true,
			errContains: "memoryMb must be greater than 0",
		},
		{
			name: "Non-positive StorageSizeGb",
			sources: []rdbmsmodel.SourceRDBMSProperty{
				{
					InstanceName:  "db-01",
					Engine:        "mysql",
					EngineVersion: "8.0",
					Vcpu:          2,
					MemoryMb:      4096,
					StorageSizeGb: 0,
				},
			},
			expectErr:   true,
			errContains: "storageSizeGb must be greater than 0",
		},
		{
			name: "Missing DatabaseName in inner database list",
			sources: []rdbmsmodel.SourceRDBMSProperty{
				{
					InstanceName:  "db-01",
					Engine:        "mysql",
					EngineVersion: "8.0",
					Vcpu:          2,
					MemoryMb:      4096,
					StorageSizeGb: 50,
					Databases: []rdbmsmodel.SourceDatabaseProperty{
						{DatabaseName: ""},
					},
				},
			},
			expectErr:   true,
			errContains: "databaseName is required",
		},
		{
			name: "Valid Source Instance",
			sources: []rdbmsmodel.SourceRDBMSProperty{
				{
					InstanceName:  "db-01",
					Engine:        "mysql",
					EngineVersion: "8.0",
					Vcpu:          4,
					MemoryMb:      8192,
					StorageSizeGb: 100,
					Databases: []rdbmsmodel.SourceDatabaseProperty{
						{DatabaseName: "app_db"},
					},
				},
			},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSourceRDBMS(tt.sources)
			if tt.expectErr {
				if err == nil {
					t.Fatalf("expected error containing '%s', got nil", tt.errContains)
				}
				if !strings.Contains(err.Error(), tt.errContains) {
					t.Fatalf("expected error containing '%s', got '%s'", tt.errContains, err.Error())
				}
			} else if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestRecommendRDBMS_ValidationFailures(t *testing.T) {
	validSource := []rdbmsmodel.SourceRDBMSProperty{
		{
			InstanceName:  "src-01",
			Engine:        "mysql",
			EngineVersion: "8.0",
			Vcpu:          2,
			MemoryMb:      4096,
			StorageSizeGb: 50,
		},
	}

	// Missing CSP
	_, err := RecommendRDBMS("", "ap-northeast-2", validSource)
	if err == nil || !strings.Contains(err.Error(), "desiredCsp and desiredRegion are required") {
		t.Errorf("expected error for missing CSP, got %v", err)
	}

	// Missing Region
	_, err = RecommendRDBMS("aws", "", validSource)
	if err == nil || !strings.Contains(err.Error(), "desiredCsp and desiredRegion are required") {
		t.Errorf("expected error for missing Region, got %v", err)
	}

	// Invalid Source (empty sources)
	_, err = RecommendRDBMS("aws", "ap-northeast-2", []rdbmsmodel.SourceRDBMSProperty{})
	if err == nil || !strings.Contains(err.Error(), "at least one source RDBMS instance is required") {
		t.Errorf("expected error for empty sources, got %v", err)
	}
}

func TestRecommendDBInstanceSpec_Proximity(t *testing.T) {
	capa := rdbmsmodel.RDBMSMetaInfo{
		DBInstanceSpecs: []rdbmsmodel.RDBMSDBInstanceSpecInfo{
			{Name: "db.t3.small", VCpuCount: "2", MemSizeMiB: "2048"},
			{Name: "db.t3.medium", VCpuCount: "2", MemSizeMiB: "4096"},
			{Name: "db.t3.large", VCpuCount: "2", MemSizeMiB: "8192"},
			{Name: "db.t3.xlarge", VCpuCount: "4", MemSizeMiB: "16384"},
			{Name: "db.m5.large", VCpuCount: "2", MemSizeMiB: "8192"},
			{Name: "db.c5.xlarge", VCpuCount: "4", MemSizeMiB: "8192"},
		},
	}

	// 1. Exact match (2 vCPU, 4096 MiB, 100 GB) -> should pick db.t3.medium
	spec, err := recommendDBInstanceSpec(2, 4096, 100, "mysql", capa)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if spec != "db.t3.medium" {
		t.Errorf("expected 'db.t3.medium', got '%s'", spec)
	}

	// 2. Compute-intensive workload (4 vCPU, 4096 MiB = ratio 1.0) -> should favor vCPU
	spec, err = recommendDBInstanceSpec(4, 4096, 100, "mysql", capa)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if spec != "db.c5.xlarge" {
		t.Errorf("expected 'db.c5.xlarge', got '%s'", spec)
	}

	// 3. Exact larger match (4 vCPU, 16384 MiB) -> should pick db.t3.xlarge
	spec, err = recommendDBInstanceSpec(4, 16384, 100, "mysql", capa)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if spec != "db.t3.xlarge" {
		t.Errorf("expected 'db.t3.xlarge', got '%s'", spec)
	}

	// 4. Empty specs list with no fallback should return error
	emptyCapa := rdbmsmodel.RDBMSMetaInfo{
		DBInstanceSpecs: []rdbmsmodel.RDBMSDBInstanceSpecInfo{},
	}
	_, err = recommendDBInstanceSpec(2, 4096, 100, "mysql", emptyCapa)
	if err == nil {
		t.Fatalf("expected error when DBInstanceSpecs is empty, got nil")
	}
}

func TestRecommendDBInstanceSpec_DiskAware(t *testing.T) {
	capa := rdbmsmodel.RDBMSMetaInfo{
		DBInstanceSpecs: []rdbmsmodel.RDBMSDBInstanceSpecInfo{
			{
				Name:               "db.t3.medium.limited_disk",
				VCpuCount:          "2",
				MemSizeMiB:         "4096",
				StorageSizeRangeGB: rdbmsmodel.StorageSizeRange{Min: 10, Max: 50},
			},
			{
				Name:               "db.t3.medium.standard_disk",
				VCpuCount:          "2",
				MemSizeMiB:         "4096",
				StorageSizeRangeGB: rdbmsmodel.StorageSizeRange{Min: 10, Max: 1000},
			},
			{
				Name:               "db.m5.large",
				VCpuCount:          "2",
				MemSizeMiB:         "8192",
				StorageSizeRangeGB: rdbmsmodel.StorageSizeRange{Min: 10, Max: 5000},
			},
		},
	}

	// 1. Source database requires 100 GB storage with 2 vCPU, 4096 MiB RAM.
	// db.t3.medium.limited_disk only supports up to 50 GB, so db.t3.medium.standard_disk should be selected.
	spec, err := recommendDBInstanceSpec(2, 4096, 100, "mysql", capa)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if spec != "db.t3.medium.standard_disk" {
		t.Errorf("expected 'db.t3.medium.standard_disk', got '%s'", spec)
	}

	// 2. Source database requires 2000 GB storage with 2 vCPU, 4096 MiB RAM.
	// Both db.t3.medium specs have Max <= 1000, so db.m5.large should be selected to satisfy disk.
	spec, err = recommendDBInstanceSpec(2, 4096, 2000, "mysql", capa)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if spec != "db.m5.large" {
		t.Errorf("expected 'db.m5.large', got '%s'", spec)
	}
}

func TestSelectEngineVersion_Proximity(t *testing.T) {
	supported := []string{"8.0.32", "8.0.35", "8.4.0"}
	warnings := make([]string, 0)

	// Exact match
	v, err := selectEngineVersion("mysql", "8.0.32", supported, &warnings, "test-inst")
	if err != nil || v != "8.0.32" {
		t.Errorf("expected '8.0.32', got '%s', err: %v", v, err)
	}

	// Prefix match
	v, err = selectEngineVersion("mysql", "8.4", supported, &warnings, "test-inst")
	if err != nil || v != "8.4.0" {
		t.Errorf("expected '8.4.0', got '%s', err: %v", v, err)
	}

	// IBM Scenario: Requested 8.0 on target cloud offering only [8.4.0] -> recommends closest version 8.4.0 with warning
	ibmSupported := []string{"8.4.0"}
	warnings = nil
	v, err = selectEngineVersion("mysql", "8.0", ibmSupported, &warnings, "test-inst")
	if err != nil || v != "8.4.0" {
		t.Errorf("expected closest version '8.4.0', got '%s', err: %v", v, err)
	}
	if len(warnings) == 0 || !strings.Contains(warnings[0], "recommended closest available version '8.4.0'") {
		t.Errorf("expected warning about recommending closest version, got: %v", warnings)
	}

	// OpenStack Scenario: Requested 8.0 on target cloud offering only [5.7.29] -> recommends 5.7.29 with warning
	openstackSupported := []string{"5.7.29"}
	warnings = nil
	v, err = selectEngineVersion("mysql", "8.0", openstackSupported, &warnings, "test-inst")
	if err != nil || v != "5.7.29" {
		t.Errorf("expected fallback version '5.7.29', got '%s', err: %v", v, err)
	}
	if len(warnings) == 0 || !strings.Contains(warnings[0], "recommended closest available version '5.7.29'") {
		t.Errorf("expected warning about recommending closest version, got: %v", warnings)
	}
}

func TestIsDBEngineSupported(t *testing.T) {
	support := rdbmsmodel.RDBMSCSPSupportInfo{
		Supported:          true,
		SupportedDBEngines: []string{"mysql", "mariadb"},
	}

	if !isDBEngineSupported(support, true, "mysql") {
		t.Errorf("expected mysql to be supported")
	}
	if !isDBEngineSupported(support, true, "mariadb") {
		t.Errorf("expected mariadb to be supported")
	}
	if isDBEngineSupported(support, true, "postgresql") {
		t.Errorf("expected postgresql to be unsupported")
	}

	unsupportedCSP := rdbmsmodel.RDBMSCSPSupportInfo{
		Supported:          true,
		SupportedDBEngines: []string{"mysql"},
	}
	if isDBEngineSupported(unsupportedCSP, true, "mariadb") {
		t.Errorf("expected mariadb to be unsupported when not in SupportedDBEngines")
	}
}

func TestSelectStorageType_NoFallback(t *testing.T) {
	capa := rdbmsmodel.RDBMSMetaInfo{
		SupportsStorageTypeSelection: true,
		StorageTypeOptions:           []string{"gp2", "gp3", "io1"},
		Notes: &rdbmsmodel.RDBMSNotes{
			StorageTypes: []rdbmsmodel.StorageTypeNote{
				{StorageType: "gp3", Recommended: true},
				{StorageType: "gp2", Recommended: false},
			},
		},
	}

	var warnings []string

	// 1. Exact match
	st, note, err := selectStorageType("gp2", capa, &warnings, "inst-01")
	if err != nil || st != "gp2" {
		t.Errorf("expected 'gp2', got '%s', err: %v", st, err)
	}

	// 2. Generic SSD mapping to recommended SSD (gp3)
	st, note, err = selectStorageType("SSD", capa, &warnings, "inst-01")
	if err != nil || st != "gp3" || note == nil || !note.Recommended {
		t.Errorf("expected 'gp3', got '%s', err: %v", st, err)
	}

	// 3. Unsupported specific storage type -> Recommends available with warning
	warnings = nil
	st, _, err = selectStorageType("non_existent_storage", capa, &warnings, "inst-01")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if st != "gp3" {
		t.Errorf("expected fallback to recommended 'gp3', got '%s'", st)
	}
	if len(warnings) == 0 || !strings.Contains(warnings[0], "not supported on target cloud") {
		t.Errorf("expected warning for unsupported storage type, got: %v", warnings)
	}

	// 4. Unselectable storage type on CSP (e.g. Azure/NCP) -> returns "" and warns if requested
	unselectableCapa := rdbmsmodel.RDBMSMetaInfo{
		ProviderName:                 "azure",
		SupportsStorageTypeSelection: false,
		StorageTypeOptions:           []string{"auto"},
	}
	st, _, err = selectStorageType("SSD", unselectableCapa, &warnings, "inst-01")
	if err != nil || st != "" {
		t.Errorf("expected empty string for unselectable storage, got '%s', err: %v", st, err)
	}

	// 5. Target cloud with non-SSD named storage (e.g. OpenStack [RBD, __DEFAULT__])
	// Generic "SSD" request should map to "RBD" with an informative warning
	openstackCapa := rdbmsmodel.RDBMSMetaInfo{
		ProviderName:                 "openstack",
		SupportsStorageTypeSelection: true,
		StorageTypeOptions:           []string{"__DEFAULT__", "RBD"},
	}
	st, _, err = selectStorageType("SSD", openstackCapa, &warnings, "inst-01")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if st != "RBD" {
		t.Errorf("expected 'RBD', got '%s'", st)
	}
	if len(warnings) == 0 || !strings.Contains(warnings[len(warnings)-1], "recommended closest available storage type 'RBD'") {
		t.Errorf("expected warning about mapping SSD to RBD, got warnings: %v", warnings)
	}

	// 6. Alibaba scenario: options contain [cloud_auto, cloud_essd, cloud_essd2, cloud_essd3, local_ssd]
	// cloud_auto is Recommended. Generic "SSD" request MUST pick "cloud_auto" and NOT "cloud_essd2" (premium)
	alibabaCapa := rdbmsmodel.RDBMSMetaInfo{
		ProviderName:                 "alibaba",
		SupportsStorageTypeSelection: true,
		StorageTypeOptions:           []string{"cloud_auto", "cloud_essd", "cloud_essd2", "cloud_essd3", "local_ssd"},
		Notes: &rdbmsmodel.RDBMSNotes{
			StorageTypes: []rdbmsmodel.StorageTypeNote{
				{StorageType: "cloud_auto", DisplayName: "Auto-selected Storage Type (SSD)", Recommended: true},
				{StorageType: "cloud_essd", DisplayName: "Enhanced SSD (ESSD PL1)", RecommendationLevel: "standard"},
				{StorageType: "cloud_essd2", DisplayName: "Enhanced SSD (ESSD PL2)", RecommendationLevel: "premium", MinSize: 500},
				{StorageType: "cloud_essd3", DisplayName: "Enhanced SSD (ESSD PL3)", RecommendationLevel: "premium", MinSize: 1500},
			},
		},
	}
	st, note, err = selectStorageType("SSD", alibabaCapa, &warnings, "inst-01")
	if err != nil {
		t.Fatalf("unexpected error for alibaba: %v", err)
	}
	if st != "cloud_auto" {
		t.Errorf("expected 'cloud_auto', got '%s'", st)
	}
	if note == nil || !note.Recommended {
		t.Errorf("expected note to be recommended")
	}
}
