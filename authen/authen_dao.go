package authen

import(
	"database/sql"
	"log"
	"neversitup/db"
	"bytes"
)

func prepareSQLGetUserByUsername(user string) (result string, bindValue []interface{}) {
	var buffer bytes.Buffer
	buffer.WriteString("SELECT id, username, password FROM users WHERE username = $1 ")
	bindValue = append(bindValue, user)
	result = buffer.String()
	return result, bindValue
}

func getUserByUsername(user string) (id, username, password string, err error) {
	dbPool := db.GetConnectPostgresDB()
	statement, bindValue := prepareSQLGetUserByUsername(user)
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
			&id,
			&username,
			&password,
		)
	}
	return id, username, password, err
}