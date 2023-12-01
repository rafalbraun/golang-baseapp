package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/*
https://stackoverflow.com/questions/73827790/struct-literal-uses-unkeyed-fields
https://stackoverflow.com/questions/54548441/composite-literal-uses-unkeyed-fields/63747469#63747469
*/

type Conn struct {
	desc            *gorm.DB
	audit           *logger.Interface
	Error           error
	RowsAffected    int64
	//auditLog     io.Writer
}

func New(db *gorm.DB, auditLogger *logger.Interface) *Conn {
	return &Conn{
		desc:  db,
		audit: auditLogger,
	}
}

func cast(db *gorm.DB, auditLogger *logger.Interface) *Conn {
	return &Conn{
		desc:         db,
		audit:        auditLogger,
		Error:        db.Error,
		RowsAffected: db.RowsAffected,
	}
}

func (conn *Conn) WriteAuditLog() *Conn {
	tx := conn.desc.Session(&gorm.Session{
		Logger: *conn.audit,
	})
	return cast(tx, conn.audit)
}

func (conn *Conn) Create(value interface{}) *Conn {
	tx := conn.desc.Create(value)
	return cast(tx, conn.audit)
}

func (conn *Conn) Model(value interface{}) *Conn {
	tx := conn.desc.Model(value)
	return cast(tx, conn.audit)
}

func (conn *Conn) Preload(query string, args ...interface{}) *Conn {
	tx := conn.desc.Preload(query, args...)
	return cast(tx, conn.audit)
}

func (conn *Conn) Where(query interface{}, args ...interface{}) *Conn {
	tx := conn.desc.Where(query, args...)
	return cast(tx, conn.audit)
}

func (conn *Conn) Joins(query string, args ...interface{}) *Conn {
	tx := conn.desc.Joins(query, args...)
	return cast(tx, conn.audit)
}

func (conn *Conn) Order(value interface{}) *Conn {
	tx := conn.desc.Order(value)
	return cast(tx, conn.audit)
}

func (conn *Conn) First(dest interface{}, conds ...interface{}) *Conn {
	tx := conn.desc.First(dest, conds...)
	return cast(tx, conn.audit)
}

func (conn *Conn) Delete(value interface{}, conds ...interface{}) *Conn {
	tx := conn.desc.Delete(value, conds...)
	return cast(tx, conn.audit)
}

func (conn *Conn) Unscoped() *Conn {
	tx := conn.desc.Unscoped()
	return cast(tx, conn.audit)
}

func (conn *Conn) Save(value interface{}) *Conn {
	tx := conn.desc.Save(value)
	return cast(tx, conn.audit)
}

func (conn *Conn) Find(dest interface{}, conds ...interface{}) *Conn {
	tx := conn.desc.Find(dest, conds...)
	return cast(tx, conn.audit)
}

func (conn *Conn) Debug() *Conn {
	tx := conn.desc.Debug()
	return cast(tx, conn.audit)
}
