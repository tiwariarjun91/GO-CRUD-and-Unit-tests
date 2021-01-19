package main

import(
	//"net/http"
	//"html/template"
	//"crud"
	//"config"
	//"strconv"
	//"log"
	"fmt"
	"testing"
	"database/sql"
	//"context"
)

type User struct{ // s of struct should be lower case
	Id int
	Name string
	Password string
}

func Connect() (*sql.DB, error){
	var db *sql.DB // its a variable so please add var, and it is sql not SQL, u urself imported the package

	db, err := sql.Open("mysql", "root:Qwerty@2412@tcp(127.0.0.1:3306)/applicationproject")

	return db, err}


func CreateNewAccount(user User) (User, error){
	name := user.Name
	password := user.Password
	db1, err := Connect()
	if err!= nil{
		panic(err.Error)
	}




	sql := "INSERT INTO vendor(name, password) VALUES(?, ?)"
	res, err := db1.Exec(sql, name, password)
	if err!= nil{
		panic(err.Error)
	}
	Ven_ID, err := res.LastInsertId()
	if err!= nil{
		panic(err.Error)
	}
	

	
	sql1 := "SELECT name,password FROM vendor WHERE id = ?"
	row, err := db1.Query(sql1, Ven_ID)
	if err!= nil{
		panic(err.Error)
	}

	row.Next()

	var i User
	
	err1:= row.Scan(&i.Name,&i.Password)

	fmt.Println(i)

	return i, err1
}

func TestCreateNewAccount(t *testing.T){
	user := User{Name : "Arjun",Password : "12"}

	account, err:= CreateNewAccount(user)

	if err!= nil{
		t.Fatalf(" error: %s", err)
	} else{
		fmt.Println("No error")
	}

	if account.Name!= user.Name || account.Password != user.Password{
		t.Fatalf("Did not create account")
	} else{
		fmt.Println("Account created successfully")
	}
}