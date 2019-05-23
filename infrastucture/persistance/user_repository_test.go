package persistance

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/soerjadi/GoBlog/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var err error

func TestSaveUser(t *testing.T) {

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func() {
		err = db.Close()
		require.NoError(t, err)
	}()

	user := &domain.User{
		ID:       12,
		Username: "name",
		Passhash: "123123",
		Email:    "email@mail.com",
		FullName: "Administrator",
	}

	query := "insert into users\\(username, passhash, email, fullname\\) values\\(\\?, \\?, \\?, \\?\\) returning id"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(user.Username, user.Passhash, user.Email, user.FullName).WillReturnResult(sqlmock.NewResult(12, 1))

	repo := UserRepositoryWithRDB(db)

	err = repo.Save(user)

	assert.NoError(t, err)
	assert.Equal(t, int64(12), user.ID)
}
