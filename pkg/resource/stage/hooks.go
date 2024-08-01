package stage

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aws-controllers-k8s/runtime/pkg/compare"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"

	svcapitypes "github.com/aws-controllers-k8s/apigateway-controller/apis/v1alpha1"
	"github.com/aws-controllers-k8s/apigateway-controller/pkg/tags"
	"github.com/aws-controllers-k8s/apigateway-controller/pkg/util"
	"github.com/aws-controllers-k8s/apigateway-controller/pkg/util/patch"
)

var syncTags = tags.SyncTags

func arnForResource(desired *svcapitypes.Stage) (string, error) {
	return util.ARNForResource(desired.Status.ACKResourceMetadata,
		fmt.Sprintf("/restapis/%s/stages/%s", *desired.Spec.RestAPIID, *desired.Spec.StageName))
}

func updateStageInput(desired, latest *resource, input *apigateway.UpdateStageInput, delta *compare.Delta) {
	latestSpec := latest.ko.Spec
	desiredSpec := desired.ko.Spec

	var patchSet patch.Set
	if delta.DifferentAt("Spec.CacheClusterEnabled") {
		var val *string
		if desiredSpec.CacheClusterEnabled != nil {
			val = aws.String(strconv.FormatBool(*desiredSpec.CacheClusterEnabled))
		}
		patchSet.Replace("/cacheClusterEnabled", val)
	}
	if delta.DifferentAt("Spec.CacheClusterSize") {
		patchSet.Replace("/cacheClusterSize", desiredSpec.CacheClusterSize)
	}
	if delta.DifferentAt("Spec.CanarySettings") {
		makeCanarySettingsPatch(delta, desiredSpec, latestSpec, &patchSet)
	}
	if delta.DifferentAt("Spec.DeploymentID") {
		patchSet.Replace("/deploymentId", desiredSpec.DeploymentID)
	}
	if delta.DifferentAt("Spec.Description") {
		patchSet.Replace("/description", desiredSpec.Description)
	}
	if delta.DifferentAt("Spec.DocumentationVersion") {
		patchSet.Replace("/documentationVersion", desiredSpec.DocumentationVersion)
	}
	if delta.DifferentAt("Spec.Variables") {
		patchSet.ForMap("/variables", latestSpec.Variables, desiredSpec.Variables, false)
	}
	if delta.DifferentAt("Spec.TracingEnabled") {
		var val *string
		if desiredSpec.TracingEnabled != nil {
			val = aws.String(strconv.FormatBool(*desiredSpec.TracingEnabled))
		}
		patchSet.Replace("/tracingEnabled", val)
	}
	input.PatchOperations = patchSet.GetPatchOperations()
}

func makeCanarySettingsPatch(delta *compare.Delta, desiredSpec, latestSpec svcapitypes.StageSpec, patchSet *patch.Set) {
	canary := desiredSpec.CanarySettings
	diff, _ := json.Marshal(delta.Differences)
	fmt.Println("makep", string(diff), canary, desiredSpec)
	const rootKey = "/canarySettings"
	if canary == nil {
		patchSet.Remove(rootKey)
		return
	}

	prefixRootKey := func(key string) string {
		return fmt.Sprintf("%s/%s", rootKey, key)
	}
	if delta.DifferentAt("Spec.CanarySettings.DeploymentID") {
		patchSet.Replace(prefixRootKey("deploymentId"), canary.DeploymentID)
	}
	if delta.DifferentAt("Spec.CanarySettings.PercentTraffic") {
		var val *string
		if canary.PercentTraffic != nil {
			val = aws.String(fmt.Sprintf("%f", *canary.PercentTraffic))
		}
		patchSet.Replace(prefixRootKey("percentTraffic"), val)
	}
	fmt.Println("melta", delta.DifferentAt("Spec.CanarySettings.stageVariableOverrides"))
	if delta.DifferentAt("Spec.CanarySettings.StageVariableOverrides") {
		desiredValues := canary.StageVariableOverrides
		if desiredValues == nil {
			desiredValues = map[string]*string{}
		}
		var currValues map[string]*string
		if latestSpec.CanarySettings != nil && latestSpec.CanarySettings.StageVariableOverrides != nil {
			currValues = latestSpec.CanarySettings.StageVariableOverrides
		} else {
			currValues = map[string]*string{}
		}
		fmt.Println("melta2", currValues, desiredValues)
		patchSet.ForMap(prefixRootKey("stageVariableOverrides"), currValues, desiredValues, false)
	}
	if delta.DifferentAt("Spec.CanarySettings.UseStageCache") {
		var val *string
		if canary.UseStageCache != nil {
			val = aws.String(strconv.FormatBool(*canary.UseStageCache))
		}
		patchSet.Replace(prefixRootKey("useStageCache"), val)
	}
}

func customPreCompare(a, b *resource) {
	if a.ko.Spec.Variables == nil && b.ko.Spec.Variables != nil {
		a.ko.Spec.Variables = map[string]*string{}
	} else if a.ko.Spec.Variables != nil && b.ko.Spec.Variables == nil {
		b.ko.Spec.Variables = map[string]*string{}
	}
	if a.ko.Spec.CanarySettings == nil && b.ko.Spec.CanarySettings == nil {
		return
	}
	if a.ko.Spec.CanarySettings == nil {
		a.ko.Spec.CanarySettings = &svcapitypes.CanarySettings{}
	}
	if b.ko.Spec.CanarySettings == nil {
		b.ko.Spec.CanarySettings = &svcapitypes.CanarySettings{}
	}
	if a.ko.Spec.CanarySettings.StageVariableOverrides == nil && b.ko.Spec.CanarySettings.StageVariableOverrides != nil {
		a.ko.Spec.CanarySettings.StageVariableOverrides = map[string]*string{}
	} else if a.ko.Spec.CanarySettings.StageVariableOverrides != nil && b.ko.Spec.CanarySettings.StageVariableOverrides == nil {
		b.ko.Spec.CanarySettings.StageVariableOverrides = map[string]*string{}
	}
}
