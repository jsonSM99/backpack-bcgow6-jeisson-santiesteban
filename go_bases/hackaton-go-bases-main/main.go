package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	var tickets []*service.Ticket
	var err error

	fileService := file.File{
		Path: "./tickets.csv",
	}

	tickets, err = fileService.Read()

	if err != nil {
		fmt.Println(err)
	}
	// Funcion para obtener tickets del archivo csv
	fmt.Print(tickets)
	bookings := service.NewBookings(tickets)

	bookings.Create(&service.Ticket{
		Id:          1001,
		Names:       "jeisson",
		Email:       "jeisson@gmail.com",
		Destination: "yopal",
		Date:        "01/01/1999",
		Price:       1,
	})
	fileService.Write(tickets)

	bookings.Update(1001, &service.Ticket{
		Id:          1001,
		Names:       "jeisson",
		Email:       "jeisson2@gmail.com",
		Destination: "yopal",
		Date:        "01/01/1999",
		Price:       1,
	})
	fileService.Write(tickets)

	if err != nil {
		fmt.Println(err)
	}
	for _, x := range tickets {
		fmt.Println(*x)
	}

	bookings.Delete(1001)
	fileService.Write(tickets)

}
