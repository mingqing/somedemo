package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/mingqing/textimg"
	"golang.org/x/net/html"
)

type imageObject struct {
	width   int
	height  int
	format  string
	content string
	point   image.Point
}

type textLine struct {
	index int
	imgs  []imageObject
	lines []string
}

func main() {
	htmlContent, err := ioutil.ReadFile("./data/demo3.html")
	if err != nil {
		fmt.Printf("err {%s}", err.Error())
		return
	}

	//fmt.Println("html:", string(htmlContent))
	doc, err := html.Parse(bytes.NewReader(htmlContent))
	if err != nil {
		fmt.Println("parse err:", err)
		return
	}

	//for c := doc.FirstChild; c != nil; c = c.NextSibling {

	text := &textLine{index: 0, lines: make([]string, 0), imgs: make([]imageObject, 0)}
	text.lines = append(text.lines, "")

	parse(doc, text)

	rgba := image.NewRGBA(image.Rect(0, 0, 1099, 35*len(text.lines)))
	timg := textimg.New(rgba, image.White)
	timg.SetFontFromPath("./fonts/simsun.ttc")
	timg.SetFontSize(14)
	timg.DrawDstimg(image.Black, text.lines)

	for _, v := range text.imgs {
		timg.AddImageFromHtmlSrcBase64(v.point, v.content)
	}

	//temp := bytes.NewBuffer(make([]byte, 0))
	//temp = timg.PNG()
	//fmt.Println("temp:", temp)

	outFile, _ := os.Create("out.png")
	defer outFile.Close()

	b := bufio.NewWriter(outFile)
	png.Encode(b, rgba)
	b.Flush()

	for _, v := range text.lines {
		fmt.Println("v:", v)
	}
}

func parse(n *html.Node, text *textLine) {
	switch n.Type {
	case html.TextNode:
		if strings.TrimSpace(n.Data) != "" {
			//fmt.Printf("data {%s}\n", n.Data)

			text.lines[text.index] += n.Data
		}
	case html.ElementNode:
		switch n.Data {
		case "html", "head", "body":
		case "table", "tbody", "td", "tr":
		case "img":
			//fmt.Printf("data {%s}\n", "<img>")
			imgobj := imageObject{}

			width := 0
			height := 0

			for _, v := range n.Attr {
				//fmt.Println("key:", v.Key, "value:", v.Val)

				if v.Key == "width" {
					width, _ = strconv.Atoi(v.Val)
				}
				if v.Key == "height" {
					height, _ = strconv.Atoi(v.Val)
				}
				if v.Key == "src" {
					imgobj.content = v.Val
				}
			}

			imgX := 0
			imgY := text.index*14*3 - 7
			if (height / 14) == 1 {
				imgY += 7
			}

			str := text.lines[text.index]
			for len(str) > 0 {
				_, size := utf8.DecodeRuneInString(str)
				switch size {
				case 1:
					imgX += 14 / 2
				case 2, 3:
					imgX += 14
				default:
					imgX += 14 / 2
				}

				str = str[size:]
			}

			for i := 0; i < (width / 14); i++ {
				text.lines[text.index] += "  "
			}
			if (width % 14) != 0 {
				text.lines[text.index] += "  "
			}

			fmt.Printf("img width {%d} height {%d} X {%d} Y {%d}\n", width, height, imgX, imgY)

			imgobj.point = image.Point{X: -imgX, Y: -imgY}
			text.imgs = append(text.imgs, imgobj)
			//text.lines[text.index] += "<img>"
		case "span":
		case "p":
			text.index += 1
			text.lines = append(text.lines, "")

			//fmt.Println("---new line---")
		default:
			if strings.TrimSpace(n.Data) != "" {
				//fmt.Printf("n.Data {%s}\n", n.Data)

				text.lines[text.index] += n.Data
			}
		}
		//default:
		//fmt.Printf("default type {%d} data {%s}\n", n.Type, n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parse(c, text)
	}
}
