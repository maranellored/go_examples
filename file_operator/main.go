package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func main() {
	// Map to hold file contents
	contents := make(map[string]string)

	fmt.Println("Enter directory to store creds")
	dirName := ""
	fmt.Scanln(&dirName)

	fmt.Println("Enter file to store creds in")
	fileName := ""
	fmt.Scanln(&fileName)

	fmt.Println("Enter access key")
	accessKey := ""
	fmt.Scanln(&accessKey)

	fmt.Println("Enter secret key")
	secretKey := ""
	fmt.Scanln(&secretKey)

	absPath := filepath.Join(dirName, fileName)
	contents["access-key"] = accessKey
	contents["secret-key"] = secretKey

	d, err := yaml.Marshal(&contents)
	if err != nil {
		panic(err)
	}

	fd, err := os.Create(absPath)
	if err != nil {
		panic(err)
	}
	// Close after we're done
	defer func() {
		if err := fd.Close(); err != nil {
			panic(err)
		}
	}()

	fmt.Println(d)
	if _, err := fd.Write(d); err != nil {
		panic(err)
	}

	reader, err := os.Open(absPath)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := reader.Close(); err != nil {
			panic(err)
		}
	}()

	buffer, err := ioutil.ReadFile(absPath)
	if err != nil {
		panic(err)
	}

	fmt.Println(buffer)
	readMap := make(map[string]string)
	err = yaml.Unmarshal(buffer, readMap)
	if err != nil {
		panic(err)
	}

	for key, value := range readMap {
		fmt.Printf("%s => %s", key, value)
	}
}
