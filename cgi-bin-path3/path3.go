package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	path3Handler()
}

func path3Handler() {
	defer func() {
		fmt.Println("HTTP/1.1 200 OK")
		fmt.Println("Content-Type: application/json")
		fmt.Println("")
	}()

	type CgiReqBody struct {
		Title string `json:"title"`
	}
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
	fmt.Println("Path3 Request body: ", cgiReqBody)
}
