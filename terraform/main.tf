provider "aws" {
  region     = "us-east-1"
  access_key = "AKIAWNDDO4TBAKX3254G"
  secret_key = "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"
}

# Define an AWS RDS instance
resource "aws_db_instance" "foldpsdb" {
 allocated_storage    = 20
  storage_type         = "gp2"
  engine               = "postgres"
  engine_version       = "15.3" # Choose a supported engine version
  instance_class       = "db.t3.micro" # Choose a supported instance class
  username             = "folduser"
  password             = "foldpassword"
}

# Define an Elasticsearch domain
resource "aws_elasticsearch_domain" "foldes" {
  domain_name           = "my-elasticsearch-domain"
  elasticsearch_version = "7.10"
  cluster_config {
    instance_type        = "t2.small.elasticsearch"
  }
  ebs_options { 
    ebs_enabled = true
    volume_size = 10 # Specify the desired volume size in GB
  }
}
