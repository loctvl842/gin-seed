variable "envfile" {
  type    = string
  default = ".env"
}

locals {
  envfile = {
    for line in split("\n", file(var.envfile)) :
    trimspace(split("=", line)[0]) => regex("=\\s*\"?([^\"]*)\"?", line)[0]
    if !startswith(trimspace(line), "#") && length(split("=", line)) > 1
  }
}

data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./adapter/model",
    "--dialect", "postgres",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url

  // Define the URL of the database which is managed
  // in this environment.
  url = "postgres://${local.envfile.PG_USER}:${local.envfile.PG_PASSWORD}@${local.envfile.PG_HOST}:${local.envfile.PG_PORT}/${local.envfile.PG_DB_NAME}?search_path=public&sslmode=disable"

  // Define the URL of the Dev Database for this environment
  // See: https://atlasgo.io/concepts/dev-database
  dev = "docker://postgres/latest/dev"  // Update Docker image for PostgreSQL

  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
