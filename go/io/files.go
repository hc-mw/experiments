package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

func WriteToFile(fileName string, content any) error {
	n := rand.Intn(100)
	tmp := fmt.Sprintf("%s.tmp.%d", fileName, n)
	// open file for writing
	file, err := os.OpenFile(tmp, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
	if err != nil {
		return err
	}

	defer func() {
		file.Close()
		if err != nil {
			os.Remove(tmp)
		}
	}()

	// write content to file
	if err := json.NewEncoder(file).Encode(content); err != nil {
		return err
	}
	if err := file.Sync(); err != nil {
		return err
	}

	return os.Rename(tmp, fileName)
}
