package crypto

import(
  // "bytes"
  // "fmt"
  // "encoding/binary"
  "crypto/sha256"
)

func DoubleSha256(header []byte) []byte {
  hash        := sha256.Sum256(header)
  finalHashed := sha256.Sum256(hash[:])

  return finalHashed[:]
}


