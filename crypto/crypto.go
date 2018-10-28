package crypto

import(
  // "encoding/binary"
  "crypto/sha256"
)

func DoubleSha256(index string, transaction string, prevHash []byte, nonce string) []byte {
  record := []byte(index + transaction + string(prevHash) + nonce)
  // uintRecord, _ := binary.Uvarint(record)
  // recordEndian := binary.LittleEndian.Uint64(record)

  hash := sha256.Sum256(record[:])
  finalHashed := sha256.Sum256(hash[:])
  return finalHashed[:]
}

