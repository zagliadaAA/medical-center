package main

import (
	"fmt"
	"project2/adapter/repository/clientRepository"
	"project2/internal/usecase/client_usecase"
)

func main() {

	fmt.Println("-----\n")

	repo := clientRepository.NewRepo() //создаем новое хранилище которое содержит мапу типа клиент + ID
	uc := client_usecase.NewUseCase(repo)

	err := uc.Create(client_usecase.CreateClientReq{
		Name:        "Artem",
		BDate:       "30.12.1999",
		PhoneNumber: "89085538251",
	})
	if err != nil {
		panic(err)
	}

	err = uc.Create(client_usecase.CreateClientReq{
		Name:        "Boba",
		BDate:       "01.01.1999",
		PhoneNumber: "89087732819",
	})
	if err != nil {
		panic(err)
	}

	err = uc.Create(client_usecase.CreateClientReq{
		Name:        "Liza",
		BDate:       "02.02.1998",
		PhoneNumber: "8908846195",
	})
	if err != nil {
		panic(err)
	}

	// Печатаем всех клиентов
	clients := repo.GetAll()
	fmt.Println(clients)

	err = uc.Delete(2)
	if err != nil {
		panic(err)
	}

	// Печатаем всех клиентов
	clients = repo.GetAll()
	fmt.Println(clients)

}

/*
ws := &CarWashStation
honda := &Honda{}
ws.Wash(honda)

type CarWashStation struct {}

func(s *CarWashStation) Wash(car Car) {
	fmt.Println("Моем машину:")
}

type Car interface {
	BeepBeep() string
}

type Honda struct {}
type Chevrolet struct {}

func(c *Honda) BeepBeep() string {
	return fmt.Sprint("honda")
}

func(c *Chevrolet) BeepBeep() string {
	return fmt.Sprint("chevrolet")
}
*/
