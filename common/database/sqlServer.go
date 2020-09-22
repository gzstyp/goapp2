package database

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"time"
)

func InitSqlServer() {
	var isdebug = true
	var server = "192.168.3.209"
	var port = 1433
	var user = "sa"
	var password = "GZS5630typ"
	var database = "gsomewsd2020820"

	//连接字符串
	connString := fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", server, port, database, user, password)
	if isdebug {
		fmt.Println(connString)
	}
	//建立连接
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
	}
	defer conn.Close()

	sql := `SELECT cl.DeviceID,cl.Val value,cl.SeqType,de.Name name from Channel cl LEFT JOIN Device de on de.ID = cl.DeviceID`

	//产生查询语句的Statement
	stmt, err := conn.Prepare(sql)
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()
	//通过Statement执行查询
	rows, err := stmt.Query(sql)
	if err != nil {
		log.Fatal("Query failed:", err.Error())
	}
	//建立一个列数组
	cols, err := rows.Columns()
	var colsdata = make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		colsdata[i] = new(interface{})
		fmt.Print(cols[i])
		fmt.Print("\t")
	}
	fmt.Println()

	//遍历每一行
	for rows.Next() {
		rows.Scan(colsdata...) //将查到的数据写入到这行中
		PrintRow(colsdata)     //打印此行
	}
	defer rows.Close()
}

//打印一行记录，传入一个行的所有列信息
func PrintRow(colsdata []interface{}) {
	for _, val := range colsdata {
		switch v := (*(val.(*interface{}))).(type) {
		case nil:
			fmt.Print("NULL")
		case bool:
			if v {
				fmt.Print("True")
			} else {
				fmt.Print("False")
			}
		case []byte:
			fmt.Print(string(v))
		case time.Time:
			fmt.Print(v.Format("2016-01-02 15:05:05.999"))
		default:
			fmt.Print(v)
		}
		fmt.Print("\t")
	}
	fmt.Println()
}
