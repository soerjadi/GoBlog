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
	row := repo.Conn.QueryRow("select * from users where id=$1", id)

	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.FullName, &user.Passhash)

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
		err = rows.Scan(&user.ID, &user.Email, &user.FullName, &user.Passhash, &user.Username)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Save save user to storage
func (repo *UserRepositoryImpl) Save(u *domain.User) (*domain.User, error) {
	stmt, err := repo.Conn.Prepare("insert into users(username, passhash, email, fullname) values ($1, $2, $3, $4) returning id")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var ID int64
	err = stmt.QueryRow(u.Username, u.Passhash, u.Email, u.FullName).Scan(&ID)

	if err != nil {
		return nil, err
	}

	newUser, err := repo.GetByID(ID)

	return newUser, err
}

// Delete user from storage
func (repo *UserRepositoryImpl) Delete(id int64) error {
	stmt, err := repo.Conn.Prepare("delete from users wehe id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}

func (repo *UserRepositoryImpl) Update(user *domain.User) (*domain.User, error) {
	stmt, err := repo.Conn.Prepare("update users set fullname=$1, username=$2, email=$3, passhash=$4 where id=$5 returning id")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var ID int64
	err = stmt.QueryRow(user.FullName, user.Username, user.Email, user.Passhash, user.ID).Scan(&ID)

	if err != nil {
		return nil, err
	}

	return user, err
}
