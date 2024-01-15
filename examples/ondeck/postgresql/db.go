// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package ondeck

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createCityStmt, err = db.PrepareContext(ctx, createCity); err != nil {
		return nil, fmt.Errorf("error preparing query CreateCity: %w", err)
	}
	if q.createVenueStmt, err = db.PrepareContext(ctx, createVenue); err != nil {
		return nil, fmt.Errorf("error preparing query CreateVenue: %w", err)
	}
	if q.deleteVenueStmt, err = db.PrepareContext(ctx, deleteVenue); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteVenue: %w", err)
	}
	if q.getCityStmt, err = db.PrepareContext(ctx, getCity); err != nil {
		return nil, fmt.Errorf("error preparing query GetCity: %w", err)
	}
	if q.getVenueStmt, err = db.PrepareContext(ctx, getVenue); err != nil {
		return nil, fmt.Errorf("error preparing query GetVenue: %w", err)
	}
	if q.listCitiesStmt, err = db.PrepareContext(ctx, listCities); err != nil {
		return nil, fmt.Errorf("error preparing query ListCities: %w", err)
	}
	if q.listVenuesStmt, err = db.PrepareContext(ctx, listVenues); err != nil {
		return nil, fmt.Errorf("error preparing query ListVenues: %w", err)
	}
	if q.updateCityNameStmt, err = db.PrepareContext(ctx, updateCityName); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateCityName: %w", err)
	}
	if q.updateVenueNameStmt, err = db.PrepareContext(ctx, updateVenueName); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateVenueName: %w", err)
	}
	if q.venueCountByCityStmt, err = db.PrepareContext(ctx, venueCountByCity); err != nil {
		return nil, fmt.Errorf("error preparing query VenueCountByCity: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createCityStmt != nil {
		if cerr := q.createCityStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createCityStmt: %w", cerr)
		}
	}
	if q.createVenueStmt != nil {
		if cerr := q.createVenueStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createVenueStmt: %w", cerr)
		}
	}
	if q.deleteVenueStmt != nil {
		if cerr := q.deleteVenueStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteVenueStmt: %w", cerr)
		}
	}
	if q.getCityStmt != nil {
		if cerr := q.getCityStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCityStmt: %w", cerr)
		}
	}
	if q.getVenueStmt != nil {
		if cerr := q.getVenueStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getVenueStmt: %w", cerr)
		}
	}
	if q.listCitiesStmt != nil {
		if cerr := q.listCitiesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listCitiesStmt: %w", cerr)
		}
	}
	if q.listVenuesStmt != nil {
		if cerr := q.listVenuesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listVenuesStmt: %w", cerr)
		}
	}
	if q.updateCityNameStmt != nil {
		if cerr := q.updateCityNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateCityNameStmt: %w", cerr)
		}
	}
	if q.updateVenueNameStmt != nil {
		if cerr := q.updateVenueNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateVenueNameStmt: %w", cerr)
		}
	}
	if q.venueCountByCityStmt != nil {
		if cerr := q.venueCountByCityStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing venueCountByCityStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                   DBTX
	tx                   *sql.Tx
	createCityStmt       *sql.Stmt
	createVenueStmt      *sql.Stmt
	deleteVenueStmt      *sql.Stmt
	getCityStmt          *sql.Stmt
	getVenueStmt         *sql.Stmt
	listCitiesStmt       *sql.Stmt
	listVenuesStmt       *sql.Stmt
	updateCityNameStmt   *sql.Stmt
	updateVenueNameStmt  *sql.Stmt
	venueCountByCityStmt *sql.Stmt
}

func (q *Queries) PrepareCreateCity(ctx context.Context) error {
	var err error
	if q.createCityStmt, err = q.db.PrepareContext(ctx, createCity); err != nil {
		return fmt.Errorf("error preparing query CreateCity: %w", err)
	}
	return nil
}

func (q *Queries) PrepareCreateVenue(ctx context.Context) error {
	var err error
	if q.createVenueStmt, err = q.db.PrepareContext(ctx, createVenue); err != nil {
		return fmt.Errorf("error preparing query CreateVenue: %w", err)
	}
	return nil
}

func (q *Queries) PrepareDeleteVenue(ctx context.Context) error {
	var err error
	if q.deleteVenueStmt, err = q.db.PrepareContext(ctx, deleteVenue); err != nil {
		return fmt.Errorf("error preparing query DeleteVenue: %w", err)
	}
	return nil
}

func (q *Queries) PrepareGetCity(ctx context.Context) error {
	var err error
	if q.getCityStmt, err = q.db.PrepareContext(ctx, getCity); err != nil {
		return fmt.Errorf("error preparing query GetCity: %w", err)
	}
	return nil
}

func (q *Queries) PrepareGetVenue(ctx context.Context) error {
	var err error
	if q.getVenueStmt, err = q.db.PrepareContext(ctx, getVenue); err != nil {
		return fmt.Errorf("error preparing query GetVenue: %w", err)
	}
	return nil
}

func (q *Queries) PrepareListCities(ctx context.Context) error {
	var err error
	if q.listCitiesStmt, err = q.db.PrepareContext(ctx, listCities); err != nil {
		return fmt.Errorf("error preparing query ListCities: %w", err)
	}
	return nil
}

func (q *Queries) PrepareListVenues(ctx context.Context) error {
	var err error
	if q.listVenuesStmt, err = q.db.PrepareContext(ctx, listVenues); err != nil {
		return fmt.Errorf("error preparing query ListVenues: %w", err)
	}
	return nil
}

func (q *Queries) PrepareUpdateCityName(ctx context.Context) error {
	var err error
	if q.updateCityNameStmt, err = q.db.PrepareContext(ctx, updateCityName); err != nil {
		return fmt.Errorf("error preparing query UpdateCityName: %w", err)
	}
	return nil
}

func (q *Queries) PrepareUpdateVenueName(ctx context.Context) error {
	var err error
	if q.updateVenueNameStmt, err = q.db.PrepareContext(ctx, updateVenueName); err != nil {
		return fmt.Errorf("error preparing query UpdateVenueName: %w", err)
	}
	return nil
}

func (q *Queries) PrepareVenueCountByCity(ctx context.Context) error {
	var err error
	if q.venueCountByCityStmt, err = q.db.PrepareContext(ctx, venueCountByCity); err != nil {
		return fmt.Errorf("error preparing query VenueCountByCity: %w", err)
	}
	return nil
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                   tx,
		tx:                   tx,
		createCityStmt:       q.createCityStmt,
		createVenueStmt:      q.createVenueStmt,
		deleteVenueStmt:      q.deleteVenueStmt,
		getCityStmt:          q.getCityStmt,
		getVenueStmt:         q.getVenueStmt,
		listCitiesStmt:       q.listCitiesStmt,
		listVenuesStmt:       q.listVenuesStmt,
		updateCityNameStmt:   q.updateCityNameStmt,
		updateVenueNameStmt:  q.updateVenueNameStmt,
		venueCountByCityStmt: q.venueCountByCityStmt,
	}
}
