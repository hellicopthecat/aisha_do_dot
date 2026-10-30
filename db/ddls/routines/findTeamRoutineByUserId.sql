SELECT
  r.*
FROM
  routines r
WHERE EXISTS (
  SELECT 1
  FROM routines_access ra
  WHERE ra.routine_id = r.id
    AND ra.access_user_id = $1
    AND ra.can_access = 'ACCESS'
);
