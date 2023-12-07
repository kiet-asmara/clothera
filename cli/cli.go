package cli

import (
	"errors"
	"pair-project/pkg/table"

	"fmt"

	"database/sql"
	"pair-project/entity"
	"pair-project/handler"

	"pair-project/pkg/validator"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func ShowMainMenu() {
	fmt.Println("")
	fmt.Println("---------")
	fmt.Println("Main Menu")
	fmt.Println("---------")

	fmt.Println("1 -> Login")
	fmt.Println("2 -> Register")
	fmt.Println("3 -> Exit\n")
}

func ShowCustomerMenu() {
	fmt.Println("-------------")
	fmt.Println("Customer Menu")
	fmt.Println("-------------")

	fmt.Println("1 -> Beli")
	fmt.Println("2 -> Rental Pakaian")
	fmt.Println("3 -> Pesanan")
	fmt.Println("4 -> Profile")
	fmt.Println("5 -> Back to Main Menu\n")
}

func ShowAdminMenu() {
	fmt.Println("----------")
	fmt.Println("Admin Menu")
	fmt.Println("----------")

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
	fmt.Println("------------")
	fmt.Println("Reports Menu")
	fmt.Println("------------")

	fmt.Println("1 -> User Report")
	fmt.Println("2 -> Order Report")
	fmt.Println("3 -> Back\n")
}

func ShowProfileMenu() {
	fmt.Println("1 -> Show Profile")
	fmt.Println("2 -> Edit Profile")
	fmt.Println("3 -> Back\n")
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

	return existingCustomer, nil
}

func ShowProfile(db *sql.DB, customer *entity.Customer) error {
	var err error

	customer, err = handler.GetCustomerByEmail(db, customer.CustomerEmail)
	if err != nil {
		switch {
		case errors.Is(err, handler.ErrorRecordNotFound):
			panic("Error: invalid system error") // at this point user should exists
		default:
			return err
		}
	}

	address, err := handler.GetAddressByID(db, customer.Address.AddressID)
	if err != nil {
		switch {
		case errors.Is(err, handler.ErrorRecordNotFound):
			panic("Error: invalid system error") // at this point address should exists
		default:
			return err
		}
	}

	customer.Address.AddressCountry = address.AddressCountry
	customer.Address.AddressCity = address.AddressCity
	customer.Address.AddressStreet = address.AddressStreet

	table.Render(table.RenderParam{
		Title:      "User Profile",
		TitleAlign: table.AlignCenter,
		Header:     []string{"ID", "Name", "Email", "Country", "City", "Street"},
		DataSiggle: table.Row{
			customer.CustomerID,
			customer.CustomerName,
			customer.CustomerEmail,
			customer.Address.AddressCountry,
			customer.Address.AddressCity,
			customer.Address.AddressStreet,
		},
		DataALign: table.AlignCenter,
	})
	return nil
}

func UpdateProfile(db *sql.DB, customer *entity.Customer) (*entity.Customer, error) {
	var newcustomer = &entity.Customer{
		CustomerID:       customer.CustomerID,
		CustomerEmail:    customer.CustomerEmail,
		CustomerName:     customer.CustomerName,
		CustomerPassword: customer.CustomerPassword,
		Address:          customer.Address,
	}
	v := validator.New()

	var count int
	for {
		if count > 4 {
			fmt.Println("You exceed the limited chance. Please try again later!")
			return customer, nil
		}

		oldpassword, _ := promptword("Please input your old password before updating")
		err := bcrypt.CompareHashAndPassword([]byte(customer.CustomerPassword), []byte(oldpassword))
		if err != nil {
			count++
			fmt.Println("err: incorrect password. try again!")
			continue
		}
		break
	}

	fmt.Println("\ntype '-' for skip!")
	newname := inputUpdateUsername(v, "New name")
	newemail := inputUpdateEmail(v, "New email")
	newpassword := inputUpdatePassword(v, "New password")
	newcountry := inputUpdateCountry(v, "New country")
	newcity := inputUpdateCity(v, "New city")
	newstreet := inputUpdateStreet(v, "New street")

	if len(newname) > 0 && newname != "-" {
		newcustomer.CustomerName = newname
	}

	if len(newemail) > 0 && newemail != "-" {
		newcustomer.CustomerEmail = newemail
	}

	if len(newpassword) > 0 && newpassword != "-" {
		newcustomer.CustomerPassword = newpassword
	}

	if len(newcountry) > 0 && newcountry != "-" {
		newcustomer.Address.AddressCountry = newcountry
	}

	if len(newcity) > 0 && newcity != "-" {
		newcustomer.Address.AddressCity = newcity
	}

	if len(newstreet) > 0 && newstreet != "-" {
		newcustomer.Address.AddressStreet = newstreet
	}

	err := handler.UpdateAdressByID(db, &newcustomer.Address)
	if err != nil {
		switch {
		case errors.Is(err, handler.ErrorRecordNotFound):
			panic(err)
		default:
			return customer, err
		}
	}
	customer.Address = newcustomer.Address

	err = handler.UpdateCustomerByID(db, newcustomer, newcustomer.CustomerPassword != customer.CustomerPassword)
	if err != nil {
		switch {
		case errors.Is(err, handler.ErrorRecordNotFound):
			panic(err)
		default:
			return customer, err
		}
	}
	customer.CustomerName = newcustomer.CustomerName
	customer.CustomerEmail = newcustomer.CustomerEmail
	customer.CustomerPassword = newcustomer.CustomerPassword

	return customer, nil
}
