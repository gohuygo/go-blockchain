package main

import (
  "io"
  "log"
  "os"
  "net"
  "bufio"
  "strings"

  "github.com/gohuygo/go-blockchain/block"

  "github.com/joho/godotenv"
  "github.com/davecgh/go-spew/spew"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal(err)
  }

  // Overzealous use of goroutine?
  go func() {
    block.GenerateGenesisBlock()
  }()

  server, err := net.Listen("tcp", ":"+os.Getenv("PORT"))

  if err != nil {
    log.Fatal(err)
  }

  log.Println("HTTP Server Listening on port: ", os.Getenv("PORT"))

  defer server.Close()

  for {
    conn, err := server.Accept()
    if err != nil {
      log.Fatal(err)
    }
    go handleConn(conn)
  }
}

// Refactor to use channels?
func handleConn(conn net.Conn) {
  defer conn.Close()
  conn.Write([]byte("Enter transactions (seperated by return): "))

  for {
    netData, err := bufio.NewReader(conn).ReadString('\n')
    if err == io.EOF {
      conn.Write([]byte("Session ended."))
      log.Println("Connection closed by client.")
      break
    }

    transaction := strings.TrimSpace(string(netData))

    // TODO: Send transaction to mempool instead
    newBlock, err := block.GenerateBlock(block.Blockchain[len(block.Blockchain)-1], transaction)
    if err != nil {
      io.WriteString(conn, "(500) Internal Server Error")
      return
    }

    if block.IsBlockValid(newBlock, block.Blockchain[len(block.Blockchain)-1]) {
      newBlockchain := append(block.Blockchain, newBlock)
      block.ReplaceChain(newBlockchain)
    }
  }
  spew.Dump(block.Blockchain)
}

