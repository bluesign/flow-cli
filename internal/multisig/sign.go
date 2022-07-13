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
	"io/ioutil"
	"net/http"

	"github.com/onflow/flow-cli/pkg/flowkit"

	"github.com/spf13/cobra"

	"github.com/onflow/flow-cli/internal/command"
	"github.com/onflow/flow-cli/pkg/flowkit/services"
	"github.com/onflow/flow-cli/pkg/flowkit/util"
)

type flagsSign struct {
	Signer  string   `default:"emulator-account" flag:"signer" info:"name of the account used to sign"`
	Include []string `default:"" flag:"include" info:"Fields to include in the output. Valid values: signatures, code, payload."`
}

var signFlags = flagsSign{}

var SignCommand = &command.Command{
	Cmd: &cobra.Command{
		Use:     "sign <multisig_url>",
		Short:   "Sign <multisig_url>",
		Example: "flow multisig sign <multisig_url> --signer alice",
		Args:    cobra.ExactArgs(1),
	},
	Flags: &signFlags,
	RunS:  sign,
}

func sign(
	args []string,
	readerWriter flowkit.ReaderWriter,
	globalFlags command.GlobalFlags,
	services *services.Services,
	state *flowkit.State,
) (command.Result, error) {
	multisig_url := args[0]
	if multisig_url == "" {
		return nil, fmt.Errorf("multisig url is empty")
	}
	payload, err := retrieve(multisig_url)

	if err != nil {
		return nil, fmt.Errorf("failed to read multisig RLP from %s: %v", multisig_url, err)
	}

	signer, err := state.Accounts().ByName(signFlags.Signer)
	if err != nil {
		return nil, fmt.Errorf("signer account: [%s] doesn't exists in configuration", signFlags.Signer)
	}

	signed, err := services.Transactions.Sign(signer, payload, globalFlags.Yes)

	if err != nil {
		return nil, err
	}

	tx := signed.FlowTransaction()
	msg := tx.Encode()

	var b bytes.Buffer
	writer := util.CreateTabWriter(&b)
	_, _ = fmt.Fprintf(writer, "%x", string(msg))
	_ = writer.Flush()

	err = post(multisig_url, b.String())

	return &MultisigResult{
		tx:     tx,
		rlp:    string(payload),
		signed: b.String(),
	}, err
}

func retrieve(rlpUrl string) ([]byte, error) {

	if rlpUrl == "" {
		return nil, fmt.Errorf("multisig url is empty")
	}

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	resp, err := client.Get(rlpUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error downloading multisig identifier")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func post(rlpUrl string, signed string) error {

	resp, err := http.Post(rlpUrl, "application/text", bytes.NewBufferString(signed))

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error posting signed RLP")
	}

	return nil
}
