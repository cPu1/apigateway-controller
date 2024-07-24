package stage

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

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
	//latestSpec := latest.ko.Spec
	desiredSpec := desired.ko.Spec

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
			/*if delta.DifferentAt("Spec.CanarySettings.StageVariableOverrides") {
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
			}*/

		}
		/*desiredSpec.CanarySettings.
		addPatch(makeReplaceOp("canarySettings"))*/
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
