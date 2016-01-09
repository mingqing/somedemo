package main

import (
	"log"

	"github.com/mingqing/godocx"
)

func main() {
	//d := godocx.NewDocXml()

	/*
		d.Test()
		err := d.Save("./data/save/")
		if err != nil {
			log.Println("err:", err)
		}
	*/

	unpackDocx()
}

func packDocx() {
	docx, err := godocx.NewDocxFile("demo2.docx")
	if err != nil {
		log.Println("err:", err)
	}

	docxParentDir := "./data/unpack/"
	err = docx.CombineTo(docxParentDir, "./data/pack/")
	if err != nil {
		log.Println("err:", err)
	}
}

func unpackDocx() {
	path := "./data/demo1.docx"
	docx, err := godocx.NewDocxFileFromPath(path)
	if err != nil {
		log.Println("err:", err)
	}

	err = docx.DecomposeTo("./data/unpack/")
	if err != nil {
		log.Println("err:", err)
	}
}
