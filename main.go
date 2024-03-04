// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	fmt.Println("Server is running on http://localhost:8080")
// 	http.ListenAndServe(":8080", nil)
// }



package main

import (
	"fmt"
	"net/http"
	"time"
)

// Middleware function to log incoming requests
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		fmt.Printf("[%s] %s %s (%v)\n", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path, duration)
	})
}

// Handler function for the home route
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the home page!")
}

// Handler function for the about route
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the about page.")
}

// // Handler function for serving static files
// func staticHandler(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, r.URL.Path[1:])
// }

func main() {
	// Use the default ServeMux provided by net/http
	mux := http.NewServeMux()

	// Attach middleware to log requests
	http.Handle("/", logMiddleware(mux))

	// Define different routes and their corresponding handlers
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/about", aboutHandler)

	// Serve static files from the "static" directory
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	
}
