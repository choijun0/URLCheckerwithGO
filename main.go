package main

import (
  "errors"
  "fmt"
  "net/http"
)

var (
  errorRequestFailed = errors.New("Http request is failed"); 
)

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

  //Make empty map
  //var results map[string]string 이렇게 초기화하면 map값이 nil이되 제기능X
  results := map[string]string{}

  for _, url := range urls {
    fmt.Println("Checking " ,url)
    err := hitUrl(url);
    result := "SUCCESS";
    if err == nil {
      results[url] = result;
    } else if err == errorRequestFailed {
      result = "FAILED"
      results[url] = result;
    }
  }

  //Print results
  for url, result := range results{
    fmt.Println(url, result);
  }
}

func hitUrl(url string) error{
  response, err := http.Get(url)
  if err != nil || response.StatusCode >= 400 {
    fmt.Println(response.StatusCode);
    return errorRequestFailed
  } else{
    return nil
  }
}