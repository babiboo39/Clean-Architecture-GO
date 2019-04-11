package database

import (
	"golang.org/x/crypto/bcrypt"
	"modul1/domain"
)

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO users (first_name, last_name, username, email, password) VALUES (?,?,?,?,?)", u.FirstName, u.LastName, u.Username, u.Email,
		hashPassword(u.Password),
	)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

func (repo *UserRepository) FindById(identifier int) (user domain.User, err error) {
	row, err := repo.Query("SELECT id, first_name, last_name, username, email, password FROM users WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int
	var firstName string
	var lastName string
	var username string
	var email string
	var password string
	row.Next()
	if err = row.Scan(&id, &firstName, &lastName, &username, &email, &password); err != nil {
		return
	}
	user.ID = id
	user.FirstName = firstName
	user.LastName = lastName
	user.Username = username
	user.Email = email
	user.Password = password
	return
}

func (repo *UserRepository) FindAll() (users domain.Users, err error) {
	rows, err := repo.Query("SELECT id, first_name, last_name FROM users")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var firstName string
		var lastName string
		if err := rows.Scan(&id, &firstName, &lastName); err != nil {
			continue
		}
		user := domain.User{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
		}
		users = append(users, user)
	}
	return
}


func (repo *UserRepository) FindByEmail(email string, password string) (user domain.User, err error) {
	//var req *http.Request

	var databaseEmail string
	var databasePassword string

	rows, err := repo.Query("Select email, password from users where email=?", email)
	if err != nil {
	}
	for rows.Next() {
		if err := rows.Scan(&databaseEmail, &databasePassword); err != nil {
			continue
		}
		err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	}
	defer rows.Close()

	return
}


func hashPassword(input string) string {
	password := []byte(input)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}
