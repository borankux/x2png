package main

import (
	"fmt"
	"golang.org/x/image/webp"
	"image/png"
	"log"
	"os"
	"regexp"
	"sync"
)

func ConvertImage(in []string, out string, delete bool) {
	for i := range in {
		entry := in[i]
		dir, err := isDir(entry)
		if err != nil {
			log.Printf("%s\n", err)
			continue
		}

		if dir {
			if err = convertFromDirectory(entry, out, delete); err != nil {
				log.Printf("convert dir:%s", err)
			}
			continue
		}

		if err = convertFromFile(entry, out, delete); err != nil {
			log.Printf("convert file:%s， is dir:%v", err.Error(), dir)
		}
	}
}

func convertFromFile(entry string, out string, delete bool) error {
	file, err := os.Open(entry)
	if err != nil {
		return err
	}
	decoded, err := webp.Decode(file)
	_ = file.Close()
	if err != nil {
		return err
	}

	destination, err := os.Create(out + "/" + entry + ".png")
	defer destination.Close()

	if delete {
		_ = os.Remove(entry)
	}
	fmt.Printf("converted :%s\n", entry+".png")
	return png.Encode(destination, decoded)
}

func convertFromDirectory(dir string, out string, delete bool) error {

	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	wg := sync.WaitGroup{}

	for _, e := range entries {
		isMatch, _ := regexp.MatchString(".webp$", e.Name())
		if !isMatch {
			continue
		}

		wg.Add(1)
		go func(fileName string) {
			err = convertFromFile(fileName, out, delete)
			if err != nil {
				log.Printf("%s:%s", fileName, err.Error())
			}
			wg.Done()
		}(e.Name())
	}
	wg.Wait()
	return nil
}

func isDir(entry string) (bool, error) {
	fileInfo, err := os.Stat(entry)
	if err != nil {
		return false, fmt.Errorf("failed to open dir:%w", err)
	}
	return fileInfo.IsDir(), err
}
