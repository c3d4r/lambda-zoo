terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

resource "aws_cloudwatch_log_group" "basic_lambda_python_terraform_log_group" {
  name              = "/aws/lambda/${aws_lambda_function.lambda_function.function_name}"
  retention_in_days = 14
}

data "aws_iam_policy_document" "lambda_logging" {
  statement {
    effect = "Allow"

    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]

    resources = ["arn:aws:logs:*:*:*"]
  }
}

resource "aws_iam_policy" "lambda_logging" {
  name        = "basic_lambda_python_terraform_logging_policy"
  path        = "/"
  description = "IAM policy for logging from a lambda"
  policy      = data.aws_iam_policy_document.lambda_logging.json
}

resource "aws_iam_role_policy_attachment" "lambda_logs" {
  role       = aws_iam_role.iam_for_lambda.name
  policy_arn = aws_iam_policy.lambda_logging.arn
}

data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "iam_for_lambda" {
  name               = "basic-lambda-python-terraform-role"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

data "archive_file" "function_archive" {
  type        = "zip"
  source_file = "../app/lambda_handler.py"
  output_path = "../app/handler.zip"
}

resource "aws_lambda_function" "lambda_function" {
  filename      = "../app/handler.zip"
  function_name = "basic-lambda-python-terraform"
  role          = aws_iam_role.iam_for_lambda.arn
  handler       = "lambda_handler.lambda_handler"
  source_code_hash = data.archive_file.function_archive.output_base64sha256
  runtime = "python3.13"
}

resource "aws_lambda_function_url" "lambda_function_url" {
  function_name      = aws_lambda_function.lambda_function.function_name
  authorization_type = "NONE"
}

output "lambda_function_url" {
  value = aws_lambda_function_url.lambda_function_url.function_url
  description = "The URL endpoint for the Lambda function"
}

