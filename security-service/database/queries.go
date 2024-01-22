package database

import (
	"database/sql"
	"fmt"
	"subproblem/security-service/models"

	_ "github.com/lib/pq"
)


func (pg *PostgresDb) CreateUserTable() error {
	
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			first_name VARCHAR(50) NOT NULL,
			last_name VARCHAR(50) NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL,
			password VARCHAR(100) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`

	_, err := pg.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (pg *PostgresDb) AddUser(u *models.User) error {

	query := `
		INSERT INTO users
		(first_name, last_name, email, password)
		VALUES
		($1, $2, $3, $4)
	`
	resp, err := pg.db.Query(
		query, 
		u.Firstname, u.Lastname, u.Email, u.Password)
	
	if err != nil {
		return err
	}
	
	fmt.Printf("%v\n", resp)

	return nil
}

func (pg *PostgresDb) ReadAllUsers() ([]*models.User, error) {

	query := `
		SELECT id, first_name, last_name, email FROM users
	`

	res, err := pg.db.Query(query)

	if err != nil {
		return nil, err
	}

	var users []*models.User

	for res.Next() {
		user := &models.User{}
		
		err := res.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	
	if err := res.Err(); err != nil {
		return nil, err
	}

	return users, nil;

}



func (pg *PostgresDb) GetUserById(id int) (*models.User, error) {

	query := `
		SELECT id, first_name, last_name, email FROM users
		WHERE id = $1
	`
	user := &models.User{}
	err := pg.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Firstname,
		&user.Lastname,
		&user.Email,
	)

	if err != nil {
		return nil, err
	}
	
	return user, nil
}

func (pg *PostgresDb) DeleteUserById(id int) (*models.User, error) {

	user, err := pg.GetUserById(id)

	if err != nil {
		return nil, err
	}

	
	query := `
		DELETE FROM users 
		WHERE id = $1
	`
	_, err2 := pg.db.Exec(query, id)

	if err2 != nil {
		return nil, err2
	}

	return user, nil
}

func (pg *PostgresDb) FindUserByEmail(email string) (*models.User, error) {

	user := &models.User{}

	query := `
		SELECT id, email, password FROM users 
		WHERE email = $1
	`
	
	err := pg.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email, 
		&user.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}
