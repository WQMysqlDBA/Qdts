package mpkg

import (
	"fmt"
	"log"
	"strconv"
)

func AllRowsCount(ip, port, user, passwd string, t int, ignorecountdb []string) (ret int) {
	_, _ = ConnectDB(ip, port, user, passwd)
	defer CloseDB()

	args := "("
	i := 1
	//for _, _ = range ignorecountdb {
	//
	//	if i < len(ignorecountdb) {
	//		args = args + "?,"
	//	} else {
	//		args = args + "?"
	//	}
	//	i++
	//}

	for _, v := range ignorecountdb {

		if i < len(ignorecountdb) {
			args = args + "'" + v + "',"
		} else {
			args = args + "'" + v + "'"
		}
		i++
	}

	f := "SELECT SUM(TABLE_ROWS) as tablerows FROM `information_schema`.`TABLES` WHERE TABLE_SCHEMA NOT IN "

	ignoreRegx := args + ")"
	sql := f + ignoreRegx
	// fmt.Println(sql)
	// 这里 要把 slice转为interface
	//var new interface{}
	fmt.Println(sql)
	//result, err := DB.Query(sql,ignorecountdb)
	result, err := DB.Query(sql)
	if err != nil {
		queryErr := err.Error()
		PrintLog("query of select DB name for db incur error" + queryErr)
	}
	Allrows := ""
	for result.Next() {
		_ = result.Scan(&Allrows)
	}

	ret, err = strconv.Atoi(Allrows)
	if err != nil {
		log.Println(err.Error())
	}
	if ret == 0 {
		PrintLog(fmt.Sprintf("[Info]: All db rows count is %v", "No values. No tables to dump"))
		return 1
	} else {
		PrintLog(fmt.Sprintf("[Info]: All db rows count is %v", ret))
	}

	return ret
}

func AllRowsCountold(ip, port, user, passwd string, t int, nocountdb []string) (ret int) {
	_, _ = ConnectDB(ip, port, user, passwd)
	defer CloseDB()
	GetDB(nocountdb)
	tablecount := Process(t)
	result := 0
	for _, v := range tablecount {
		result = result + v
	}
	PrintLog(fmt.Sprintf("[Info]: All db rows count is %v", result))
	//if result == 0 {
	//	result = 1
	//}
	return result

}
