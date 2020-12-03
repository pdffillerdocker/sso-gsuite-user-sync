// Copyright (c) 2020, airSlate, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aws

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Get file from s3 to buffer
func S3getobject(bucket string, item string, region string) (response []byte, err error) {

    buff := &aws.WriteAtBuffer{}

    sess, _ := session.NewSession(&aws.Config{
        Region: aws.String(region)},
    )

    downloader := s3manager.NewDownloader(sess)

    numBytes, err := downloader.Download(buff,
        &s3.GetObjectInput{
            Bucket: aws.String(bucket),
            Key: aws.String(item),
        })

    if err != nil || numBytes == 0 {
        return nil, err
    }

    return buff.Bytes(), nil
}
