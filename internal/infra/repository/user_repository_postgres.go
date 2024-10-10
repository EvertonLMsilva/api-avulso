package repository

import (
	"database/sql"
	"fmt"

	"github.com/EvertonLMsilva/api-avulso/internal/entity"
)

type UserRepositoryPG struct {
	DB *sql.DB
}

func NewUserRepositoryPg(db *sql.DB) *UserRepositoryPG {
	return &UserRepositoryPG{DB: db}
}

func (r *UserRepositoryPG) Create(user *entity.User) error {
	var active int8 = 0

	if user.Active {
		active = 1
	}

	if user.Name == "" {
		return fmt.Errorf("Name empty!")
	}

	if user.Birthday == "" {
		return fmt.Errorf("Birthday empty!")
	}

	_, err := r.DB.Exec(
		"INSERT INTO api_avulso.profile.users (id, name, birthday, active) VALUES ($1, $2, $3, $4)",
		user.ID, user.Name, user.Birthday, active)

	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryPG) FindAll() ([]*entity.User, error) {
	rows, err := r.DB.Query("SELECT * FROM api_avulso.profile.users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.ID, &user.Name, &user.Birthday, &user.Active)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r *UserRepositoryPG) Disable(id string) error {
	_, err := r.DB.Exec("UPDATE api_avulso.profile.users SET active='0' WHERE id=$1", id)
	if err != nil {
		return err
	}

	return nil
}
