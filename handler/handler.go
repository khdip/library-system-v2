package handler

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/jmoiron/sqlx"
)

type Books struct {
	ID          int    `db:"id" json:"id"`
	BookName    string `db:"book_name" json:"book_name"`
	Author      string `db:"author" json:"author"`
	Category    string `db:"category" json:"category"`
	BookDesc    string `db:"book_description" json:"book_description"`
	BookCover   string `db:"book_cover" json:"book_cover"`
	IsAvailable bool   `db:"is_available" json:"is_available"`
}

type Handler struct {
	templates *template.Template
	db        *sqlx.DB
	decoder   *schema.Decoder
}

func GetHandler(db *sqlx.DB, decoder *schema.Decoder) *mux.Router {
	hand := &Handler{
		db:      db,
		decoder: decoder,
	}
	hand.GetTemplate()

	r := mux.NewRouter()
	r.HandleFunc("/", hand.GetBooks)
	r.HandleFunc("/create", hand.createBook)
	r.HandleFunc("/store", hand.storeBook)
	r.HandleFunc("/q", hand.searchBook)
	r.HandleFunc("/{id:[0-9]+}/edit", hand.editBook)
	r.HandleFunc("/{id:[0-9]+}/Update", hand.updateBook)
	r.HandleFunc("/{id:[0-9]+}/delete", hand.deleteBook)
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := hand.templates.ExecuteTemplate(w, "404.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	return r
}

func (h *Handler) GetTemplate() {
	h.templates = template.Must(template.ParseFiles(
		"templates/create-book.html",
		"templates/list-book.html",
		"templates/edit-book.html",
		"templates/search-result.html",
		"templates/no-search-result.html",
		"templates/404.html",
	))
}
