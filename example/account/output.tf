output "mnemonic_account1" {
  value = nonsensitive(algorand_account.account1.mnemonic)
}

output "mnemonic_account2" {
  value = nonsensitive(algorand_account.account2.mnemonic)
}