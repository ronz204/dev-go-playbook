package patterns

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL    string
	Err    error
	Status int
}

func worker(id int, jobs <-chan string, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for url := range jobs {
		fmt.Printf("Worker %d iniciando check: %s\n", id, url)

		client := http.Client{Timeout: 5 * time.Second}
		resp, err := client.Get(url)

		if err != nil {
			results <- Result{URL: url, Err: err}
			fmt.Printf("Worker %d error: %s\n", id, err)
			continue
		}

		results <- Result{URL: url, Status: resp.StatusCode}
		fmt.Printf("Worker %d finalizando check: %s - Status: %d\n", id, url, resp.StatusCode)
		resp.Body.Close()
	}
}

func WorkersDemo() {
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://www.invalid-url.com",
	}

	const WORKERS int = 3
	var wg sync.WaitGroup

	jobs := make(chan string, len(urls))
	results := make(chan Result, len(urls))

	for w := 1; w <= WORKERS; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		if res.Err != nil {
			fmt.Printf("Error checking %s: %s\n", res.URL, res.Err)
		} else {
			fmt.Printf("Checked %s: Status %d\n", res.URL, res.Status)
		}
	}
}
