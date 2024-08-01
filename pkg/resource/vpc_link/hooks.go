package vpc_link

import (
	"fmt"

	"github.com/aws-controllers-k8s/runtime/pkg/compare"
	"github.com/aws/aws-sdk-go/service/apigateway"

	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"

	svcapitypes "github.com/aws-controllers-k8s/apigateway-controller/apis/v1alpha1"
	"github.com/aws-controllers-k8s/apigateway-controller/pkg/tags"
	"github.com/aws-controllers-k8s/apigateway-controller/pkg/util"
	"github.com/aws-controllers-k8s/apigateway-controller/pkg/util/patch"
)

var syncTags = tags.SyncTags

func validateDeleteState(r *resource) error {
	if status := r.ko.Status.Status; status != nil {
		switch svcapitypes.VPCLinkStatus_SDK(*status) {
		case svcapitypes.VPCLinkStatus_SDK_DELETING, svcapitypes.VPCLinkStatus_SDK_PENDING:
			return ackrequeue.NeededAfter(
				fmt.Errorf("VPCLink is in %s state, it cannot be modified or deleted", *status),
				ackrequeue.DefaultRequeueAfterDuration,
			)
		}
	}
	return nil
}

func arnForResource(desired *svcapitypes.VPCLink) (string, error) {
	return util.ARNForResource(desired.Status.ACKResourceMetadata, fmt.Sprintf("/vpclinks/%s", *desired.Status.ID))
}

func updateVPCLinkInput(desired *resource, input *apigateway.UpdateVpcLinkInput, delta *compare.Delta) {
	var patchSet patch.Set
	if delta.DifferentAt("Spec.Name") {
		patchSet.Replace("/name", desired.ko.Spec.Name)
	}
	if delta.DifferentAt("Spec.Description") {
		patchSet.Replace("/description", desired.ko.Spec.Description)
	}
	input.PatchOperations = patchSet.GetPatchOperations()
}
