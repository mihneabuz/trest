package internal

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func parseResponse(resp *http.Response) string {
	var body string
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		body = err.Error()
	}
	body = strings.Replace(string(bodyBytes), "\n", "\n> ", -1)

	content := resp.Header["Content-Type"]

	res := fmt.Sprintf("=> [code: %s] [type: %s]\n> %s", resp.Status, content, body)

	return res
}

func Delete(args []string) string {
	req, err := http.NewRequest("DELETE", args[0], nil)
	if err != nil {
		return err.Error()
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	return parseResponse(resp)
}

func Post(args []string) string {
	data := bytes.NewBuffer([]byte(args[1]))

	resp, err := http.Post(args[0], "application/json", data)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	return parseResponse(resp)
}

func Get(args []string) string {
	resp, err := http.Get(args[0])
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	return parseResponse(resp)
}
