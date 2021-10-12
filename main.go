package main

import (
  "errors"
  "fmt"
  "net/http"
)

var (
  errorRequestFailed = errors.New("Http request is failed"); 
)

type RequestCheck struct{
  url string
  response string
}

func main() {
  urls := []string{
    "https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
  }
  channel := make(chan RequestCheck)
  results := map[string]string{}

  for _, url := range urls {
    fmt.Println("Checking " ,url)
    go hitUrl(url, channel);
  }

  //Get Result
  for i:=0;i<len(urls); i++{
    result := <-channel
    results[result.url] = result.response;
  }

  //Print
  for url, response := range results {
    fmt.Println(url, response);
  }
}

func hitUrl(url string, channel chan RequestCheck) {
  response, err := http.Get(url)
  if err != nil || response.StatusCode >= 400 {
    channel <- RequestCheck{url:url, response : "FAILED"}
  } else{
    channel <- RequestCheck{url:url, response : "OK"}
  }
}