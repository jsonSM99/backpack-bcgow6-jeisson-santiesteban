package service

import (
	"errors"
)

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
	GetTickets() []Ticket
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func validate(t Ticket, tickets []Ticket) bool {
	if t.Id <= 0 || t.Price <= 0 {
		return false
	}
	if t.Names == "" || t.Email == "" || t.Destination == "" || t.Date == "" {
		return false
	}
	for _, ticket := range tickets {
		if ticket.Id == t.Id {
			return false
		}
	}
	return true
}
func (b *bookings) Create(t Ticket) (ticket Ticket, err error) {

	ok := validate(t, b.Tickets)

	if !ok {
		err = errors.New("Campos invalidos")
		return
	}
	ticket = t
	b.Tickets = append(b.Tickets, ticket)
	return
}

func (b *bookings) Read(id int) (ticket Ticket, err error) {
	for _, t := range b.Tickets {
		if t.Id == id {
			ticket = t
		}
	}

	if ticket.Id == 0 {
		err = errors.New("Ticket not found")
		return
	}
	return
}
func (b *bookings) pos(id int) int {
	for p, v := range b.Tickets {
		if v.Id == id {
			return p
		}
	}
	return -1
}

func (b *bookings) Update(id int, t Ticket) (ticket Ticket, err error) {

	ticketIndex := b.pos(t.Id)

	if ticketIndex == -1 {
		err = errors.New("Ticket not found")
		return
	}
	b.Tickets[ticketIndex] = t

	return b.Tickets[ticketIndex], nil
}

func (b *bookings) Delete(id int) (deleted int, err error) {
	ticketIndex := b.pos(id)

	if ticketIndex == -1 {
		err = errors.New("Ticket not found")
		return
	}
	b.Tickets = append(b.Tickets[:ticketIndex], b.Tickets[ticketIndex+1:]...)
	deleted = id
	return
}

func (b bookings) GetTickets() (tickets []Ticket) {
	return b.Tickets
}
