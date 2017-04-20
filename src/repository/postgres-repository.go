// Copyright 2017 Kevin Bayes
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"../config"
)

var db *sql.DB = nil;

var definitions = []string{
	"CREATE TABLE IF NOT EXISTS registries " +
		"( id serial primary key," +
		"  name VARCHAR(50) not null," +
		"  description VARCHAR(1000) not null," +
		"  uri VARCHAR(255) not null," +
		"  username VARCHAR(100) not null," +
		"  password VARCHAR(100) not null," +
		"  created_on TIMESTAMP not null," +
		"  version integer not null" +
		") ",
	"CREATE TABLE IF NOT EXISTS containers " +
		"( id serial primary key," +
		"  image VARCHAR(255) not null," +
		"  registry_id BIGINT not null," +
		"  state VARCHAR(15) not null," +
		"  shield VARCHAR(1000) not null," +
		"  DateUpdated TIMESTAMP," +
		"  created_on TIMESTAMP not null," +
		"  version integer not null" +
		") ",
	"create unique index if not exists containers_image_registry_id_uindex on containers (image, registry_id)",
	"create table if not exists container_image_report " +
	"( " +
		"id bigserial not null " +
		"constraint container_image_report_pkey " +
		"primary key, " +
		"image_id bigint not null " +
		"constraint container_image_report_containers_id_fk " +
		"references containers, " +
		"layer_id varchar(100) not null, " +
		"shield varchar(1000) not null, " +
		"created_on timestamp default now() not null " +
	");",
	"create table if not exists container_image_vulnerability_counts " +
	"( " +
		"id serial not null " +
		"constraint container_image_vulnerability_counts_pkey " +
		"primary key, " +
		"vulnerability_level varchar(20) not null, " +
		"count integer not null, " +
		"image_report_id bigint not null " +
		"constraint container_image_vulnerability_counts_container_image_report_id_ " +
		"references container_image_report " +
	");",
};

func init() {

	_, err := Connect()

	if(err != nil) {
		panic(err)
	}
}


func Connect() (*sql.DB, error) {

	if( db == nil ) {

		_db, err := openDatabase()

		if(err != nil) {

			//Should die if a connection cannot be established.
			panic(err)
		}

		db = _db

		for _, def := range definitions {

			createTable(def)
		}
	}

	return db, nil;
}

func openDatabase() (*sql.DB, error) {

	config := config.GetConfig()
	database := config.Database

	_db, err := sql.Open(database.Vendor, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable statement_timeout=60000", database.Host, database.Port, database.Username,
		database.Password, database.Database))

	if(err != nil) {
		panic(err)
	}

	_db.SetMaxIdleConns(10)
	_db.SetMaxOpenConns(100)

	return _db, err
}

func createTable(definition string) {

	stmt, err := db.Prepare(definition)
	if(err != nil) {

		panic(err)
	}

	result, err2 := stmt.Exec()
	if(err2 != nil) {

		panic(err2)
	}

	rows, _ := result.RowsAffected()

	println(rows)
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}