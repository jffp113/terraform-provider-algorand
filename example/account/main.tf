terraform {
  required_providers {
    algorand = {
      source = "terraform.local/local/algorand"
      version = "1.0.0"
    }
  }
}

provider "algorand" {
  algod_endpoint = "https://node.testnet.algoexplorerapi.io"
  algod_token = ""
  indexer_endpoint = "https://algoindexer.testnet.algoexplorerapi.io"
  indexer_token = ""
}


resource "algorand_account" "account1" {
  name = var.account1_name
}

resource "algorand_account" "account2" {
  name = var.account2_name
}

data "algorand_account" "imported_account1" {
  name = var.account1_name
  mnemonic_envvar = "ALGORAND_MNEMONIC"
}

data "algorand_account" "imported_account2" {
  name = var.account1_name
  mnemonic = algorand_account.account2.mnemonic
}