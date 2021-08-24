

check-connections:
	@docker exec pgdb psql -U postgres -c 'SELECT pid, usename, state, query FROM pg_stat_activity WHERE state IS NOT NULL;'