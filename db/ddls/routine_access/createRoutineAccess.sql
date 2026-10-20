INSERT INTO routine_access (
  id,
  created_at,
  updated_at,
  created_by,
  updated_by,
  can_access,
  editable,
  routine_id,
  access_user_id
) VALUES (
  $1,
  NOW(),
  NOW(),
  $2,
  $3,
  $4,
  $5,
  $6,
  $7
);