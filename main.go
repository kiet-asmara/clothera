package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"pair-project/cli"
	"pair-project/config"
	"pair-project/entity"
	"pair-project/handler"
)

func main() {
	db, err := config.GetDB("root:@tcp(127.0.0.1:3306)/clothera")
	if err != nil {
		log.Fatal("Failed to connect")
	}
	defer db.Close()

	exitMainMenu := false
	var choiceMainMenu int

	for !exitMainMenu {
		cli.ShowMainMenu()
		choiceMainMenu = cli.PromptChoice("Choice")

		var customer *entity.Customer
	RG_OK:

		switch choiceMainMenu {
		case 1:

			if nil == customer {
				fmt.Println("customer is nil")
				os.Exit(1)
			}

			exit2 := false
			switch customer.CustomerType {
			case entity.User:
				var choiceCustomer int
				for !exit2 {
					cli.ShowCustomerMenu()
					fmt.Print("Choice: ")
					fmt.Scan(&choiceCustomer)

					switch choiceCustomer {
					case 1:
						fmt.Println("Beli")
					case 2:
						fmt.Println("Rental Pakaian")
					case 3:
						fmt.Println("Pesanan")
					case 4:
						fmt.Println("Edit Profil")
					case 5:
						fmt.Println("Back to Main Menu")
						exit2 = true
					default:
						fmt.Println("Invalid choice")
					}
				}
			case entity.Admin:
				for !exit2 {
					var choiceCustomer int
					cli.ShowAdminMenu()
					fmt.Print("Choice: ")
					fmt.Scan(&choiceCustomer)

					switch choiceCustomer {
					case 1:
						var productChoice int
						productExit := false

						for !productExit {
							cli.ShowAdminProdukMenu()
							fmt.Print("Choice: ")
							fmt.Scan(&productChoice)

							switch productChoice {
							case 1:
								fmt.Println("Add Produk")
							case 2:
								fmt.Println("Delete Produk")
							case 3:
								fmt.Println("Update Produk")
							case 4:
								productExit = true
							default:
								fmt.Println("Invalid choice")
							}
						}
					case 2:
						var productChoice int
						productExit := false

						for !productExit {
							cli.ShowAdminReportMenu()
							fmt.Print("Choice: ")
							fmt.Scan(&productChoice)

							switch productChoice {
							case 1:
								fmt.Println("User Report")
							case 2:
								fmt.Println("Order Report")
							case 3:
								fmt.Println("Stock Report")
							case 4:
								productExit = true
							default:
								fmt.Println("Invalid choice")
							}
						}
					case 3:
						fmt.Println("Back to Main Menu")
						exit2 = true
					default:
						fmt.Println("Invalid choice")
					}
				}
			}

		case 2:
			var err error
			customer, err = cli.Register(db)
			if err != nil {
				switch {
				case errors.Is(err, handler.ErrorDuplicateEntry):
					fmt.Println("User with this email already exists. Try login instead!")
				default:
					fmt.Println("Sorry We Have Problem in our server. Please Try Again!")
				}
				continue
			}

			fmt.Println("Register Success!")
			choiceMainMenu = 1
			goto RG_OK

		case 3:
			fmt.Println("Thank you for ordering")
			exitMainMenu = true
		default:
			fmt.Println("Invalid choice")
		}
	}
}
