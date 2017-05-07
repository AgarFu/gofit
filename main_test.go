package main

import (
  "testing"
  "net/http"
  "net/http/httptest"
  "os"
  "io/ioutil"
  "fmt"
)

func TestMain(m *testing.M) {
    mySetupFunction()
    retCode := m.Run()
    // myTeardownFunction()
    os.Exit(retCode)
}

func mySetupFunction() {
  Init(ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard)
}

/*
func myTeardownFunction() {
  fmt.Println("Tearing down")
}
*/

func TestHandler(t *testing.T) {
  var path = "tests"

  req, err := http.NewRequest("GET", fmt.Sprintf("/%s", path), nil)
  if err != nil {
    t.Fatal(err)
  }

  rr := httptest.NewRecorder()
  testHandler := http.HandlerFunc(handler)
  testHandler.ServeHTTP(rr, req)

  // Check the status code is what we expect.
  if status := rr.Code; status != http.StatusOK {
      t.Errorf("handler returned wrong status code: got %v want %v",
          status, http.StatusOK)
  }

  // Check the response body is what we expect.
  expected := fmt.Sprintf("Hi there, I love %s!", path)
  if rr.Body.String() != expected {
      t.Errorf("handler returned unexpected body: got %v want %v",
          rr.Body.String(), expected)
  }
}
