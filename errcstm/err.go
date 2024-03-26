package errcstm

const (
	DATABASE_OPEN = "can not open database: "
	REQUEST_TYPE  = "incorrect request type"
	REQUEST_BODY  = "incorrect user body request: "

	USER_SELECT = "can not read user from db: "
	USER_INSERT = "can not insert user in db: "

	ROLE_SELECT = "can not read role from db: "
	ROLE_INSERT = "can not insert role in db: "

	ADVERTISMENT_SELECT = "can not read advertisement from db: "
	ADVERTISMENT_INSERT = "can not insert advertisement in db: "

	JWT_TOKEN = "error with JWT"
)
