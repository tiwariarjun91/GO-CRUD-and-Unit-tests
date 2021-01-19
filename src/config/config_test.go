package config

import(
	"fmt"
	"database/sql"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error){
	var db *sql.DB // its a variable so please add var, and it is sql not SQL, u urself imported the package

	db, err := sql.Open("mysql", "root:Qwerty@2412@tcp(127.0.0.1:3306)/applicationproject")

	return db, err

}

func ConnectingTest(t *testing.T) {
	//var db *sql.DB

	db1, err := Connect()
	
	if err!= nil{
		t.Fatalf("SQL connection error: %s", err)
	} else{
		fmt.Println("Database connected")
	}

	if db1 == nil{
		t.Fatalf("SQL connection error)")
	} else{
		fmt.Println("Database connected")
	}

	db1.Close()
}

func OpenTest(t *testing.T){
	db1, _ := Connect()
	err := db1.Ping()
	if err!= nil{
		t.Fatalf("SQL connection error)")
	} else{
		fmt.Println("Database connected")
	}
	
	db1.Close()
	
}

func CloseTest(t *testing.T){
	db1, _ := Connect()
	db1.Close()
	err := db1.Ping()
	if err == nil{
		t.Fatalf("SQL connection error)")
	} else{
		fmt.Println("Database connection terminted")
	}
}