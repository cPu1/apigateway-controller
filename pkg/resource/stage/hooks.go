package stage

import (
	"fmt"
	"strconv"

	"github.com/aws-controllers-k8s/runtime/pkg/compare"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"
)

func updateStageInput(desired, latest *resource, input *apigateway.UpdateRestApiInput, delta *compare.Delta) error {
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

	desiredSpec.CanarySettings.DeploymentID
	if delta.DifferentAt("Spec.CacheClusterEnabled") {
		var val *string
		if desiredSpec.CacheClusterEnabled != nil {
			val = aws.String(strconv.FormatBool(*desiredSpec.CacheClusterEnabled))
		}
		addPatch(makeReplaceOp("/cacheClusterEnabled", val))
	}
	if delta.DifferentAt("Spec.CacheClusterSize") {
		addPatch(makeReplaceOp("/cacheClusterSize", desiredSpec.CacheClusterSize))
	}
	if delta.DifferentAt("Spec.CanarySettings") {
		if desiredSpec.CanarySettings == nil {
			addPatch(&apigateway.PatchOperation{
				Op:   aws.String(apigateway.OpRemove),
				Path: aws.String("/canarySettings"),
			})
		} else {
			canary := desiredSpec.CanarySettings
			if delta.DifferentAt("Spec.DeploymentID") {
				addPatch(makeReplaceOp("/deploymentId", canary.DeploymentID))
			}
			if delta.DifferentAt("Spec.PercentTraffic") {
				var val *string
				if canary.PercentTraffic != nil {
					val = aws.String(fmt.Sprintf("%f", *canary.PercentTraffic))
				}
				addPatch(makeReplaceOp("/percentTraffic", val))
			}
			if delta.DifferentAt("Spec.CanarySettings.StageVariableOverrides") {
				var (
					desiredVal map[string]*string
					currVal    map[string]*string
				)
				if canary.StageVariableOverrides == nil {
					desiredVal = map[string]*string{}
				}
				if latestSpec.CanarySettings != nil && latestSpec.CanarySettings.StageVariableOverrides != nil {
					currVal = latestSpec.CanarySettings.StageVariableOverrides
				}
				makeMapPatch(canary.StageVariableOverrides)
			}

		}
		desiredSpec.CanarySettings.
			addPatch(makeReplaceOp("canarySettings"))
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
