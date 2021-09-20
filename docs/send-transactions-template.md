---
title: Send a Transaction template with the Flow CLI
sidebar_title: Send a Transaction Template
description: How to send a Flow transaction template from the command line
---

The Flow CLI provides a command to send common transaction templates to any Flow Access API.

```shell
flow transactions send-template <template name> [<argument> <argument>...] [flags]
```

## Example Usage

```shell
> flow transactions send-template fusd-transfer 10.0 0x01cf0e2f2f715450
    
Status		✅ SEALED
ID		b04b6bcc3164f5ee6b77fa502c3a682e0db57fc47e5b8a8ef3b56aae50ad49c8
Payer		f8d6e0586b0a20c7
Authorizers	[f8d6e0586b0a20c7]

Proposal Key:	
    Address	f8d6e0586b0a20c7
    Index	0
    Sequence	0

No Payload Signatures

Envelope Signature 0: f8d6e0586b0a20c7
Signatures (minimized, use --include signatures)

Events:	 None

Code (hidden, use --include code)

Payload (hidden, use --include payload)

```

## Arguments

### Template Name
- Name: `template name`
- Valid inputs:
  - `fusd-transfer` [source](https://github.com/onflow/fusd/blob/main/transactions/transfer_fusd.cdc)
  - `fusd-setup` [source](https://github.com/onflow/fusd/blob/main/transactions/setup_fusd_vault.cdc)

### Arguments
- Name: `argument`
- Valid inputs: valid [cadence values](https://docs.onflow.org/cadence/json-cadence-spec/) 
  matching argument type in transaction code.

Input arguments values matching corresponding types in the source code and passed in the same order.

## Flags

### Include Fields

- Flag: `--include`
- Valid inputs: `code`, `payload`

Specify fields to include in the result output. Applies only to the text output.

### Code

- Flag: `--code`

⚠️  No longer supported: use filename argument.

### Results

- Flag: `--results`

⚠️  No longer supported: all transactions will provide result.

### Exclude Fields

- Flag: `--exclude`
- Valid inputs: `events`

Specify fields to exclude from the result output. Applies only to the text output.

### Signer

- Flag: `--signer`
- Valid inputs: the name of an account defined in the configuration (`flow.json`)

Specify the name of the account that will be used to sign the transaction.

### Arguments

- Flag: `--arg`
- Valid inputs: argument in `Type:Value` format.

Arguments passed to the Cadence transaction in `Type:Value` format.
The `Type` must be the same as type in the transaction source code for that argument.

⚠️  Deprecated: use command arguments instead.

### Arguments JSON

- Flag: `--args-json`
- Valid inputs: arguments in JSON-Cadence form.

Arguments passed to the Cadence transaction in Cadence JSON format.
Cadence JSON format contains `type` and `value` keys and is 
[documented here](https://docs.onflow.org/cadence/json-cadence-spec/).

### Gas Limit

- Flag: `--gas-limit`
- Valid inputs: an integer greater than zero.
- Default: `1000`

Specify the gas limit for this transaction.

### Host

- Flag: `--host`
- Valid inputs: an IP address or hostname.
- Default: `127.0.0.1:3569` (Flow Emulator)

Specify the hostname of the Access API that will be
used to execute the command. This flag overrides
any host defined by the `--network` flag.

### Network

- Flag: `--network`
- Short Flag: `-n`
- Valid inputs: the name of a network defined in the configuration (`flow.json`)
- Default: `emulator`

Specify which network you want the command to use for execution.

### Filter

- Flag: `--filter`
- Short Flag: `-x`
- Valid inputs: a case-sensitive name of the result property.

Specify any property name from the result you want to return as the only value.

### Output

- Flag: `--output`
- Short Flag: `-o`
- Valid inputs: `json`, `inline`

Specify the format of the command results.

### Save

- Flag: `--save`
- Short Flag: `-s`
- Valid inputs: a path in the current filesystem.

Specify the filename where you want the result to be saved

### Log

- Flag: `--log`
- Short Flag: `-l`
- Valid inputs: `none`, `error`, `debug`
- Default: `info`

Specify the log level. Control how much output you want to see during command execution.

### Configuration

- Flag: `--config-path`
- Short Flag: `-f`
- Valid inputs: a path in the current filesystem.
- Default: `flow.json`

Specify the path to the `flow.json` configuration file.
You can use the `-f` flag multiple times to merge
several configuration files.