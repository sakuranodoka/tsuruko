package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ordergo.com/database"
)

type API struct {
	Data *database.Database
}

func (api *API) InsertCustomer(context *gin.Context) {
	var request struct {
		// Customer_name    string `json:"name"`
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

	var response struct {
		Status        string
		StatusMessage string
		Customer      database.Customer
	}
	err := context.BindJSON(&request)
	if err != nil {
		response.Status = "error"
		response.StatusMessage = err.Error()
		context.JSON(http.StatusInternalServerError, response)
		return
	}

	//! map request something substring other
	customer := database.Customer{
		Customer_name:    request.Customer_name,
		Customer_surname: request.Customer_surname,
		Address_name:     request.Address_name,
		Birthday:         request.Birthday,
		Address:          request.Address,
		Village:          request.Village,
		Soi:              request.Soi,
		Road:             request.Road,
		Subdistrict:      request.Subdistrict,
		District:         request.District,
		Province:         request.Province,
		Post:             request.Post,
		Tel:              request.Tel,
	}
	//! map close
	err = api.Data.InsertCustomerDB(&customer)
	if err != nil {
		response.Status = "error"
		response.StatusMessage = err.Error()
		context.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Customer = customer
	context.JSON(http.StatusOK, response)
	return
}

func (api *API) SearchCustomer(context *gin.Context) {
	var request struct {
		// Customer_name    string `json:"name"`
		Customer_name    string
		Customer_surname string
		// Address_name     string
		// Birthday         string
		// Address          string
		// Village          string
		// Soi              string
		// Road             string
		// Subdistrict      string
		// District         string
		// Province         string
		// Post             string
		// Tel              string
	}

	var response struct {
		Status        string
		StatusMessage string
		Customer      []database.Customer
	}
	err := context.BindJSON(&request)
	if err != nil {
		response.Status = "error"
		response.StatusMessage = err.Error()
		context.JSON(http.StatusInternalServerError, response)
		return
	}

	//! map request something substring other
	// customer := database.Customer{
	// 	Customer_name:    request.Customer_name,
	// 	Customer_surname: request.Customer_surname,
	// }
	//! map close
	customer, err := api.Data.SearchCustomerDB(request.Customer_name)
	if err != nil {
		response.Status = "error"
		response.StatusMessage = err.Error()
		context.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Customer = customer
	context.JSON(http.StatusOK, response)
	return
}
