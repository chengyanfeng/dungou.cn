package datasource

import (
	_ "github.com/denisenkom/go-mssqldb"
	"database/sql"
	"fmt"
	"log"
	"time"
)

var DbMssql *sql.DB

type Mssql struct {
}

func (this *Mssql) Init() {
	var err error
	connString := fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;port=%d;encrypt=disable",
		SqlServerConn["host"],
		SqlServerConn["name"],
		SqlServerConn["username"],
		SqlServerConn["password"],
		SqlServerConn["port"],
	)
	DbMssql,err = sql.Open("mssql", connString)
	if err != nil {
		panic(err.Error())
	}
}

func (this *Mssql)Query(str string) []map[string]interface{} {
	//产生查询语句的Statement
	stmt, err := DbMssql.Prepare(str)
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

func printRow(c []interface{},colsdata []interface{}) map[string]interface{}{
	a :=make(map[string]interface{})
	for i, val := range colsdata {
		switch v := (*(val.(*interface{}))).(type) {
		case nil:
			fmt.Print("NULL")
		case bool:
			if v {
				a[c[i].(string)]="1"
			} else {
				a[c[i].(string)]="0"
			}
		case []byte:
			a[c[i].(string)]=string(v)
		case time.Time:
			a[c[i].(string)]=v.Format("2006-01-02 15:04:05")
		default:
			a[c[i].(string)]=v
		}
	}
	return a
}
