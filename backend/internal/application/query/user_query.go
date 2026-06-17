package query

const (
	UserCreateNew string = `WITH user_count AS (SELECT COUNT(*) AS total FROM users)
	INSERT INTO users (username, password, role, status_id) VALUES
	('%s', '%s', (CASE WHEN (SELECT total FROM user_count) = 0 THEN '%s' ELSE '%s' END)::user_role,
	(SELECT _id FROM status WHERE slug = CASE WHEN (SELECT total FROM user_count) = 0 THEN '%s' ELSE '%s' END LIMIT 1))`
)
