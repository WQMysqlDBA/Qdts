package mpkg

import (
	"Qdts/rwfile"
	"fmt"
	"strconv"
	"strings"
)

//var dstconn *sql.DB

//Modify data adaptively
func ModifyData(dumpdir string) {
	text := "############# Modify data adaptively #############\n" +
		"  1、自动将MyISAM、MEMORY引擎表转换为InnoDB表\n" +
		"  2、对于Mysql 5.6 版本 `ROW_FORMAT=FIXED` 表属性已经在5.7以后的版本废弃，自动修改为5.7以后的默认值"
	PrintLog(text)
	modifyRowFixed(dumpdir)
	modifyInnodbtable(dumpdir) // 执行完这一步，mydumper读取修改的文件有问题，直接在这里把对应的内容在数据库执行一下吧。
}

func modifyRowFixed(dir string) {
	cmd1 := "ls " + dir + "  | grep 'schema.sql'|wc -l"
	out1, _ := strconv.Atoi(string(Cmd(cmd1, true, false)))
	//fmt.Println(cmd1, out1)

	if out1 == 1 {
		//out1 恒为0  所以这里怎么改？ 直接注释掉对out1的判断逻辑即可
		//if out1 == 0 {
		PrintLog("当前备份的数据库不存在表")
	} else {
		cmd2 := "ls " + dir + "  | grep 'schema.sql'"
		out2 := string(Cmd(cmd2, true, false))
		tableschemafile := strings.Split(out2, "\n")
		//fmt.Println(tableschemafile,"debug")
		const (
			_engineflag    = "ENGINE="
			_rowformatflag = "ROW_FORMAT=FIXED"
			_newflag       = " "
		)
		f := 0

		for _, v := range tableschemafile {
			var rowsContext []string
			needModify := false
			if v != "" {
				//fmt.Println(v)
				schemaFile := dir + "/" + v
				//fmt.Println(schemaFile)
				//flag := "s/ROW_FORMAT=FIXED//g"
				//cmd := "sed -i -e " + "\"" + flag + "\"" + " " + schemaFile
				//fmt.Println(cmd)
				//fmt.Println(schemaFile)
				rows := rwfile.ReadFile(schemaFile)
				//f1, _ :=os.Open(schemaFile)
				//rows := bufio.NewScanner(f1)
				for rows.Scan() {
					msg := rows.Text()

					// ENGINE ROW_FORMAT=FIXED 一定是在最后一行的
					if strings.Contains(msg, _engineflag) && strings.Contains(msg, _rowformatflag) {
						f++
						needModify = true
						//Cmd(cmd, true,false)
						msg = strings.Replace(msg, _rowformatflag, _newflag, 1)
						PrintLog(fmt.Sprintf("调整schema-file%s", schemaFile))
					}

					//if 0 == len(msg) || msg == "\r\n"{
					//	fmt.Println("这是个空行",schemaFile,a)
					//}else {
					//	rowsContext = append(rowsContext, msg)
					//}
					/* 这块儿文件里有个^M,不知道怎么处理*/
					rowsContext = append(rowsContext, msg)
				}
				if needModify {
					schemaSQL := strings.Join(rowsContext, "\n")
					schemaSQL = schemaSQL + "\n"
					// 真相了，C++ 代码中 根据';\n来拆解dump文件'，我这么拼肯定不对。。。。。
					if ok := rwfile.WriteFile(schemaFile, schemaSQL); ok {
						PrintLog(fmt.Sprintf("修改table schema文件: %s success，内容为: %s", schemaFile, schemaSQL))
					}
				}
			}
		}
		modifyRowFixedresult := fmt.Sprintf("调整`ROW_FORMAT=FIXED`参数完毕，共有%v个符合条件的表\n", f)
		//modifyRowFixedresult := fmt.Sprintf("调整`ROW_FORMAT=FIXED`参数完毕")
		PrintLog(modifyRowFixedresult)
	}
}

func modifyInnodbtable(dir string) {
	info := "############# 调整非innodb表为innodb表 #############"
	if len(UnInnodbTableInfo) == 0 {
		PrintLog("备份文件中没有非InnoDB表")
	} else {
		//var txt string
		//for _, v := range UnInnodbTableInfo {
		//	txt = txt + v
		//}
		info = info + "\n" + "  1、需要调整的表对应的文件为: "
		PrintLog(info)

		for _, v := range UnInnodbTableInfo {
			tbSchema := strings.Split(v, "===")[0]
			tbNAME := strings.Split(v, "===")[1]
			tbEngine := strings.Split(v, "===")[2]
			//schemaFile := tbSchema + "." + tbNAME + "-schema.sql"
			//text := "    " + fileName
			//PrintLog(text)
			//flag := "s/" + tbEngine + "/InnoDB/g"
			//
			//cmd := "sed -i -e " + "\"" + flag + "\"" + " " + dir + "/" + fileName
			////fmt.Println(cmd)
			//Cmd(cmd, true, false)

			_engineflag := "ENGINE="
			_noinoodbflag := "ENGINE=" + tbEngine
			_innodbflag := "ENGINE=InnoDB"

			needModify := false
			var rowsContext []string
			schemaFile := dir + "/" + tbSchema + "." + tbNAME + "-schema.sql"
			rows := rwfile.ReadFile(schemaFile)
			fmt.Println(schemaFile)

			//f1, _ :=os.Open(schemaFile)
			//rows := bufio.NewScanner(f1)
			for rows.Scan() {

				msg := rows.Text()
				if strings.Contains(msg, _engineflag) {
					needModify = true
					msg = strings.Replace(msg, _noinoodbflag, _innodbflag, 1)
					PrintLog(fmt.Sprintf("调整schema-file %s", schemaFile))
				}
				//if 0 == len(msg) || msg == "\r\n" {
				//	continue
				//}
				rowsContext = append(rowsContext, msg)
			}

			if needModify {
				//var newrowsContext []string
				//copy(newrowsContext, rowsContext)
				//for k, v := range newrowsContext {
				//	fmt.Println(k, v)
				//}
				//for index, value := range rowsContext {
				//	if len(value) == 0 {
				//		//a = append(a[:index], a[index+1:]...)
				//		newrowsContext = append(newrowsContext[:index], newrowsContext[index+1:]...)
				//	}
				//
				//}
				//for k, v := range newrowsContext {
				//	fmt.Println(k, v)
				//}

				schemaSQL := strings.Join(rowsContext, "\n")
				schemaSQL = schemaSQL + "\n"
				if ok := rwfile.WriteFile(schemaFile, schemaSQL); ok {
					PrintLog(fmt.Sprintf("修改table schema文件: %s success，内容为:\n %s", schemaFile, schemaSQL))
				}
			}
		}
	}

}

/* 如下是对备份文件对说明 */
/* 1、按照schema来看，每个DB都有一个 DBNAME-schema-create.sql 文件 即create database ... */
/* 2、每个表各有两个文件，即表定义文件: DBNAME.TABLENAME-schema.sql和 数据文件: DBNAME.TABLENAME.sql */
/* 对表存储引擎和ROW_FORMAT的修改即修改表定义文件*/
//test1.chtest-schema.sql
//test1.chtest-schema-triggers.sql
//test1.myview-schema.sql
//test1.myview-schema-view.sql
//test1-schema-create.sql
//test1-schema-post.sql
