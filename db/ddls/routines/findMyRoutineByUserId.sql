SELECT
  r.id,
  r.created_at,
  r.updated_at,
  r.created_by,
  r.updated_by,
  r.routine_title,
  r.routine_desc,
  r.priority_one_depth,
  r.category,
  r.user_id
FROM
  routines r
WHERE
  r.user_id = $1;
