package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type PutObjectRequest struct {
    bucketName string
    objectName string
    content networking.ReaderWithSizeDecorator
    checksum networking.Checksum
    queryParams *url.Values
}

func NewPutObjectRequest(bucketName string, objectName string, content networking.ReaderWithSizeDecorator) *PutObjectRequest {
    return &PutObjectRequest{
        bucketName: bucketName,
        objectName:objectName,
        content:content,
        checksum: networking.NewNoneChecksum(), //Default checksum type of None
        queryParams: &url.Values{},
    }
}

func (putObjectRequest *PutObjectRequest) WithChecksum(contentHash string, checksumType networking.ChecksumType) *PutObjectRequest {
    putObjectRequest.checksum.ContentHash = contentHash
    putObjectRequest.checksum.Type = checksumType
    return putObjectRequest
}

func (PutObjectRequest) Verb() networking.HttpVerb {
    return networking.PUT
}

func (putObjectRequest *PutObjectRequest) Path() string {
    return "/" + putObjectRequest.bucketName + "/" + putObjectRequest.objectName
}

func (putObjectRequest *PutObjectRequest) QueryParams() *url.Values {
    return putObjectRequest.queryParams
}

func (PutObjectRequest) Header() *http.Header {
    return &http.Header{}
}

func (putObjectRequest *PutObjectRequest) GetContentStream() networking.ReaderWithSizeDecorator {
    return putObjectRequest.content
}

func (putObjectRequest *PutObjectRequest) GetChecksum() networking.Checksum {
    return putObjectRequest.checksum
}