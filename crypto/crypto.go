package crypto

import(
  "bytes"
  "fmt"
  "encoding/binary"
  "crypto/sha256"
)

func DoubleSha256(index string, transaction string, prevHash []byte, nonce string) []byte {
  header := []byte(index + transaction + string(prevHash) + nonce)
  buf := new(bytes.Buffer)

  err := binary.Write(buf, binary.LittleEndian, header)
  if err != nil {
    fmt.Println("binary.Write failed:", err)
  }

  hash := sha256.Sum256(buf.Bytes())
  finalHashed := sha256.Sum256(hash[:])

  return finalHashed[:]
}

