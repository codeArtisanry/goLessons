package mysql2json

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ToJson(username string, password string, address string, database string, sqlstatement string) (string, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", username, password, address, database)
	db, err := sql.Open("mysql", dsn)

	rows, err := db.Query(sqlstatement)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	//
	tableDatas := make([]map[string]interface{}, 0)
	//1. 获取所有列
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	// 存储
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	// 遍历所有记录
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		// 存储单行数据的对象entry
		entry := make(map[string]interface{})
		// 遍历每一列
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableDatas = append(tableDatas, entry)
	}

	jsonData, err := json.Marshal(tableDatas)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
