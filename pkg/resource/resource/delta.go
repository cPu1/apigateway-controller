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

package resource

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

	if ackcompare.HasNilDifference(a.ko.Spec.ParentID, b.ko.Spec.ParentID) {
		delta.Add("Spec.ParentID", a.ko.Spec.ParentID, b.ko.Spec.ParentID)
	} else if a.ko.Spec.ParentID != nil && b.ko.Spec.ParentID != nil {
		if *a.ko.Spec.ParentID != *b.ko.Spec.ParentID {
			delta.Add("Spec.ParentID", a.ko.Spec.ParentID, b.ko.Spec.ParentID)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.ParentRef, b.ko.Spec.ParentRef) {
		delta.Add("Spec.ParentRef", a.ko.Spec.ParentRef, b.ko.Spec.ParentRef)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PathPart, b.ko.Spec.PathPart) {
		delta.Add("Spec.PathPart", a.ko.Spec.PathPart, b.ko.Spec.PathPart)
	} else if a.ko.Spec.PathPart != nil && b.ko.Spec.PathPart != nil {
		if *a.ko.Spec.PathPart != *b.ko.Spec.PathPart {
			delta.Add("Spec.PathPart", a.ko.Spec.PathPart, b.ko.Spec.PathPart)
		}
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

	return delta
}
