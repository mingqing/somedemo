package main

import (
	"encoding/xml"
	"fmt"

	"github.com/mingqing/godocx"
)

func main() {
	//printXml()
	//packDocx()
	documentTest()
}

func documentTest() {
	d, err := godocx.NewDocXml("./data/example1/test.docs")
	if err != nil {
		fmt.Printf("create docxml err {%s}\n", err.Error())
		return
	}

	document := d.Document()
	paragh := document.AddParagraph()

	ppr := paragh.AddProperties()
	rpr := ppr.AddRunProperties()
	rpr.Bold(true)
	font := rpr.AddFont()
	font.EastAsia = "黑体"
	run := paragh.AddRunContent()
	rpr2 := run.AddRunProperties()
	rpr2.Bold(true)
	font2 := rpr2.AddFont()
	font2.Ascii = "黑体"
	font2.EastAsia = "黑体"
	font2.Hint = "eastAsia"
	run.Text("绝密★启用前")

	docByte, _ := xml.MarshalIndent(document, "", "  ")
	fmt.Println(xml.Header + string(docByte))
}

func printXml() {
	d, _ := godocx.NewDocXml("./data/example1/test.docs")
	d.Test()
}

func packDocx() {
	docx, err := godocx.NewDocxFile("demo2.docx")
	if err != nil {
		fmt.Println("err:", err)
	}

	docxParentDir := "./data/unpack/"
	err = docx.CombineTo(docxParentDir, "./data/pack/")
	if err != nil {
		fmt.Println("err:", err)
	}
}

func unpackDocx() {
	path := "./data/demo2.docx"
	docx, err := godocx.NewDocxFileFromPath(path)
	if err != nil {
		fmt.Println("err:", err)
	}

	err = docx.DecomposeTo("./data/unpack/")
	if err != nil {
		fmt.Println("err:", err)
	}
}
