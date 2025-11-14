# 定义变量（开发/生产共用）
variable "db_user" {
  type    = string
  default = "dev"
}

variable "db_password" {
  type    = string
  default = "dev"
}

variable "db_host" {
  type    = string
  default = "localhost"
}

variable "db_port" {
  type    = string
  default = "3306"
}

variable "db_name" {
  type    = string
  default = "cds_infra"
}

locals {
  mysql_url = "mysql://${var.db_user}:${var.db_password}@${var.db_host}:${var.db_port}/${var.db_name}?parseTime=True&loc=Asia%2FShanghai"
}

diff {
#   skip {
#     // By default, none of the changes are skipped.
#     drop_schema = true
#     drop_table  = true
#   }
  concurrent_index {
    create = true
    drop   = true
  }
}

env {
  name = atlas.env
  url = "${local.mysql_url}"

  format {
    migrate {
      # apply = format(
      #   "{{ json . | json_merge %q }}",
      #   jsonencode({
      #     EnvName : atlas.env
      #   })
      # )
      # diff = format(
      # )
    }
  }
}

# 开发环境（允许自动修复小变更）
env "dev" {
  // Declare where the schema definition resides.
  src  = "ent://ent/schema"
  // Define the URL of the database which is managed in this environment.
  # url  = "mysql://dev:dev@localhost:3306/ccl_base?parseTime=True&loc=Asia%2FShanghai"
  url = "${local.mysql_url}"
  // Define the URL of the Dev Database for this environment
  # dev  = "mysql://dev:dev@localhost:3306/ccl_base?parseTime=True&loc=Asia%2FShanghai"
  dev = "docker://mysql/8/dev"

  migration {
    dir = "file://migrations"
      # exclude = []
  }
}

env "test" {
  src = "file://sql/migrations"
  url = "${local.mysql_url}"
}

# atlas migrate diff create_blog_posts --dir "file://migrations" --to "file://schema.hcl"  --dev-url "mysql://dev:dev@localhost:3306/cds_infra"
# 生产环境（严格校验，禁止删表）
env "prod" {
  src = "file://sql/migrations"
  url = "${local.mysql_url}"
  lint = {
    # deny = ["DROP TABLE", "DROP COLUMN"]  # 禁止危险操作
  }
}
