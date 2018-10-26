package main

import (
  // "encoding/json"
  "io"
  "log"
  // "net/http"
  "os"
  // "time"
  // "strconv"
  "net"
  "bufio"

  "github.com/gohuygo/go-blockchain/block"

  // "github.com/gorilla/mux"
  "github.com/joho/godotenv"
)

var bcServer chan []block.Block
var strings = make(chan string)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal(err)
  }

  // bcServer = make(chan []block.Block)

  // t := time.Now()
  // genesisBlock := block.Block{0, t.String(), 0, "", ""}
  // block.Blockchain = append(block.Blockchain, genesisBlock)


  server, err := net.Listen("tcp", ":"+os.Getenv("PORT"))

  if err != nil {
    log.Fatal(err)
  }

  log.Println("HTTP Server Listening on port :", os.Getenv("PORT"))
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
  io.WriteString(conn, "Enter a transaction:")

  // Take input and add it to blockchain
  // TODO: Check for newly validated blocks instead

  scanner := bufio.NewScanner(conn)

  go func() {
    for scanner.Scan() {
      log.Println("User entered: ")
      strings <- scanner.Text()
    }

    // log.Println(block.Blockchain)
  }()

  for c := range strings {
    log.Println(c)
  }
}

