package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"strconv"
	"testing"
)

func TestReadExcel(t *testing.T) {
	path := "./testData.xlsx"
	file, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Printf("读取文件[%s]发生错误[%v]", path, err)
		os.Exit(-1)
	}

	sheetMap := file.GetSheetMap()
	for _, v := range sheetMap {
		row := file.GetRows(v)
		var sum = 0.0
		for _, cell := range row {
			c, err := strconv.ParseFloat(cell[0], 64)
			if err != nil {
				fmt.Println("转换失败")
				c = 0.0
			}
			sum += c
		}
		fmt.Println(sum)
	}
}
