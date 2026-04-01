package webchecker

import "sync"

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

var wg sync.WaitGroup

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	c := make(chan result)

	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c <- result{url, wc(url)}
		}()
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for r := range c {
		results[r.string] = r.bool
	}

	return results
}
