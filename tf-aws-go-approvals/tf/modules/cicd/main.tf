# create ci/cd user with access keys (for build system)
resource "aws_iam_user" "cicd" {
  name = "${var.app}-cicd"
}

resource "aws_iam_access_key" "cicd_keys" {
  user = aws_iam_user.cicd.name
}

# grant required permissions to deploy
data "aws_iam_policy_document" "cicd_policy" {
  statement {
    sid = "assumerole"

    actions = [
      "lambda:UpdateFunctionCode",
    ]

    resources = [
      "arn:aws:lambda:${var.region}:${var.account_id}:function:${var.app}-*"
    ]
  }
}

resource "aws_iam_user_policy" "cicd_user_policy" {
  name    = "${var.app}-cicd"
  user    = aws_iam_user.cicd.name
  policy  = data.aws_iam_policy_document.cicd_policy.json
}
