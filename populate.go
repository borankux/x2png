package main

import (
	"embed"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

//go:embed images/cat.webp
var fs embed.FS

func Populate(n int) {
	if n > 1000 {
		fmt.Println("max n is 1000")
		return
	}

	sampleImage, _ := fs.ReadFile("images/cat.webp")
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			timeStamp := time.Now().UnixNano()
			rnd := rand.Int63()
			target, _ := os.Create(fmt.Sprintf("populated.%d.%d.webp", timeStamp, rnd))
			_, _ = target.Write(sampleImage)
			_ = target.Close()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("populated %d pictures\n", n)
}
