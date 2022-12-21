package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"io/ioutil"
	"log"
	"misa-generator/helper"
	"os"
	"path"
	"strings"
)

const (
	TitleSplitBySpaceThreshold = 25
	TextSplitBySpaceThreshold  = 40
	TextMaximumLine            = 5
)

func main() {
	preprocess()
	titles, texts := textProcessing()
	processTextToImage(titles, texts)
}

func preprocess() {
	dir, err := ioutil.ReadDir("output")
	if err != nil {
		panic(err)
	}
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{"output", d.Name()}...))
	}
}

func splitByNearestSpace(s string, threshold int) []string {
	parts := strings.Split(s, " ")
	results := make([]string, 0)
	temp := ""
	for _, data := range parts {
		if len(temp)+len(data) < threshold {
			if len(temp) == 0 {
				temp = data
			} else {
				temp = temp + " " + data
			}
		} else {
			results = append(results, temp)
			temp = data
		}
	}

	if len(temp) != 0 {
		results = append(results, temp)
	}

	return results
}

func textProcessing() ([]string, []string) {
	content, err := ioutil.ReadFile("assets\\lyric\\test.txt")
	if err != nil {
		log.Fatal(err)
	}
	contents := strings.Split(string(content), "\r\n")

	titles := []string{}
	res := []string{}
	for i, data := range contents {
		if len(data) == 0 {
			continue
		}
		if i != 0 {
			texts := splitByNearestSpace(data, TextSplitBySpaceThreshold)
			res = append(res, texts...)
		} else {
			texts := splitByNearestSpace(data, TitleSplitBySpaceThreshold)
			titles = append(titles, texts...)
		}
	}

	for _, data := range res {
		fmt.Println(len(data), "  -  ", data)
	}
	return titles, res
}

func processTextToImage(titles []string, texts []string) {
	fileNum := 0
	tempTexts := make([]string, 0)
	for i := 0; i < len(texts); i++ {
		if i != 0 && i%TextMaximumLine == 0 {
			filePath := fmt.Sprintf("output\\test-%d.png", fileNum)
			generateImage(titles, tempTexts, filePath)
			tempTexts = make([]string, 0)
			fileNum++
		}
		tempTexts = append(tempTexts, texts[i])
	}
	if len(tempTexts) != 0 {
		filePath := fmt.Sprintf("output\\test-%d.png", fileNum)
		generateImage(titles, tempTexts, filePath)
		tempTexts = make([]string, 0)
		fileNum++
	}
}

func generateImage(titles []string, texts []string, filePath string) {
	const W = 1280
	const H = 720
	dc := gg.NewContext(W, H)

	im, err := gg.LoadPNG("template\\template-lagu.png")
	if err != nil {
		panic(err)
	}
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.DrawImage(im, 0, 0)

	// Draw title
	titleConfig := &helper.FontConfig{}
	titleConfig = titleConfig.ToTitleConfig(len(titles))
	if err = dc.LoadFontFace("assets\\font\\roboto\\RobotoCondensed-Bold.ttf", titleConfig.FontSize); err != nil {
		panic(err)
	}
	for i, data := range titles {
		y := titleConfig.BaseY + (float64(i) * titleConfig.IncrementY)
		dc.DrawStringAnchored(data, titleConfig.BaseX, y, 0, 0)
	}

	textConfig := &helper.FontConfig{}
	textConfig = titleConfig.ToTextConfig()
	if err = dc.LoadFontFace("assets\\font\\roboto\\Roboto-Light.ttf", textConfig.FontSize); err != nil {
		panic(err)
	}
	for i, data := range texts {
		y := textConfig.BaseY + (float64(i) * textConfig.IncrementY)
		dc.DrawStringAnchored(data, textConfig.BaseX, y, 0, 0)
	}
	dc.SavePNG(filePath)
}
