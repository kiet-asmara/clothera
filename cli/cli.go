package cli

import (
	"database/sql"
	"fmt"
	"pair-project/entity"
	"pair-project/handler"
	"pair-project/pkg/validator"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func ShowMainMenu() {
	fmt.Println("1 -> Login")
	fmt.Println("2 -> Register")
	fmt.Println("3 -> Exit\n")
}

func ShowCustomerMenu() {
	fmt.Println("\n1 -> Beli")
	fmt.Println("2 -> Rental Pakaian")
	fmt.Println("3 -> Pesanan")
	fmt.Println("4 -> Edit Profil")
	fmt.Println("5 -> Back to Main Menu\n")
}

func ShowAdminMenu() {
	fmt.Println("1 -> Produk")
	fmt.Println("2 -> Report")
	fmt.Println("3 -> Back to Main Menu\n")
}

func ShowAdminProdukMenu() {
	fmt.Println("1 -> Add Produk")
	fmt.Println("2 -> Delete Produk")
	fmt.Println("3 -> Update Produk")
	fmt.Println("4 -> Back\n")
}

func ShowAdminReportMenu() {
	fmt.Println("1 -> User Report")
	fmt.Println("2 -> Order Report")
	fmt.Println("3 -> Stock Report")
	fmt.Println("4 -> Back\n")
}

func PromptChoice(prompt string) int {
	input, err := promptline(prompt)
	if err != nil {
		return -1
	}

	input = strings.TrimSpace(input)

	num, err := strconv.Atoi(input)
	if err != nil {
		return -1
	}
	return num
}

func Register(db *sql.DB) (*entity.Customer, error) {
	var customer = &entity.Customer{CustomerType: entity.User}
	var address = &entity.Address{}
	v := validator.New()

	customer.CustomerName = inputUsername(v, "name")
	customer.CustomerEmail = inputEmail(v, "email")
	customer.CustomerPassword = inputPassword(v, "password")
	address.AddressCountry = inputCountry(v, "country")
	address.AddressCity = inputCity(v, "city")
	address.AddressStreet = inputStreet(v, "street")

	err := handler.InsertAddress(db, address)
	if err != nil {
		return nil, err
	}

	customer.Address = *address
	err = handler.InsertCustomer(db, customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func Login(db *sql.DB) (*entity.Customer, error) {
	var customer = &entity.Customer{CustomerType: entity.User}
	v := validator.New()

	customer.CustomerEmail = inputEmail(v, "email")
	customer.CustomerPassword = inputPassword(v, "password")

	existingCustomer, err := handler.GetCustomerByEmail(db, customer.CustomerEmail)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingCustomer.CustomerPassword), []byte(customer.CustomerPassword))
	if err != nil {
		return nil, err
	}

	return customer, nil
}
