package controller

import (
	"crud-example/initial/pkg/types"
	"crud-example/initial/services"
	"html/template"
	"net/http"
)

type UserController interface {
	Index(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Store(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	Edit(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	service services.UserService
}

func NewUserController(
	service services.UserService,
) UserController {
	return &userController{
		service: service,
	}
}

func (c *userController) Index(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles(
		"public/product/index.html", "public/includes/header.html", "public/includes/footer.html",
	)

	users, err := c.service.GetAllUsers(r)

	if err != nil {
		http.Redirect(w, r, "/users", 302)
	}

	view.ExecuteTemplate(w, "PRODUCT_INDEX", types.DataUsersIndex{Users: users, Count: len(users)})
}

func (c *userController) Create(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles(
		"public/users/create.html", "public/includes/header.html", "public/includes/footer.html",
	)

	view.ExecuteTemplate(w, "PRODUCT_CREAT", nil)
}

func (c *userController) Show(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles(
		"public/users/show.html", "public/includes/header.html", "public/includes/footer.html",
	)

	user, err := c.service.FirstUser(r)

	if err != nil {
		http.Redirect(w, r, "/users", 302)
	}

	view.ExecuteTemplate(w, "PRODUCT_SHOW", user)
}

func (users *userController) Edit(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles(
		"public/users/edit.html", "public/includes/header.html", "public/includes/footer.html",
	)

	user, err := users.service.FirstUser(r)

	if err != nil {
		http.Redirect(w, r, "/users", 302)
	}

	view.ExecuteTemplate(w, "PRODUCT_EDIT", user)
}

func (users *userController) Store(w http.ResponseWriter, r *http.Request) {
	users.service.CreateUsers(r)
	http.Redirect(w, r, "/users/create", 302)
}

func (users *userController) Update(w http.ResponseWriter, r *http.Request) {

	users.service.UpdateUsers(r)

	http.Redirect(w, r, "/users", 302)
}

func (users *userController) Delete(w http.ResponseWriter, r *http.Request) {

	users.service.DeleteUsersById(r)

	http.Redirect(w, r, "/users", 302)
}
