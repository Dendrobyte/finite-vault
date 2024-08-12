package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
 * Shit that needs to get done
 * [ ] 0. Choose a framework (e.g. fasthttp)
 * [ ] 1. Set up the OAuth stuff as mirrored in the JS side of things. Auth guards??
 * [ ] 2. Get the login data working, which means pulling data from mongo and sending it back
 * [ ] 3. Incrementing daily value when user logs in based on timestamp (this is beyond the JS stuff)
 * [ ] 4. Full transaction support via mongo, etc. (the actual meaty part, not sure if it'll take longer than auth)
 */

func main() {
	fmt.Println("Server starting on port 5000")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		respLine := fmt.Sprintf("%v\n", r)
		fmt.Printf("Response line: %v\n", respLine)
		f, err := os.OpenFile("./data/test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			f, _ = os.Create("./data/test.txt")
		}
		defer f.Close()

		done, err := io.WriteString(f, respLine)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Wrote bytes: %d | Resp Line length: %d", done, len(respLine))
		f.Sync()

		w.Write([]byte("Data written to file! >:)"))
	})

	http.ListenAndServe("localhost:5000", nil)
}
