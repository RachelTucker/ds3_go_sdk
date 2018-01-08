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
    "net/http"
)

type DeleteS3TargetFailureNotificationRegistrationSpectraS3Response struct {
    
    Headers *http.Header
}



func NewDeleteS3TargetFailureNotificationRegistrationSpectraS3Response(webResponse WebResponse) (*DeleteS3TargetFailureNotificationRegistrationSpectraS3Response, error) {
    expectedStatusCodes := []int { 204 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &DeleteS3TargetFailureNotificationRegistrationSpectraS3Response{Headers: webResponse.Header()}, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}