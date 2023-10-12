package repository

import (
	"database/sql"
	_ "embed"

	"github.com/robertobouses/http-framework_ejercicio8/entity"
)

type REPOSITORY interface {
	InsertMeasurement(medicion entity.Measurement) error
}

type repository struct {
	db         *sql.DB
	insertStmt *sql.Stmt
	//printStmt  *sql.Stmt
}

//go:embed sql/insert_measurement.sql
var insertMeasurementQuery string

func NewRepository(db *sql.DB) (*repository, error) {
	insertStmt, err := db.Prepare(insertMeasurementQuery)

	if err != nil {
		return nil, err
	}
	return &repository{
		db:         db,
		insertStmt: insertStmt,
	}, nil
}

func (r *repository) InsertMeasurement(medicion entity.Measurement) error {
	_, err := r.insertStmt.Exec(medicion.ValorX, medicion.ValorY, medicion.ValorZ)
	return err
}

/* insertMedicionQuery
const insertMeasurementQuery = "INSERT INTO arithmetic.measurements(valorx, valory, valorz) VALUES($1, $2, $3)"
*/
