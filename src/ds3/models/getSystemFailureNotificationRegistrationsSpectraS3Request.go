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

type GetSystemFailureNotificationRegistrationsSpectraS3Request struct {
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    UserId *string
}

func NewGetSystemFailureNotificationRegistrationsSpectraS3Request() *GetSystemFailureNotificationRegistrationsSpectraS3Request {
    return &GetSystemFailureNotificationRegistrationsSpectraS3Request{
    }
}

func (getSystemFailureNotificationRegistrationsSpectraS3Request *GetSystemFailureNotificationRegistrationsSpectraS3Request) WithLastPage() *GetSystemFailureNotificationRegistrationsSpectraS3Request {
    getSystemFailureNotificationRegistrationsSpectraS3Request.LastPage = true
    return getSystemFailureNotificationRegistrationsSpectraS3Request
}

func (getSystemFailureNotificationRegistrationsSpectraS3Request *GetSystemFailureNotificationRegistrationsSpectraS3Request) WithPageLength(pageLength int) *GetSystemFailureNotificationRegistrationsSpectraS3Request {
    getSystemFailureNotificationRegistrationsSpectraS3Request.PageLength = &pageLength
    return getSystemFailureNotificationRegistrationsSpectraS3Request
}

func (getSystemFailureNotificationRegistrationsSpectraS3Request *GetSystemFailureNotificationRegistrationsSpectraS3Request) WithPageOffset(pageOffset int) *GetSystemFailureNotificationRegistrationsSpectraS3Request {
    getSystemFailureNotificationRegistrationsSpectraS3Request.PageOffset = &pageOffset
    return getSystemFailureNotificationRegistrationsSpectraS3Request
}

func (getSystemFailureNotificationRegistrationsSpectraS3Request *GetSystemFailureNotificationRegistrationsSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetSystemFailureNotificationRegistrationsSpectraS3Request {
    getSystemFailureNotificationRegistrationsSpectraS3Request.PageStartMarker = &pageStartMarker
    return getSystemFailureNotificationRegistrationsSpectraS3Request
}

func (getSystemFailureNotificationRegistrationsSpectraS3Request *GetSystemFailureNotificationRegistrationsSpectraS3Request) WithUserId(userId string) *GetSystemFailureNotificationRegistrationsSpectraS3Request {
    getSystemFailureNotificationRegistrationsSpectraS3Request.UserId = &userId
    return getSystemFailureNotificationRegistrationsSpectraS3Request
}

