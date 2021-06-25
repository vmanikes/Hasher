package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

const DefaultParallelRequests = 10

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

func main()  {
	parallel := flag.Int("parallel", DefaultParallelRequests, "Number of parallel requests to make")

	flag.Parse()

	var urls []string

	if *parallel == DefaultParallelRequests {
		urls = os.Args[1:]
	} else {
		urls = os.Args[3:]
	}

	var wg sync.WaitGroup

	for idx, url := range urls {
		if idx == *parallel {
			wg.Wait()
		}

		wg.Add(1)

		go func(url string, wg *sync.WaitGroup) {
			defer wg.Done()

			hash, err := Hasher(url)
			if err != nil {
				log.Println(err)
				return
			}

			fmt.Println(url + " " + hash)
		}(url, &wg)
	}

	wg.Wait()
}

// Hasher takes in an url, makes http request and returns the md5 has of http response and error (if exists)
func Hasher(url string) (string, error){
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return "", err
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			log.Println("unable to close response body")
		}
	}()

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", md5.Sum(responseBytes)), nil
}