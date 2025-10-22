package newapp

const tmplConfigDatabaseYAML = `development:
  - nick: default
    name: {{ .AppName }}
    username: postgres
    password: postgres
    hostname: localhost
    port: "5432"
    max_conn: 20
    max_idle: 5
    read_only: false
    main: true
    transaction_timeout: 30
    ssl_mode: disable
    ssl_client:
      path_cert: /etc/ssl/certs/client.crt
      path_key: /etc/ssl/private/client.key
      path_ca: /etc/ssl/certs/ca.crt

test:
  - nick: default
    name: {{ .AppName }}
    username: postgres
    password: postgres
    hostname: localhost
    port: "5432"
    max_conn: 20
    max_idle: 5
    read_only: false
    main: true
    transaction_timeout: 30
    ssl_mode: disable
    ssl_client:
      path_cert: /etc/ssl/certs/client.crt
      path_key: /etc/ssl/private/client.key
      path_ca: /etc/ssl/certs/ca.crt

production:
  - nick: default
    name: {{ .AppName }}
    username: postgres
    password: postgres
    hostname: localhost
    port: "5432"
    max_conn: 20
    max_idle: 5
    read_only: false
    main: true
    transaction_timeout: 30
    ssl_mode: enable
    ssl_client:
      path_cert: /etc/ssl/certs/client.crt
      path_key: /etc/ssl/private/client.key
      path_ca: /etc/ssl/certs/ca.crt
`
