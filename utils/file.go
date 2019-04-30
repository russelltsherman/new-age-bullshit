package utils

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
)

// FileRead - read a file
func FileRead(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("FileOpenError: %v", err)
	}
	defer f.Close()

	stat, err := os.Stat(path)
	if err != nil {
		log.Fatalf("FileStatError: %v", err)
	}

	b1 := make([]byte, stat.Size())
	n1, err := f.Read(b1)
	if err != nil {
		log.Fatalf("FileReadError: %v", err)
	}

	log.Tracef("read %d bytes from %s\n", n1, path)

	// dat, err := ioutil.ReadFile(path)
	// if err != nil {
	// 	log.Fatalf("ReadfileError: %v", err)
	// }

	// log.Printf("% bytes written to %v \n", n2, path)
	// return dat
	return b1
}

// FileWrite - write a file
func FileWrite(path string, fileBytes []byte) {
	// ensure written file is indented
	var fileInterface map[string]interface{}
	json.Unmarshal([]byte(fileBytes), &fileInterface)
	fileIndentedBytes, _ := json.MarshalIndent(fileInterface, "", "  ")

	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("CreateFileError: %v", err)
	}
	defer f.Close()

	n2, err := f.Write(fileIndentedBytes)
	if err != nil {
		log.Fatalf("WriteFileError: %v", err)
	}
	log.Tracef("% bytes written to %v \n", n2, path)
}
