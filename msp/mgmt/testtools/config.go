/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msptesttools

import (
	"deepchain/common/util"
	"deepchain/core/config/configtest"
	"deepchain/msp"
	"deepchain/msp/mgmt"
)

// LoadTestMSPSetup sets up the local MSP
// and a chain MSP for the default chain
func LoadMSPSetupForTesting() error {
	dir, err := configtest.GetDevMspDir()
	if err != nil {
		return err
	}
	conf, err := msp.GetLocalMspConfig(dir, nil, "SampleOrg")
	if err != nil {
		return err
	}

	err = mgmt.GetLocalMSP().Setup(conf)
	if err != nil {
		return err
	}

	err = mgmt.GetManagerForChain(util.GetTestChainID()).Setup([]msp.MSP{mgmt.GetLocalMSP()})
	if err != nil {
		return err
	}

	return nil
}

// Loads the development local MSP for use in testing.  Not valid for production/runtime context
func LoadDevMsp() error {
	mspDir, err := configtest.GetDevMspDir()
	if err != nil {
		return err
	}

	return mgmt.LoadLocalMsp(mspDir, nil, "SampleOrg")
}
