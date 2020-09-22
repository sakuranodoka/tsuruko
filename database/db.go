package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	client *sql.DB
	url    string
}

func NewDB(connectionURL string) *Database {
	client, err := sql.Open("mysql", connectionURL)
	if err != nil {
		panic(err.Error())
	}
	name := "Chaiyut"
	log.Println("connect")
	fmt.Println("console", name)

	clientDB := new(Database)
	clientDB.client = client
	clientDB.url = connectionURL
	return clientDB
}

func (db *Database) Close() {
	db.client.Close()
}

type Customer struct {
	Customer_name    string
	Customer_surname string
	Address_name     string
	Birthday         string
	Address          string
	Village          string
	Soi              string
	Road             string
	Subdistrict      string
	District         string
	Province         string
	Post             string
	Tel              string
}

func (db *Database) InsertCustomerDB(customer *Customer) error {

	tx, err := db.client.Begin()
	if err != nil {
		return err
	}

	statement, err := tx.Prepare(`INSERT INTO pc_customer (customer_name, customer_surname, address_name, birthday)VALUES (?,?,?,?)`)
	if err != nil {
		tx.Rollback()
		return err
	}

	res, err := statement.Exec(
		customer.Customer_name,
		customer.Customer_surname,
		customer.Address_name,
		customer.Birthday)
	if err != nil {
		tx.Rollback()
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	statement2, err := tx.Prepare(`INSERT INTO pc_customer_address (customer_id, address_name, address, village, soi, road, subdistrict, district, province, post, tel) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = statement2.Exec(
		id,
		customer.Address_name,
		customer.Address,
		customer.Village,
		customer.Soi,
		customer.Road,
		customer.Subdistrict,
		customer.District,
		customer.Province,
		customer.Post,
		customer.Tel)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (db *Database) SearchCustomerDB(name string) ([]Customer, error) {
	namecat := name + "%"
	// log.Println("connect", namecat)
	results, err := db.client.Query("SELECT Customer_name,Customer_surname,Address_name,Birthday,Address,Village,Soi,Road,Subdistrict,District,Province,Post,Tel FROM pc_customer WHERE customer_name like ?", namecat)
	if err != nil {
		return nil, err
	}
	var customerDataList []Customer
	for results.Next() {
		var customer Customer
		err = results.Scan(&customer.Customer_name, &customer.Customer_surname, &customer.Address_name, &customer.Birthday, &customer.Address, &customer.Village, &customer.Soi, &customer.Road, &customer.Subdistrict, &customer.District, &customer.Province, &customer.Post, &customer.Tel)
		if err != nil {
			// continue
			// return nil, err
		}
		customerDataList = append(customerDataList, customer)
	}
	return customerDataList, nil

}
