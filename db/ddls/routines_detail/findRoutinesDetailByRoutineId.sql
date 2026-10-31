SELECT
  rd.id,
  rd.created_at,
  rd.updated_at,
  rd.created_by,
  rd.updated_by,
  rd.priority_two_depth,
  rd.pre_event_start_at,
  rd.pre_event_end_at,
  rd.start_at,
  rd.end_at,
  rd.routine_id,
  COALESCE(rdt.tags, JSON_ARRAY()) as rdt,
  COALESCE(rdd.decriptions, JSON_ARRAY()) as rdd
FROM 
  routines_detail rd
LEFT JOIN
  (
    SELECT 
      id,
      routine_detail_id,
      JSON_ARRAYAGG(tag) as tags
    FROM routines_detail_tags 
    GROUP BY routine_detail_id
  ) rdt
  ON rdt.routine_detail_id = rd.id
LEFT JOIN
  (
    SELECT
      id,
      routine_detail_id,
      JSON_ARRAYAGG(detail_desc) as descriptions
    FROM routines_detail_desc
    GROUP BY routine_detail_id
  ) rdd
  ON rdd.routine_detail_id = rd.id
WHERE 
  rd.routine_id = $1;