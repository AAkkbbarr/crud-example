package repositories

import (
	"crud-example/initial/entities"
	"database/sql"
)

type UsersRepository interface {
	Create(r *entities.User) bool
	Update(r *entities.User) bool
	DeleteById(id int) bool
	GetAll() (entities.Users, error)
	First(id int) (entities.User, error)
}

type usersRepository struct {
	db *sql.DB
}

func NewUsersRepository(database *sql.DB) UsersRepository {
	return &usersRepository{
		db: database,
	}
}

func (r *usersRepository) GetAll() (entities.Users, error) {

	rows, err := r.db.Query("Select * from users")

	if err != nil {
		return nil, err
	}

	var users entities.Users

	for rows.Next() {
		var user entities.User
		rows.Scan(&user.ID, &user.Name, &user.Username, &user.Email)
		users = append(users, user)
	}

	return users, nil
}

func (r *usersRepository) Create(user *entities.User) bool {

	result, err := r.db.Exec("insert into users(name, text, price, image) values(?,?,?,?)", user.Name, user.Username, user.Email)

	if err != nil {
		return false
	}

	rowAffected, _ := result.RowsAffected()

	return rowAffected > 0
}

func (r *usersRepository) Update(user *entities.User) bool {

	result, err := r.db.Exec("update users set name = ?, username = ?, email = ? where id = ?", user.Name, user.Username, user.ID)

	if err != nil {
		return false
	}

	rowAffected, _ := result.RowsAffected()

	return rowAffected > 0
}

func (r *usersRepository) DeleteById(id int) bool {
	_, err := r.db.Exec("delete from products where id = ?", id)

	if err != nil {
		return false
	}

	return false
}

func (r *usersRepository) First(id int) (entities.User, error) {

	result := r.db.QueryRow("select * from users where id = ?", id)

	var user entities.User

	err := result.Scan(&user.ID, &user.Name, &user.Username, &user.Email)

	if err != nil {
		return user, err
	}

	return user, nil
}
