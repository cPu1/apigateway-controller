// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package integration

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}
	customPreCompare(a, b)

	if len(a.ko.Spec.CacheKeyParameters) != len(b.ko.Spec.CacheKeyParameters) {
		delta.Add("Spec.CacheKeyParameters", a.ko.Spec.CacheKeyParameters, b.ko.Spec.CacheKeyParameters)
	} else if len(a.ko.Spec.CacheKeyParameters) > 0 {
		if !ackcompare.SliceStringPEqual(a.ko.Spec.CacheKeyParameters, b.ko.Spec.CacheKeyParameters) {
			delta.Add("Spec.CacheKeyParameters", a.ko.Spec.CacheKeyParameters, b.ko.Spec.CacheKeyParameters)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CacheNamespace, b.ko.Spec.CacheNamespace) {
		delta.Add("Spec.CacheNamespace", a.ko.Spec.CacheNamespace, b.ko.Spec.CacheNamespace)
	} else if a.ko.Spec.CacheNamespace != nil && b.ko.Spec.CacheNamespace != nil {
		if *a.ko.Spec.CacheNamespace != *b.ko.Spec.CacheNamespace {
			delta.Add("Spec.CacheNamespace", a.ko.Spec.CacheNamespace, b.ko.Spec.CacheNamespace)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ConnectionID, b.ko.Spec.ConnectionID) {
		delta.Add("Spec.ConnectionID", a.ko.Spec.ConnectionID, b.ko.Spec.ConnectionID)
	} else if a.ko.Spec.ConnectionID != nil && b.ko.Spec.ConnectionID != nil {
		if *a.ko.Spec.ConnectionID != *b.ko.Spec.ConnectionID {
			delta.Add("Spec.ConnectionID", a.ko.Spec.ConnectionID, b.ko.Spec.ConnectionID)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.ConnectionRef, b.ko.Spec.ConnectionRef) {
		delta.Add("Spec.ConnectionRef", a.ko.Spec.ConnectionRef, b.ko.Spec.ConnectionRef)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ConnectionType, b.ko.Spec.ConnectionType) {
		delta.Add("Spec.ConnectionType", a.ko.Spec.ConnectionType, b.ko.Spec.ConnectionType)
	} else if a.ko.Spec.ConnectionType != nil && b.ko.Spec.ConnectionType != nil {
		if *a.ko.Spec.ConnectionType != *b.ko.Spec.ConnectionType {
			delta.Add("Spec.ConnectionType", a.ko.Spec.ConnectionType, b.ko.Spec.ConnectionType)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ContentHandling, b.ko.Spec.ContentHandling) {
		delta.Add("Spec.ContentHandling", a.ko.Spec.ContentHandling, b.ko.Spec.ContentHandling)
	} else if a.ko.Spec.ContentHandling != nil && b.ko.Spec.ContentHandling != nil {
		if *a.ko.Spec.ContentHandling != *b.ko.Spec.ContentHandling {
			delta.Add("Spec.ContentHandling", a.ko.Spec.ContentHandling, b.ko.Spec.ContentHandling)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Credentials, b.ko.Spec.Credentials) {
		delta.Add("Spec.Credentials", a.ko.Spec.Credentials, b.ko.Spec.Credentials)
	} else if a.ko.Spec.Credentials != nil && b.ko.Spec.Credentials != nil {
		if *a.ko.Spec.Credentials != *b.ko.Spec.Credentials {
			delta.Add("Spec.Credentials", a.ko.Spec.Credentials, b.ko.Spec.Credentials)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.HTTPMethod, b.ko.Spec.HTTPMethod) {
		delta.Add("Spec.HTTPMethod", a.ko.Spec.HTTPMethod, b.ko.Spec.HTTPMethod)
	} else if a.ko.Spec.HTTPMethod != nil && b.ko.Spec.HTTPMethod != nil {
		if *a.ko.Spec.HTTPMethod != *b.ko.Spec.HTTPMethod {
			delta.Add("Spec.HTTPMethod", a.ko.Spec.HTTPMethod, b.ko.Spec.HTTPMethod)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.IntegrationHTTPMethod, b.ko.Spec.IntegrationHTTPMethod) {
		delta.Add("Spec.IntegrationHTTPMethod", a.ko.Spec.IntegrationHTTPMethod, b.ko.Spec.IntegrationHTTPMethod)
	} else if a.ko.Spec.IntegrationHTTPMethod != nil && b.ko.Spec.IntegrationHTTPMethod != nil {
		if *a.ko.Spec.IntegrationHTTPMethod != *b.ko.Spec.IntegrationHTTPMethod {
			delta.Add("Spec.IntegrationHTTPMethod", a.ko.Spec.IntegrationHTTPMethod, b.ko.Spec.IntegrationHTTPMethod)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PassthroughBehavior, b.ko.Spec.PassthroughBehavior) {
		delta.Add("Spec.PassthroughBehavior", a.ko.Spec.PassthroughBehavior, b.ko.Spec.PassthroughBehavior)
	} else if a.ko.Spec.PassthroughBehavior != nil && b.ko.Spec.PassthroughBehavior != nil {
		if *a.ko.Spec.PassthroughBehavior != *b.ko.Spec.PassthroughBehavior {
			delta.Add("Spec.PassthroughBehavior", a.ko.Spec.PassthroughBehavior, b.ko.Spec.PassthroughBehavior)
		}
	}
	if len(a.ko.Spec.RequestParameters) != len(b.ko.Spec.RequestParameters) {
		delta.Add("Spec.RequestParameters", a.ko.Spec.RequestParameters, b.ko.Spec.RequestParameters)
	} else if len(a.ko.Spec.RequestParameters) > 0 {
		if !ackcompare.MapStringStringPEqual(a.ko.Spec.RequestParameters, b.ko.Spec.RequestParameters) {
			delta.Add("Spec.RequestParameters", a.ko.Spec.RequestParameters, b.ko.Spec.RequestParameters)
		}
	}
	if len(a.ko.Spec.RequestTemplates) != len(b.ko.Spec.RequestTemplates) {
		delta.Add("Spec.RequestTemplates", a.ko.Spec.RequestTemplates, b.ko.Spec.RequestTemplates)
	} else if len(a.ko.Spec.RequestTemplates) > 0 {
		if !ackcompare.MapStringStringPEqual(a.ko.Spec.RequestTemplates, b.ko.Spec.RequestTemplates) {
			delta.Add("Spec.RequestTemplates", a.ko.Spec.RequestTemplates, b.ko.Spec.RequestTemplates)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ResourceID, b.ko.Spec.ResourceID) {
		delta.Add("Spec.ResourceID", a.ko.Spec.ResourceID, b.ko.Spec.ResourceID)
	} else if a.ko.Spec.ResourceID != nil && b.ko.Spec.ResourceID != nil {
		if *a.ko.Spec.ResourceID != *b.ko.Spec.ResourceID {
			delta.Add("Spec.ResourceID", a.ko.Spec.ResourceID, b.ko.Spec.ResourceID)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.ResourceRef, b.ko.Spec.ResourceRef) {
		delta.Add("Spec.ResourceRef", a.ko.Spec.ResourceRef, b.ko.Spec.ResourceRef)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RestAPIID, b.ko.Spec.RestAPIID) {
		delta.Add("Spec.RestAPIID", a.ko.Spec.RestAPIID, b.ko.Spec.RestAPIID)
	} else if a.ko.Spec.RestAPIID != nil && b.ko.Spec.RestAPIID != nil {
		if *a.ko.Spec.RestAPIID != *b.ko.Spec.RestAPIID {
			delta.Add("Spec.RestAPIID", a.ko.Spec.RestAPIID, b.ko.Spec.RestAPIID)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.RestAPIRef, b.ko.Spec.RestAPIRef) {
		delta.Add("Spec.RestAPIRef", a.ko.Spec.RestAPIRef, b.ko.Spec.RestAPIRef)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.TimeoutInMillis, b.ko.Spec.TimeoutInMillis) {
		delta.Add("Spec.TimeoutInMillis", a.ko.Spec.TimeoutInMillis, b.ko.Spec.TimeoutInMillis)
	} else if a.ko.Spec.TimeoutInMillis != nil && b.ko.Spec.TimeoutInMillis != nil {
		if *a.ko.Spec.TimeoutInMillis != *b.ko.Spec.TimeoutInMillis {
			delta.Add("Spec.TimeoutInMillis", a.ko.Spec.TimeoutInMillis, b.ko.Spec.TimeoutInMillis)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.TLSConfig, b.ko.Spec.TLSConfig) {
		delta.Add("Spec.TLSConfig", a.ko.Spec.TLSConfig, b.ko.Spec.TLSConfig)
	} else if a.ko.Spec.TLSConfig != nil && b.ko.Spec.TLSConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.TLSConfig.InsecureSkipVerification, b.ko.Spec.TLSConfig.InsecureSkipVerification) {
			delta.Add("Spec.TLSConfig.InsecureSkipVerification", a.ko.Spec.TLSConfig.InsecureSkipVerification, b.ko.Spec.TLSConfig.InsecureSkipVerification)
		} else if a.ko.Spec.TLSConfig.InsecureSkipVerification != nil && b.ko.Spec.TLSConfig.InsecureSkipVerification != nil {
			if *a.ko.Spec.TLSConfig.InsecureSkipVerification != *b.ko.Spec.TLSConfig.InsecureSkipVerification {
				delta.Add("Spec.TLSConfig.InsecureSkipVerification", a.ko.Spec.TLSConfig.InsecureSkipVerification, b.ko.Spec.TLSConfig.InsecureSkipVerification)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Type, b.ko.Spec.Type) {
		delta.Add("Spec.Type", a.ko.Spec.Type, b.ko.Spec.Type)
	} else if a.ko.Spec.Type != nil && b.ko.Spec.Type != nil {
		if *a.ko.Spec.Type != *b.ko.Spec.Type {
			delta.Add("Spec.Type", a.ko.Spec.Type, b.ko.Spec.Type)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.URI, b.ko.Spec.URI) {
		delta.Add("Spec.URI", a.ko.Spec.URI, b.ko.Spec.URI)
	} else if a.ko.Spec.URI != nil && b.ko.Spec.URI != nil {
		if *a.ko.Spec.URI != *b.ko.Spec.URI {
			delta.Add("Spec.URI", a.ko.Spec.URI, b.ko.Spec.URI)
		}
	}

	return delta
}
