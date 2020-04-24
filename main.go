package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	readCurrentDir()

}

func readCurrentDir() {
	dirname := os.Getenv("FOLDER_PATH")
	file, err := os.Open(dirname)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0)

	for _, files := range list { // list all files in directory and calcuate md5sum
		apsoluteFilPath := path.Join(dirname, files)
		files, err := os.Open(apsoluteFilPath)
		if err != nil {
			log.Fatal(err)
		}
		defer files.Close()
		h := md5.New()
		// if _, err := io.Copy(h, f); err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Printf("%x\n", h.Sum(nil))
		fmt.Println("File: ", files, "md5sum: ", hex.EncodeToString(h.Sum(nil)))
	}
}
