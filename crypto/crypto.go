package crypto

import(
  "crypto/sha256"
)

func DoubleSha256(index string, transaction string, prevHash []byte, nonce string) []byte {
  record := []byte(index + transaction + string(prevHash) + nonce)

  hash := sha256.Sum256(record[:])
  finalHashed := sha256.Sum256(hash[:])
  return finalHashed[:]
}

