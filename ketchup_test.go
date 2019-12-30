package ketchup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func TestParseDocument(t *testing.T) {
	files, err := ioutil.ReadDir("tests")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(files); i += 2 {
		fileName := files[i].Name()
		testName := strings.Split(fileName, ".")[0]

		sourceBuffer, sourceErr := ioutil.ReadFile("tests/" + testName + ".html")
		if sourceErr != nil {
			errorString := "Error loading test" + testName
			t.Error(errorString)
		}

		source := string(sourceBuffer)
		parsingResult := ParseDocument(source)

		ret, err := json.MarshalIndent(parsingResult, "", " ")
		if err != nil {
			t.Log(err)
		}

		t.Log(string(ret))
	}
}

func subTestResult(result bool, testName string) {
	if result {
		fmt.Println("  └── PASS: " + testName)
	} else {
		fmt.Println("  └── FAIL: " + testName)
	}
}
