terraform {
    required_providers {
        aws = {
            source = "hashicorp/aws"
            version = "~> 5.0"
        }
    }
    backend "s3" {
      bucket  = "weproov-testweproov-shared-tfstates"
    #   key     = "eu-west-1/lambdas.tfstate"
      region  = "eu-west-3"
      # profile = "terraform" # you have to give the profile name here. not the variable("${var.AWS_PROFILE}")
      # Replace this with your DynamoDB table name!
      dynamodb_table = "weproov-testweproov-locks"
      encrypt        = true
  }
    # required_versions = ">= 1.2.0"
}