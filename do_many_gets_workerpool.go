package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime"
"sync"
)

func deliver(payload string) (delivered string, err error) {
	params := url.Values{}
	params.Set("payload", payload)
	result, e := http.Get("http://127.0.0.1/?" + params.Encode())
	if e != nil {
		log.Print(err)
	} else {
		result.Body.Close()
	}
	return payload, err
}

func worker(id int, jobs chan string, results chan<- string) {
	for j := range jobs {
		_, err := deliver(j)
		if err != nil {
			log.Print("Worker ", id, " retrying ", j, " ", err)
			jobs <- j
		} else {
			// log.Print("Worker ", id, " sent ", j)
		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)

	jobs := make(chan string, 10000)
	results := make(chan string, 10000)

	var wg sync.WaitGroup

	for w := 1; w <= 200; w++ {
		wg.Add(1)
		go func () {
			worker(w, jobs, results)
			wg.Done()
		}()
	}

	payload_file := os.Args[1]

	f, err := os.Open(payload_file)
	if err != nil {
		log.Fatal(err)
	}
	bf := bufio.NewReader(f)
	for {
		payload, isPrefix, err := bf.ReadLine()

		if err == io.EOF {
			// Already read all the contents in the file
			break
		} else if err != nil {
			// Some other error
			log.Fatal(err)
		} else if isPrefix {
			// Line is too long for the default buffer size (4K)
			log.Fatal("Error: Unexpected long line reading ", f.Name())
		}
		jobs <- string(payload)
	}
	close(jobs)
	f.Close()
	wg.Wait()
	log.Print("Finished")
}
