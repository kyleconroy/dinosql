package dinosql

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParserErrors(t *testing.T) {
	for _, tc := range []struct {
		query string
		err   Error
	}{
		{
			"SELECT foo FROM bar WHERE baz = $4;",
			Error{Code: "42P18", Message: "could not determine data type of parameter $1"},
		},
		{
			"SELECT foo FROM bar WHERE baz = $1 AND baz = $3;",
			Error{Code: "42P18", Message: "could not determine data type of parameter $2"},
		},
		{
			"ALTER TABLE unknown RENAME TO known;",
			Error{Code: "42P01", Message: "relation \"unknown\" does not exist"},
		},
		{
			"ALTER TABLE unknown DROP COLUMN dropped;",
			Error{Code: "42P01", Message: "relation \"unknown\" does not exist"},
		},
		{
			`
			CREATE TABLE bar (id serial not null);

			-- name: foo :one
			SELECT foo FROM bar;
			`,
			Error{Code: "42703", Message: "column \"foo\" does not exist"},
		},
	} {
		test := tc
		t.Run(test.query, func(t *testing.T) {
			_, err := parseSQL(test.query)

			var actual Error
			if err != nil {
				actual = err.(Error)
			}

			if diff := cmp.Diff(test.err, actual); diff != "" {
				t.Errorf("error mismatch: \n%s", diff)
			}
		})
	}
}
