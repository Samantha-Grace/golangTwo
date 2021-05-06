//url shortener in Go
package main

import (
	"log"
	"net/http"

	"github.com/Samantha-Grace/golangTwo/web"
)

func main() {
	mux := http.NewServeMux()
	web.RegisterHandlers(mux)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("unable to start server: %v\n", err)
	}
}

// func Hello(target string) string {
// 	return fmt.Sprintf("Hello, %s!", target)
// }
// func main() {
// 	target := "World"
// 	if len(os.Args) >= 2 {
// 		target = os.Args[1]
// 	}
// 	fmt.Println(Hello(target))
// }
