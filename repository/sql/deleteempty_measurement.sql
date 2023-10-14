DELETE FROM arithmetic.measurements

WHERE (entity IS NULL OR entity = '') AND (id IS NULL OR id = '');
