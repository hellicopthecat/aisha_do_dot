SELECT
  rd.*,
  rdt.*,
  rdd.*
FROM 
  routines_detail rd
LEFT JOIN
  routines_detail_tags rdt
  ON rdt.routine_detail_id = rd.id
LEFT JOIN
  routines_detail_desc rdd
  ON rdd.routine_detail_id = rd.id
WHERE 
  rd.routine_id = $1;