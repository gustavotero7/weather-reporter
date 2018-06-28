package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTables_20180628_100431 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTables_20180628_100431{}
	m.Created = "20180628_100431"

	migration.Register("CreateTables_20180628_100431", m)
}

// Run the migrations
func (m *CreateTables_20180628_100431) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`CREATE TABLE reports (
		id int(11) NOT NULL AUTO_INCREMENT,
		location_name varchar(100) DEFAULT NULL,
		temperature varchar(100) DEFAULT NULL,
		wind varchar(100) DEFAULT NULL,
		pressure varchar(100) DEFAULT NULL,
		humidity varchar(100) DEFAULT NULL,
		sunrise varchar(100) DEFAULT NULL,
		sunset varchar(100) DEFAULT NULL,
		geo_coordinates varchar(100) DEFAULT NULL,
		requested_time datetime DEFAULT NULL,
		code_name varchar(100) NOT NULL,
		PRIMARY KEY (id)
	  ) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=latin1;`)

}

// Reverse the migrations
func (m *CreateTables_20180628_100431) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL(`DROP TABLE reports`)
}
