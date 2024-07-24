package vpc_link

import (
	"fmt"

	"github.com/aws-controllers-k8s/runtime/pkg/compare"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"

	"github.com/aws-controllers-k8s/apigateway-controller/pkg/tags"
)

var syncTags = tags.SyncTags

func makeARN(vpcLinkName string) string {
	return fmt.Sprintf("arn:aws:apigateway:us-west-2::/vpclinks/%s", vpcLinkName)
}

func updateVPCLinkInput(desired *resource, input *apigateway.UpdateVpcLinkInput, delta *compare.Delta) error {
	makePatchOp := func(path string, desiredVal *string) *apigateway.PatchOperation {
		return &apigateway.PatchOperation{
			Op:    aws.String(apigateway.OpReplace),
			Path:  aws.String(fmt.Sprintf("/%s", path)),
			Value: desiredVal,
		}
	}
	if delta.DifferentAt("Spec.Name") {
		input.PatchOperations = append(input.PatchOperations, makePatchOp("name", desired.ko.Spec.Name))
	}
	if delta.DifferentAt("Spec.Description") {
		input.PatchOperations = append(input.PatchOperations, makePatchOp("description", desired.ko.Spec.Description))
	}
	return nil
}
