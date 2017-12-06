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

type ImportPoolSpectraS3Response struct {
    Pool Pool
    Headers *http.Header
}

func (importPoolSpectraS3Response *ImportPoolSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &importPoolSpectraS3Response.Pool)
}

func NewImportPoolSpectraS3Response(webResponse WebResponse) (*ImportPoolSpectraS3Response, error) {
    expectedStatusCodes := []int { 200 }

    switch code := webResponse.StatusCode(); code {
    case 200:
        var body ImportPoolSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}
