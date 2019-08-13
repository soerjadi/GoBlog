package persistance

import (
	"github.com/soerjadi/GoBlog/database"
	"reflect"
	"testing"

	"github.com/soerjadi/GoBlog/domain"
	"github.com/soerjadi/GoBlog/infrastructure/persistance"
	"github.com/stretchr/testify/assert"
)

func TestSaveUser(t *testing.T) {

	conn := database.RDB().InitTestDB()
	db := database.DBTestRepository(conn)

	db.Clean("users")

	user := &domain.User{
		ID:       1,
		Username: "name",
		Passhash: "123123",
		Email:    "email@mail.com",
		FullName: "Administrator",
	}

	repo := persistance.UserRepositoryWithRDB(conn)

	newUser, err := repo.Save(user)

	assert.NoError(t, err)
	assert.Equal(t, newUser, user)
}

func TestGetListUser(t *testing.T) {

	conn := database.RDB().InitTestDB()
	db := database.DBTestRepository(conn)

	db.Clean("users")

	repo := persistance.UserRepositoryWithRDB(conn)

	users := []*domain.User{
		{
			ID:       1,
			Username: "name",
			Passhash: "123123",
			Email:    "email@mail.com",
			FullName: "Administrator",
		},
		{
			ID:       2,
			Username: "name2",
			Passhash: "123123",
			Email:    "email@mail.com",
			FullName: "Administrator",
		},
	}

	for _, user := range users {
		_, _ = repo.Save(user)
	}

	users, err := repo.GetList(0, 10)

	assert.NoError(t, err)
	assert.Equal(t, len(users), 2)

}

func TestGetById(t *testing.T) {

	conn := database.RDB().InitTestDB()
	db := database.DBTestRepository(conn)

	db.Clean("users")

	user := &domain.User{
		ID:       1,
		Username: "name",
		Passhash: "123123",
		Email:    "email@mail.com",
		FullName: "Administrator",
	}

	repo := persistance.UserRepositoryWithRDB(conn)

	_, err := repo.Save(user)

	if err != nil {
		t.Fatalf("error when save user")
	}

	u, err := repo.GetByID(int64(1))

	if err != nil {
		t.Fatalf("want error %#v, got %#v", nil, err)
	}

	if !reflect.DeepEqual(u.ID, user.ID) {
		t.Errorf("want %d, got %d", user.ID, u.ID)
	}

}

func TestUpdateUser(t *testing.T) {

	conn := database.RDB().InitTestDB()
	db := database.DBTestRepository(conn)

	db.Clean("users")
	repo := persistance.UserRepositoryWithRDB(conn)

	user := &domain.User{
		ID:       1,
		Username: "name",
		Passhash: "123123",
		Email:    "email@mail.com",
		FullName: "Administrator",
	}

	_, err := repo.Save(user)

	if err != nil {
		t.Fatal("error when save user")
	}

	updateUser := &domain.User{
		ID:       1,
		Username: "name",
		Passhash: "123123",
		Email:    "update@mail.com",
		FullName: "Administrator",
	}
	newUser, err := repo.Update(updateUser)

	if err != nil {
		t.Fatalf("got error when update %#v", err)
	}

	if reflect.DeepEqual(newUser, user) {
		t.Errorf("user not updated old : %#v, new : %#v", user, newUser)
	}

	assert.NoError(t, err)
}
