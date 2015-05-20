package models

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "gopkg.in/gorp.v1"
    "github.com/copperhead-security/android_ota_server/lib"
)

var dbmap *gorp.DbMap

func InitDb(dbPath string) *gorp.DbMap {
  db, err := sql.Open("sqlite3", dbPath)
  lib.CheckErr(err, "sql.Open failed")

  dbmap = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
  dbmap.AddTableWithName(Release{}, "releases").SetKeys(true, "Id")
  dbmap.AddTableWithName(User{}, "users").SetKeys(true, "Id")

  // dbmap.DropTables()
  // err = dbmap.TruncateTables()

  err = dbmap.CreateTablesIfNotExists()
  lib.CheckErr(err, "Create tables failed")
  return dbmap
}