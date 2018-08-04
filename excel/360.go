package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"

	"errors"
	"sort"
	"strconv"
	"strings"
)

const excelFile = `E:\work\e\Task.xlsx`

var column = []string{"display", "field", "type", "null", "index", "use"}

type ExcelParse struct {
	xlsx   *excelize.File
	heads  []string
	sheets map[string]SheetMeta
}

type SheetMeta struct {
	sheetName string
	isNew     bool
}

func (this *ExcelParse) read() {
	xlsx, err := excelize.OpenFile(excelFile)
	if err != nil {
		fmt.Println("open error", err)
		return
	}
	this.xlsx = xlsx

	shetMap := xlsx.GetSheetMap()
	var sheetNames = make(map[string]int)

	for k, v := range shetMap {
		sheetNames[v] = k
		fmt.Println(k, v)
	}
	nameSheetMap := make(map[string]string)

	for name := range sheetNames {
		if !strings.HasPrefix(name, "_") {
			if _, ok := nameSheetMap[name]; !ok {
				nameSheetMap[name] = ""
			}

		} else {
			sheetName := strings.TrimPrefix(name, "_")
			if _, ok := nameSheetMap[name]; ok {

				panic(errors.New(fmt.Sprintf("已经存在相同的sheet名称:%s", name)))
			}

			nameSheetMap[sheetName] = name
		}
	}
	this.sheets = make(map[string]SheetMeta)
	for key, value := range nameSheetMap {
		if value == "" {
			newSheet := "_" + key
			this.sheets[key] = SheetMeta{newSheet, true}
			fmt.Println("Start create ", newSheet)
			nameSheetMap[key] = newSheet
			//sheet := xlsx.Sheet[key]
		} else {
			this.sheets[key] = SheetMeta{value, false}
		}
	}

	fmt.Println(this.sheets)
	this.copyCellValue()

	save(this.xlsx)
	fmt.Println(nameSheetMap)

	fmt.Println(sheetNames)
}

func (this *ExcelParse) copyCellValue() {
	xlsx := this.xlsx

	for sheet, meta := range this.sheets {
		row := xlsx.GetRows(sheet)[0]
		heads := make(map[string]int)
		fields := make(map[string]int)

		for i, value := range row {
			fmt.Println(value)
			if _, exist := heads[value]; !exist {
				head := strings.TrimSpace(value)
				if head == "" {
					continue
				}
				heads[head] = i
			}
		}

		row = xlsx.GetRows(sheet)[1]

		for i, value := range row {
			fmt.Println(value)
			if _, exist := heads[value]; !exist {
				head := strings.TrimSpace(value)
				if head == "" {
					continue
				}
				fields[head] = i
			}
		}
		this.makeMeta(sortMap(heads), sortMap(fields), meta)
	}
}

func sortMap(m map[string]int) []string {
	values := []int{}
	tempMap := make(map[int]string)

	for k, v := range m {
		values = append(values, v)
		tempMap[v] = k
	}

	sort.Ints(values)
	result := []string{}

	for _, i := range values {
		result = append(result, tempMap[i])
	}
	return result

}

func (this *ExcelParse) makeMeta(heads []string, fields []string, meta SheetMeta) {
	fmt.Println(heads, fields)
	if meta.isNew {
		this.xlsx.NewSheet(meta.sheetName)
		this.xlsx.SetSheetRow(meta.sheetName, "A1", &column)

		var i = 2
		for _, key := range heads {
			s := "A" + strconv.Itoa(i)
			i++
			this.xlsx.SetCellValue(meta.sheetName, s, key)
		}
		i = 2
		for _, key := range fields {
			s := "B" + strconv.Itoa(i)
			i++
			this.xlsx.SetCellValue(meta.sheetName, s, key)
		}

	} else {
		rows := this.xlsx.GetRows(meta.sheetName)

		rowMap := make(map[string][]string)
		for i, row := range rows {
			if i == 0 {
				continue
			}
			if _, e := rowMap[row[0]]; !e {
				rowMap[row[0]] = row
			}
		}

		rowLen := len(rows)

		println("----", rowLen)
		// 清空
		temp := &[]string{""}
		for i := 0; i < rowLen; i++ {
			this.xlsx.SetSheetRow(meta.sheetName, "A"+strconv.Itoa(i+1), temp)
		}

		this.xlsx.SetSheetRow(meta.sheetName, "A1", &column)
		var i = 2

		for _, key := range heads {
			s := "A" + strconv.Itoa(i)
			i++
			if row, exist := rowMap[key]; exist {
				this.xlsx.SetSheetRow(meta.sheetName, s, &row)
			} else {
				this.xlsx.SetCellValue(meta.sheetName, s, key)
			}
		}

		i = 2
		for _, key := range fields {
			s := "B" + strconv.Itoa(i)
			i++
			this.xlsx.SetCellValue(meta.sheetName, s, key)
		}
	}
}

func save(xlsx *excelize.File) {
	err := xlsx.Save()
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	e := &ExcelParse{}
	e.read()

}
