package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"rest/database"
	"rest/errcstm"
	"rest/jwt"
	"rest/model"
)

func CreateUser(r *http.Request, userService *UserService, roleService *RoleService) (model.User, int, error) {
	var user model.User

	db, err := database.DbInit()
	if err != nil {
		return user, http.StatusInternalServerError, errors.New(errcstm.DATABASE_OPEN + err.Error())
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return user, http.StatusInternalServerError, errors.New(err.Error())
	}

	if r.Method != "POST" {
		return user, http.StatusBadRequest, errors.New(errcstm.REQUEST_TYPE)
	}

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return user, http.StatusInternalServerError, errors.New(errcstm.REQUEST_BODY + err.Error())
	}

	userSaved, err := userService.CreateUser(&user, tx)
	if err != nil && err.Error() != "user has already registered" {
		return user, http.StatusInternalServerError, err
	}
	if err = roleService.CreateRole(userSaved, err, tx); err != nil {
		return user, http.StatusInternalServerError, err
	}

	if err = tx.Commit(); err != nil {
		return user, http.StatusInternalServerError, err
	}

	_, err = jwt.GenerateJWT(userSaved)
	if err != nil {
		return user, http.StatusNonAuthoritativeInfo, errors.New(errcstm.JWT_TOKEN + err.Error())
	}
	return *userSaved, http.StatusOK, nil
}
