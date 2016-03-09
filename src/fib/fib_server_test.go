package fib

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var server *httptest.Server

func TestMain(m *testing.M) {
	server = httptest.NewServer(http.HandlerFunc(FibServerHandler))
	os.Exit(m.Run())
}

func TestServerReturnsJsonArray(t *testing.T) {
	t.Log("The server should return a JSON array")
	resp, err := http.Get(server.URL + "/fib.json?n=0")

	if err != nil {
		t.Error(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}

	var respObj []interface{}
	err = json.Unmarshal(body, &respObj)

	if err != nil {
		t.Error(err)
	}
}

func TestServerGiven0(t *testing.T) {
	t.Log("The server should return a JSON array with the element 0")

	resp, _ := http.Get(server.URL + "/fib.json?n=0")

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var respObj []int
	err := json.Unmarshal(body, &respObj)

	if err != nil {
		t.Error(err)
	} else if len(respObj) != 1 {
		t.Errorf("response should == [0]. Got %s instead", body)
	}
}

func TestServerGiven1(t *testing.T) {
	t.Log("The server should return a JSON array with length 2")

	resp, _ := http.Get(server.URL + "/fib.json?n=1")

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var respObj []int
	err := json.Unmarshal(body, &respObj)

	if err != nil {
		t.Error(err)
	} else if len(respObj) != 2 {
		t.Errorf("response should == [0,1]. Got %s instead", body)
	}
}

func TestNonexistentUrl(t *testing.T) {
	t.Log("The server should return a 404 (Not Found) error for /")

	resp, _ := http.Get(server.URL)
	if resp.StatusCode != http.StatusNotFound {
		t.Error("Missing a 404 error from the server")
	}
}

func TestNegativeParam(t *testing.T) {
	t.Log("The server should return a 400 (Bad Request) error for n=-1")

	resp, _ := http.Get(server.URL + "/fib.json?n=-1")
	if resp.StatusCode != http.StatusBadRequest {
		t.Error("Missing a 400 error from the server")
	}
}

func TestFloatParam(t *testing.T) {
	t.Log("The server should return a 400 (Bad Request) error for n=1.0")

	resp, _ := http.Get(server.URL + "/fib.json?n=1.0")
	if resp.StatusCode != http.StatusBadRequest {
		t.Error("Missing a 400 error from the server")
	}
}
