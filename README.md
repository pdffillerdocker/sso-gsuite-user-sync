# sso-gsuite-user-sync
Utilite for Synchronization users from GSuite to AWS SSO

#### Maintainer Dmitry Teikovtsev teikovtsev.dmitry@pdffiller.team

Forked from awslabs/ssosync v1.0.0-rc.8 original project in https://github.com/awslabs/ssosync

Changes in https://github.com/pdffillerdocker/sso-gsuite-user-sync/blob/main/CHANGELOG.md

As per the [AWS SSO](https://aws.amazon.com/single-sign-on/) Homepage:

> AWS Single Sign-On (SSO) makes it easy to centrally manage access
> to multiple AWS accounts and business applications and provide users
> with single sign-on access to all their assigned accounts and applications
> from one place.

Key part further down:

> With AWS SSO, you can create and manage user identities in AWS SSO’s
>identity store, or easily connect to your existing identity source including
> Microsoft Active Directory and **Azure Active Directory (Azure AD)**.

AWS SSO can use other Identity Providers as well... such as Google Apps for Domains. Although AWS SSO
supports a subset of the SCIM protocol for populating users, it currently only has support for Azure AD.

This project provides a CLI tool to pull users and groups from Google and push them into AWS SSO.
`sso-gsuite-user-sync` deals with removing users as well. The heavily commented code provides you with the detail of
what it is going to do.

### References

 * [SCIM Protocol RFC](https://tools.ietf.org/html/rfc7644)
 * [AWS SSO - Connect to Your External Identity Provider](https://docs.aws.amazon.com/singlesignon/latest/userguide/manage-your-identity-source-idp.html)
 * [AWS SSO - Automatic Provisioning](https://docs.aws.amazon.com/singlesignon/latest/userguide/provision-automatically.html)

## Installation

 - install golang version >= go1.15.5 or update if need see https://github.com/udhos/update-golang
 - run in CLI:
```bash
go build github.com/pdffillerdocker/sso-gsuite-user-sync
sso-gsuite-user-sync --help
```

## Configuration

You need a few items of configuration. One side from AWS, and the other
from Google Cloud to allow for API access to each. You should have configured
Google as your Identity Provider for AWS SSO already.

You will need the files produced by these steps for AWS Lambda deployment as well
as locally running the sso-gsuite-user-sync tool.

### Google

First, you have to setup your API. In the project you want to use go to the [Console](https://console.developers.google.com/apis) and select *API & Services* > *Enable APIs and Services*. Search for *Admin SDK* and *Enable* the API.

You have to perform this [tutorial](https://developers.google.com/admin-sdk/directory/v1/guides/delegation) to create a service account that you use to sync your users. Save the JSON file you create during the process and rename it to `credentials.json`.

> you can also use the `--google-credentials` parameter to explicitly specify the file on S3 with the service credentials. Please, keep this file safe.

In the domain-wide delegation for the Admin API, you have to specify the following scopes for the user.

`https://www.googleapis.com/auth/admin.directory.user.readonly`

Back in the Console go to the Dashboard for the API & Services and select "Enable API and Services".
In the Search box type `Admin` and select the `Admin SDK` option. Click the `Enable` button.

You will have to specify the email address of an admin via `--google-admin` to assume this users role in the Directory.

Specific these as environment variables.
```bash
SSOSYNC_GOOGLE_CREDENTIALS=<PATH_TO_S3_FILE_WITH_GOOGLE_CREDENTIALS>
SSOSYNC_GOOGLE_ADMIN=<GOOGLE_ADMIN_EMAIL>

```

### AWS

Go to the AWS Single Sign-On console in the region you have set up AWS SSO and select
Settings. Click `Enable automatic provisioning`.

A pop up will appear with URL and the Access Token. The Access Token will only appear
at this stage. Specific these as environment variables.

```bash
SSOSYNC_SCIM_ACCESS_TOKEN=<PATH_TO_S3_FILE_WITH_YOUR_TOKEN>
SSOSYNC_SCIM_ENDPOINT=<YOUR_ENDPOINT>
```

### Dependesies:
 - create aws s3 secret bucket.
 - create folder like 'ssosync' on s3 secret bucket.
 - prepare file credentials.json with Google service account credentials.
 - create file like scim_access_token.txt contains scim access token ( !!! no set CR/LF on end of file !!! ).
 - copy credentials.json and scim_access_token.txt to folder on s3.
 - for run local: modify file test-environment.sh and run. Please, keep this file safe.

## Local Usage

```bash
Usage:
  sso-gsuite-user-sync [flags]

Flags:
  -t, --access-token string           The Path to file with SCIM Access Token
  -d, --debug                         Enable verbose / debug logging
  -e, --endpoint string               SCIM Endpoint
  -u, --google-admin string           Google Admin Email
  -c, --google-credentials string     The path to file with credentials for Google
  -h, --help                          help for sso-gsuite-user-sync
      --ignore-users strings          ignores these users
      --log-format string             log format (default "text")
      --log-level string              log level (default "warn")
  -b, --secret_bucket string          Secret Bucket name
  -r, --secret_bucket_region string   Secret Bucket region
  -v, --version                       version for sso-gsuite-user-sync
```

## License

[Apache-2.0](/LICENSE)
