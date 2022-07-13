---
title: Sign a Multisig transaction with the Flow CLI
sidebar_title: Sign a Multisig Transaction
description: How to sign a Multisig transaction from the command line
---

The Flow CLI provides a command to sign a multsi transaction with options to specify
signer account.

Use this functionality after a multisig transaction has been generated:
 - When a multisig transaction the keys needed to sign have already been identified
 - This signing method will first pull down the transaction RLP from url to allow the user to sign
 - After the transaction has been signed the resulting signed RLP will be posted to the server

```shell
flow multisig sign <transaction id ulr>
```

## Example Usage

```shell
> flow multisig sign https://path/to/tx/identifier --signer alice 

🙏 RLP retrieved successfully
✨ Signed RLP Posted successfully

```

By not including `--yes` the transaction will be displayed so the user can verify what is being signed

## Arguments

### Url to transaction RLP
- Name: `identifier`
- Valid inputs: RLP server URL to the transaction.

Specify the full url containing transaction identifier.

## Flags

### Signer

- Flag: `--signer`
- Valid inputs: the name of an account defined in the configuration (`flow.json`)

Specify the name of the account that will be used to sign the transaction.


### Configuration

- Flag: `--conf`
- Short Flag: `-f`
- Valid inputs: valid filename

Specify a filename for the configuration files, you can provide multiple configuration
files by using `-f` flag multiple times.
