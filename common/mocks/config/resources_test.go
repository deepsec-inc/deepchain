/*
Copyright IBM Corp. 2016 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"testing"

	"deepchain/common/channelconfig"
)

func TestConfigtxResourcesInterface(t *testing.T) {
	_ = channelconfig.Resources(&Resources{})
}
