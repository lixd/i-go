package orm

import (
	"database/sql"

	dialect2 "i-go/7days/orm/dialect"
	log2 "i-go/7days/orm/log"
	session2 "i-go/7days/orm/session"
)

// Engine is the main struct of geeorm, manages all db sessions and transactions.
type Engine struct {
	db      *sql.DB
	dialect dialect2.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log2.Error(err)
		return
	}
	// Send a ping to make sure the database connection is alive.
	if err = db.Ping(); err != nil {
		log2.Error(err)
		return
	}
	// make sure the specific dialect exists
	dial, ok := dialect2.GetDialect(driver)
	if !ok {
		log2.Errorf("dialect %s Not Found", driver)
		return
	}
	e = &Engine{db: db, dialect: dial}
	log2.Info("Connect database success")
	return
}

// Close database connection
func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log2.Error("Failed to close database")
	}
	log2.Info("Close database success")
}

func (engine *Engine) NewSession() *session2.Session {
	return session2.New(engine.db, engine.dialect)
}
