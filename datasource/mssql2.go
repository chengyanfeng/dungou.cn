package datasource

import (
	_ "github.com/denisenkom/go-mssqldb"
	"database/sql"
	"fmt"
	"log"
)

var DbTjMssql *sql.DB

type TjMssql struct {
}

func (this *TjMssql) Init() {
	var err error
	connString := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;port=%d;encrypt=disable",
		TjConn["host"],
		TjConn["name"],
		TjConn["username"],
		TjConn["password"],
		TjConn["port"],
	)
	DbTjMssql,err = sql.Open("mssql", connString)
	if err != nil {
		panic(err.Error())
	}

}

func (this *TjMssql)Query(str string) []map[string]interface{} {
	//产生查询语句的Statement
	stmt, err := DbTjMssql.Prepare(str)
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()

	//通过Statement执行查询
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("Query failed:", err.Error())
	}

	//建立一个列数组
	cols, err := rows.Columns()
	a:=make([]interface{},0)
	var colsdata = make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		colsdata[i] = new(interface{})
		a = append(a,cols[i])
	}
	//遍历每一行
	d := make([]map[string]interface{},0)
	for rows.Next() {
		rows.Scan(colsdata...)            //将查到的数据写入到这行中
		d=append(d, printRow(a,colsdata)) //打印此行
	}
	defer rows.Close()
	return d
}

