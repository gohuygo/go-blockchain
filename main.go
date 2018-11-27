package main

import (
  "io"
  "log"
  "os"
  "net"
  "bufio"
  "strings"

  "github.com/joho/godotenv"
  "github.com/davecgh/go-spew/spew"

  "github.com/gohuygo/go-blockchain/src/block"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal(err)
  }

  // Overzealous use of goroutine?
  go func() {
    block.GenerateGenesis()
  }()

  server, err := net.Listen("tcp", ":"+os.Getenv("PORT"))

  if err != nil {
    log.Fatal(err)
  }

  log.Println("TCP Server Listening on port: ", os.Getenv("PORT"))

  defer server.Close()

  for {
    conn, err := server.Accept()
    if err != nil {
      log.Fatal(err)
    }
    go handleConn(conn)
  }
}

func handleConn(conn net.Conn) {
  defer conn.Close()

  for {
    conn.Write([]byte("Enter transactions (seperated by return): "))

    netData, err := bufio.NewReader(conn).ReadString('\n')
    if err == io.EOF {
      conn.Write([]byte("Session ended."))
      log.Println("Connection closed by client.")
      break
    }

    log.Println("Transaction received.")
    log.Println("Mining...")
    transaction := strings.TrimSpace(string(netData))

    // TODO: Send transaction to mempool instead
    newBlock := *block.New(transaction)

    log.Println("isNewblockValid?", block.IsBlockValid(newBlock))
    if block.IsBlockValid(newBlock) {
      newBlockchain := append(block.Blockchain, newBlock)
      block.ReplaceChain(newBlockchain)
    }
  }
  spew.Dump(block.Blockchain)
}

