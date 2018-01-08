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

type PutDs3DataReplicationRuleSpectraS3Response struct {
    Ds3DataReplicationRule Ds3DataReplicationRule
    Headers *http.Header
}

func (putDs3DataReplicationRuleSpectraS3Response *PutDs3DataReplicationRuleSpectraS3Response) parse(webResponse WebResponse) error {
        return parseResponsePayload(webResponse, &putDs3DataReplicationRuleSpectraS3Response.Ds3DataReplicationRule)
}

func NewPutDs3DataReplicationRuleSpectraS3Response(webResponse WebResponse) (*PutDs3DataReplicationRuleSpectraS3Response, error) {
    expectedStatusCodes := []int { 201 }

    switch code := webResponse.StatusCode(); code {
    case 201:
        var body PutDs3DataReplicationRuleSpectraS3Response
        if err := body.parse(webResponse); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}