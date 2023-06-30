package product

import(
	"bytes"
	"neversitup/db"
	"database/sql"
	"log"
)

func prepareSQLInsertProduct(req Product) (result string, bindValue []interface{}) {
	var buffer bytes.Buffer
	buffer.WriteString("INSERT INTO products (product_name,product_desc,product_type,price,created_date) VALUES ($1, $2, $3,$4,now()) ")
	bindValue = append(bindValue, req.ProductName, req.ProductDesc, req.ProductType, req.Price)
	result = buffer.String()
	return result, bindValue
}

func prepareSQLGetProductByName(name string) (result string, bindValue []interface{}) {
	var buffer bytes.Buffer
	buffer.WriteString("SELECT 1 FROM products WHERE product_name = $1 ")
	bindValue = append(bindValue, name)
	result = buffer.String()
	return result, bindValue
}

func getProductByName(name string) (chk string, err error) {
	dbPool := db.GetConnectPostgresDB()
	statement, bindValue := prepareSQLGetProductByName(name)
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

func getAllProduct(proType, id string, page, size int) (products []Product, err error) {
	dbPool := db.GetConnectPostgresDB()
	statement, bindValue := prepareSQLGetAllProduct(proType, id, page, size)
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
		var p Product
		rows.Scan(
			&p.Id,
			&p.ProductName,
			&p.ProductDesc,
			&p.ProductType,
			&p.Price,
			&p.CreatedDate,
		)
		products = append(products, p)
	}
	return products, err
}

func prepareSQLGetAllProduct(proType, id string, page, size int) (result string, bindValue []interface{}) {
	var buffer bytes.Buffer
	buffer.WriteString("SELECT id, product_name, product_desc, product_type, price, created_date FROM products WHERE 1=1 ")
	if len(id) > 0 {
		buffer.WriteString("AND id = $1 ")
		bindValue = append(bindValue, id)
	}
	if len(proType) > 0 {
		buffer.WriteString("AND product_type = $1 ")
		bindValue = append(bindValue, proType)
	}
	if page >= 0 && size > 0 {
		buffer.WriteString("LIMIT $2 OFFSET $1 ")
		bindValue = append(bindValue, page, size)
	}
	result = buffer.String()
	return result, bindValue
}