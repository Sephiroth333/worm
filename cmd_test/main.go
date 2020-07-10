package main

import (
	"fmt"
	"worm"
	//"worm/log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, _ := worm.NewEngine("sqlite3", "geetest.db")
	defer engine.Close()
	//实际一次会话中的sql都是由session实际执行的
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	//raw返回的是设置好的
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
