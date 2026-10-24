INSERT INTO routines_detail_desc (
  id,
  created_at,
  updated_at,
  created_by,
  updated_by,
  detail_desc,
  routine_detail_id
)
VALUES (
  $1,
  NOW(),
  NOW(),
  $2,
  $3,
  $4,
  $5
);