package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("starting converter program..")

	data, _ := ioutil.ReadFile("/Users/pivotal/Desktop/data.csv")

	lines := strings.Split(string(data), "\n")
	lines = lines[1:len(lines)]

	var rowsAndColumns [][]string

	for _, line := range lines {
		rowsAndColumns = append(rowsAndColumns, strings.Split(line, ","))
	}

	var sql string

	for _, columns := range rowsAndColumns {
		sql += fmt.Sprintf(`INSERT INTO product_service_attribute (product_id, service_id, service_type, service_name, service_reference) VALUES ('%s', '%s', '%s', '%s', '%s');
`, columns[1], columns[3], columns[4], columns[5], columns[6])
	}

	fmt.Println(sql)

	fmt.Println("converter program done")
}
