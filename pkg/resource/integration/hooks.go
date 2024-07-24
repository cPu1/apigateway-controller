package integration

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/aws-controllers-k8s/runtime/pkg/compare"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

func updateIntegrationInput(desired, latest *resource, input *apigateway.UpdateIntegrationInput, delta *compare.Delta) error {
	makeReplaceOp := func(path string, desiredVal *string) *apigateway.PatchOperation {
		return &apigateway.PatchOperation{
			Op:    aws.String(apigateway.OpReplace),
			Path:  aws.String(path),
			Value: desiredVal,
		}
	}
	addPatch := func(operations ...*apigateway.PatchOperation) {
		input.PatchOperations = append(input.PatchOperations, operations...)
	}
	latestSpec := latest.ko.Spec
	desiredSpec := desired.ko.Spec

	if delta.DifferentAt("Spec.CacheKeyParameters") {
		addPatch(makeSlicePatch(latestSpec.CacheKeyParameters, desiredSpec.CacheKeyParameters, "/cacheKeyParameters")...)
	}
	if delta.DifferentAt("Spec.CacheNamespace") {
		addPatch(makeReplaceOp("/cacheNamespace", desiredSpec.CacheNamespace))
	}
	if delta.DifferentAt("Spec.ConnectionID") {
		addPatch(makeReplaceOp("/connectionId", desiredSpec.ConnectionID))
	}
	if delta.DifferentAt("Spec.ConnectionType") {
		addPatch(makeReplaceOp("/connectionType", desiredSpec.ConnectionType))
	}
	if delta.DifferentAt("Spec.ContentHandling") {
		addPatch(makeReplaceOp("/contentHandling", desiredSpec.ContentHandling))
	}
	if delta.DifferentAt("Spec.HTTPMethod") {
		addPatch(makeReplaceOp("/httpMethod", desiredSpec.HTTPMethod))
	}
	if delta.DifferentAt("Spec.PassthroughBehavior") {
		addPatch(makeReplaceOp("/passthroughBehavior", desiredSpec.PassthroughBehavior))
	}
	if delta.DifferentAt("Spec.RequestParameters") {
		addPatch(makeMapPatch(latestSpec.RequestParameters, desiredSpec.RequestParameters, "/requestParameters")...)
	}
	if delta.DifferentAt("Spec.RequestTemplates") {
		addPatch(makeMapPatch(latestSpec.RequestTemplates, desiredSpec.RequestTemplates, "/requestTemplates")...)
	}
	if delta.DifferentAt("Spec.TimeoutInMillis") {
		var val *string
		if desiredSpec.TimeoutInMillis != nil {
			val = aws.String(strconv.FormatInt(*desiredSpec.TimeoutInMillis, 10))
		}
		addPatch(makeReplaceOp("/timeoutInMillis", val))
	}
	if delta.DifferentAt("Spec.TLSConfig.InsecureSkipVerification") {
		var val *string
		if desiredSpec.TLSConfig != nil && desiredSpec.TLSConfig.InsecureSkipVerification != nil {
			val = aws.String(strconv.FormatBool(*desiredSpec.TLSConfig.InsecureSkipVerification))
		}
		addPatch(makeReplaceOp("/tlsConfig/insecureSkipVerification", val))
	}

	if delta.DifferentAt("Spec.URI") {
		addPatch(makeReplaceOp("/uri", desiredSpec.URI))
	}
	return nil
}

var patchKeyEncoder = strings.NewReplacer("~", "~0", "/", "~1")

func makeSlicePatch(currValues, desiredValues []*string, path string) []*apigateway.PatchOperation {
	current := aws.StringValueSlice(currValues)
	desired := aws.StringValueSlice(desiredValues)
	var patchOps []*apigateway.PatchOperation
	for _, val := range current {
		if !slices.Contains(desired, val) {
			patchOps = append(patchOps, &apigateway.PatchOperation{
				Op:   aws.String(apigateway.OpRemove),
				Path: aws.String(fmt.Sprintf("%s/%s", path, patchKeyEncoder.Replace(val))),
			})
		}
	}
	for _, val := range desired {
		if !slices.Contains(current, val) {
			patchOps = append(patchOps, &apigateway.PatchOperation{
				Op:   aws.String(apigateway.OpAdd),
				Path: aws.String(fmt.Sprintf("%s/%s", path, patchKeyEncoder.Replace(val))),
			})
		}
	}
	return patchOps
}

func makeMapPatch(currValues, desiredValues map[string]*string, path string) []*apigateway.PatchOperation {
	var patchOps []*apigateway.PatchOperation
	for k, v := range currValues {
		if _, ok := desiredValues[k]; !ok {
			patchOps = append(patchOps, &apigateway.PatchOperation{
				Op:   aws.String(apigateway.OpRemove),
				Path: aws.String(fmt.Sprintf("%s/%s", path, patchKeyEncoder.Replace(k))),
			})
		} else {
			patchOps = append(patchOps, &apigateway.PatchOperation{
				Op:    aws.String(apigateway.OpReplace),
				Path:  aws.String(fmt.Sprintf("%s/%s", path, patchKeyEncoder.Replace(k))),
				Value: v,
			})
		}
	}
	return patchOps
}
