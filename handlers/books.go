package handlers
import (
	"github.com/bookrepositorygo/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Books struct {
	l *log.Logger
}

func NewBooks(l *log.Logger) *Books {
	return &Books{l}
}

func (p *Books) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getBooks(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.addBook(rw, r)
		return
	}
	if r.Method == http.MethodPut {
		// expect the id in the URI
		regex := regexp.MustCompile(`/([0-9]+)`)
		group := regex.FindAllStringSubmatch(r.URL.Path, -1)

		if len(group) != 1 || len(group[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := group[0][1]
		// Ignore the error for now
		id, _ := strconv.Atoi(idString)

		p.updateBook(id, rw, r)
	}
	// catch all other http verb with 405
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Books) getBooks(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET books")

	listOfBooks := data.GetBooks()
	err := listOfBooks.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Books) addBook(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST book")

	prod := &data.Book{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	data.AddBook(prod)
}

func (p *Books) updateBook(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Put book")

	prod := &data.Book{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateBook(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Book not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Book not found", http.StatusInternalServerError)
		return
	}

}
