package main

import(
	"net/http"
	"html/template"
	//"crud"
	"config"
	"strconv"
	"log"
	"fmt"
)

type Product struct{
	ProductId string
	VenderId string
	ProductName string
	ProductQuantity string 
	ProductPrice string
}
type Product1 struct{
	ProductId int
	VenderId int
	ProductName string
	ProductQuantity float64
	ProductPrice float64
}

var tp *template.Template
func Login(w http.ResponseWriter, r *http.Request){
	tp, _ = template.ParseGlob("*.html") // u used : 

	tp.ExecuteTemplate(w, "login.html", nil)
}

func DeleteData(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	name1 := r.FormValue("username")
	password1 := r.FormValue("password")
	vendor_form_id := r.FormValue("vendor_id")
	prod_id := r.FormValue("prod_id")
	prod_name := r.FormValue("name")
	fmt.Println(name1, password1)
	id_vendor, err := strconv.ParseInt(vendor_form_id, 10, 64)
	id_prod, err := strconv.ParseInt(prod_id, 10, 64)
	id_prod_del := int(id_prod)
	ven_ID, err := config.Db.Query("SELECT id FROM vendor WHERE name = ? AND password = ?", name1, password1)
	fmt.Println(ven_ID)
	if err!= nil{
			panic(err.Error)}
	var id int
	ven_ID.Next()
	err = ven_ID.Scan(&id)
	if err!= nil{
		panic(err.Error)}
	ven_ID1 := id
	fmt.Println(ven_ID1, "this is ven id")
	ven_id_comp := int64(ven_ID1)
	if(id_vendor == ven_id_comp){
		res, err := config.Db.Query("DELETE FROM product WHERE product_id = ? AND product_name = ?", id_prod_del, prod_name)
		if err!= nil{
			panic(err.Error())
		}
		res.Close()
		tp.ExecuteTemplate(w, "showdata.html", nil)


	}else{
	
	tp.ExecuteTemplate(w, "login.html", nil)}
}
func Createuser(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	type Data struct{
		Id int64
	}
	name := r.FormValue("username")
	password := r.FormValue("password")
	sql := "INSERT INTO vendor(name, password) VALUES(?, ?)"
	res, err := config.Db.Exec(sql, name, password)
	if err!= nil{
		panic(err.Error)
	}
	Ven_ID, err := res.LastInsertId()

	//sql1 := "SELECT id FROM vendor WHERE name = ? AND password = ?"
	//ven_ID, err := config.Db.Exec(sql1, name, password)
	if err!= nil{
		panic(err.Error)
	}
	fmt.Println(Ven_ID)

	Ven := Data{}
	Ven.Id = Ven_ID
	tp.ExecuteTemplate(w, "add_data1.html", Ven)

}
func ShowData(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}
	name1 := r.FormValue("username")
	password1 := r.FormValue("password")
	vendor_form_id := r.FormValue("vendor_id")
	fmt.Println(name1, password1)
	id_vendor, err := strconv.ParseInt(vendor_form_id, 10, 64)
	/*sql1 := "SELECT id FROM vendor WHERE name = ? AND password = ?"
	ven_ID, err := config.Db.Exec(sql1, name1, password1)*/
	//sql1 := "SELECT id FROM vendor WHERE name = ? AND password = ?"
	//stmt, err:= config.Db.Prepare(sql1)
	if err!= nil{
		panic(err.Error())
	}
	//defer stmt.Close()
	//ven_ID, err := config.Db.Exec(sql1,name1, password1)
	//if err!= nil{
	//	panic(err.Error)
	//}
	ven_ID, err := config.Db.Query("SELECT id FROM vendor WHERE name = ? AND password = ?", name1, password1)
	fmt.Println(ven_ID)
	if err!= nil{
			panic(err.Error)}
	var id int
	ven_ID.Next()
	err = ven_ID.Scan(&id)
	if err!= nil{
		panic(err.Error)}
	ven_ID1 := id
	fmt.Println(ven_ID1, "this is ven id")
	ven_id_comp := int64(ven_ID1)
	if(id_vendor == ven_id_comp){
	//sql := "SELECT * FROM product WHERE vendor_id= ?"
	//result, err := config.Db.Query("SELECT * FROM product WHERE vendor_id = ?", ven_ID)
	sql := "SELECT * FROM product WHERE vendor_id = ?"
	result, err := config.Db.Query(sql, ven_ID1)

	//sql := "SELECT * FROM product WHERE vendor_id = ?"
	//stmt2, err := config.Db.Prepare(sql)
	if err!= nil{
		panic(err.Error())
	}
	//defer stmt2.Close()
	
	//result := stmt2.QueryRow( ven_ID)
	//result,err := config.Db.Query(sql) 
	/*if err!= nil{
		panic(err.Error)
	}*/
	prod := Product1{}
	rows := []Product1{}
	fmt.Fprintln(w,` <table><tr><td>`,"Product Id",`</td><td>`,"VenderId",`</td><td>`,"Product Name",`</td><td>`,"Product Quantity",`</td><td>`,"Product Price",`</td>`)
	for result.Next(){
		var productid, venderid int 
		var name string
		var quantity, price float64
		err := result.Scan(&productid,&venderid,&name, &quantity, &price)
		if err!= nil{
			panic(err.Error)
		}
		fmt.Fprintln(w,` <table><tr><td>`,productid,`</td><td>`,venderid,`</td><td>`,name,`</td><td>`,quantity,`</td><td>`,price,`</td>`)
		prod.ProductId = productid
		prod.VenderId = venderid
		prod.ProductName = name
		prod.ProductQuantity = quantity
		prod.ProductPrice = price
		rows = append(rows, prod)
		

	}

	tp.ExecuteTemplate(w, "showdata.html", nil)}else {
		fmt.Fprintf(w,"passowrd/id/missmatch")
		tp.ExecuteTemplate(w, "login.html", nil)
	}

}

func AddProduct(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	
	var p = Product{}
	//var id1, quantity1, price1 string
	//id := r.FormValue("vender_id") // u got this name wrong
	name := r.FormValue("name")
	quantity := r.FormValue("quantity")
	price := r.FormValue("price")
	p.ProductName = name 
	id := r.FormValue("vend_id")
	
	/*id, _:= strconv.ParseInt(id1, 10, 32)
	quantity, _:= strconv.ParseFloat(quantity1,64)
	price, _:= strconv.ParseFloat(price1, 64)*/


	//p.VenderId = id
	p.ProductQuantity = quantity
	p.ProductPrice = price
	p.VenderId = id

	tp.ExecuteTemplate(w, "add_data.html", p)
	
	id1, _ := strconv.ParseFloat(id, 64)
	//id1 := 1
	quantity1, _:= strconv.ParseFloat(quantity,64)
	price1, _:= strconv.ParseFloat(price, 64)

	sql := "INSERT INTO product(vendor_id,product_name, product_quantity, product_price) VALUES(?,?,?,?)"
	res, err := config.Db.Exec(sql, id1, name, quantity1, price1)

	if err!= nil{
		panic(err.Error())
	}
	

    lastId, err := res.LastInsertId()

    if err != nil {
        log.Fatal(err)
    }

	fmt.Printf("The last inserted row id: %d\n", lastId)
	fmt.Println(id1, name, quantity1, price1)

	
}

func UpdateProduct(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	name1 := r.FormValue("username")
	password1 := r.FormValue("password")
	vendor_form_id := r.FormValue("vendor_id")
	prod_id := r.FormValue("prod_id")

	prod_q := r.FormValue("quantity")
	
	prod_p := r.FormValue("price")
	product_quant, err := strconv.ParseInt(prod_q, 10, 64)
	if err!= nil{
		panic(err.Error())
	}
	product_quantity := int(product_quant)

	product_price, err := strconv.ParseFloat(prod_p, 64)
	if err!= nil{
		panic(err.Error())
	}
	fmt.Println(name1, password1)
	id_vendor, err := strconv.ParseInt(vendor_form_id, 10, 64)
	id_prod, err := strconv.ParseInt(prod_id, 10, 64)
	id_prod_del := int(id_prod)
	ven_ID, err := config.Db.Query("SELECT id FROM vendor WHERE name = ? AND password = ?", name1, password1)
	fmt.Println(ven_ID)
	if err!= nil{
			panic(err.Error)}
	var id int
	ven_ID.Next()
	err = ven_ID.Scan(&id)
	if err!= nil{
		panic(err.Error)}
	ven_ID1 := id
	ven_id_comp := int64(ven_ID1)

	fmt.Println(ven_id_comp, "this is ven id from db")
	fmt.Println(id_vendor, "this is ven id from form")
	if(id_vendor == ven_id_comp){
		res, err := config.Db.Query("UPDATE product SET product_quantity = ?, product_price = ? WHERE product_id = ? ",product_quantity, product_price, id_prod_del )
		if err!= nil{
			panic(err.Error())
		}
		res.Close()
		tp.ExecuteTemplate(w, "showdata.html", nil)
	}else{
		tp.ExecuteTemplate(w, "login.html", nil)
	}
}

func main(){
	
	config.Connect_db()
	
	http.HandleFunc("/", Login)
	http.HandleFunc("/ShowData", ShowData)
	http.HandleFunc("/Createuser", Createuser)
	http.HandleFunc("/AddProduct", AddProduct)//you did not write /show in action 
	http.HandleFunc("/DeleteData", DeleteData)
	http.HandleFunc("/UpdateProduct", UpdateProduct)

	http.ListenAndServe(":8000", nil)
	defer config.Db.Close()
	
}