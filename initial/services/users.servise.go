package services

import (
	"crud-example/initial/dto"
	"crud-example/initial/entities"
	"crud-example/initial/repositories"
	"encoding/json"
	"net/http"
	"strconv"
)

type UserService interface {
	CreateUsers(req *http.Request)
	UpdateUsers(req *http.Request)
	DeleteUsersById(req *http.Request)
	GetAllUsers(req *http.Request) (entities.Users, error)
	FirstUser(req *http.Request) (entities.User, error)
}

type usersService struct {
	repository repositories.UsersRepository
}

func NewUsersService(
	respository repositories.UsersRepository,
) UserService {
	return &usersService{
		repository: respository,
	}
}

func (s *usersService) CreateUsers(req *http.Request) {

	var userDTO = dto.CreateUsersDTO{}

	err := json.NewDecoder(req.Body).Decode(&userDTO)

	if err != nil {

		return
	}

	name := req.FormValue("name")
	username := req.FormValue("username")
	email := req.FormValue("email")

	var p entities.User = entities.User{
		Name:     name,
		Username: username,
		Email:    email,
	}

	s.repository.Create(&p)
}

func (s *usersService) UpdateUsers(req *http.Request) {

	var userDTO dto.CreateUsersDTO

	err := json.NewDecoder(req.Body).Decode(&userDTO)

	if err != nil {
		return
	}

	userID, err := strconv.Atoi(req.URL.Query().Get("id"))

	if err != nil {
		return
	}

	_, err = s.repository.First(userID)

	if err != nil {
		return
	}

	name := req.FormValue("name")
	username := req.FormValue("username")
	email := req.FormValue("email")

	data := &entities.User{
		ID:       userID,
		Name:     name,
		Username: username,
		Email:    email,
	}

	s.repository.Update(data)
}

func (s *usersService) DeleteUsersById(req *http.Request) {

	userID, err := strconv.Atoi(req.URL.Query().Get("id"))

	if err != nil {
		return
	}

	s.repository.DeleteById(userID)
}

func (s *usersService) GetAllUsers(req *http.Request) (entities.Users, error) {

	users, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return users, err
}

func (s *usersService) FirstUser(req *http.Request) (entities.User, error) {

	id, err := strconv.Atoi(req.URL.Query().Get("id"))

	if err != nil {
		return entities.User{}, err
	}

	user, err := s.repository.First(id)

	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}
