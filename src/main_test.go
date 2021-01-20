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

type Product2 struct{
	ProductId float64
	VendorId float64
	ProductName string
	ProductQuantity float64 
	ProductPrice float64

}

func Connect() (*sql.DB, error){
	var db *sql.DB // its a variable so please add var, and it is sql not SQL, u urself imported the package

	db, err := sql.Open("mysql", "root:Qwerty@2412@tcp(127.0.0.1:3306)/applicationproject")

	return db, err
}

func CreateNewProduct(product Product2) (Product2, error){
	vid := product.VendorId
	name := product.ProductName
	quantity := product.ProductQuantity
	price := product.ProductPrice

	/*id1, _ := strconv.ParseFloat(vid, 64)
	quantity1, _:= strconv.ParseFloat(quantity,64)
	price1, _:= strconv.ParseFloat(price, 64)*/

	db1, err := Connect()
	if err!= nil{
		panic(err.Error)
	}

	sql := "INSERT INTO product(vendor_id,product_name, product_quantity, product_price) VALUES(?,?,?,?)"
	res, err := db1.Exec(sql, vid, name, quantity, price)

	if err!= nil{
		panic(err.Error())
	}

	Prod_ID, err := res.LastInsertId()
	if err!= nil{
		panic(err.Error)
	}

	sql1 := "SELECT vendor_id, product_name,product_quantity, product_price FROM product WHERE product_id = ?"
	row, err := db1.Query(sql1, Prod_ID)
	if err!= nil{
		panic(err.Error)
	}

	row.Next()

	var p Product2
	
	err1:= row.Scan(&p.VendorId,&p.ProductName, &p.ProductQuantity, &p.ProductPrice)

	fmt.Println(p)

	return p, err1

}
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
func UpdateProductt(product Product2)(Product2, error){
	productId := product.ProductId
	vendorId := product.VendorId
	quantity := product.ProductQuantity
	price := product.ProductPrice

	var p Product2

	db1,err:= Connect()
	if err!= nil{
		panic(err.Error)
	}

	res, err := db1.Query("UPDATE product SET product_quantity = ?, product_price = ? WHERE product_id = ? AND vendor_id = ?",quantity,price, productId,vendorId )
		if err!= nil{
			panic(err.Error())
		}

	res.Close()

	row, err := db1.Query("SELECT product_quantity, product_price FROM product WHERE product_id = ? ", productId)
	if err!= nil{
		panic(err.Error())
	}
	row.Next()

	err1:= row.Scan(&p.ProductQuantity, &p.ProductPrice)

	return p, err1

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

func TestCreateNewProduct(t *testing.T){
	product := Product2{VendorId : 1, ProductName : "Test Product", ProductQuantity : 10, ProductPrice: 100}
	productInserted, err := CreateNewProduct(product) 

	if err!= nil{
		t.Fatalf(" error: %s", err)
	} else{
		fmt.Println("No error")
	}

	if product.VendorId!= productInserted.VendorId || product.ProductName != productInserted.ProductName{
		t.Fatalf("Did not create new product")
	} else{
		fmt.Println("Product created successfully")
	}

}

func TestUpdateProduct(t *testing.T){
	product := Product2{VendorId : 1, ProductId : 28, ProductName : "Test Product", ProductQuantity : 1, ProductPrice: 10}
	updatedProduct, err:= UpdateProductt(product)
	if err!= nil{
		t.Fatalf(" error: %s", err)
	} else{
		fmt.Println("No error")
	}

	if product.ProductQuantity!= updatedProduct.ProductQuantity || product.ProductPrice != updatedProduct.ProductPrice{
		t.Fatalf("Did not update th  product")
	} else{
		fmt.Println("Product updated successfully")
	}
}

func TestDeleteProduct(t *testing.T){
	
	name := "arjun"
	password := "1234"
	Id := 1
	pid := 2
	var productName string


	db1,err:= Connect()

	if err!= nil{
		panic(err.Error)
	}

	ven_ID, err := db1.Query("SELECT id FROM vendor WHERE name = ? AND password = ?", name, password)
	if err != nil{
		panic(err.Error)
	}
	var id int
	ven_ID.Next()
	err = ven_ID.Scan(&id)
	if err!= nil{
		panic(err.Error)
	}
	if id == Id{
		res, err := db1.Query("DELETE FROM product WHERE product_id = ?", pid)
		if err!= nil{
			panic(err.Error())
		}
		res.Close()
	}

	row, _:= db1.Query("Select product_name FROM product WHERE product_id = ?",id)
	row.Next()
	err1:= row.Scan(productName)

	if err1 !=nil{
		fmt.Println("Row deleted succesfully")
	}else{
		t.Fatalf("Did not delete the  product")
	}

	

	
}