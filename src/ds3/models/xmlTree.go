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
    "encoding/xml"
)

type XmlNode struct {
    XMLName  xml.Name
    Attrs    []xml.Attr `xml:"-"`
    Content  []byte     `xml:",chardata"`
    Children []XmlNode     `xml:",any"`
}

// Uses the Go library to unmarshal XML into a simple node structure
func (xmlNode *XmlNode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    xmlNode.Attrs = start.Attr

    type n XmlNode

    return d.DecodeElement((*n)(xmlNode), &start)
}