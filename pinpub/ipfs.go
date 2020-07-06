package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
)

// IPFSAPI to handle requests
const IPFSAPI = "http://127.0.0.1:5001/api/v0"

// IpfsConfig JSON like struct to access the node ID and Key
type IpfsConfig struct {
	ID           string `json:"ID"`
	PublicKey    string `json:"PublicKey"`
	AgentVersion string `json:"AgentVersion"`
}

func getIPFSconfig() (IpfsConfig, error) {
	resp, err := http.Post(IPFSAPI+"/id", "application/json", nil)
	var config IpfsConfig
	if err != nil {
		return config, err
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return config, err
	}

	json.Unmarshal(buf, &config)
	return config, err
}

func addWithPost() error {
	resp, err := http.Post(IPFSAPI+"/add", "multipart/form-data", strings.NewReader("This is a clue!"))
	if err != nil {
		return err
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("This is the response: %s", buf)
	return err
}

func addThings() error {
	fileContents := []byte("is this real life")

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "catboy.txt")
	if err != nil {
		return err
	}
	part.Write(fileContents)

	err = writer.Close()
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", IPFSAPI+"/add", body)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	fmt.Print("Here is the response from the add: ", resp)
	//json.Unmarshal(buf, &config)
	return nil
}
