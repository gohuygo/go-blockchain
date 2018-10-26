package main

import (
  // "encoding/json"
  "io"
  "log"
  "os"
  "net"
  "bufio"

  "github.com/gohuygo/go-blockchain/block"

  // "github.com/gorilla/mux"
  "github.com/joho/godotenv"
  "github.com/davecgh/go-spew/spew"
)

// Request body that is serialized/marshalled from input body.
type RequestBody struct {
  Data int
}

var bcServer chan []block.Block
var strings = make(chan string)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal(err)
  }

  bcServer = make(chan []block.Block)

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

func handleConn(conn net.Conn) {
  defer conn.Close()
  io.WriteString(conn, "Enter a transaction: ")

  // Take input and add it to blockchain
  // TODO: Check for newly validated blocks instead
  scanner := bufio.NewScanner(conn)
  scanner.Scan()


  // TODO: Send transaction to mempool instead
  newBlock, err := block.GenerateBlock(block.Blockchain[len(block.Blockchain)-1], scanner.Text())

  if err != nil {
    io.WriteString(conn, "(500) Internal Server Error")
    return
  }

  if block.IsBlockValid(newBlock, block.Blockchain[len(block.Blockchain)-1]) {
    newBlockchain := append(block.Blockchain, newBlock)
    block.ReplaceChain(newBlockchain)
  }

  spew.Dump(block.Blockchain)
}

