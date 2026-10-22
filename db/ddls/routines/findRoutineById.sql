SELECT
  id,
  created_at,
  updated_at,
  created_by,
  updated_by,
  routine_title,
  routine_desc,
  priorityOneDepth,
  category,
  user_id
FROM
  routines
WHERE
  id = $1;