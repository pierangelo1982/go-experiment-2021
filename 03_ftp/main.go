package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jlaffaye/ftp"
)

type Header struct {
	Comment string    // comment
	Extra   []byte    // "extra data"
	ModTime time.Time // modification time
	Name    string    // file name
	OS      byte      // operating system type
}

func zipit(source, target string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})

	return err
}

func main() {
	/* err := zipit("/var/www/www.balloonssrl.com", "/tmp/balloons_settimanale.zip")
	if err != nil {
		log.Println(err)
	} */

	c, err := ftp.Dial("xxxxx:21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	// login
	err = c.Login("xxxxx", "xxxxxx")
	if err != nil {
		log.Fatal(err)
	}
	// Do something with the FTP conn
	// data := bytes.NewBufferString("Hello World")
	/*
		data, err := ioutil.ReadFile("/tmp/balloons_settimanale.zip")
		if err != nil {
			panic(err)
		}
	*/
	file, err := os.Open("/tmp/balloons_settimanale.zip")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 1024*1024)
	scanner.Buffer(buf, 10*1024*1024)

	for scanner.Scan() {
		print(scanner.Bytes())
		err = c.Append("/Backup_VPS/balloons_settimanale.zip", bytes.NewBuffer(scanner.Bytes()))
		if err != nil {
			panic(err)
		}
	}

	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}
