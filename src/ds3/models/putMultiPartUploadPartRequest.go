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

type PutMultiPartUploadPartRequest struct {
    BucketName string
    ObjectName string
    Content ReaderWithSizeDecorator
    PartNumber int
    UploadId string
}

func NewPutMultiPartUploadPartRequest(bucketName string, objectName string, content ReaderWithSizeDecorator, partNumber int, uploadId string) *PutMultiPartUploadPartRequest {
    return &PutMultiPartUploadPartRequest{
        BucketName: bucketName,
        ObjectName: objectName,
        PartNumber: partNumber,
        UploadId: uploadId,
        Content: content,
    }
}

