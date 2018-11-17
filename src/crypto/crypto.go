package crypto

import(
  "crypto/sha256"
)

func DoubleSha256(header []byte) []byte {
  hash        := sha256.Sum256(header[:])
  finalHashed := sha256.Sum256(hash[:])

  return finalHashed[:]
}


