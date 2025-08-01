# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

resource "aws_comprehend_document_classifier" "test" {
  name = var.rName

  data_access_role_arn = aws_iam_role.test.arn

  language_code = "en"
  input_data_config {
    s3_uri = "s3://${aws_s3_object.documents.bucket}/${aws_s3_object.documents.key}"
  }

  depends_on = [
    aws_iam_role_policy.test,
  ]
}

data "aws_partition" "current" {}

# testAccDocumentClassifierBasicRoleConfig

resource "aws_iam_role" "test" {
  name = var.rName

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "comprehend.${data.aws_partition.current.dns_suffix}"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "test" {
  role = aws_iam_role.test.name

  policy = data.aws_iam_policy_document.role.json
}

data "aws_iam_policy_document" "role" {
  statement {
    actions = [
      "s3:GetObject",
    ]

    resources = [
      "${aws_s3_bucket.test.arn}/*",
    ]
  }
  statement {
    actions = [
      "s3:ListBucket",
    ]

    resources = [
      aws_s3_bucket.test.arn,
    ]
  }
}

# testAccDocumentClassifierS3BucketConfig

resource "aws_s3_bucket" "test" {
  bucket = var.rName

  force_destroy = true
}

resource "aws_s3_bucket_public_access_block" "test" {
  bucket = aws_s3_bucket.test.bucket

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket_ownership_controls" "test" {
  bucket = aws_s3_bucket.test.bucket

  rule {
    object_ownership = "BucketOwnerEnforced"
  }
}

# testAccDocumentClassifierConfig_S3_documents

resource "aws_s3_object" "documents" {
  bucket = aws_s3_bucket.test.bucket
  key    = "documents.csv"
  source = "test-fixtures/document_classifier/documents.csv"
}

variable "rName" {
  description = "Name for resource"
  type        = string
  nullable    = false
}
