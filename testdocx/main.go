package main

import (
	"log"

	"github.com/mingqing/godocx"
)

func main() {
	d := godocx.NewDocXml()
	d.Test()
}

func testDocxFile() {
	path := "./data/demo1.docx"
	docx, err := godocx.NewDocxFileFromPath(path)
	if err != nil {
		log.Println("err:", err)
	}

	/*
		err = docx.DecomposeTo("./data/unpack/")
		if err != nil {
			log.Println("err:", err)
		}
	*/

	docxParentDir := "./data/unpack/"
	err = docx.CombineTo(docxParentDir, "./data/pack/")
	if err != nil {
		log.Println("err:", err)
	}
}
