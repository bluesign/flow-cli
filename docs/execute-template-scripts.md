---
title: Execute a Script Template with the Flow CLI
sidebar_title: Execute a Script Template
description: How to execute a template Cadence script on Flow from the command line
---

The Flow CLI provides a command to execute a template Cadence script on the Flow execution state with any Flow Access API.

```shell
flow scripts execute-template <template name> [<argument> <argument>...] [flags]
```

## Example Usage

```shell
> flow scripts execute-template 0x01cf0e2f2f715450

1000.0
```

## Arguments

### Template Name
- Name: `template name`
- Valid inputs:
  - `fusd-balance` [source](https://github.com/onflow/fusd/blob/main/transactions/scripts/get_fusd_balance.cdc)
  
### Arguments
- Name: `argument`
- Valid inputs: valid [cadence values](https://docs.onflow.org/cadence/json-cadence-spec/)
  matching argument type in script code.

Input arguments values matching corresponding types in the source code and passed in the same order.

## Flags

### Arguments

- Flag: `--arg`
- Valid inputs: argument in `Type:Value` format.

Arguments passed to the Cadence script in `Type:Value` format. 
The `Type` must be the same as type in the script source code for that argument.  

For passing complex argument values see [send transaction](https://docs.onflow.org/flow-cli/send-transactions/#example-usage) document. 

⚠️  Deprecated: use command arguments instead.

### Arguments JSON

- Flag: `--args-json`
- Valid inputs: arguments in JSON-Cadence form.

Arguments passed to the Cadence script in the Cadence JSON format.
Cadence JSON format contains `type` and `value` keys and is 
[documented here](https://docs.onflow.org/cadence/json-cadence-spec/).

### Code

- Flag: `--code`

⚠️  No longer supported: use filename argument.

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