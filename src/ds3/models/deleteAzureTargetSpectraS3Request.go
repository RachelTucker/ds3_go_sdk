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

type DeleteAzureTargetSpectraS3Request struct {
    azureTarget string
    queryParams *url.Values
}

func NewDeleteAzureTargetSpectraS3Request(azureTarget string) *DeleteAzureTargetSpectraS3Request {
    queryParams := &url.Values{}

    return &DeleteAzureTargetSpectraS3Request{
        azureTarget: azureTarget,
        queryParams: queryParams,
    }
}




func (DeleteAzureTargetSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteAzureTargetSpectraS3Request *DeleteAzureTargetSpectraS3Request) Path() string {
    return "/_rest_/azure_target/" + deleteAzureTargetSpectraS3Request.azureTarget
}

func (deleteAzureTargetSpectraS3Request *DeleteAzureTargetSpectraS3Request) QueryParams() *url.Values {
    return deleteAzureTargetSpectraS3Request.queryParams
}

func (DeleteAzureTargetSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteAzureTargetSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeleteAzureTargetSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
