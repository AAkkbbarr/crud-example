package controller

import (
	"html/template"
	"net/http"
)

type MainController interface {
	Index(w http.ResponseWriter, r *http.Request)
}

type mainController struct {
}

func NewMainController() MainController {
	return &mainController{}
}

func (main *mainController) Index(w http.ResponseWriter, r *http.Request) {

	view, _ := template.ParseFiles(
		"public/index.html",
		"public/includes/header.html",
		"public/includes/footer.html",
	)

	view.ExecuteTemplate(w, "INDEX", nil)
}
