package mysql

import (
  "database/sql"
  "fmt"
  "strconv"
  "time"
)

var db *sql.DB

const (
  USERNAME = "miniboom"
  PASSWORD = "li1860123/-"
  SERVER   = "localhost"
  PORT     = 3306
  DBNAME = "go-nuomi"
  CHARSET = "utf8"
  PARSETIME = "true"
)

const (
  connMaxLifetime = 4 * time.Hour
  maxIdleConns    = 10
  maxOpenConns    = 10
)

func GetGlobalDB ()(*sql.DB){
  return db
}

func ActiveMysqlDb(){
  db, err := openDb()
  if err != nil{
    fmt.Errorf("open db failed:%v",err)
    return
  }
  db.SetConnMaxLifetime(connMaxLifetime)
  db.SetMaxIdleConns(maxIdleConns)
  db.SetMaxOpenConns(maxOpenConns)
}

// db, err := sql.Open("mysql", "user:password@tcp(localhost:5555)/dbname?charset=utf8&parseTime=true")
func openDb() (db *sql.DB, err error) {
  dbSql := USERNAME + ":" + PASSWORD + "@tcp(" + SERVER + ":" +
      strconv.Itoa(PORT) + ")/" + DBNAME + "?charset=" + CHARSET +
      "&parseTime=" + PARSETIME
  if db, err = sql.Open("mysql", dbSql); err != nil {
    return
  }
  return
}

