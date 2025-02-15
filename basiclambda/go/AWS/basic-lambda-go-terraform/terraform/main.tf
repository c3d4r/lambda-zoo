provider "aws" {
  region = "eu-west-1"
}

locals {
  archive_path  = "../lambda/bootstrap.zip"
}

// build the binary for the lambda function in a specified path
resource "null_resource" "function_binary" {
  provisioner "local-exec" {
    working_dir = "../lambda"
    command = "GOOS=linux GOARCH=arm64 CGO_ENABLED=0 GOFLAGS=-trimpath go build -mod=readonly -ldflags='-s -w' -o bootstrap main.go"
  }
}

// zip the binary, as we can use only zip files to AWS lambda
data "archive_file" "function_archive" {
  depends_on = [null_resource.function_binary]

  type        = "zip"
  source_file = "../lambda/bootstrap"
  output_path = local.archive_path
}

# allow lambda service to assume (use) the role with such policy
data "aws_iam_policy_document" "assume_lambda_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

# create lambda role, that lambda function can assume (use)
resource "aws_iam_role" "basic_lambda_go_terraform_role" {
  name               = "BasicLambdaGoTerraformRole"
  description        = "Role for lambda to assume lambda"
  assume_role_policy = data.aws_iam_policy_document.assume_lambda_role.json
}

data "aws_iam_policy_document" "allow_lambda_logging" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]

    resources = [
      "arn:aws:logs:*:*:*",
    ]
  }
}

resource "aws_iam_policy" "function_logging_policy" {
  name        = "BasicLambdaGoTerraformAllowLambdaLoggingPolicy"
  description = "Policy for lambda cloudwatch logging"
  policy      = data.aws_iam_policy_document.allow_lambda_logging.json
}

resource "aws_iam_role_policy_attachment" "lambda_logging_policy_attachment" {
  role       = aws_iam_role.basic_lambda_go_terraform_role.id
  policy_arn = aws_iam_policy.function_logging_policy.arn
}

// create the lambda function from zip file
resource "aws_lambda_function" "function" {
  function_name = "basic-lambda-go-terraform"
  description   = "basic-lambda-go-terraform"
  role          = aws_iam_role.basic_lambda_go_terraform_role.arn
  handler       = "bootstrap"
  memory_size   = 128
  architectures  = ["arm64"]


  filename         = local.archive_path
  source_code_hash = data.archive_file.function_archive.output_base64sha256

  runtime = "provided.al2"
}

// create log group in cloudwatch to gather logs of our lambda function
resource "aws_cloudwatch_log_group" "log_group" {
  name              = "/aws/lambda/${aws_lambda_function.function.function_name}"
  retention_in_days = 7
}

resource "aws_lambda_function_url" "function_url" {
  function_name      = aws_lambda_function.function.function_name
  authorization_type = "NONE" # Change to "AWS_IAM" for restricted access
}

output "function_url" {
  value = aws_lambda_function_url.function_url.function_url
}