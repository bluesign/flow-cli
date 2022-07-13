/*
 * Flow CLI
 *
 * Copyright 2019 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package multisig

import (
	"bytes"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/onflow/flow-cli/pkg/flowkit/output"
	"github.com/onflow/flow-cli/pkg/flowkit/util"
	"github.com/onflow/flow-go-sdk"
)

var Cmd = &cobra.Command{
	Use:   "multisig",
	Short: "Multisig Utilities",
}

func init() {
	SignCommand.AddToParent(Cmd)
}

type MultisigResult struct {
	tx     *flow.Transaction
	rlp    string
	signed string
}

func (k *MultisigResult) JSON() interface{} {
	result := make(map[string]string)
	result["rlp"] = k.rlp
	result["signed"] = k.signed

	return result
}

func (k *MultisigResult) String() string {
	var b bytes.Buffer
	writer := util.CreateTabWriter(&b)
	if k.rlp != "" {
		_, _ = fmt.Fprintf(writer, "%s RLP retrieved successfully\n", output.TryEmoji())
	}

	if k.signed != "" {
		_, _ = fmt.Fprintf(writer, "%s Signed RLP Posted successfully\n", output.SuccessEmoji())
	}

	_ = writer.Flush()
	return b.String()
}

func (k *MultisigResult) Oneliner() string {
	result := fmt.Sprintf("Done: %v", k.signed != "")

	return result
}
