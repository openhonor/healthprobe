package request

import (
	"crypto/tls"
	"net"
	"net/http"
	"strings"
	"time"
)

var tr = &http.Transport{
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
	},
	DialContext: (&net.Dialer{
		Timeout: time.Second * 8,
	}).DialContext,
}

var Client = &http.Client{Transport: tr}

func DoRequest(url, addr string, requestHeader map[string]string) (bool, error) {

	url = strings.Replace(url, " ", "", -1)
	addr = strings.Replace(addr, " ", "", -1)
	requst, err := http.NewRequest("GET",
		url,
		nil)
	if err != nil {
		return false, err
	}

	for k, v := range requestHeader {
		requst.Header.Add(k, v)
	}
	requst.URL.Host = addr
	response, err := Client.Do(requst)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		//fmt.Println(response.Header)
		return true, nil
	}
	return false, err
}
