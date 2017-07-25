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

type TapeDrive struct {
    CleaningRequired bool `xml:"CleaningRequired"`
    ErrorMessage *string `xml:"ErrorMessage"`
    ForceTapeRemoval bool `xml:"ForceTapeRemoval"`
    Id string `xml:"Id"`
    LastCleaned *string `xml:"LastCleaned"`
    MfgSerialNumber *string `xml:"MfgSerialNumber"`
    PartitionId string `xml:"PartitionId"`
    Quiesced Quiesced `xml:"Quiesced"`
    SerialNumber *string `xml:"SerialNumber"`
    State TapeDriveState `xml:"State"`
    TapeId *string `xml:"TapeId"`
    Type TapeDriveType `xml:"Type"`
}