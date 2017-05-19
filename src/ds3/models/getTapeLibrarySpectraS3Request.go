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

type GetTapeLibrarySpectraS3Request struct {
    tapeLibraryId string
    queryParams *url.Values
}

func NewGetTapeLibrarySpectraS3Request(tapeLibraryId string) *GetTapeLibrarySpectraS3Request {
    queryParams := &url.Values{}

    return &GetTapeLibrarySpectraS3Request{
        tapeLibraryId: tapeLibraryId,
        queryParams: queryParams,
    }
}




func (GetTapeLibrarySpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getTapeLibrarySpectraS3Request *GetTapeLibrarySpectraS3Request) Path() string {
    return "/_rest_/tape_library/" + getTapeLibrarySpectraS3Request.tapeLibraryId
}

func (getTapeLibrarySpectraS3Request *GetTapeLibrarySpectraS3Request) QueryParams() *url.Values {
    return getTapeLibrarySpectraS3Request.queryParams
}

func (GetTapeLibrarySpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetTapeLibrarySpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetTapeLibrarySpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}
