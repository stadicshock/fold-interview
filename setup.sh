psql postgres -c "\i sql/schema.sql"
psql postgres -c "\i sql/seed.sql"


curl -X PUT "http://localhost:9200/users" -H "Content-Type: application/json" -d @esmapping/users_index_mapping.json
curl -X PUT "http://localhost:9200/hashtags" -H "Content-Type: application/json" -d @esmapping/hashtags_index_mapping.json
curl -X PUT "http://localhost:9200/projects" -H "Content-Type: application/json" -d @esmapping/projects_index_mapping.json
