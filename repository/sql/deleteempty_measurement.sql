--deleteempty_measurement.sql:

DELETE FROM arithmetic.measurements
WHERE id IS NULL OR id = '';
