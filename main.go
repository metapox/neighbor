package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	proxyURL, err := url.Parse(req.URL.String())
	if err != nil {
		log.Fatal(err)
	}

	proxy, err := http.NewRequest(req.Method, req.URL.String(), req.Body)
	if err != nil {
		log.Fatal(err)
	}

	proxy.Header = req.Header

	client := &http.Client{}
	response, err := client.Do(proxy)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadGateway)
		return
	}
	defer response.Body.Close()

	for key, value := range response.Header {
		for _, v := range value {
			res.Header().Add(key, v)
		}
	}
	res.WriteHeader(response.StatusCode)
	io.Copy(res, response.Body)
}

func main() {
	http.HandleFunc("/", handleRequestAndRedirect)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
