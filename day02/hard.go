package main

import (
	"fmt"
	"os"
	"log"
)

//интерфейс, представляющий запись данных
type Writer interface {
	Write(data[]byte)(int, error)
}

type FileWriter struct {
	file *os.File // предоставление файла
}

//структура для вывода данных в консоль
type Console struct {}

func newFile (name string) (*FileWriter, error) {
	file, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	return &FileWriter{file: file}, nil
}

func (fw *FileWriter) Write(data []byte)(int, error) {
	return fw.file.Write(data)
}

func (c Console) Write(data []byte)(int, error) {
	cons, err := fmt.Println(string(data))
	return cons, err
}

func main() {
	var writer Writer = Console{}

	writer, err := newFile("file.txt")
	if err != nil {
		log.Fatal("Error creating", err)
		return
	}

	data := "Hello, GO!"
	_, err = writer.Write([]byte(data))
	if err != nil {
		log.Fatal("Error writing", err)
		return
	}
	fmt.Println("Success")
}