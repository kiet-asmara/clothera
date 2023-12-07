package main

import (
	"errors"
	"fmt"
	"log"
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

		// user yang udah authenticated disimpan disini
		var customer *entity.Customer
	RG_OK:

		switch choiceMainMenu {
		case 1:

			if nil == customer {
				customer, err = cli.Login(db)
				if err != nil {
					fmt.Printf("Sorry your crendential is not valid. Please try again!\n\n")
					continue
				}
				fmt.Printf("Login Success\n\n")
			}

			exit2 := false

			var choiceCustomer int
			switch customer.CustomerType {
			case entity.User:
				// create order
				orderID, err := handler.CreateOrder(db, customer.CustomerID)
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
									temp := handler.DisplayClothesByCategory(db, handler.Categories[choiceCategory-1])
									fmt.Print("Silahkan pilih barang yang ingin dibeli (0 untuk kembali): ")
									fmt.Scan(&choiceProdukID)

									var temp1 int
									if choiceProdukID != 0 {
										temp1, err = handler.ByName(db, temp[choiceProdukID-1])
										if err != nil {
											panic("Error getting clothes by name!")
										}
									}

									if choiceProdukID < 0 || choiceProdukID > len(handler.Categories) {
										fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
										continue
									} else if choiceProdukID == 0 {
										exit4 = true
										handler.ListCategory(db)
									} else {
										selectedClothes, err := handler.GetClothesByID(db, temp1)
										if err != nil {
											log.Fatal(err)
										}
										price, err := handler.GetPriceClothes(db, selectedClothes.ClothesID)
										if err != nil {
											log.Fatal(err)
										}

										fmt.Print("Enter the quantity: ")
										var quantity int
										fmt.Scan(&quantity)

										err = handler.AddClothes(db, *selectedClothes, *customer, orderID, quantity)
										if err != nil {
											log.Fatal(err)
										}

										totalPrice += price * float64(quantity)
										fmt.Printf("Added %d %s to your order.\n", quantity, selectedClothes.ClothesName)
										fmt.Printf("Total price: %.2f.\n", totalPrice)
									}
								}
							}

						}

					case 2:
						fmt.Println("--------------")
						fmt.Println("Rental Pakaian")
						fmt.Println("--------------")

						// tampilkan kategori
						kategori := handler.CategoryCostume()

						// tampilkan list produk
						err := handler.ListCostumes(db, kategori)
						if err != nil {
							log.Fatalln(err)
						}

						// rent function
						price, err := handler.Rent(db, orderID)
						if price == 0 {
							continue // if out of stock
						}
						if err != nil {
							log.Fatalln(err)
						}

						fmt.Printf("\nRental price: %.2f\n", price)

						// add to total
						totalPrice += price
						fmt.Printf("\nYour total is now: $%.2f\n", totalPrice)
					case 3:
						fmt.Println("")
						fmt.Println("--------")
						fmt.Println("Pesanan")
						fmt.Println("--------")
						// check if totalprice = 0
						if totalPrice == 0 {
							fmt.Println("Order is empty. Total price is $0.00")
							continue
						}

						// list barang pesanan
						err := handler.ListPesanan(db, orderID)
						if err != nil {
							log.Fatalln(err)
						}

						// hitung diskon & pajak
						fmt.Println("------------------------------------")
						totalPrice = handler.CalcDiscount(totalPrice)
						fmt.Printf("Your total with tax (11%%) is: $%.2f.\n", totalPrice)

						// insert total price ke tabel orders
						err = handler.InsertTotal(db, totalPrice, orderID)
						if err != nil {
							log.Fatalln(err)
						}
						return

					// Update Profile
					case 4:
						var exit bool
						for !exit {

							cli.ShowProfileMenu()
							choice := cli.PromptChoice("Choice")

							switch choice {
							case 1:
								err := cli.ShowProfile(db, customer)
								if err != nil {
									fmt.Printf("Sorry We Have Problem in our server. Please Try Again!\n\n")
								}

							case 2:
								updatedCustomer, err := cli.UpdateProfile(db, customer)
								if err != nil {
									fmt.Printf("Sorry We Have Problem in our server. Please Try Again!\n\n")
									fmt.Println(err)
									continue
								}
								customer = updatedCustomer
								fmt.Printf("Profile updated sucessfully!\n\n")

							case 3:
								exit = true
							default:
								fmt.Println("Invalid choice")
							}
						}
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
								fmt.Println("-----------")
								fmt.Println("User Report")
								fmt.Println("-----------")

								err := handler.UserReportMenu(db)
								if err != nil {
									log.Fatalln(err)
								}

							case 2:
								fmt.Println("------------")
								fmt.Println("Order Report")
								fmt.Println("------------")

								err := handler.OrderReportMenu(db)
								if err != nil {
									log.Fatalln(err)
								}

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
				fmt.Println(err)
				switch {
				case errors.Is(err, handler.ErrorDuplicateEntry):
					fmt.Printf("User with this email already exists. Try login instead!\n\n")
				default:
					fmt.Printf("Sorry We Have Problem in our server. Please Try Again!\n\n")
				}
				continue
			}

			fmt.Printf("Register Success!\n\n")
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
