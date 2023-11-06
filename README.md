# fold-interview
Take home task from fold.money
REST API service that will query Elasticsearch.

# Prerequisites
Install postgres and logstash on local

Steps to start REST api service

1) Setup infra (Optional) // TODO : Pass config/secrets to env such as hostname,username and password
    - run `terraform init`
    - run `terraform apply` # This will setup postgres in AWS and ES in ES cloud

2) run `sh setup.sh` 
    - This will create table schemas, triggers and functions which are required.
    - And it will also create ES index in ES cloud

3) Setup and run logstash
    - Download logstash
    - Download jdbc driver library for postgre
    - Copy scripts/logstash/jdbc.conf to logstash directory
    - Update jdbc_driver_library field in jdbc.conf
    - run `bin/logstash -f jdbc.conf` # This will start logstash service, which will sync data from postgres to ES cloud

4) run `go run cmd/main.go`
    - This will start golang rest api service

Sample curl for all apis

1) Search for projects created by a particular user
curl 'localhost:8080/projects/created-by/1' 
Where 1 is the user ID


2) Search for projects that use specific hashtags**
curl 'localhost:8080/projects/search-with-hashtags' \
--header 'Content-Type: application/json' \
--data '{
    "hashtags":["programming","design"]
}'


3) Full-text fuzzy search for projects
curl 'localhost:8080/projects/search?q=project'


Pending tasks:
1) AWS postgres connectivity issue (alternatively using local postgres)
2) Pass secrets to env from terraform out and use them in app/scripts
3) Move logstash setup to aws via terraform
4) Handle logstash for delete operation in postgres
5) Code documentation and unit test cases
