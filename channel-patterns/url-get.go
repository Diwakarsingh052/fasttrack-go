package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

var url = []string{
	`https://pkg.go.dev/`,
	`https://github.com/`,
	`abc.com/1234`,
}

type response struct {
	url  string
	resp *http.Response
	err  error
}

var wg = &sync.WaitGroup{}

func main() {
	doGetRequest(url)
	wg.Wait()
}

func doGetRequest(urls []string) {
	respChan := make(chan response, len(url)) // buffered channel

	for _, v := range urls {
		wg.Add(1)
		//fanning out go routines // one task = one goroutine
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			r := response{
				url:  url,
				resp: resp,
				err:  err,
			}
			respChan <- r //sending the resp struct to respCh

		}(v)

		wg.Add(1)
		go func() {
			defer wg.Done()
			r := <-respChan

			if r.err != nil {
				log.Println(r.err)
				return
			}

			b, err := io.ReadAll(r.resp.Body)
			if err != nil {
				log.Println(err)
				return

			}
			defer r.resp.Body.Close()
			fmt.Println(r.url)
			//fmt.Println(string(b))
			_ = b
			fmt.Println(r.resp.Status)

		}()

	}

}
