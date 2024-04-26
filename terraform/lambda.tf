module "myawesomelambda" {
#   source = "../../../../modules/lambda/terraform-aws-lambda"
  source = "github.com/ProovGroup/infrastructure/iac/modules/lambda/terraform-aws-lambda"
  function_name = "myawesomelambda"
    handler                        = "main"
  runtime                        = "provided.al2"
    source_path = "function.zip"
      attach_sqs_policy = false
      attach_ssm_policy = false
  timeout = 30
  attach_s3_policy     = false
 tags = {
    Project = "test"
    Terraform = "true"
  }
  providers = {
    aws = myawesomeprovider
  }
}