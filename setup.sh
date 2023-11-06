# Setup DB
psql postgres -c "\i scripts/sql/schema.sql"
psql postgres -c "\i scripts/sql/seed.sql"

# Setup ES
ELASTICSEARCH_URL="https://fold-money-test.es.europe-west3.gcp.cloud.es.io"
ECE_API_KEY="SUhha3BZc0JGZlRBUzN3LWFTbG46ZVBvTUhnYlRSLVNiMzRJQXNNVTZiUQ=="
# todo get URL and api key via terraform 
curl -X PUT "$ELASTICSEARCH_URL/projects" -H "Content-Type: application/json" -H "Authorization: ApiKey $ECE_API_KEY" -d @scripts/esmapping/projects_index_mapping.json