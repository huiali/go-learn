package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getLocation() (string, error) {
	resp, err := http.Get("http://www.msftconnecttest.com/redirect")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	url, err := resp.Location()
	if err != nil {
		return "", err
	}
	return url.String(), nil
}

func main() {
	url, err := getLocation()
	if err != nil {
		return
	}
	fmt.Println(url)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true, VerifyPeerCertificate: nil},
	}
	var client = &http.Client{Transport: tr}
	resp, err := client.Get("https://192.168.10.9:8443/portal/PortalSetup.action?portal=ee753630-a594-11e5-bc65-005056aa7bee&sessionId=c0a80c01004c06435d3831cb&action=cwa")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
