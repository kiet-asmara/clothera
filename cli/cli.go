package cli

import (
	"fmt"
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
