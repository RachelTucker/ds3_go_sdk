package models

import "net/http"

type GetServiceResponse struct {
    Owner Owner
    Buckets []Bucket
}

func NewServiceResponse(response *http.Response) *GetServiceResponse {
    return &GetServiceResponse{}
}

type Owner struct {
    Id, DisplayName string
}

type Bucket struct {
    Name, CreationDate string
}

