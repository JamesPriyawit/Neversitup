package user

import (
	"bytes"
	"log"
	"neversitup/db"
	"database/sql"
)

func prepareSQLInsertUser(req User, password []byte) (result string, bindValue []interface{}) {
	var buffer bytes.Buffer
	buffer.WriteString("INSERT INTO users (username,password,firstname,lastname,created_date) VALUES ($1, $2, $3, $4,now()) ")
	bindValue = append(bindValue, req.Username, password, req.Firstname, req.Lastname)
	result = buffer.String()
	return result, bindValue
}

func prepareSQLUserByUsername(username string) (result string, bindValue []interface{}) {
	var buffer bytes.Buffer
	buffer.WriteString("SELECT 1 FROM users WHERE username = $1 ")
	bindValue = append(bindValue, username)
	result = buffer.String()
	return result, bindValue
}

func getUserByUsername(username string) (chk string, err error) {
	dbPool := db.GetConnectPostgresDB()
	statement, bindValue := prepareSQLUserByUsername(username)
	prepare, err := dbPool.Prepare(statement)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer prepare.Close()
	var rows *sql.Rows
	rows, err = prepare.Query(bindValue...)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&chk)
	}
	return chk, err
}

func prepareSQLGetUserById(id string) (result string, bindValue []interface{}) {
	var buffer bytes.Buffer
	buffer.WriteString("SELECT id,username,firstname,lastname,created_date FROM users WHERE id = $1 ")
	bindValue = append(bindValue, id)
	result = buffer.String()
	return result, bindValue
}

func getUserById(id string) (user User, err error) {
	dbPool := db.GetConnectPostgresDB()
	statement, bindValue := prepareSQLGetUserById(id)
	prepare, err := dbPool.Prepare(statement)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer prepare.Close()
	var rows *sql.Rows
	rows, err = prepare.Query(bindValue...)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(
			&user.Id,
			&user.Username,
			&user.Firstname,
			&user.Lastname,
			&user.CreatedDate,
		)
	}
	return user, err
}