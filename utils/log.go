package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func LogInfo(v ...interface{}) {
	folder := "logs"
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		if err := os.Mkdir(folder, os.FileMode(0755)); err != nil {
			return
		}
	}

	dt := time.Now()
	fileName := dt.Format("20060102") + "_log.log"
	path := fmt.Sprintf("%s/%s", folder, fileName)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println("Could not open file: ", path)
		return
	}

	writer := io.MultiWriter(os.Stdout, f)
	log.SetOutput(writer)
	log.Println(fmt.Sprintln(v...))
}

func LogError(v ...interface{}) {
	folder := "logs"
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		if err := os.Mkdir(folder, os.FileMode(0755)); err != nil {
			return
		}
	}

	dt := time.Now()
	fileName := dt.Format("20060102") + "_error.log"
	path := fmt.Sprintf("%s/%s", folder, fileName)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println("Could not open file: ", path)
		return
	}

	writer := io.MultiWriter(os.Stdout, f)
	log.SetOutput(writer)
	log.Println(fmt.Sprintln(v...))
}