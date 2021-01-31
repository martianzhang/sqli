package sqli

import (
	"testing"
)

func TestSQLi(t *testing.T) {
	sqls := []string{
		"asdf asd ; -1' and 1=1 union/* foo */select load_file('/etc/passwd')--",
		"SELECT * FROM users WHERE username='asdsad' AND password='e2a521bc01c1ca09e173bcf65bcc97e9'",
	}
	for _, sql := range sqls {
		if fp, is := (SQLi(sql)); !is {
			t.Error(fp, is)
		}
	}
}
