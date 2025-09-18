provider "aws" {
  region = "us-east-1" # change if needed
}

resource "aws_iam_role" "lambda_exec" {
  name = "golambda_exec_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_basic_execution" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_lambda_function" "golambda" {
  function_name = "math_operations"
  role          = aws_iam_role.lambda_exec.arn
  handler       = "main.handler"
  runtime       = "provided.al2023"
  architectures = ["arm64"]

  filename         = "math_operations.zip" # must exist before apply
  source_code_hash = filebase64sha256("math_operations.zip")
}

resource "aws_lambda_function_url" "golambda_url" {
  function_name      = aws_lambda_function.golambda.function_name
  authorization_type = "NONE"
}

output "golambda_function_url" {
  value = aws_lambda_function_url.golambda_url.function_url
}
