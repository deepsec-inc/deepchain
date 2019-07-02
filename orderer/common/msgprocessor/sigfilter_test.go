/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msgprocessor

import (
	"fmt"
	"testing"

	"deepchain/common/flogging"
	mockchannelconfig "deepchain/common/mocks/config"
	mockpolicies "deepchain/common/mocks/policies"
	cb "deepchain/protos/common"
	"deepchain/protos/utils"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func init() {
	flogging.ActivateSpec("orderer.common.msgprocessor=DEBUG")
}

func makeEnvelope() *cb.Envelope {
	return &cb.Envelope{
		Payload: utils.MarshalOrPanic(&cb.Payload{
			Header: &cb.Header{
				SignatureHeader: utils.MarshalOrPanic(&cb.SignatureHeader{}),
			},
		}),
	}
}

func TestAccept(t *testing.T) {
	mpm := &mockchannelconfig.Resources{
		PolicyManagerVal: &mockpolicies.Manager{Policy: &mockpolicies.Policy{}},
	}
	assert.Nil(t, NewSigFilter("foo", mpm).Apply(makeEnvelope()), "Valid envelope and good policy")
}

func TestMissingPolicy(t *testing.T) {
	mpm := &mockchannelconfig.Resources{
		PolicyManagerVal: &mockpolicies.Manager{},
	}
	err := NewSigFilter("foo", mpm).Apply(makeEnvelope())
	assert.NotNil(t, err)
	assert.Regexp(t, "could not find policy", err.Error())
}

func TestEmptyPayload(t *testing.T) {
	mpm := &mockchannelconfig.Resources{
		PolicyManagerVal: &mockpolicies.Manager{Policy: &mockpolicies.Policy{}},
	}
	err := NewSigFilter("foo", mpm).Apply(&cb.Envelope{})
	assert.NotNil(t, err)
	assert.Regexp(t, "could not convert message to signedData", err.Error())
}

func TestErrorOnPolicy(t *testing.T) {
	mpm := &mockchannelconfig.Resources{
		PolicyManagerVal: &mockpolicies.Manager{Policy: &mockpolicies.Policy{Err: fmt.Errorf("Error")}},
	}
	err := NewSigFilter("foo", mpm).Apply(makeEnvelope())
	assert.NotNil(t, err)
	assert.Equal(t, ErrPermissionDenied, errors.Cause(err))
}
