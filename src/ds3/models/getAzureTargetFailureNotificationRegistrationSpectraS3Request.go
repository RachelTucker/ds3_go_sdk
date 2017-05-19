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

type GetAzureTargetFailureNotificationRegistrationSpectraS3Request struct {
    azureTargetFailureNotificationRegistration string
    queryParams *url.Values
}

func NewGetAzureTargetFailureNotificationRegistrationSpectraS3Request(azureTargetFailureNotificationRegistration string) *GetAzureTargetFailureNotificationRegistrationSpectraS3Request {
    queryParams := &url.Values{}

    return &GetAzureTargetFailureNotificationRegistrationSpectraS3Request{
        azureTargetFailureNotificationRegistration: azureTargetFailureNotificationRegistration,
        queryParams: queryParams,
    }
}




func (GetAzureTargetFailureNotificationRegistrationSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getAzureTargetFailureNotificationRegistrationSpectraS3Request *GetAzureTargetFailureNotificationRegistrationSpectraS3Request) Path() string {
    return "/_rest_/azure_target_failure_notification_registration/" + getAzureTargetFailureNotificationRegistrationSpectraS3Request.azureTargetFailureNotificationRegistration
}

func (getAzureTargetFailureNotificationRegistrationSpectraS3Request *GetAzureTargetFailureNotificationRegistrationSpectraS3Request) QueryParams() *url.Values {
    return getAzureTargetFailureNotificationRegistrationSpectraS3Request.queryParams
}

func (GetAzureTargetFailureNotificationRegistrationSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetAzureTargetFailureNotificationRegistrationSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetAzureTargetFailureNotificationRegistrationSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
