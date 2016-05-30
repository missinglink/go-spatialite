package spatialite

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
)

type entrypoint struct {
	lib  string
	proc string
}

var LibNames = []entrypoint{
	{"mod_spatialite", "sqlite3_modspatialite_init"},
	{"libspatialite7", "sqlite3_spatialite_init"},
	{"libspatialite5", "sqlite3_spatialite_init"},
	{"libspatialite3", "sqlite3_spatialite_init"},
	{"libspatialite", "sqlite3_modspatialite_init"},
}

var ErrSpatialiteNotFound = errors.New("shaxbee/go-spatialite: spatialite extension not found.")

func init() {
	sql.Register("spatialite", &sqlite3.SQLiteDriver{
		ConnectHook: func(conn *sqlite3.SQLiteConn) error {
			for _, v := range LibNames {
				if err := conn.LoadExtension(v.lib, v.proc); err == nil {
					return nil
				}
			}
			return ErrSpatialiteNotFound
		},
	})
}
