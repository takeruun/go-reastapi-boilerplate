package config

import (
	"net/http"

	"github.com/wader/gormstore/v2"
	"gorm.io/gorm"
)

const (
	SESSION_KEY_NAME = "_goreset_session"
)

type SESSION_STORE struct {
	SECRET_HASH_KEY string
}

func NewSessionStore(db *DB) *gormstore.Store {
	c := NewConfig()
	return newSessionStore(db.DB, &SESSION_STORE{
		SECRET_HASH_KEY: c.SESSION_STORE.SecretHashKey,
	})
}

func newSessionStore(db *gorm.DB, ss *SESSION_STORE) *gormstore.Store {
	store := gormstore.NewOptions(
		db,
		gormstore.Options{},
		[]byte(ss.SECRET_HASH_KEY),
	)

	store.SessionOpts.Secure = false // postman での使用は false に
	store.SessionOpts.HttpOnly = true
	store.SessionOpts.MaxAge = 60 * 60 * 24
	store.SessionOpts.SameSite = http.SameSiteNoneMode

	return store
}
