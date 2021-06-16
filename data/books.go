package data

import (
	"encoding/json"
	"fmt"
	"io"
)

type Book struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Writer      string 	`json:"writer"`
	Type      	string 	`json:"type"`
	Description string  `json:"description"`
}

type Books []*Book

func (p *Books) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func (p *Book) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

func GetBooks() Books {
	return bookList
}

func AddBook(p *Book) {
	p.ID = getNextID()
	bookList = append(bookList, p)
}

func UpdateBook(id int, p *Book) error {
	_, pos, err := findBook(id)
	if err != nil {
		return err
	}

	p.ID = id
	bookList[pos] = p

	return nil
}

func findBook(id int) (*Book, int, error) {
	for i, p := range bookList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

var ErrProductNotFound = fmt.Errorf("Book is not found")

func getNextID() int {
	lastProduct := bookList[len(bookList)-1]
	return lastProduct.ID + 1
}

var bookList = []*Book{
	&Book{
		ID:           1,
		Name:         "The Great Gatsby",
		Writer:       "F. Scott Fitzgerald",
		Type: 		  "Classics",
		Description:  "The story is of the fabulously wealthy Jay Gatsby and his new love for the beautiful Daisy Buchanan. It is an exquisitely crafted tale of America in the 1920s.",
	},
	&Book{
		ID:           2,
		Name:         "The Diary of a Young Girl",
		Writer:       "Anne Frank",
		Type: 		  "Autobiography",
		Description:  "Discovered in the attic in which she spent the last years of her life, Anne Frank’s remarkable diary has become a world classic.",
	},
}
