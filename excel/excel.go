package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func SetTitle(xlsx *excelize.File, columnList ...string) {
	for i, axis := 0, 'A'; i < len(columnList); i++ {
		xlsx.SetCellValue("Sheet1", string(axis)+"1", columnList[i])
		axis = axis + 1
	}
}

func AddRow(xlsx *excelize.File, rowNum int, rowList ...interface{}) {
	for i, axis := 0, 'A'; i < len(rowList); i++ {
		index := fmt.Sprintf("%c%d", axis, rowNum)
		xlsx.SetCellValue("Sheet1", index, rowList[i])
		axis = axis + 1
	}
}
