# drops all table
psql postgres -c "\i scripts/sql/cleanup.sql"

# delete all indexes
ELASTICSEARCH_URL="https://fold-money-test.es.europe-west3.gcp.cloud.es.io"
ECE_API_KEY="SUhha3BZc0JGZlRBUzN3LWFTbG46ZVBvTUhnYlRSLVNiMzRJQXNNVTZiUQ=="

curl -X DELETE -H "Content-Type: application/json" -H "Authorization: ApiKey $ECE_API_KEY" "$ELASTICSEARCH_URL/_all"
