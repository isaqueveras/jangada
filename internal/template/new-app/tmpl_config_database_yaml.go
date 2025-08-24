package newapp

const tmplConfigDatabaseYAML = `development:
  driver: sqlite3
  url: db/jangada_development.db?cache=shared&mode=memory

test:
  driver: sqlite3
  url: db/jangada_test.db?cache=shared&mode=memory

production:
  driver: ${DATABASE_DRIVER}
  url: ${DATABASE_URL}

staging:
  driver: ${DATABASE_DRIVER}
  url: ${DATABASE_URL}
`
