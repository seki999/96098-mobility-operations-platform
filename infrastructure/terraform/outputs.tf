output "vpc_id" {
  value = aws_vpc.main.id
}

output "backend_ecr_repository_url" {
  value = aws_ecr_repository.backend.repository_url
}

output "apprunner_service_url" {
  value = aws_apprunner_service.backend.service_url
}

output "rds_endpoint" {
  value     = aws_db_instance.postgres.endpoint
  sensitive = true
}
