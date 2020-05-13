package main

import (
    "fmt"
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
    "encoding/base64"
)

func main() {
    // Setup proxy transport
    proxyStr := "___YOUR_PROXY_URL___"
    proxyURL, _ := url.Parse(proxyStr)
    auth := "___YOUR_AUTH_KEY___:"
    basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
    hdr := http.Header{}
    hdr.Add("Proxy-Authorization", basicAuth)
    transport := &http.Transport{
        Proxy: http.ProxyURL(proxyURL),
        ProxyConnectHeader: hdr,
    }

    // Create HTTP client
    client := &http.Client {
        Transport: transport,
    }

    // Create request
    req, err := http.NewRequest("GET", "http://httpbin.scrapinghub.com/get", nil)

    // Add proxy authorization header
    // @TODO: figure out why this doesn't work with the transport
    req.Header.Add("Proxy-Authorization", basicAuth)

    // Log errors
    if err != nil {
        log.Fatalln(err)
    }

    // Perform request
    resp, err := client.Do(req)

    // Check for error response
    if err != nil {

        fmt.Println(err)
    }

    defer resp.Body.Close()

    // Output the request
    /*
    requestDump, err := httputil.DumpRequest(req, true)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(requestDump))
    */

    // Output the response
    responseDump, err := httputil.DumpResponse(resp, true)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(responseDump))
}
