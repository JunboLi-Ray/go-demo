package test

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestRequest(t *testing.T) {
	testUrl("Test 1", "http://localhost:8080/users", "GET", "", t)
	testUrl("Test 2", "http://localhost:8080/users", "POST", "{\"name\":\"Alice\"}", t)
	testUrl("Test 2", "http://localhost:8080/users", "POST", "{\"name\":\"Paul\"}", t)
	testUrl("Test 3", "http://localhost:8080/users/1/relationships", "GET", "", t)
	testUrl("Test 4", "http://localhost:8080/users/1/relationships/2", "PUT", "{\"state\":\"liked\"}", t)
}

func testUrl(name string, url string, reqType string, data string, t *testing.T) {
	t.Log(name)
	{
		t.Logf("check Url: \"%s\"", url)
		{
			reqest, _ := http.NewRequest(reqType, url, strings.NewReader(data))
			client := &http.Client{}
			resp, err := client.Do(reqest)
			if err != nil {
				t.Fatal("Call Fail", ballotX, err)
			}
			defer resp.Body.Close()
			result, err := ioutil.ReadAll(resp.Body)
			if resp.StatusCode != http.StatusNotFound {
				t.Logf("Result Check %v", checkMark)
			} else {
				t.Errorf("Result Check %v, Response: \"%s\"", ballotX, result)
			}
		}
	}
}
