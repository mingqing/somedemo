package main

import (
	"encoding/xml"
	"fmt"

	"github.com/mingqing/godocx"
	"github.com/pborman/uuid"
)

func main() {
	//printXml()
	//packDocx("/tmp/data/9240893b-bcbd-11e5-9aa0-0021cc684b34", "/tmp/data", "demo3.docx")
	//packDocx("./data/demo2/fail", "./", "fail-test.docx")
	packDocx("./data/unpack", "./", "fail-test.docx")
	//unpackDocx()
	//documentTest()
}

func documentTest() {
	randomId := uuid.NewUUID()
	d, err := godocx.NewDocXml("./data/example1/"+randomId.String(), randomId.String()+".docs")
	if err != nil {
		fmt.Printf("create docxml err {%s}\n", err.Error())
		return
	}

	document := d.Document()
	paragh := document.AddParagraph()
	document.Save(d.Dir)

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
	d, _ := godocx.NewDocXml("./data/example1", "test.docs")
	d.Test()
}

func packDocx(docxParentDir, packTo, name string) {
	docx, err := godocx.NewDocxFile(name)
	if err != nil {
		fmt.Println("err:", err)
	}

	err = docx.CombineTo(docxParentDir, packTo)
	if err != nil {
		fmt.Println("err:", err)
	}
}

func unpackDocx() {
	path := "./data/demo2/5be1eae1-bd8b-11e5-896e-782bcb24d315.docx"
	docx, err := godocx.NewDocxFileFromPath(path)
	if err != nil {
		fmt.Println("err:", err)
	}

	err = docx.DecomposeTo("./data/unpack/")
	if err != nil {
		fmt.Println("err:", err)
	}
}
