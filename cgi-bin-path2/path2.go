package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	path1Handler()
}

func path1Handler() {

	// CgiReqBody define the cgi-bin request body struct
	type CgiReqBody struct {
		Input string `json:"input"`
	}

	// CgiResBody define the cgi-bin request body struct
	type CgiResBody struct {
		Output string `json:"output"`
	}
	var responseJson []byte

	defer func() {
		fmt.Println("HTTP/1.1 200 OK")
		fmt.Println("Content-Type: application/json")
		fmt.Println("")
		fmt.Println(string(responseJson))
	}()

	fmt.Println("Content-Type: application/json")
	fmt.Println("")
	reqBody, err := io.ReadAll(os.Stdin)
	if err != nil || len(reqBody) == 0 {
		fmt.Println("request body no content")
		return
	}
	var cgiReqBody CgiReqBody
	err = json.Unmarshal([]byte(reqBody), &cgiReqBody)
	if err != nil {
		fmt.Println("parse request body failed")
		return
	}

	cgiResBody := CgiResBody{
		Output: "path2, " + cgiReqBody.Input,
	}
	response, err := json.Marshal(cgiResBody)
	if err != nil {
		fmt.Println("marshal response body failed")
		return
	}
	responseJson = response
}
