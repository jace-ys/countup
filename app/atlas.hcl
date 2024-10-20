env "local" {
  src = "file://schema/schema.sql"
  dev = "docker://postgres/15/dev"

  migration {
    dir = "file://schema/migrations"
    format = "goose"
  }
}