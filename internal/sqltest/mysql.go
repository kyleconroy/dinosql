package sqltest

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/kyleconroy/sqlc/internal/sql/sqlpath"

	_ "github.com/go-sql-driver/mysql"
)

func MySQL(t *testing.T, migrations []string) (*sql.DB, func()) {
	t.Helper()

	data := os.Getenv("MYSQL_DATABASE")
	host := os.Getenv("MYSQL_HOST")
	pass := os.Getenv("MYSQL_ROOT_PASSWORD")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")

	if user == "" {
		user = "root"
	}

	if pass == "" {
		pass = "mysecretpassword"
	}

	if port == "" {
		port = "3306"
	}

	if host == "" {
		host = "127.0.0.1"
	}

	if data == "" {
		data = "dinotest"
	}

	source := fmt.Sprintf("%s:%s@%s:%s/%s", user, pass, host, port, data)
	t.Logf("db: %s", source)

	db, err := sql.Open("mysql", source)
	if err != nil {
		t.Fatal(err)
	}

	files, err := sqlpath.Glob(migrations)
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		blob, err := ioutil.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := db.Exec(string(blob)); err != nil {
			t.Fatalf("%s: %s", filepath.Base(f), err)
		}
	}

	return db, func() {}
}
