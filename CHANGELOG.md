## Next release

### [v0.1.1](https://github.com/pdffiller/terraform-modules-hcl2/compare/v0.1.1...HEAD)

*Add :*

*Changes :*

*Remove :*

*Improve :*

*Fix :*

## History

### [v0.1.0](https://github.com/pdffiller/terraform-modules-hcl2/compare/v0.1.0...v0.0.0)

- Forked from awslabs/ssosync v1.0.0-rc.8.
- Rename module.

*Add :*
  - variables "secret_bucket", "secret_bucket_region".
  - code for get stcret files from AWS S3 KMS encrypted bucket.
  - script for test run module localy as AWS Lambda function.

*Remove :*
  - code for syncing GSuite group to AWS SSO.
  - code for save credentials in AWS Secrets
  - file for deploy via AWS SAR.
  - file for deploy via SAM and goreleaser.
  - file for deploy via AWS CloudFormation Console.
  - deleted not used files .editorconfig, .envrc, .envsh, internal/aws/config.go, internal/aws/groups.go, internal/aws/groups_test.go, internal/config/secrets.go, template.yaml
