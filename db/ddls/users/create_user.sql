INSERT INTO users (
  id,
  created_at,
  updated_at,
  created_by,
  updated_by,
  email,
  social,
  provider_id,
  name,
  refresh_token
)VALUES (
  $1,
  NOW(),
  NOW(),
  $2, -- created by
  $3, -- updated by
  $4, -- email
  $5, -- social
  $6, -- provider_id,
  $7, -- name
  $8 -- refresh token
);