#!/bin/bash

export SSOSYNC_SECRET_BUCKET="the-secret-backet-name"
export SSOSYNC_SECRET_BUCKET_REGION="us-east-1"
export SSOSYNC_SCIM_ENDPOINT="https://scim.us-east-1.amazonaws.com/xxxxxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/scim/v2/"
export SSOSYNC_SCIM_ACCESS_TOKEN="ssosync/scim_access_token.txt"
export SSOSYNC_GOOGLE_CREDENTIALS="ssosync/google_credentials.json"
export SSOSYNC_GOOGLE_ADMIN="admin@example.team"
export SSOSYNC_DEBUG="false"
export SSOSYNC_LOG_LEVEL="debug"
export SSOSYNC_LOG_FORMAT="text"

./sso-gsuite-user-sync

#end of file
