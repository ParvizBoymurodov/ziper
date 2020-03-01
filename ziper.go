package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"sync"
)

func main() {
	fil, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := fil.Close()
		if err != nil {
			log.Printf("Can't close file; %v", err)
		}
	}()
	log.SetOutput(fil)

	file := os.Args[1:]
	archivator(file)
	conArchivator(file)
}

func zipFlies(fileName, files string) {
	fileZip, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("can't create file %v", err)
	}
	defer func() {
		err := fileZip.Close()
		if err != nil {
			log.Fatalf("can't close file %v", err)
		}
	}()
	writerZip := zip.NewWriter(fileZip)
	defer func() {
		err := writerZip.Close()
		if err != nil {
			log.Printf("can't write zip %v", err)
		}
	}()
	zipFile, err := os.Open(fileWay + files)
	if err != nil {
		log.Printf("can't open  file: %v", err)
	}
	defer func() {
		err := zipFile.Close()
		if err != nil {
			log.Printf("can't close zip file: %v", err)
		}
	}()

	fileInfo, err := zipFile.Stat()
	if err != nil {
		log.Printf("cannot get information  %v", err)
	}
	infoHeader, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		log.Printf("can't get header info %v", err)
	}
	infoHeader.Name = files
	infoHeader.Method = zip.Deflate
	writer, err := writerZip.CreateHeader(infoHeader)
	if err != nil {
		log.Printf("can't  create header %v", err)
	}
	if _, err := io.Copy(writer, zipFile); err != nil {
		log.Printf("can't copy %v", err)
	}
}

func archivator(file []string) {
	for _, files := range file {
		fileName := fileWay + archiveWay + files + ziper
		zipFlies(fileName, files)
	}
}

func conArchivator(file []string) {
	wait := sync.WaitGroup{}
	for _, files := range file {
		filename := fileWay + conArchiveWay + files + ziper
		wait.Add(1)
		go func( fileName string) {
			defer wait.Done()
			zipFlies(filename, files)
		}( files)

	}
	wait.Wait()
}
