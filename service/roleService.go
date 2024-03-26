package service

import (
	"database/sql"
	"errors"
	constan "rest/constant"
	"rest/errcstm"
	"rest/model"
	"rest/queries"
)

type RoleService struct {
	db *sql.DB
}

func NewRoleService(db *sql.DB) *RoleService {
	return &RoleService{
		db: db,
	}
}

func (rs *RoleService) CreateRole(user *model.User, errUser error, tx *sql.Tx) error {
	var roleDb model.Role
	row := tx.QueryRow(queries.SELECT_ROLE, constan.USER)
	err := row.Scan(&roleDb.ID, &roleDb.Role)
	if err != nil && err != sql.ErrNoRows {
		return errors.New(errcstm.ROLE_SELECT + err.Error())
	}

	if errors.Is(err, sql.ErrNoRows) {
		roleDb.CreateId()
		roleDb.Role = constan.USER
		_, err = tx.Exec(queries.INSERT_ROLE, roleDb.ID, roleDb.Role)
		if err != nil {
			tx.Rollback()
			return errors.New(errcstm.ROLE_INSERT + err.Error())
		}
	}

	if errUser == nil || errUser.Error() != "user has already registered" {
		_, err = tx.Exec(queries.INSERT_ROLE_X_USER, user.ID, roleDb.ID)
		if err != nil {
			tx.Rollback()
			return errors.New("can not insert in user_x_roles" + err.Error())
		}
	}
	return nil
}
