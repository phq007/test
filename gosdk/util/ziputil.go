package util

import (
	"archive/zip"
	"log"
	"os"
	"io"
)

//Upzipfile unzip file
func Upzipfile(srcfile string, dstdir string) {
	rc, err := zip.OpenReader(srcfile)
	log.Println("upzipfile..")
	os.Mkdir(dstdir, os.ModePerm)
	if err != nil {
		log.Println(err.Error())
		defer rc.Close()
	}
	for _, _file := range rc.File {
		log.Println(_file.Name)

		f, _ := _file.Open()

		desfile, err1 := os.OpenFile(dstdir+_file.Name, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err1 == nil {
			log.Println("OK")
			io.CopyN(desfile, f, int64(_file.UncompressedSize64))
			desfile.Close()
		} else {
			log.Println(err1.Error())
			defer desfile.Close()
		}
	}
}