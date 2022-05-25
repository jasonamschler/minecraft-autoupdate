package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	cmdInput := os.Args[1:]
	fmt.Println(cmdInput)

	argFileInput := os.Args[1]
	argFileVersion := os.Args[2]

	workDir := "C:\\home\\minecraft\\game\\"
	workName := workDir + "server.jar"
	archiveDir := "C:\\home\\minecraft\\game\\downloads\\"
	archiveName := "minecraft_server." + argFileVersion + ".jar"

	err := DownloadFile(workName, argFileInput)
	if err != nil {
		panic(err)
	}
	fmt.Println("Downloaded: " + argFileInput)
	fmt.Println(argFileVersion)

	bytesCopied, errCopy := customCopy(workName, archiveDir+archiveName)
	if errCopy != nil {
		panic(errCopy)
	}
	println("Copied", bytesCopied)
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func customCopy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
