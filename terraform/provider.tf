provider "aws" {
    region = "eu-west-1"
}

provider "aws" {
    alias = "paris"
    region = "eu-west-3"
}

provider "aws" {
    alias = "us"
    region = "us-east-1"
}

provider "aws" {
    alias = "canada"
    region = "ca-central-1"
}