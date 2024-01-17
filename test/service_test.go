package test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"

	"filelineserve/data/response"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	lineIndex      int
	expectedString string
	expectedStatus int
}

func TestGetLine(t *testing.T) {
	testCases := []testCase{
		{
			lineIndex:      1,
			expectedString: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
			expectedStatus: http.StatusOK,
		},
		{
			lineIndex:      3,
			expectedString: "Egestas sed sed risus pretium quam vulputate dignissim suspendisse in.",
			expectedStatus: http.StatusOK,
		},
		{
			lineIndex:      5,
			expectedString: "Lorem sed risus ultricies tristique nulla aliquet enim tortor.",
			expectedStatus: http.StatusOK,
		},
		{
			lineIndex:      8,
			expectedString: "",
			expectedStatus: 413,
		},
	}

	for _, test := range testCases {
		requesteURL := fmt.Sprintf("http://localhost:8080/lines/%d", test.lineIndex)
		resp, err := http.Get(requesteURL)
		if err != nil {
			fmt.Println(err)
			t.Error("Failed to make http request")
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error("Response body reading error")
		}
		var result response.ReadLine
		err = json.Unmarshal(body, &result)
		if err != nil {
			t.Error("response unmarshall error")
		}

		assert.Equal(t, test.expectedStatus, resp.StatusCode, "Unexpected status code")
		if test.expectedStatus == http.StatusOK {
			assert.Equal(t, test.expectedString, result.Line, "Unexpected response")
		}

	}

}

func TestSwaggerPage(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/swagger/index.html")
	if err != nil {
		t.Errorf("Got an unexpected error making API request on test Swagger page up\nerror: %s", err.Error())
	}
	_ = resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Unexpected status code")
}

func TestMain(m *testing.M) {
	runCMD := exec.Command("go", "run", "../main.go", "../files/example.txt")
	runCMD.Env = os.Environ()
	runCMD.Env = append(runCMD.Env, "API_PORT=8080")
	err := runCMD.Start()
	if err != nil {
		log.Fatal("error running api main file: ", err)
	}

	// need to make sure API had time to start before making requests
	time.Sleep(time.Second * 10)

	code := m.Run()
	_ = runCMD.Process.Kill()
	os.Exit(code)

}
