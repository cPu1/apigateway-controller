package rest_api

import (
	"errors"
	"fmt"
	"slices"
	"strconv"

	"github.com/aws-controllers-k8s/runtime/pkg/compare"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"

	"github.com/aws-controllers-k8s/apigateway-controller/pkg/tags"
)

var syncTags = tags.SyncTags

func makeARN(vpcLinkName string) string {
	// TODO: TESTME.
	return fmt.Sprintf("arn:aws:apigateway:us-west-2::/restapis/%s", vpcLinkName)
}

func updateRestAPIInput(desired, latest *resource, input *apigateway.UpdateRestApiInput, delta *compare.Delta) error {
	makeReplaceOp := func(path string, desiredVal *string) *apigateway.PatchOperation {
		return &apigateway.PatchOperation{
			Op:    aws.String(apigateway.OpReplace),
			Path:  aws.String(path),
			Value: desiredVal,
		}
	}
	emptyIfNil := func(val *string) *string {
		if val == nil {
			return aws.String("")
		}
		return val
	}
	latestSpec := latest.ko.Spec
	desiredSpec := desired.ko.Spec

	if delta.DifferentAt("Spec.APIKeySource") {
		input.PatchOperations = append(input.PatchOperations, makeReplaceOp("/apiKeySource", emptyIfNil(desiredSpec.APIKeySource)))
	}
	if delta.DifferentAt("Spec.BinaryMediaTypes") {
		patchOps := makeSlicePatch(aws.StringValueSlice(latestSpec.BinaryMediaTypes), aws.StringValueSlice(desiredSpec.BinaryMediaTypes),
			"/binaryMediaTypes")
		input.PatchOperations = append(input.PatchOperations, patchOps...)
	}
	if delta.DifferentAt("Spec.Description") {
		input.PatchOperations = append(input.PatchOperations, makeReplaceOp("/description", emptyIfNil(desiredSpec.Description)))
	}
	if delta.DifferentAt("Spec.DisableExecuteAPIEndpoint") {
		var disable bool
		if desiredSpec.DisableExecuteAPIEndpoint != nil {
			disable = *desiredSpec.DisableExecuteAPIEndpoint
		}
		input.PatchOperations = append(input.PatchOperations, makeReplaceOp("/disableExecuteApiEndpoint",
			aws.String(strconv.FormatBool(disable))))
	}
	if delta.DifferentAt("Spec.EndpointConfiguration.Types") {
		if desiredSpec.EndpointConfiguration == nil {
			return errors.New("spec.endpointConfiguration.types is required")
		}
		if len(desiredSpec.EndpointConfiguration.Types) != 1 {
			return errors.New("spec.endpointConfiguration.types must contain exactly one element")
		}
		input.PatchOperations = append(input.PatchOperations, &apigateway.PatchOperation{
			Op:    aws.String(apigateway.OpReplace),
			Path:  aws.String("/endpointConfiguration/types/0"),
			Value: desiredSpec.EndpointConfiguration.Types[0],
		})

	}
	if delta.DifferentAt("Spec.EndpointConfiguration.VPCEndpointIDs") {
		var (
			currEndpointIDs    []string
			desiredEndpointIDs []string
		)
		if latestSpec.EndpointConfiguration != nil {
			currEndpointIDs = aws.StringValueSlice(latestSpec.EndpointConfiguration.VPCEndpointIDs)
		}
		if desiredSpec.EndpointConfiguration != nil {
			desiredEndpointIDs = aws.StringValueSlice(desiredSpec.EndpointConfiguration.VPCEndpointIDs)
		}
		patchOps := makeSlicePatch(currEndpointIDs, desiredEndpointIDs, "/endpointConfiguration/vpcEndpointIds")
		input.PatchOperations = append(input.PatchOperations, patchOps...)
	}
	if delta.DifferentAt("Spec.MinimumCompressionSize") {
		var val *string
		if desiredSpec.MinimumCompressionSize != nil {
			val = aws.String(strconv.FormatInt(*desiredSpec.MinimumCompressionSize, 10))
		}
		input.PatchOperations = append(input.PatchOperations, makeReplaceOp("/minimumCompressionSize", val))
	}
	if delta.DifferentAt("Spec.Name") {
		input.PatchOperations = append(input.PatchOperations, makeReplaceOp("/name", desiredSpec.Name))
	}
	if delta.DifferentAt("Spec.Policy") {
		input.PatchOperations = append(input.PatchOperations, makeReplaceOp("/policy", emptyIfNil(desiredSpec.Policy)))
	}
	return nil
}

func makeSlicePatch(currValues, desiredValues []string, path string) []*apigateway.PatchOperation {
	var patchOps []*apigateway.PatchOperation
	for _, val := range currValues {
		if !slices.Contains(desiredValues, val) {
			patchOps = append(patchOps, &apigateway.PatchOperation{
				Op:    aws.String(apigateway.OpRemove),
				Path:  aws.String(path),
				Value: aws.String(val),
			})
		}
	}
	for _, val := range desiredValues {
		if !slices.Contains(currValues, val) {
			patchOps = append(patchOps, &apigateway.PatchOperation{
				Op:    aws.String(apigateway.OpAdd),
				Path:  aws.String(path),
				Value: aws.String(val),
			})
		}
	}
	return patchOps
}
