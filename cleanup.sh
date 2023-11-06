# drops all table
psql postgres -c "\i sql/cleanup.sql"

# delete all indexes
curl -X DELETE 'http://localhost:9200/_all'
