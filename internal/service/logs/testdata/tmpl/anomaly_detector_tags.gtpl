resource "aws_cloudwatch_log_group" "test" {
  count = 2
  name  = "${var.rName}-${count.index}"
}

resource "aws_cloudwatch_log_anomaly_detector" "test" {
  detector_name           = var.rName
  log_group_arn_list      = [aws_cloudwatch_log_group.test[0].arn]
  anomaly_visibility_time = 7
  evaluation_frequency    = "TEN_MIN"
  enabled                 = "false"

{{- template "tags" . }}
}
