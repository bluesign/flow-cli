/*
 * Flow CLI
 *
 * Copyright 2019-2021 Dapper Labs, Inc.
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

package flowkit

import (
	"fmt"

	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"

	"github.com/onflow/flow-cli/pkg/flowkit/config"
	"github.com/onflow/flow-cli/pkg/flowkit/util"
)

type Account struct {
	name    string
	address flow.Address
	key     AccountKey
}

func (a *Account) Address() flow.Address {
	return a.address
}

func (a *Account) Name() string {
	return a.name
}

func (a *Account) Key() AccountKey {
	return a.key
}

func (a *Account) SetKey(key AccountKey) {
	a.key = key
}

func accountsFromConfig(conf *config.Config) ([]*Account, error) {
	accounts := make([]*Account, 0, len(conf.Accounts))

	for _, accountConf := range conf.Accounts {
		account, err := accountFromConfig(accountConf)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func AccountFromAddressAndKey(address flow.Address, privateKey crypto.PrivateKey) *Account {
	return &Account{
		name:    "",
		address: address,
		key:     NewHexAccountKeyFromPrivateKey(0, crypto.SHA3_256, privateKey),
	}
}

func accountFromConfig(account config.Account) (*Account, error) {
	key, err := NewAccountKey(account.Key)
	if err != nil {
		return nil, err
	}

	return &Account{
		name:    account.Name,
		address: account.Address,
		key:     key,
	}, nil
}

func accountsToConfig(accounts []*Account) config.Accounts {
	accountConfs := make([]config.Account, 0)

	for _, account := range accounts {
		accountConfs = append(accountConfs, accountToConfig(account))
	}

	return accountConfs
}

func accountToConfig(account *Account) config.Account {
	return config.Account{
		Name:    account.name,
		Address: account.address,
		Key:     account.key.ToConfig(),
	}
}

func generateEmulatorServiceAccount(sigAlgo crypto.SignatureAlgorithm, hashAlgo crypto.HashAlgorithm) (*Account, error) {
	seed, err := util.RandomSeed(crypto.MinSeedLength)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.GeneratePrivateKey(sigAlgo, seed)
	if err != nil {
		return nil, fmt.Errorf("failed to generate emulator service key: %v", err)
	}

	return &Account{
		name:    config.DefaultEmulatorServiceAccountName,
		address: flow.ServiceAddress(flow.Emulator),
		key:     NewHexAccountKeyFromPrivateKey(0, hashAlgo, privateKey),
	}, nil
}
