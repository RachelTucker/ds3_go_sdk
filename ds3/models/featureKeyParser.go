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

import "log"

func (featureKey *FeatureKey) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CurrentValue":
            featureKey.CurrentValue = parseNullableInt64(child.Content, aggErr)
        case "ErrorMessage":
            featureKey.ErrorMessage = parseNullableString(child.Content)
        case "ExpirationDate":
            featureKey.ExpirationDate = parseNullableString(child.Content)
        case "Id":
            featureKey.Id = parseString(child.Content)
        case "Key":
            parseEnum(child.Content, &featureKey.Key, aggErr)
        case "LimitValue":
            featureKey.LimitValue = parseNullableInt64(child.Content, aggErr)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing FeatureKey.", child.XMLName.Local)
        }
    }
}