package concurrency

import (
	"testing"
	"time"
)

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.Run("non concurrent", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CheckWebsites(slowStubWebsiteChecker, urls)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CheckWebsitesConcurrent(slowStubWebsiteChecker, urls)
		}
	})

}
