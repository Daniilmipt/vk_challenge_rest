package service

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"rest/errcstm"
	"rest/jwt"
	"rest/model"
	"rest/queries"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (us *UserService) CreateUser(user *model.User, tx *sql.Tx) (*model.User, error) {
	var userDb model.User
	row := tx.QueryRow(queries.SELECT_USER_BY_LOGIN, user.Login)
	err := row.Scan(&userDb.ID, &userDb.Login, &userDb.Password)
	if err != nil && err != sql.ErrNoRows {
		return &userDb, errors.New(errcstm.USER_SELECT + err.Error())
	}

	if errors.Is(err, sql.ErrNoRows) {
		user.CreateId()
		_, err = tx.Exec(queries.INSERT_USER, user.ID, user.Login, user.Password)
		if err != nil {
			tx.Rollback()
			return user, errors.New(errcstm.USER_INSERT + err.Error())
		}
		return user, nil
	}
	return &userDb, errors.New("user has already registered")
}

func (us *UserService) LoginUser(r *http.Request) (model.Token, int, error) {
	var token model.Token
	if r.Method != "POST" {
		return token, http.StatusBadRequest, errors.New(errcstm.REQUEST_TYPE)
	}

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return token, http.StatusInternalServerError, errors.New(errcstm.REQUEST_BODY + err.Error())
	}

	var userDb model.User
	row := us.db.QueryRow(queries.SELECT_USER_BY_LOGIN_PASSWORD, user.Login, user.Password)
	err = row.Scan(&userDb.ID, &userDb.Login, &userDb.Password)
	if err != nil {
		return token, http.StatusInternalServerError, errors.New(errcstm.USER_SELECT + err.Error())
	}

	jwtStr, err := jwt.GenerateJWT(&userDb)
	if err != nil {
		return token, http.StatusNonAuthoritativeInfo, errors.New(errcstm.JWT_TOKEN + err.Error())
	}
	return model.Token{Value: jwtStr}, http.StatusOK, nil
}
