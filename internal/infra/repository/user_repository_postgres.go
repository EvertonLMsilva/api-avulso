package repository

import (
	"database/sql"

	"github.com/EvertonLMsilva/api-avulso/internal/entity"
)

type UserRepositoryPG struct {
	DB *sql.DB
}

func NewUserRepositoryPg(db *sql.DB) *UserRepositoryPG {
	return &UserRepositoryPG{DB: db}
}

func (r *UserRepositoryPG) Create(user *entity.User) error {
	_, err := r.DB.Exec("insert into api-avulso.users (id, name, birthday) values(?,?,?)",
		user.ID, user.Name, user.Birthday)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryPG) FindAll() ([]*entity.User, error) {
	rows, err := r.DB.Query("SELECT id, name, birthday FROM api-avulso.users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var user entity.User

		err = rows.Scan(&user.ID, &user.Name, &user.Birthday)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}
