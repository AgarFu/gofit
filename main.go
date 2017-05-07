package main

import (
  "os"
  "strconv"
  "io"
  "io/ioutil"
  "fmt"
  "log"
  "net/http"
)

// Different logging levels
var (
  Trace   *log.Logger
  Info    *log.Logger
  Warning *log.Logger
  Error   *log.Logger
)

// Init Initializes the logging system
func Init(
  traceHandle io.Writer,
  infoHandle io.Writer,
  warningHandle io.Writer,
  errorHandle io.Writer) {

  Trace = log.New(traceHandle,
    "TRACE: ",
    log.Ldate|log.Ltime|log.Lshortfile)

  Info = log.New(infoHandle,
    "INFO: ",
    log.Ldate|log.Ltime|log.Lshortfile)

  Warning = log.New(warningHandle,
    "WARNING: ",
    log.Ldate|log.Ltime|log.Lshortfile)

  Error = log.New(errorHandle,
    "ERROR: ",
    log.Ldate|log.Ltime|log.Lshortfile)
}

func handler(w http.ResponseWriter, r *http.Request) {
  Info.Println(r.RemoteAddr, r.Method, r.Proto, r.ContentLength, r.Host, r.URL)
  fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func getenv(key, fallback string) string {

  value := os.Getenv(key)
  if len(value) == 0 {
    return fallback
  }
  return value
}

func main() {
  const defaultListeningPort = "8080"
  var listeningPort int

  Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

  listeningPort, err := strconv.Atoi(getenv("GO_TEST_PORT", defaultListeningPort))
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  fmt.Println("PORT:", listeningPort)
  http.HandleFunc("/", handler)
  http.ListenAndServe(fmt.Sprintf(":%d", listeningPort), nil)
}
