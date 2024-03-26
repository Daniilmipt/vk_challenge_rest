package queries

const (
	INSERT_ADV = "INSERT INTO advertisement (id, datetime, user_id, title, content, image_path, price) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	SELECT_ADV = "SELECT u.login, a.* FROM advertisement as a inner join users as u on u.id = a.user_id"

	SELECT_ROLE = "SELECT * FROM role WHERE role = $1"
	INSERT_ROLE = "INSERT INTO role (id, role) VALUES ($1, $2)"

	INSERT_ROLE_X_USER = "INSERT INTO user_x_roles (user_id, role_id) VALUES ($1, $2)"

	SELECT_USER_BY_ID             = "SELECT * FROM users WHERE id = $1"
	SELECT_USER_BY_LOGIN          = "SELECT * FROM users WHERE login = $1"
	SELECT_USER_BY_LOGIN_PASSWORD = "SELECT * FROM users WHERE login = $1 AND password = $2"
	INSERT_USER                   = "INSERT INTO users (id, login, password) VALUES ($1, $2, $3)"
)
