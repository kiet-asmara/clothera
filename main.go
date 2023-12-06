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
			fmt.Print("\nPilih Tipe Customer (Admin / Customer): ")
			fmt.Scan(&CustomerType)

			exit2 := false
			var choiceCustomer int
			switch CustomerType {
			case "Customer":
				for !exit2 {
					cli.ShowCustomerMenu()
					fmt.Print("Choice: ")
					fmt.Scan(&choiceCustomer)

					switch choiceCustomer {
					case 1:
						handler.ListCategory(db)
						exit3 := false
						var choiceCategory int
						for !exit3 {
							fmt.Print("Silahkan pilih kategori (0 untuk kembali): ")
							fmt.Scan(&choiceCategory)

							if choiceCategory < 0 || choiceCategory > len(handler.Categories) {
								fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
								continue
							} else if choiceCategory == 0 {
								exit3 = true
							} else {
								exit4 := false
								var choiceProdukID int

								for !exit4 {
									handler.DisplayClothesByCategory(db, handler.Categories[choiceCategory-1])
									fmt.Print("Silahkan pilih barang yang ingin dibeli (0 untuk kembali): ")
									fmt.Scan(&choiceProdukID)

									if choiceProdukID < 0 || choiceProdukID > len(handler.Categories) {
										fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
										continue
									} else if choiceProdukID == 0 {
										exit4 = true
										handler.ListCategory(db)
									} else {
										// function getClothes
									}
								}
							}

						}

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
