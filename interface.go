package gorm

import "database/sql"

// SQLCommon is the minimal database connection functionality gorm requires.  Implemented by *sql.DB.
type SQLCommon interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	// 请求多行数据
	Query(query string, args ...interface{}) (*sql.Rows, error)
	// 请求单行数据
	QueryRow(query string, args ...interface{}) *sql.Row
}

type sqlDb interface {
	// 开始事务
	Begin() (*sql.Tx, error)
}

type sqlTx interface {
	Commit() error
	Rollback() error
}
