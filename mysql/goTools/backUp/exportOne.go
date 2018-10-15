package backUp

import (
	"fmt"
	"time"
	"../common"
	"../db"
)

func ExportOne(fields common.DbConnFields, workDir string, ch chan <- string, flag common.OpFlag) {
	var fileName string
	if fields.FileAlias == "" {
		fileName = workDir + fields.DbName + "-" + time.Now().Format("2006-01-02") + ".sql"
	}else{
		fileName = workDir + fields.FileAlias + "-" + time.Now().Format("2006-01-02") + ".sql"
	}
	fmt.Println("Export ", fields.DbName , "\t start at \t", time.Now().Format("2006-01-02 15:04:05"))

	if err := db.TestDbConn(fields); err != nil {
		ch <- fmt.Sprintln("Error: ", fields.DbName, "\t test db connect throw, \t", err)
		return
	}

	setSqlHeader(fields, fileName)

	if flag.Tables {
		err := exportTables(fileName, fields, flag)
		if err != nil {
			ch <- fmt.Sprintln("Error: ", fields.DbName, "\t export tables throw, \t", err)
			return
		}
	}

	if flag.Views {
		err := exportViews(fileName, fields)
		if err != nil {
			ch <- fmt.Sprintln("Error: ", fields.DbName, "\t export views throw, \t", err)
			return
		}
	}

	if flag.Funcs {
		err := exportFuncs(fileName, fields)
		if err != nil {
			ch <- fmt.Sprintln("Error: ", fields.DbName, "\t export funcs throw, \t", err)
			return
		}
	}

	ch <- fmt.Sprintln("Export ", fields.DbName, "\t success at \t", time.Now().Format("2006-01-02 15:04:05"))
}