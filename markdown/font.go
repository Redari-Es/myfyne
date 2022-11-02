package main

// import (
// 	"os"
// 	"strings"
//
// 	"github.com/flopp/go-findfont"
// )
//
// // 查找当前目录库
// func init() {
// 	fontPaths := findfont.List()
// 	for _, path := range fontPaths {
// 		if strings.Contains(path, "msyh.ttf") || strings.Contains(path, "simhei.ttf") || strings.Contains(path, "simsun.ttc") || strings.Contains(path, "simkai.ttf") {
// 			os.Setenv("FYNE_FONT", path)
// 			break
// 		}
// 	}
// }
//
import (
	"fmt"
	"os"

	"github.com/flopp/go-findfont"
	"github.com/golang/freetype/truetype"
)

// 查找系统fonts库
func font() {
	fontPath, err := findfont.Find("simsun.ttc")
	if err != nil {
		panic(err)
	}
	//	fmt.Printf("Found 'arial.ttf' in '%s'\n", fontPath)

	// load the font with the freetype library
	//	fontData, err := ioutil.ReadFile(fontPath)
	fontData, err := os.ReadFile(fontPath)
	if err != nil {
		panic(err)
	}
	font, err := truetype.Parse(fontData)
	if err != nil {
		fmt.Println(font)
		panic(err)
	}
	os.Setenv("FYNE_FONT", fontPath)

	// use the font...
}
