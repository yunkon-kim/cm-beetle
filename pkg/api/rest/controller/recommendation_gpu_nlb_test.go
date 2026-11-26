package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupControllerMockTumblebug(t *testing.T) {
	t.Helper()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/provider/"):
			_ = json.NewEncoder(w).Encode(tbmodel.RegionDetail{
				RegionId:   "ap-northeast-2",
				RegionName: "ap-northeast-2",
				Zones:      []string{"ap-northeast-2a", "ap-northeast-2c"},
			})
		case r.Method == http.MethodPost && r.URL.Path == "/recommendSpec":
			bodyBytes := make([]byte, 0, 4096)
			buf := new(bytes.Buffer)
			_, _ = buf.ReadFrom(r.Body)
			bodyBytes = buf.Bytes()
			bodyStr := string(bodyBytes)

			var specs []tbmodel.SpecInfo
			if strings.Contains(bodyStr, "acceleratorCount") {
				specs = []tbmodel.SpecInfo{
					{
						Id:                  "aws+ap-northeast-2+g4dn.2xlarge",
						Name:                "g4dn.2xlarge",
						CspSpecName:         "g4dn.2xlarge",
						ProviderName:        "aws",
						RegionName:          "ap-northeast-2",
						Architecture:        "x86_64",
						VCPU:                8,
						MemoryGiB:           32,
						AcceleratorCount:    1,
						AcceleratorMemoryGB: 16,
						AcceleratorModel:    "NVIDIA T4",
						CostPerHour:         0.752,
						AvailabilityStatus:  "Available",
						AvailableZones:      []string{"ap-northeast-2a"},
					},
					{
						Id:                  "aws+ap-northeast-2+g4ad.2xlarge",
						Name:                "g4ad.2xlarge",
						CspSpecName:         "g4ad.2xlarge",
						ProviderName:        "aws",
						RegionName:          "ap-northeast-2",
						Architecture:        "x86_64",
						VCPU:                8,
						MemoryGiB:           32,
						AcceleratorCount:    1,
						AcceleratorMemoryGB: 8,
						AcceleratorModel:    "AMD Radeon Pro V520",
						CostPerHour:         0.379,
						AvailabilityStatus:  "Available",
						AvailableZones:      []string{"ap-northeast-2a"},
					},
				}
			} else {
				specs = []tbmodel.SpecInfo{
					{
						Id:                 "aws+ap-northeast-2+m6i.xlarge",
						Name:               "m6i.xlarge",
						CspSpecName:        "m6i.xlarge",
						ProviderName:       "aws",
						RegionName:         "ap-northeast-2",
						Architecture:       "x86_64",
						VCPU:               4,
						MemoryGiB:          16,
						CostPerHour:        0.192,
						AvailabilityStatus: "Available",
						AvailableZones:     []string{"ap-northeast-2a"},
					},
				}
			}
			_ = json.NewEncoder(w).Encode(specs)
		case r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/searchImage"):
			_ = json.NewEncoder(w).Encode(tbmodel.SearchImageResponse{
				ImageCount: 1,
				ImageList: []tbmodel.ImageInfo{
					{
						Id:             "aws+ap-northeast-2+ubuntu-2204",
						Name:           "ubuntu-22.04",
						ProviderName:   "aws",
						CspImageName:   "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server",
						OSType:         "ubuntu 22.04",
						OSArchitecture: "x86_64",
						OSDistribution: "Ubuntu 22.04 LTS",
						OSDiskType:     "SSD",
					},
				},
			})
		case r.Method == http.MethodPost && r.URL.Path == "/specImagePairReview":
			var req tbmodel.SpecImagePairReviewReq
			_ = json.NewDecoder(r.Body).Decode(&req)
			_ = json.NewEncoder(w).Encode(tbmodel.SpecImagePairReviewResult{
				IsValid: true,
				Status:  "OK",
				SpecId:  req.SpecId,
				ImageId: req.ImageId,
				ImageValidation: tbmodel.ReviewResourceValidation{
					ResourceId:    req.ImageId,
					IsAvailable:   true,
					Status:        "Available",
					CspResourceId: req.ImageId + "-latest",
				},
				SuggestedSystemDisk: "gp3",
			})
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))

	restore := tbclient.SetClient(tbclient.NewClient(tbclient.ApiConfig{
		RestUrl:  srv.URL,
		Username: "default",
		Password: "default",
	}))
	t.Cleanup(func() {
		restore()
		srv.Close()
	})
}

func sampleCrossVendorOnpremInfra() onpremmodel.OnpremInfra {
	return onpremmodel.OnpremInfra{
		Nodes: []onpremmodel.NodeProperty{
			{
				Hostname:  "nvidia-1",
				MachineId: "nvidia-1",
				CPU:       onpremmodel.CpuProperty{Architecture: "x86_64", Cpus: 8, Cores: 8, Threads: 1},
				Memory:    onpremmodel.MemoryProperty{Type: "DDR4", TotalSize: 32},
				RootDisk:  onpremmodel.DiskProperty{Label: "boot", Type: "SSD", TotalSize: 100},
				Interfaces: []onpremmodel.NetworkInterfaceProperty{
					{Name: "eth0", IPv4CidrBlocks: []string{"10.0.10.1/24"}},
				},
				OS:       onpremmodel.OsProperty{PrettyName: "Ubuntu 22.04", ID: "ubuntu", VersionID: "22.04"},
				GPUCards: []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "NVIDIA T4", MemoryTotalGB: 16}},
			},
			{
				Hostname:  "amd-1",
				MachineId: "amd-1",
				CPU:       onpremmodel.CpuProperty{Architecture: "x86_64", Cpus: 8, Cores: 8, Threads: 1},
				Memory:    onpremmodel.MemoryProperty{Type: "DDR4", TotalSize: 32},
				RootDisk:  onpremmodel.DiskProperty{Label: "boot", Type: "SSD", TotalSize: 100},
				Interfaces: []onpremmodel.NetworkInterfaceProperty{
					{Name: "eth0", IPv4CidrBlocks: []string{"10.0.10.2/24"}},
				},
				OS:       onpremmodel.OsProperty{PrettyName: "Ubuntu 22.04", ID: "ubuntu", VersionID: "22.04"},
				GPUCards: []onpremmodel.GpuCardProperty{{Vendor: "AMD", Model: "Radeon Pro V520", MemoryTotalGB: 8}},
			},
		},
		NLBs: []onpremmodel.NlbProperty{
			{
				Software: "haproxy",
				Listener: onpremmodel.NlbListenerProperty{BindAddress: "*", Port: 8000, Protocol: "tcp"},
				Backend: onpremmodel.NlbBackendProperty{
					Name: "infer-pool",
					Servers: []onpremmodel.NlbServerProperty{
						{Name: "n1", IP: "10.0.10.1", Port: 8000},
						{Name: "a1", IP: "10.0.10.2", Port: 8000},
					},
				},
			},
		},
	}
}

func TestRecommendInfraCandidates_InvalidSizingPolicy_Returns400(t *testing.T) {
	e := echo.New()
	body, err := json.Marshal(RecommendInfraRequest{
		DesiredCspAndRegionPair: cloudmodel.CloudProperty{Csp: "aws", Region: "ap-northeast-2"},
		OnpremiseInfraModel:     sampleCrossVendorOnpremInfra(),
	})
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/beetle/recommendation/infra?sizingPolicy=invalidMode", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	require.NoError(t, RecommendInfraCandidates(e.NewContext(req, rec)))
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	nlbBody, err := json.Marshal(RecommendInfraWithNlbRequest{
		DesiredCsp:  "aws",
		DesiredRegion: "ap-northeast-2",
		SourceInfra: sampleCrossVendorOnpremInfra(),
	})
	require.NoError(t, err)

	req2 := httptest.NewRequest(http.MethodPost, "/beetle/recommendation/infraWithNlb?sizingPolicy=bad", bytes.NewReader(nlbBody))
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	require.NoError(t, RecommendInfraWithNlbCandidates(e.NewContext(req2, rec2)))
	assert.Equal(t, http.StatusBadRequest, rec2.Code)

	multiBody, err := json.Marshal(RecommendMultiInfraWithNlbRequest{
		DesiredCspAndRegionPairs: []cloudmodel.CloudProperty{
			{Csp: "aws", Region: "ap-northeast-2"},
			{Csp: "azure", Region: "koreacentral"},
		},
		SourceInfra: sampleCrossVendorOnpremInfra(),
	})
	require.NoError(t, err)

	req3 := httptest.NewRequest(http.MethodPost, "/beetle/recommendation/multiInfraWithNlb?sizingPolicy=bad", bytes.NewReader(multiBody))
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	require.NoError(t, RecommendMultiInfraWithNlbCandidates(e.NewContext(req3, rec3)))
	assert.Equal(t, http.StatusBadRequest, rec3.Code)
}

func TestRecommendInfraCandidates_UnifiedDispatch_CrossVendor422AndAuto200(t *testing.T) {
	setupControllerMockTumblebug(t)
	e := echo.New()

	body, err := json.Marshal(RecommendInfraRequest{
		DesiredCspAndRegionPair: cloudmodel.CloudProperty{Csp: "aws", Region: "ap-northeast-2"},
		OnpremiseInfraModel:     sampleCrossVendorOnpremInfra(),
	})
	require.NoError(t, err)

	// Default rightSizeUp on cross-vendor GPU+NLB returns HTTP 422
	req422 := httptest.NewRequest(http.MethodPost, "/beetle/recommendation/infra", bytes.NewReader(body))
	req422.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec422 := httptest.NewRecorder()
	require.NoError(t, RecommendInfraCandidates(e.NewContext(req422, rec422)))
	assert.Equal(t, http.StatusUnprocessableEntity, rec422.Code)

	// Unified endpoint with sizingPolicy=auto automatically resolves into 2 vendor-divided NodeGroups and 2 NLBs
	reqAuto := httptest.NewRequest(http.MethodPost, "/beetle/recommendation/infra?sizingPolicy=auto", bytes.NewReader(body))
	reqAuto.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recAuto := httptest.NewRecorder()
	require.NoError(t, RecommendInfraCandidates(e.NewContext(reqAuto, recAuto)))
	require.Equal(t, http.StatusOK, recAuto.Code)

	var resp model.ApiResponse[[]cloudmodel.RecommendedInfra]
	require.NoError(t, json.Unmarshal(recAuto.Body.Bytes(), &resp))
	assert.True(t, resp.Success)
	require.NotEmpty(t, resp.Data)
	assert.Len(t, resp.Data[0].TargetInfra.NodeGroups, 2)
	assert.Len(t, resp.Data[0].TargetNlbList, 2)
}
