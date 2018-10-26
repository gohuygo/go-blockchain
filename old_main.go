

// type RequestBody struct {
//   Data int
// }

// func main() {
//   err := godotenv.Load()
//   if err != nil {
//     log.Fatal(err)
//   }

//   go func() {
//     t := time.Now()
//     genesisBlock := block.Block{0, t.String(), 0, "", ""}
//     block.Blockchain = append(block.Blockchain, genesisBlock)
//   }()

//   log.Fatal(run())
// }

// func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
//   var requestBody RequestBody
//   decoder := json.NewDecoder(r.Body)
//   if err := decoder.Decode(&requestBody); err != nil {
//     respondWithJSON(w, r, http.StatusBadRequest, r.Body)
//     return
//   }
//   defer r.Body.Close()

//   newBlock, err := block.GenerateBlock(block.Blockchain[len(block.Blockchain)-1], requestBody.Data)

//   if err != nil {
//     respondWithJSON(w, r, http.StatusInternalServerError, requestBody)
//     return
//   }

//   if block.IsBlockValid(newBlock, block.Blockchain[len(block.Blockchain)-1]) {
//     newBlockchain := append(block.Blockchain, newBlock)
//     block.ReplaceChain(newBlockchain)
//   }

//   respondWithJSON(w, r, http.StatusCreated, newBlock)

// }


// func makeMuxRouter() http.Handler {
//   muxRouter := mux.NewRouter()
//   muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
//   muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
//   return muxRouter
// }

// func run() error {
//   mux := makeMuxRouter()
//   httpAddr := os.Getenv("PORT")
//   log.Println("Listening on ", os.Getenv("PORT"))
//   s := &http.Server{
//     Addr:           ":" + httpAddr,
//     Handler:        mux,
//     ReadTimeout:    10 * time.Second,
//     WriteTimeout:   10 * time.Second,
//     MaxHeaderBytes: 1 << 20,
//   }

//   if err := s.ListenAndServe(); err != nil {
//     return err
//   }

//   return nil
// }

// func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
//   bytes, err := json.MarshalIndent(block.Blockchain, "", "  ")
//   if err != nil {
//     http.Error(w, err.Error(), http.StatusInternalServerError)
//     return
//   }
//   io.WriteString(w, string(bytes))
// }

// func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
//   response, err := json.MarshalIndent(payload, "", "  ")
//   if err != nil {
//     w.WriteHeader(http.StatusInternalServerError)
//     w.Write([]byte("HTTP 500: Internal Server Error"))
//     return
//   }
//   w.WriteHeader(code)
//   w.Write(response)
// }


