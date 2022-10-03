package file

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type File struct {
	Path string
}

func openFile(path string, flag int, perm fs.FileMode) (document *os.File, err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			err = recovered.(error)
		}
	}()

	document, err = os.OpenFile(path, flag, perm)

	if err != nil {
		panic(err)
	}
	return
}

func (f *File) Read() (tickets []service.Ticket, err error) {

	var document *os.File
	document, err = openFile(f.Path, os.O_RDONLY, 0755)

	if err != nil {
		return
	}

	scanner := bufio.NewScanner(document)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	document.Close()

	for _, field := range text {
		fields := strings.Split(field, ",")
		// fields = fields[1:]
		for i := 0; i <= len(fields)-5; i += 5 {
			Id, Price := 0, 0
			_, err = fmt.Sscan(fields[i], &Id)
			_, err = fmt.Sscan(fields[i+5], &Price)
			var ticket = service.Ticket{
				Id:          Id,
				Names:       fields[i+1],
				Email:       fields[i+2],
				Destination: fields[i+3],
				Date:        fields[i+4],
				Price:       Price,
			}
			tickets = append(tickets, ticket)
		}
	}

	return
}

func appendToFile(file *os.File, ticket service.Ticket) (err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			err = recovered.(error)
		}
	}()
	_, err = file.WriteString(fmt.Sprintf("%d,%s,%s,%s,%s,%d\n", ticket.Id, ticket.Names, ticket.Email, ticket.Destination, ticket.Date, ticket.Price))
	if err != nil {
		panic(err)
	}
	return
}

func writeLines(file string, tickets []service.Ticket) (err error) {

	return
}

func ticketsToLines(tickets []service.Ticket) (lines [][]string) {

	return
}
func (f *File) Write(tickets []service.Ticket) error {

	file, err := os.Create(f.Path)
	defer file.Close()

	if err != nil {
		return err
	}

	csvwriter := csv.NewWriter(file)
	defer csvwriter.Flush()

	for _, ticket := range tickets {
		record := []string{
			strconv.Itoa(ticket.Id),
			ticket.Names,
			ticket.Email,
			ticket.Destination,
			ticket.Date,
			strconv.Itoa(ticket.Price)}

		if err = csvwriter.Write(record); err != nil {
			fmt.Println(err)
		}
	}
	return err
}
