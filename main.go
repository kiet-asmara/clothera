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

										quantity, err := handler.AddClothes(db, *selectedClothes, *customer, orderID)
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
						// list barang pesanan

						// hitung diskon & pajak
						totalPrice = handler.CalcDiscount(totalPrice)
						fmt.Printf("Your total with tax (11%%) is: $%.2f.\n", totalPrice)

						// insert total price ke tabel orders
						err := handler.InsertTotal(db, totalPrice, orderID)
						if err != nil {
							log.Fatal(err)
						}
						os.Exit(1)

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
								var addProdukAdmin int
								exit3 := false

								for !exit3 {
									cli.ShowAdminAddProductMenu()
									fmt.Print("Choice: ")
									fmt.Scan(&addProdukAdmin)

									switch addProdukAdmin {
									case 1:
										categories := handler.FetchAllCategoriesFromDatabase(db)
										handler.PrintCategoriesClothes(categories)

										selectedCategory := handler.GetSelectedCategoryFromUser(categories)

										newProductClothes := handler.GetProductDetailsFromAdmin(selectedCategory)

										err := handler.InsertProductIntoDatabase(db, newProductClothes)
										if err != nil {
											fmt.Println("Error adding product:", err)
										} else {
											fmt.Println("Product added successfully!")
										}
									case 2:
										categories := handler.FetchAllCategoriesFromDatabaseCostumes(db)
										handler.PrintCategoriesCostumes(categories)

										selectedCategory := handler.GetSelectedCategoryFromUserCostumes(categories)

										newProductCostumes := handler.GetProductDetailsFromAdminCostumes(selectedCategory)

										err := handler.InsertProductIntoDatabaseCostumes(db, newProductCostumes)
										if err != nil {
											fmt.Println("Error adding product:", err)
										} else {
											fmt.Println("Product added successfully!")
										}
									case 3:
										exit3 = true
									default:
										fmt.Println("Invalid choice")
									}
								}

							case 2:
								var addProdukAdmin int
								exit3 := false

								for !exit3 {
									cli.ShowAdminAddProductMenu()
									fmt.Print("Choice: ")
									fmt.Scan(&addProdukAdmin)

									switch addProdukAdmin {
									case 1:
										categories := handler.FetchAllCategoriesFromDatabase(db)
										fmt.Println("Available Categories:", categories)
										handler.ShowProductsByCategory(db)
										handler.DeleteProduct(db)
									case 2:
										categories := handler.FetchAllCategoriesFromDatabaseCostumes(db)
										fmt.Println("Available Categories:", categories)
										handler.ShowProductsByCategoryCostumes(db)
										handler.DeleteProductCostumes(db)
									case 3:
										exit3 = true
									default:
										fmt.Println("Invalid choice")
									}
								}
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
