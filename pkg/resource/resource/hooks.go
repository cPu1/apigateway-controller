package resource

import (
	"github.com/aws-controllers-k8s/runtime/pkg/compare"
	"github.com/aws/aws-sdk-go/service/apigateway"

	"github.com/aws-controllers-k8s/apigateway-controller/pkg/util/patch"
)

func updateResourceInput(desired *resource, input *apigateway.UpdateResourceInput, delta *compare.Delta) {
	desiredSpec := desired.ko.Spec
	var patchSet patch.Set
	if delta.DifferentAt("Spec.ParentID") {
		patchSet.Replace("/parentID", desiredSpec.ParentID)
	}
	if delta.DifferentAt("Spec.PathPart") {
		patchSet.Replace("/pathPart", desiredSpec.PathPart)
	}
	input.PatchOperations = patchSet.GetPatchOperations()
}
