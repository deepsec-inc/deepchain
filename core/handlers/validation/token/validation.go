/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package token

import (
	"deepchain/core/handlers/validation/api"
	"deepchain/protos/common"
)

type ValidationFactory struct {
}

func (*ValidationFactory) New() validation.Plugin {
	return &ValidationPlugin{}
}

type ValidationPlugin struct {
}

func (v *ValidationPlugin) Init(dependencies ...validation.Dependency) error {
	return nil
}

func (v *ValidationPlugin) Validate(block *common.Block, namespace string, txPosition int, actionPosition int, contextData ...validation.ContextDatum) error {
	return nil
}
