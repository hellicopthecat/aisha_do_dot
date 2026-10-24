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
  r.user_id,
  u.email,
  u.name,
  COALESCE(
    json_agg(
      json_build_object(
        'can_access', ra.can_access,
        'editable', ra.editable,
        'access_user_id', ra.access_user_id
      )
    ) FILTER (WHERE ra.id IS NOT NULL),
    '[]'::json
  ) AS access_users
FROM
  routines r
LEFT JOIN
  users u
  ON u.id = r.user_id
LEFT JOIN
  routine_access ra
  ON r.id = ra.routine_id
WHERE
  r.id = $1
GROUP BY
  r.id;
