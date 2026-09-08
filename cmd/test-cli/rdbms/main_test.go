package main

import (
	"testing"

	"github.com/cloud-barista/cm-beetle/pkg/core/common"
)

func TestResolveBaseRDBMSName(t *testing.T) {
	tests := []struct {
		name         string
		tc           TestCase
		seed         string
		expectedBase string
		expectedId   string
	}{
		{
			name: "Case with pre-existing test prefix in rdbmsId",
			tc: TestCase{
				Csp:     "openstack",
				RdbmsId: "test-rdbms-openstack",
			},
			seed:         "test",
			expectedBase: "rdbms-openstack",
			expectedId:   "test-rdbms-openstack",
		},
		{
			name: "Case with base name in rdbmsId without seed prefix",
			tc: TestCase{
				Csp:     "openstack",
				RdbmsId: "rdbms-openstack",
			},
			seed:         "test",
			expectedBase: "rdbms-openstack",
			expectedId:   "test-rdbms-openstack",
		},
		{
			name: "Case with empty rdbmsId (default fallback to CSP)",
			tc: TestCase{
				Csp:     "openstack",
				RdbmsId: "",
			},
			seed:         "test",
			expectedBase: "rdbms-openstack",
			expectedId:   "test-rdbms-openstack",
		},
		{
			name: "Case with empty seed",
			tc: TestCase{
				Csp:     "openstack",
				RdbmsId: "test-rdbms-openstack",
			},
			seed:         "",
			expectedBase: "test-rdbms-openstack",
			expectedId:   "test-rdbms-openstack",
		},
		{
			name: "Case with MariaDB engine and empty rdbmsId",
			tc: TestCase{
				Csp:      "openstack",
				DBEngine: "mariadb",
				RdbmsId:  "",
			},
			seed:         "test",
			expectedBase: "rdbms-mariadb-openstack",
			expectedId:   "test-rdbms-mariadb-openstack",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			base := resolveBaseRDBMSName(tt.tc, tt.seed)
			if base != tt.expectedBase {
				t.Errorf("resolveBaseRDBMSName() base = %v, want %v", base, tt.expectedBase)
			}
			composed := common.ComposeName(base, tt.seed)
			if composed != tt.expectedId {
				t.Errorf("common.ComposeName() composed = %v, want %v", composed, tt.expectedId)
			}
		})
	}
}

func TestParallelMultiCSPNamingUniqueness(t *testing.T) {
	csps := []string{"aws", "azure", "gcp", "alibaba", "tencent", "ibm", "ncp", "nhn", "openstack"}
	seed := "test"
	seen := make(map[string]bool)

	for _, csp := range csps {
		tc := TestCase{
			Csp:     csp,
			RdbmsId: "test-rdbms-" + csp,
		}
		base := resolveBaseRDBMSName(tc, seed)
		composed := common.ComposeName(base, seed)

		if seen[composed] {
			t.Fatalf("Collision detected for CSP %s: %s already used", csp, composed)
		}
		seen[composed] = true

		expected := "test-rdbms-" + csp
		if composed != expected {
			t.Errorf("CSP %s composed name = %s, want %s", csp, composed, expected)
		}
	}
}
