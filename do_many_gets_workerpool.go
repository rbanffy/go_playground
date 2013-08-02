package main
 
import (
	"bufio"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime"
)
 
func deliver(payload string) {
	params := url.Values{}
	params.Set("payload", payload)
	result, err := http.Get("http://127.0.0.1/?" + params.Encode())
	if err != nil {
		log.Print(err)
	} else {
		result.Body.Close()
	}
}
 
func worker(id int, jobs <-chan string, results chan<- string) {
	for j := range jobs {
		deliver(j)
		log.Print("Worker ", id, " sent ", j)
	}
}
 
func main() {
	jobs := make(chan string, 100)
        results := make(chan string, 100)
 
	for w := 1; w <=1000 ; w++ {
		go worker(w, jobs, results)
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
	log.Print("Finished")
}
