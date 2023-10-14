package repository

import (
	"database/sql"
	"log"

	_ "embed"

	"github.com/robertobouses/http-framework_ejercicio8/entity"
)

type REPOSITORY interface {
	InsertMeasurement(medicion entity.Measurement) error
	ExtractMeasurement() ([]entity.Measurement, error)
	DeleteMeasurement(id string) error
	ExtractMeasurementId(id string) (entity.Measurement, error)
	DeleteAllMeasurement() error
	DeleteEmptyMeasurement() error
}

type repository struct {
	db              *sql.DB
	insertStmt      *sql.Stmt
	extractStmt     *sql.Stmt
	deleteStmt      *sql.Stmt
	extractIdStmt   *sql.Stmt
	deleteAllStmt   *sql.Stmt
	deleteEmptyStmt *sql.Stmt
}

//go:embed sql/insert_measurement.sql
var insertMeasurementQuery string

//go:embed sql/extract_measurement.sql
var extractMeasurementQuery string

//go:embed sql/delete_measurement.sql
var deleteMeasurementQuery string

//go:embed sql/extractid_measurement.sql
var extractIdMeasurementQuery string

//go:embed sql/deleteall_measurement.sql
var deleteAllMeasurementQuery string

//go:embed sql/deleteempty_measurement.sql
var deleteEmptyMeasurementQuery string

func NewRepository(db *sql.DB) (*repository, error) {
	insertStmt, err := db.Prepare(insertMeasurementQuery)
	if err != nil {
		return nil, err
	}

	extractStmt, err := db.Prepare(extractMeasurementQuery)
	if err != nil {
		return nil, err
	}

	deleteStmt, err := db.Prepare(deleteMeasurementQuery)
	if err != nil {
		return nil, err
	}

	extractIdStmt, err := db.Prepare(extractIdMeasurementQuery)
	if err != nil {
		return nil, err
	}

	deleteAllStmt, err := db.Prepare(deleteAllMeasurementQuery)
	if err != nil {
		return nil, err
	}

	deleteEmptyStmt, err := db.Prepare(deleteEmptyMeasurementQuery)
	if err != nil {
		return nil, err
	}

	return &repository{
		db:              db,
		insertStmt:      insertStmt,
		extractStmt:     extractStmt,
		deleteStmt:      deleteStmt,
		extractIdStmt:   extractIdStmt,
		deleteAllStmt:   deleteAllStmt,
		deleteEmptyStmt: deleteEmptyStmt,
	}, nil
}

func (r *repository) InsertMeasurement(medicion entity.Measurement) error {
	_, err := r.insertStmt.Exec(medicion.Id, medicion.ValorX, medicion.ValorY, medicion.ValorZ)
	return err
}

func (r *repository) ExtractMeasurement() ([]entity.Measurement, error) {
	log.Println("SQL Query", extractMeasurementQuery)
	rows, err := r.db.Query(extractMeasurementQuery)

	if err != nil {
		log.Printf("Error al ejecutar la consulta SQL", err)
		return nil, err
	}
	defer rows.Close()
	var measurements []entity.Measurement

	for rows.Next() {
		var measurement entity.Measurement
		if err := rows.Scan(&measurement.Id, &measurement.ValorX, &measurement.ValorY, &measurement.ValorZ); err != nil {
			log.Printf("Error al escanear filas", err)
			return nil, err
		}
		measurements = append(measurements, measurement)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error al recorrer filas", err)
		return nil, err
	}

	return measurements, nil
}

func (r *repository) DeleteMeasurement(id string) error {
	_, err := r.db.Exec(deleteMeasurementQuery, id)
	return err
}
func (r *repository) ExtractMeasurementId(id string) (entity.Measurement, error) {
	//query := "SELECT id, valorx, valory, valorz FROM arithmetic.measurements WHERE id = $1"

	row := r.db.QueryRow(extractIdMeasurementQuery, id)

	var measurement entity.Measurement

	if err := row.Scan(&measurement.Id, &measurement.ValorX, &measurement.ValorY, &measurement.ValorZ); err != nil {
		if err == sql.ErrNoRows {
			return entity.Measurement{}, nil
		} else {
			log.Printf("Error al escanear fila", err)
			return entity.Measurement{}, err
		}
	}

	return measurement, nil
}

func (r *repository) DeleteAllMeasurement() error {
	_, err := r.db.Exec(deleteAllMeasurementQuery)
	return err
}

func (r *repository) DeleteEmptyMeasurement() error {
	_, err := r.db.Exec(deleteEmptyMeasurementQuery)
	return err
}
