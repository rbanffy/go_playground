package main
 
import (
	"flag"
	"fmt"
)
 
var start string
 
func init() {
	flag.StringVar(&start, "start", "", "time to start")
	flag.Parse()
}
 
func main() {
	fmt.Println("start is ", start)
	fmt.Println("Other args below")
	for i, arg := range flag.Args() {
		fmt.Println(i, arg)
	}
}
