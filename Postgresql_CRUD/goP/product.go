package goP

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/sphinxql"
	_ "github.com/lib/pq"
)

func Hey() {
	fmt.Println("say hey")
}

func Hata(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Product struct {
	Id                 int
	Title, Description string
	Price              float32
}

// const (
// 	host     = "localhost"
// 	port     = "5432"
// 	user     = "postgres"
// 	password = "321654"
// 	dbname   = "Postgresql"
// )

var db *sql.DB

func init() {
	var err error

	// connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s ",
	//  host, port, user, password, dbname)

	db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=321654 dbname=Title01 sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(1 * time.Second)
	db.SetConnMaxLifetime(30 * time.Second)
	Hata(err)
}

func InsertProduct(data Product) {
	result, err := db.Exec("INSERT INTO test(id,title,description,price)VALUES($1,$2,$3,$4)", data.Id, data.Title, data.Description, data.Price)
	Hata(err)
	rowsAffected, err := result.RowsAffected()
	Hata(err)
	fmt.Printf("Etkilenen Kayıt Sayısı:(%d)", rowsAffected)
}

func UpdateProduct(data Product) {
	result, err := db.Exec("UPDATE test SET title=$2 WHERE id=$1", data.Title, data.Id)
	Hata(err)
	rowsAffected, err := result.RowsAffected()
	Hata(err)
	fmt.Printf("Etkilenen Kayıt Sayısı:(%d)", rowsAffected)
}

// func DeleteProduct(data Product) {
// 	result, err := db.Exec("DELETE FROM products WHERE id=2")
// 	Hata(err)
// 	rowsAffected, err := result.RowsAffected()
// 	Hata(err)
// 	fmt.Printf("Etkilenen Kayıt Sayısı:(%d)", rowsAffected)
// }

func GetProduct(data Product) {
	rows, err := db.Query("SELECT * FROM test")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Records Found!")
			return
		}
		log.Fatal(err)
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {
		prd := &Product{}
		err := rows.Scan(&prd.Id, &prd.Title, &prd.Description, &prd.Price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, prd)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	for _, value := range products {
		fmt.Printf("%d - %s, %s, $%.2f\n", value.Id, value.Title, value.Description, value.Price)
	}
}

func GetProductById(id int) {
	var product string
	err := db.QueryRow("SELECT title FROM test WHERE id=1", id).Scan(&product)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No product with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Product is %s\n", product)
	}
}
