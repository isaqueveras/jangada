package newapp

const tmplConfigStorageYaml string = `development:
  service: Disk
  root: tmp/dev/storage

test:
  service: Disk
  root: tmp/test/storage

production:
  service: S3
  access_key_id: ""
  secret_access_key: ""
  bucket: "your-s3-bucket-?"
  region: "us-east-1"
`
