package main

import (
	"demo/go-password/account"
	"fmt"
)

func main() {
	fmt.Println("--- Password Manager ---")
Menu:
	for {
		varinat := getMenu()
		switch varinat {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}
}

func getMenu() int {
	fmt.Println("1. создать аккаунт")
	fmt.Println("2. найти аккаунт")
	fmt.Println("3. удалить аккаунт")
	fmt.Println("4. выход")
	fmt.Print("Выберите действие: ")
	var variant int
	fmt.Scan(&variant)
	return variant
}

func createAccount() {

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите url")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный логин или URL")
		return
	}
	vault := account.NewVault()
	vault.AddAccount(*myAccount)

}

func findAccount() {

}

func deleteAccount() {

}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
