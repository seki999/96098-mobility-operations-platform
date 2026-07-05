variable "project_name" {
  description = "Portfolio project name"
  type        = string
  default     = "96098-mobility-operations-platform"
}

variable "aws_region" {
  description = "AWS region for the reference architecture"
  type        = string
  default     = "ap-northeast-1"
}

variable "vpc_cidr" {
  description = "VPC CIDR block"
  type        = string
  default     = "10.98.0.0/16"
}

variable "availability_zones" {
  description = "Two availability zones for subnet design"
  type        = list(string)
  default     = ["ap-northeast-1a", "ap-northeast-1c"]
}

variable "allowed_http_cidrs" {
  description = "CIDR blocks allowed to reach HTTP entry point"
  type        = list(string)
  default     = ["0.0.0.0/0"]
}

variable "allowed_origin" {
  description = "CORS origin for frontend"
  type        = string
  default     = "https://portfolio.local"
}

variable "db_name" {
  description = "RDS database name"
  type        = string
  default     = "mobility_ops"
}

variable "db_username" {
  description = "RDS username. Use tfvars or secret manager in real deployment."
  type        = string
  default     = "mobility_admin"
}

variable "db_password" {
  description = "RDS password. Do not commit real values."
  type        = string
  sensitive   = true
}

variable "db_instance_class" {
  description = "RDS instance class"
  type        = string
  default     = "db.t4g.micro"
}
