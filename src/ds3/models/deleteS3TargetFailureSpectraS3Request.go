// Copyright 2014-2017 Spectra Logic Corporation. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License"). You may not use
// this file except in compliance with the License. A copy of the License is located at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file.
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

// This code is auto-generated, do not modify

package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type DeleteS3TargetFailureSpectraS3Request struct {
    s3TargetFailure string
    queryParams *url.Values
}

func NewDeleteS3TargetFailureSpectraS3Request(s3TargetFailure string) *DeleteS3TargetFailureSpectraS3Request {
    queryParams := &url.Values{}

    return &DeleteS3TargetFailureSpectraS3Request{
        s3TargetFailure: s3TargetFailure,
        queryParams: queryParams,
    }
}




func (DeleteS3TargetFailureSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteS3TargetFailureSpectraS3Request *DeleteS3TargetFailureSpectraS3Request) Path() string {
    return "/_rest_/s3_target_failure/" + deleteS3TargetFailureSpectraS3Request.s3TargetFailure
}

func (deleteS3TargetFailureSpectraS3Request *DeleteS3TargetFailureSpectraS3Request) QueryParams() *url.Values {
    return deleteS3TargetFailureSpectraS3Request.queryParams
}

func (DeleteS3TargetFailureSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteS3TargetFailureSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeleteS3TargetFailureSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}