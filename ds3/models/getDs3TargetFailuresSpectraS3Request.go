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

type GetDs3TargetFailuresSpectraS3Request struct {
    ErrorMessage *string
    LastPage bool
    PageLength *int
    PageOffset *int
    PageStartMarker *string
    TargetFailureType TargetFailureType
    TargetId *string
}

func NewGetDs3TargetFailuresSpectraS3Request() *GetDs3TargetFailuresSpectraS3Request {
    return &GetDs3TargetFailuresSpectraS3Request{
    }
}

func (getDs3TargetFailuresSpectraS3Request *GetDs3TargetFailuresSpectraS3Request) WithErrorMessage(errorMessage string) *GetDs3TargetFailuresSpectraS3Request {
    getDs3TargetFailuresSpectraS3Request.ErrorMessage = &errorMessage
    return getDs3TargetFailuresSpectraS3Request
}

func (getDs3TargetFailuresSpectraS3Request *GetDs3TargetFailuresSpectraS3Request) WithLastPage() *GetDs3TargetFailuresSpectraS3Request {
    getDs3TargetFailuresSpectraS3Request.LastPage = true
    return getDs3TargetFailuresSpectraS3Request
}

func (getDs3TargetFailuresSpectraS3Request *GetDs3TargetFailuresSpectraS3Request) WithPageLength(pageLength int) *GetDs3TargetFailuresSpectraS3Request {
    getDs3TargetFailuresSpectraS3Request.PageLength = &pageLength
    return getDs3TargetFailuresSpectraS3Request
}

func (getDs3TargetFailuresSpectraS3Request *GetDs3TargetFailuresSpectraS3Request) WithPageOffset(pageOffset int) *GetDs3TargetFailuresSpectraS3Request {
    getDs3TargetFailuresSpectraS3Request.PageOffset = &pageOffset
    return getDs3TargetFailuresSpectraS3Request
}

func (getDs3TargetFailuresSpectraS3Request *GetDs3TargetFailuresSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetDs3TargetFailuresSpectraS3Request {
    getDs3TargetFailuresSpectraS3Request.PageStartMarker = &pageStartMarker
    return getDs3TargetFailuresSpectraS3Request
}

func (getDs3TargetFailuresSpectraS3Request *GetDs3TargetFailuresSpectraS3Request) WithTargetId(targetId string) *GetDs3TargetFailuresSpectraS3Request {
    getDs3TargetFailuresSpectraS3Request.TargetId = &targetId
    return getDs3TargetFailuresSpectraS3Request
}

func (getDs3TargetFailuresSpectraS3Request *GetDs3TargetFailuresSpectraS3Request) WithTargetFailureType(targetFailureType TargetFailureType) *GetDs3TargetFailuresSpectraS3Request {
    getDs3TargetFailuresSpectraS3Request.TargetFailureType = targetFailureType
    return getDs3TargetFailuresSpectraS3Request
}
