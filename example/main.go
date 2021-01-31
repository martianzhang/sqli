package main

import (
	"fmt"

	"github.com/martianzhang/sqli"
)

func main() {
	sqls := []string{
		"asdf asd ; -1' and 1=1 union/* foo */select load_file('/etc/passwd')--",
		"SELECT * FROM users WHERE username='asdsad' AND password='e2a521bc01c1ca09e173bcf65bcc97e9'",
		"SELECT * FROM tb",
		"",
		" ",
	}
	for _, sql := range sqls {
		fmt.Println(sqli.SQLi(sql))
	}
}
