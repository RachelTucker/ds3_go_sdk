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

type DeleteTapePartitionSpectraS3Request struct {
    tapePartition string
    queryParams *url.Values
}

func NewDeleteTapePartitionSpectraS3Request(tapePartition string) *DeleteTapePartitionSpectraS3Request {
    queryParams := &url.Values{}

    return &DeleteTapePartitionSpectraS3Request{
        tapePartition: tapePartition,
        queryParams: queryParams,
    }
}




func (DeleteTapePartitionSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteTapePartitionSpectraS3Request *DeleteTapePartitionSpectraS3Request) Path() string {
    return "/_rest_/tape_partition/" + deleteTapePartitionSpectraS3Request.tapePartition
}

func (deleteTapePartitionSpectraS3Request *DeleteTapePartitionSpectraS3Request) QueryParams() *url.Values {
    return deleteTapePartitionSpectraS3Request.queryParams
}

func (DeleteTapePartitionSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteTapePartitionSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeleteTapePartitionSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
