package persistance

import (
	"database/sql"

	"github.com/soerjadi/GoBlog/domain"
	"github.com/soerjadi/GoBlog/domain/repository"
)

// UserRepositoryImpl implements repository UserRepository
type UserRepositoryImpl struct {
	Conn *sql.DB
}

// UserRepositoryWithRDB return initialized UserRepositoryImpl
func UserRepositoryWithRDB(conn *sql.DB) repository.UserRepository {
	return &UserRepositoryImpl{Conn: conn}
}

// GetByID get user by id
func (repo *UserRepositoryImpl) GetByID(id int64) (*domain.User, error) {
	row, err := repo.Conn.Query("select * from users where id=?", id)

	if err != nil {
		return nil, err
	}

	user := &domain.User{}
	err = row.Scan(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetList return list of user
func (repo *UserRepositoryImpl) GetList(offset int, limit int) ([]*domain.User, error) {
	rows, err := repo.Conn.Query("select * from users")

	if err != nil {
		return nil, err
	}

	users := make([]*domain.User, 0)
	for rows.Next() {
		user := &domain.User{}
		err = rows.Scan(user)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Save save user to storage
func (repo *UserRepositoryImpl) Save(u *domain.User) error {
	stmt, err := repo.Conn.Prepare("insert into users(username, passhash, email, fullname) values(?, ?, ?, ?) returning id")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(u.Username, u.Passhash, u.Email, u.FullName)

	return err
}

// Delete user from storage
func (repo *UserRepositoryImpl) Delete(id int64) error {
	return nil
}
