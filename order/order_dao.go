package order

import (
	"bytes"
	"database/sql"
	"log"
	"neversitup/db"
	"neversitup/constant"
)

func prepareSQLInsertOrder(req Order) (result string, bindValue []interface{}) {
	var buffer bytes.Buffer
	buffer.WriteString("INSERT INTO orders (user_id,product_id,status,created_date) VALUES ($1, $2, $3, now()) ")
	bindValue = append(bindValue, req.UserId, req.ProductId, constant.Inprogress)
	result = buffer.String()
	return result, bindValue
}

func getOrderById(orderId, userId, status string, page, size int) (orders []Order, err error) {
	dbPool := db.GetConnectPostgresDB()
	statement, bindValue := prepareSQLGetOrderById(orderId, userId, status, page, size)
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
		var o Order
		rows.Scan(
			&o.Id,
			&o.Status,
			&o.CreatedDate,
		)
		orders = append(orders, o)
	}
	return orders, err
}

func prepareSQLGetOrderById(orderId, userId, status string, page, size int) (result string, bindValue []interface{}) {
	var buffer bytes.Buffer
	buffer.WriteString("SELECT id, status, created_date FROM orders WHERE 1=1 ")
	if len(userId) > 0 {
		buffer.WriteString("AND user_id = $1 ")
		bindValue = append(bindValue, userId)
		if len(status) > 0 {
			buffer.WriteString("AND status = $2 ")
			bindValue = append(bindValue, status)
			if page >= 0 && size > 0 {
				buffer.WriteString("LIMIT $4 OFFSET $3 ")
				bindValue = append(bindValue, page, size)
			}
		}else {
			if page >= 0 && size > 0 {
				buffer.WriteString("LIMIT $3 OFFSET $2 ")
				bindValue = append(bindValue, page, size)
			}
		}
	}
	if len(orderId) > 0 {
		buffer.WriteString("AND id = $1 ")
		bindValue = append(bindValue, orderId)
	}

	result = buffer.String()
	return result, bindValue
}

func prepareSQLDeleteOrder(id string) (result string, bindValue []interface{}) {
	var buffer bytes.Buffer
	buffer.WriteString("DELETE FROM orders WHERE id = $1 ")
	bindValue = append(bindValue, id)
	result = buffer.String()
	return result, bindValue
}