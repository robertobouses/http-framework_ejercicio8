--delete_measurement.sql

DELETE FROM arithmetic.measurements WHERE id = $1;
