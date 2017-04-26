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
	"log"
)

type Pagination struct {
	Page int
	Offset int
	Size int
}

type PaginationResult struct {
	Result interface{}
	Total int
}

type sqlFunction func(*sql.DB) (error)

type sqlPaginationFunction func(*sql.DB) (*PaginationResult, error)

type sqlFunctionTx func(*sql.Tx) (error)

type sqlPaginationFunctionTx func(*sql.Tx) (*PaginationResult, error)

//For no return needed except an error
func notTransaction(_func sqlFunction) (error) {

	db, err := Connect()

	if(err != nil) {

		log.Printf("Error opening DB connection: %s", err)
		return err;
	}

	return _func(db)
}

func notTransactionWithPagination(_func sqlPaginationFunction) (*PaginationResult, error) {

	db, err := Connect()

	if(err != nil) {

		log.Printf("Error opening DB connection: %s", err)
		return nil, err;
	}

	return _func(db)
}

func inTransactionWithPagination(_func sqlPaginationFunctionTx) (*PaginationResult, error) {

	db, err := Connect()

	if(err != nil) {

		log.Printf("Error opening DB connection: %s", err)
		return nil, err;
	}

	tx, err := db.Begin()

	if(err != nil) {

		log.Printf("Error starting DB transaction: %s", err)
		return nil, err;
	}

	result, err := _func(tx)

	if(err != nil) {

		log.Printf("Rollback transaction: %s", err)
		tx.Rollback()
	} else {

		log.Print("Commit transaction")
		tx.Commit()
	}

	return result, err;
}

//For no return needed except an error
func inTransaction(_func sqlFunctionTx) (error) {

	db, err := Connect()

	if(err != nil) {

		log.Printf("Error opening DB connection: %s", err)
		return err;
	}

	tx, err := db.Begin()

	if(err != nil) {

		log.Printf("Error starting DB transaction: %s", err)
		return err;
	}

	err = _func(tx)

	if(err != nil) {

		log.Printf("Rollback transaction: %s", err)
		tx.Rollback()
	} else {

		log.Print("Commit transaction")
		tx.Commit()
	}

	return err;
}

func count(tablename string) (int, error) {

	var count int = 0

	db, err := Connect()
	if(err != nil) {

		return count, err
	}

	err = db.QueryRow("select count(*) from " + tablename).Scan(&count)

	return count, err
}