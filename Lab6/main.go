package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Client struct {
	name          string
	surname       string
	accountNumber int
	cDeposit      int
	cCredit       int
	clientType    bool
}

func (p *Client) SetName(name string) {
	p.name = name
}
func (p Client) Name() string {
	return p.name
}
func (p *Client) SetSurname(surname string) {
	p.surname = surname
}
func (p Client) Surname() string {
	return p.surname
}
func (p *Client) SetAccountNumber(accountNumber int) bool {
	if accountNumber > 0 {
		p.accountNumber = accountNumber
		return true
	}
	return false
}
func (p Client) AccountNumber() int {
	return p.accountNumber
}

func (p *Client) SetCCredit(cCredit int) bool {
	if cCredit >= 0 {
		p.cCredit = cCredit
		return true
	}
	return false
}
func (p Client) CCredit() int {
	return p.cCredit
}
func (p *Client) SetCDeposit(cDeposit int) bool {
	if cDeposit >= 0 {
		p.cDeposit = cDeposit
		return true
	}
	return false
}
func (p Client) CDeposit() int {
	return p.cDeposit
}

func (p Client) SetDepositClient(clientType bool) {
	p.clientType = clientType
}

func (p Client) DepositClient() bool {
	return p.clientType
}

type Bank struct {
	name         string
	bankMoney    int
	deposit      int
	credit       int
	clients      []Client
	countClients int
}

func (p *Bank) SetName(name string) {
	p.name = name
}
func (p Bank) Name() string {
	return p.name
}
func (p *Bank) CountClients() int {
	p.countClients++
	return p.countClients
}
func (p *Bank) SetBankMoney(bankMoney int) bool {
	if bankMoney >= 0 {
		p.bankMoney = bankMoney
		return true
	}
	return false
}
func (p Bank) BankMoney() int {
	return p.bankMoney
}

func (p *Bank) SetDeposit(deposit int) bool {
	if deposit >= 0 {
		p.deposit = deposit
		return true
	}
	return false
}
func (p Bank) Deposit() int {
	return p.deposit
}

func (p *Bank) SetCredit(credit int) bool {
	if credit >= 0 {
		p.credit = credit
		return true
	}
	return false
}
func (p Bank) Credit() int {
	return p.credit
}
func (p Bank) GiveDeposit(money int) {
	p.SetBankMoney(p.BankMoney() + money)
	p.SetDeposit(p.Deposit() + money)
}

func (p *Bank) GetCredit(money int) bool {
	if p.BankMoney() > money {
		p.SetBankMoney(p.BankMoney() - money)
		p.SetCredit(p.Credit() + money)
		return true
	}
	return false
}
func (p *Bank) Clients() []Client {
	return p.clients
}

func createBank() Bank {
	var bank Bank

	fmt.Print("Введіть назву: ")
	fmt.Scan(&bank.name)
	var temp string
	for true {
		fmt.Print("Введіть к-сть власних коштів банку: ")
		fmt.Scan(&temp)
		if s, err := strconv.ParseInt(temp, 10, 32); err == nil {
			bank.SetBankMoney(int(s))
			break
		} else {
			fmt.Print("Помилка\n")
		}
	}
	bank.SetDeposit(0)
	bank.SetCredit(0)

	return bank
}
func createClient(bank *Bank, creditClient bool) Client {
	var client Client

	fmt.Print("Введіть ім'я: ")
	fmt.Scan(&client.name)
	fmt.Print("Введіть прізвище: ")
	fmt.Scan(&client.surname)

	var temp string
	if creditClient {
		for true {
			fmt.Print("Введіть суму взятого кредиту: ")
			fmt.Scan(&temp)
			if s, err := strconv.ParseInt(temp, 10, 32); err == nil {
				if bank.GetCredit(int(s)) {
					client.SetCCredit(int(s))
					break
				} else {
					fmt.Println("Сума кредиту перевищує ліміт")
				}
			} else {
				fmt.Print("Помилка\n")
			}
		}
	} else {
		for true {
			fmt.Print("Введіть суму депозиту: ")
			fmt.Scan(&temp)
			if s, err := strconv.ParseInt(temp, 10, 32); err == nil {
				bank.GiveDeposit(int(s))
				client.SetCDeposit(int(s))
				break
			} else {
				fmt.Print("Помилка\n")
			}
		}
	}
	client.SetAccountNumber(bank.CountClients())
	client.clientType = false

	return client
}

func startDepositClient(client *Client, bank *Bank, mutex *sync.Mutex) {

	deposit := client.CDeposit()
	fmt.Println("\n", "Рахунок №", client.AccountNumber(), " :", client.Surname(), client.Name(), ", Депозит на суму: ", deposit, ", Кошти банку: ", bank.BankMoney())
	for client.CDeposit() > 0 {
		var sum int
		flag := true
		for flag {
			min := -client.CDeposit()
			max := client.CDeposit()
			if min != 0 && max != 0 {
				sum = rand.Intn(max-min) + min
			}
			mutex.Lock()
			bank.GiveDeposit(sum)
			client.SetCDeposit(client.CDeposit() + sum)
			mutex.Unlock()
		}

		time.Sleep(1 * time.Second)
	}
	fmt.Println(client.Surname(), client.Name(), ", Забрано депозит: ", deposit)
	fmt.Println(bank.Name(), ":", " Вкладених депозитів:", bank.Deposit(), "( Всього коштів:", bank.bankMoney, ")")
}

func startCreditClient(client *Client, bank *Bank, mutex *sync.Mutex) {

	credit := client.CCredit()
	fmt.Println("\n", "Рахунок №", client.AccountNumber(), client.Surname(), client.Name(), ", Видано кредит: ", client.CCredit())
	for client.CCredit() > 0 {
		var sum int
		flag := true
		for flag {
			min := -client.CCredit()
			max := client.CCredit()
			if min != 0 && max != 0 {
				sum = rand.Intn(max-min) + min
			}
			mutex.Lock()
			if bank.GetCredit(sum) {
				client.SetCCredit(client.CCredit() + sum)
				flag = false
			} else {
				flag = true
			}
			mutex.Unlock()
		}

		time.Sleep(1 * time.Second)
	}
	fmt.Println("\n", "Рахунок №", client.AccountNumber(), client.Surname(), client.Name(), ", Погашено кредит: ", credit)
	fmt.Println(bank.Name(), ":", " Видано кредитів:", bank.Credit(), "( Всього коштів:", bank.bankMoney, ")")
}

func getClientInfoSurname(name string, bank *Bank) {
	clients := bank.Clients()
	fl := true
	for i := 0; i < len(clients); i++ {
		if clients[i].Surname() == name {

			printClient(&clients[i])
			fl = false
		}
	}
	if fl {
		fmt.Println("Нічого не знайдено")
	}
}

func getClientInfoAccountNumber(accountNumber int, bank *Bank) {
	clients := bank.Clients()
	fl := true
	for i := 0; i < len(clients); i++ {
		if clients[i].AccountNumber() == accountNumber {
			printClient(&clients[i])
			fl = false
		}
	}
	if fl {
		fmt.Println("Нічого не знайдено")
	}
}

func printClient(client *Client) {
	fmt.Println("№ рахунку: \t", client.AccountNumber())
	fmt.Println("Прізвище: \t", client.Surname())
	fmt.Println("Iм'я: \t", client.Name())
	fmt.Println("Сума депозитів: \t", client.CDeposit())
	fmt.Println("Сума кредитів: \t", client.CCredit())
}

func printBank(bank *Bank)  {
	fmt.Println("Назва: \t", bank.Name())
	fmt.Println("Кошти банку: \t", bank.BankMoney())
	fmt.Println("З них:")
	fmt.Println("Сума депозитів: \t", bank.Deposit())
	fmt.Println("Сума кредитів: \t", bank.Credit())
}
var main_bank Bank


func main() {

	menu := 0

	var mutex sync.Mutex
	for menu != -1 {

		fmt.Print("Меню:\n1.Створення банку\n2.Створення кліента для роботи з кредитами\n3.Створення кліента для роботи з депозитами\n4.Виведення інформації про клієнта за прізвищем (поточний стан клієнта)\n5.Виведення інформації про клієнта за номером рахунку (поточний стан клієнта)\n6. Виведення інформації про банк\n7.Вийти")
		menu = 0
		fmt.Print("\nВиберіть:")
		fmt.Scan(&menu)
		if menu != 1 && menu != 7 {
			mutex.Lock()
			if main_bank.Name() == "" {
				fmt.Println("Створіть банк!")
				menu = 0
			}
			mutex.Unlock()
		}
		switch menu {
		case 1:
			main_bank = createBank()
			break
		case 2:
			main_bank.clients = append(main_bank.clients, createClient(&main_bank, true))
			go startCreditClient(&main_bank.clients[len(main_bank.clients)-1], &main_bank, &mutex)
			break
		case 3:
			main_bank.clients = append(main_bank.clients, createClient(&main_bank, false))
			go startDepositClient(&main_bank.clients[len(main_bank.clients)-1], &main_bank, &mutex)
			break
		case 4:
			var name string
			fmt.Println("Пошук по прізвищу")

			fmt.Print("Введіть прізвище: ")
			fmt.Scan(&name)
			getClientInfoSurname(name, &main_bank)
		case 5:
			var accNumber int
			fmt.Println("Пошук по номеру рахунку")

			fmt.Print("Введіть номер рахунку: ")
			fmt.Scan(&accNumber)
			getClientInfoAccountNumber(accNumber, &main_bank)
		case 6:
			printBank(&main_bank);
			break;
			case 7:
			menu = -1
			break
		default:
			{
				fmt.Println("Даного пункту немає в меню, повторіть спробу")
			}
		}

	}
}
