package main

import (
	"demo/go-password/account"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("--- Password Manager ---")
	vault := account.NewVault()
Menu:
	for {
		varinat := getMenu()
		switch varinat {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
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

func createAccount(vault *account.Vault) {

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите url")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный логин или URL")
		return
	}
	vault.AddAccount(*myAccount)
}

func findAccount(vault *account.Vault) {
	url := promptData("Введите url для поиска")
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}

}

func deleteAccount(vault *account.Vault) {
	url := promptData("Введите url для удаления")
	isDel := vault.DeleteAccountsByUrl(url)
	if isDel {
		color.Green("Запись удалена")
	} else {
		color.Red("Не удалось удалить")
	}
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
