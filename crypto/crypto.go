package crypto

import(
  "crypto/sha256"
  "encoding/hex"
)


func DoubleSha(index string, transaction string, prevHash string, nonce string) string {
  hash := sha256.New()
  record := index + transaction + prevHash + nonce
  hash.Write([]byte(record))
  hashed := hash.Sum(nil)

  hash.Write([]byte(hashed))
  secondHashed := hash.Sum(nil)

  return hex.EncodeToString(secondHashed)
}
