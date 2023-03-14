package main

import (
	controllers "crud-example/initial/controllers"
	repositories "crud-example/initial/repositories"
	servises "crud-example/initial/services"
	conn "crud-example/pkg/db"
	"database/sql"
	"net/http"
)

var (
	db                *sql.DB                         = conn.ConnectDB()
	productRepository repositories.IProductRepository = repositories.NewProductRepository(db)
	productService    servises.IProductService        = servises.NewProductService(productRepository)
	mainController    controllers.MainController      = controllers.NewMainController()
	productController controllers.IProductController  = controllers.NewProductController(productService)

	usersRepository repositories.UsersRepository = repositories.NewUsersRepository(db)
	usersService    servises.UserService         = servises.NewUsersService(usersRepository)
	usersController controllers.UserController   = controllers.NewUserController(usersService)
)

func main() {

	defer conn.CloseDB(db)

	http.HandleFunc("/", mainController.Index)

	http.HandleFunc("/products", productController.Index)
	http.HandleFunc("/products/create", productController.Create)
	http.HandleFunc("/products/edit", productController.Edit)

	http.HandleFunc("/products/show", productController.Show)
	http.HandleFunc("/products/store", productController.Store)
	http.HandleFunc("/products/update", productController.Update)
	http.HandleFunc("/products/delete", productController.Delete)

	// http.HandleFunc("/api/products", productController.Index)
	// http.HandleFunc("/api/products/show", productController.Show)
	// http.HandleFunc("/api/products/create", productController.Store)
	// http.HandleFunc("/api/products/edit", productController.Update)
	// http.HandleFunc("/api/products/delete", productController.Delete)

	http.HandleFunc("/users", usersController.Index)
	http.HandleFunc("/users/create", usersController.Create)
	http.HandleFunc("/users/edit", usersController.Edit)

	http.HandleFunc("/users/show", usersController.Show)
	http.HandleFunc("/users/store", usersController.Store)
	http.HandleFunc("/users/update", usersController.Update)
	http.HandleFunc("/users/delete", usersController.Delete)

	http.ListenAndServe(":8080", nil)
}
