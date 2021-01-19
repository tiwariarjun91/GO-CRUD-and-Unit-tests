package crud

import(
	//"fmt"
	//"database/sql"
	//"config"
	"net/http"
	"html/template"
	//"strconv"

	//_ "github.com/go-sql-driver/mysql"
)


type Product struct{
	//VenderId string
	ProductName string
	ProductQuantity string 
	ProductPrice string
}

var tp *template.Template
func AddData(w http.ResponseWriter, r *http.Request){
	tp, _ = template.ParseGlob("*.html") // u used : 
	tp.ExecuteTemplate(w, "enterdata.html", nil)
}

func Show(w http.ResponseWriter, r *http.Request){
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
	
	/*id, _:= strconv.ParseInt(id1, 10, 32)
	quantity, _:= strconv.ParseFloat(quantity1,64)
	price, _:= strconv.ParseFloat(price1, 64)*/


	//p.VenderId = id
	p.ProductQuantity = quantity
	p.ProductPrice = price

	tp.ExecuteTemplate(w, "showdata.html", p)
}
