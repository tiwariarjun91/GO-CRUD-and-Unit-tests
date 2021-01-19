package config

import(
	"fmt"
	"database/sql"



	_ "github.com/go-sql-driver/mysql"
)
var Db *sql.DB

func Connect_db() /*(db *sql.DB)*/{
	var err error

	Db, err = sql.Open("mysql", "root:Qwerty@2412@tcp(127.0.0.1:3306)/applicationproject")
	if err!= nil{
		panic(err.Error())
	}
	

	/*defer db.Close()*/
	fmt.Println("Database connected")
	//return Db
}

