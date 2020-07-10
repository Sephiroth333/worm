package worm

import (
	"database/sql"

	"worm/log"
	"worm/session"
)

//包装db对象，对连接进行准备和收尾工作
type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	// Send a ping to make sure the database connection is alive.
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	e = &Engine{db: db}
	log.Info("Connect database success")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

//每次都是一个新的session引用
func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db)
}
