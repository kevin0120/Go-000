package wraperror

import (
	"database/sql"
	"github.com/pkg/errors"
)

func GetUser(id int) (string, error) {
	var user = ""
	if user, err := getFromDB(id); errors.Is(err, sql.ErrNoRows) {
		return user, errors.Wrapf(err, "User not found")
	}
	return user, nil
}

func getFromDB(id int) (string, error) {
	return "", sql.ErrNoRows
}
