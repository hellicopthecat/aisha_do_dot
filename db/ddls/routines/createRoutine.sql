INSERT INTO routines (
  id,
  created_at,
  updated_at,
  created_by,
  updated_by,
  routine_title,
  routine_desc,
  priority_one_depth,
  category,
  user_id
) VALUES (
  $1,
  NOW(),
  NOW(),
  $2, -- created_by
  $3, -- updated_by
  $4, -- routine_title
  $5, -- routine_desc
  $6, -- priority_one_depth
  $7, -- category
  $8 -- user_id
);