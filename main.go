package main

import (
	"fmt"
	"log"
	"pair-project/cli"
	"pair-project/config"
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
		fmt.Print("Choice: ")
		fmt.Scan(&choiceMainMenu)

		switch choiceMainMenu {
		case 1:
			CustomerType := ""
			fmt.Print("Pilih Tipe Customer (Admin / Customer): ")
			fmt.Scan(&CustomerType)

			exit2 := false

			var choiceCustomer int

			switch CustomerType {
			case "Customer":
				// create order
				orderID, err := handler.CreateOrder(db, 2)
				if err != nil {
					log.Fatal(err)
				}

				// create total price
				var totalPrice float64

				for !exit2 {
					cli.ShowCustomerMenu()
					fmt.Print("Choice: ")
					fmt.Scan(&choiceCustomer)

					switch choiceCustomer {
					case 1:
						fmt.Println("Beli")
					case 2:
						fmt.Println("Rental Pakaian")
						// tampilkan kategori

						// tampilkan list produk

						// input: costumeid, quantity, enddate
						price, err := handler.Rent(db, orderID)
						if err != nil {
							log.Fatal(err)
						}
						fmt.Println("price:", price)
						totalPrice += price
						fmt.Printf("\nYour total is now: $%.2f\n\n", totalPrice)
					case 3:
						fmt.Println("Pesanan")
						// insert total price ke tabel orders

						// list beli & rental

					case 4:
						fmt.Println("Edit Profil")
					case 5:
						fmt.Println("Back to Main Menu")
						exit2 = true
					default:
						fmt.Println("Invalid choice")
					}
				}
			case "Admin":
				for !exit2 {
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
			fmt.Println("Register")
		case 3:
			fmt.Println("Thank you for ordering")
			exitMainMenu = true
		default:
			fmt.Println("Invalid choice")
		}
	}
}
