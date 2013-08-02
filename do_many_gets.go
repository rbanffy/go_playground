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
 
func deliver(payload string) {
	params := url.Values{}
	params.Set("payload", payload)
	result, err := http.Get("http://127.0.0.1/?" + params.Encode())
	if err != nil {
		log.Print(err)
	} else {
		log.Print("Sent ", payload)
		result.Body.Close()
	}
}
 
func main() {
	runtime.GOMAXPROCS(1) // We don't need more
 
	var wg sync.WaitGroup
 
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
		wg.Add(1)
		go func(payload string) {
			deliver(payload)
			wg.Done()
		}(string(payload))
	}
	wg.Wait()
	log.Print("Finished")
	defer f.Close()
}
