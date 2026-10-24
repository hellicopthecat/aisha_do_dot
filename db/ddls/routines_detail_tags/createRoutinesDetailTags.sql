INSERT INTO routines_detail_tags (
  id,
  created_at,
  updated_at,
  created_by,
  updated_by,
  tag,
  routine_detail_id
)
VALUES(
  $1,
  NOW(),
  NOW(),
  $2,
  $3,
  $4,
  $5
);