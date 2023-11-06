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
  engine_version       = "15.3" 
  instance_class       = "db.t3.micro" 
  identifier           = "folddb"
  db_name              = "folddb"
  username             = "folduser"
  password             = "foldpassword"
  publicly_accessible = true
}

## ES vis ES Cloud

terraform {
  required_version = ">= 1.0.0"

  required_providers {
    ec = {
      source  = "elastic/ec"
      version = "0.4.0"
    }
  }
}

provider "ec" {
    apikey = "essu_U21ad1ZIQlpjMEpDV25Jd1FtNHdSRkpXZWtNNk5FTnpVRjl4U25aUk5FOXZNMEpEWjFkallsRk5adz09AAAAABo+V8A="
}

resource "ec_deployment" "fold-deploy" {
  name                   = "fold money test"
  region                 = "gcp-europe-west3"
  version                = "8.1.3"
  deployment_template_id = "gcp-memory-optimized-v2"

  elasticsearch {}

  kibana {}
}