package functions

import (
	"bytes"
	"database/sql"
	"reflect"
	"strconv"

	"github.com/jmoiron/sqlx"
)

func Sqlcolsinsert(typestructstring interface{}) []string {
	var sql []string
	var sql1 bytes.Buffer //欄位
	var sql2 bytes.Buffer //:欄位

	val := reflect.ValueOf(typestructstring)
	num := val.NumField()
	for i := 0; i < num; i++ {
		field := val.Type().Field(i)
		tag := field.Tag

		//fieldType := field.Type      // get struct's variable type
		//fieldName := field.Name      //get struct variable's name
		//fmt.Println(fieldType)
		//fmt.Println(fieldName)
		//tag.Get("db")  // get struct tag's name

		if tag.Get("db") != "" {
			sql1.WriteString(tag.Get("db"))
			sql1.WriteString(", ")

			sql2.WriteString(":")
			sql2.WriteString(tag.Get("db"))
			sql2.WriteString(", ")
		}
	}

	sql = append(sql, sql1.String()[:len(sql1.String())-2]) //去掉最後的 ", "
	sql = append(sql, sql2.String()[:len(sql2.String())-2]) //去掉最後的 ", "

	return sql
}

func Sqlcolsupdate(typestructstring interface{}) string {
	var sql string
	var sql1 bytes.Buffer //欄位

	val := reflect.ValueOf(typestructstring)
	num := val.NumField()
	for i := 0; i < num; i++ {
		field := val.Type().Field(i)
		tag := field.Tag

		//fieldType := field.Type      // get struct's variable type
		//fieldName := field.Name      //get struct variable's name
		//fmt.Println(fieldType)
		//fmt.Println(fieldName)
		//tag.Get("db")  // get struct tag's name

		if tag.Get("db") != "" {
			sql1.WriteString(tag.Get("db"))
			sql1.WriteString(" = :")
			sql1.WriteString(tag.Get("db"))
			sql1.WriteString(", ")
		}
	}

	sql = sql1.String()[:len(sql1.String())-2] //去掉最後的 ", "

	return sql
}

func Sqlpost(db1 *sqlx.DB, table string, typestructstring interface{}) (int, map[string]interface{}) {
	sqlt := Sqlcolsinsert(typestructstring)
	rows, err := db1.NamedExec(`insert into `+table+` (`+sqlt[0]+`) values (`+sqlt[1]+`)`, typestructstring)
	if err != nil {
		//c.JSON(404, gin.H{"Err Msg": err})
		return 404, map[string]interface{}{"Err Msg": err.Error()}
	}
	tid, err := rows.LastInsertId()
	if err != nil {
		//c.JSON(201, gin.H{"Get New Row id": "Fail"})
		return 201, map[string]interface{}{"Get New Row id": "Fail"}
	}
	//c.JSON(201, gin.H{"New Row id": tid})
	return 201, map[string]interface{}{"New Row id": strconv.FormatInt(tid, 10)}
}

func Sqlput(db1 *sqlx.DB, table string, id string, typestructstring interface{}) (int, map[string]interface{}) {
	code, msg := sqlifexist(db1, table, id)

	if code == 404 {
		return code, msg
	}

	sqlt := Sqlcolsupdate(typestructstring)
	_, err := db1.NamedExec(`update `+table+` set `+sqlt+` where tid = `+id, typestructstring)
	if err != nil {
		//c.JSON(404, gin.H{"Err Msg": err})
		return 404, map[string]interface{}{"Err Msg": err.Error()}
	}
	//c.JSON(201, gin.H{"New Row id": tid})
	return 200, map[string]interface{}{"Upadte": "Success!"}
}

func Sqldelete(db1 *sqlx.DB, table string, id string) (int, map[string]interface{}) {
	code, msg := sqlifexist(db1, table, id)

	if code == 404 {
		return code, msg
	}
	_, err := db1.Exec(`delete from ` + table + ` where tid = ` + id)
	if err != nil {
		//c.JSON(404, gin.H{"Err Msg": err})
		return 404, map[string]interface{}{"Err Msg": err.Error()}
	}
	//c.JSON(201, gin.H{"New Row id": tid})
	return 200, map[string]interface{}{"Delete": "Success!"}
}

func sqlifexist(db1 *sqlx.DB, table string, id string) (int, map[string]interface{}) {
	row := db1.QueryRow(`SELECT tid FROM `+table+` WHERE tid=?`, id)
	var tid int
	err := row.Scan(&tid)
	if err != nil {
		//c.JSON(404, gin.H{"Err Msg": err})
		return 404, map[string]interface{}{"Err Msg": err.Error()}
	} else if err == sql.ErrNoRows {
		return 404, map[string]interface{}{"Err Msg": "No Rows."}
	}
	return 200, map[string]interface{}{"Row if exist": "Exist."}
}

/*func Sqlselectall(db1 *sqlx.DB, table string) []map[string]interface{} {

	rows, err := db1.Queryx("SELECT * FROM  " + table)
	if err != nil {
	}

	results := []map[string]interface{}{}
	for rows.Next() {
		// cols is an []interface{} of all of the column results
		//cols, err = rows.SliceScan()
		//fmt.Println("cols 1:", cols)
		result := make(map[string]interface{})
		err = rows.MapScan(result)
		fmt.Println("result:", result)
		results = append(results, result)
	}

	return results
}*/
