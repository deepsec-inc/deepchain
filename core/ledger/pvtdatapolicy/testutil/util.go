/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package testutil

import (
	"deepchain/core/ledger/pvtdatapolicy"
	"deepchain/core/ledger/pvtdatapolicy/mock"
	"deepchain/protos/common"
)

// SampleBTLPolicy helps tests create a sample BTLPolicy
// The example input entry is [2]string{ns, coll}:btl
func SampleBTLPolicy(m map[[2]string]uint64) pvtdatapolicy.BTLPolicy {
	ccInfoRetriever := &mock.CollectionInfoProvider{}
	ccInfoRetriever.CollectionInfoStub = func(ccName, collName string) (*common.StaticCollectionConfig, error) {
		btl := m[[2]string{ccName, collName}]
		return &common.StaticCollectionConfig{BlockToLive: btl}, nil
	}
	return pvtdatapolicy.ConstructBTLPolicy(ccInfoRetriever)
}
