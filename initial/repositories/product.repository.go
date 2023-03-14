package repositories

import (
	"crud-example/initial/entities"
	"database/sql"
)

type IProductRepository interface {
	Create(p *entities.Product) bool
	Update(p *entities.Product) bool
	DeleteById(id int) bool
	GetAll() (entities.Products, error)
	First(id int) (entities.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(database *sql.DB) IProductRepository {
	return &productRepository{
		db: database,
	}
}

func (product *productRepository) GetAll() (entities.Products, error) {
	rows, err := product.db.Query("Select * from products")

	if err != nil {
		return nil, err
	}

	var p entities.Products

	for rows.Next() {
		var product entities.Product
		rows.Scan(&product.ID, &product.Name, &product.Text, &product.Price, &product.Image)
		p = append(p, product)
	}

	return p, nil
}

func (product *productRepository) Create(p *entities.Product) bool {

	result, err := product.db.Exec("insert into products(name, text, price, image) values(?,?,?,?)", p.Name, p.Text, p.Price, p.Image)

	if err != nil {
		return false
	}

	rowAffected, _ := result.RowsAffected()

	return rowAffected > 0
}

func (product *productRepository) Update(p *entities.Product) bool {

	result, err := product.db.Exec("update products set name = ?, text = ?, price = ?, image = ? where id = ?", p.Name, p.Text, p.Price, p.Image, p.ID)

	if err != nil {
		return false
	}

	rowAffected, _ := result.RowsAffected()

	return rowAffected > 0
}

func (product *productRepository) DeleteById(id int) bool {
	_, err := product.db.Exec("delete from products where id = ?", id)

	if err != nil {
		return false
	}

	return false
}

func (product *productRepository) First(id int) (entities.Product, error) {

	result := product.db.QueryRow("select * from products where id = ?", id)

	var p entities.Product

	err := result.Scan(&p.ID, &p.Name, &p.Text, &p.Price, &p.Image)

	if err != nil {
		return p, err
	}

	return p, nil
}
