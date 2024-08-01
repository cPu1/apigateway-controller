package integration

import (
	"strconv"

	"github.com/aws-controllers-k8s/runtime/pkg/compare"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"

	"github.com/aws-controllers-k8s/apigateway-controller/pkg/util/patch"
)

func updateIntegrationInput(desired, latest *resource, input *apigateway.UpdateIntegrationInput, delta *compare.Delta) {
	latestSpec := latest.ko.Spec
	desiredSpec := desired.ko.Spec

	var patchSet patch.Set
	if delta.DifferentAt("Spec.CacheKeyParameters") {
		patchSet.ForSlice("/cacheKeyParameters", latestSpec.CacheKeyParameters, desiredSpec.CacheKeyParameters)
	}
	if delta.DifferentAt("Spec.CacheNamespace") {
		patchSet.Replace("/cacheNamespace", desiredSpec.CacheNamespace)
	}
	if delta.DifferentAt("Spec.ConnectionID") {
		patchSet.Replace("/connectionId", desiredSpec.ConnectionID)
	}
	if delta.DifferentAt("Spec.ConnectionType") {
		patchSet.Replace("/connectionType", desiredSpec.ConnectionType)
	}
	if delta.DifferentAt("Spec.ContentHandling") {
		patchSet.Replace("/contentHandling", desiredSpec.ContentHandling)
	}
	if delta.DifferentAt("Spec.HTTPMethod") {
		patchSet.Replace("/httpMethod", desiredSpec.HTTPMethod)
	}
	if delta.DifferentAt("Spec.PassthroughBehavior") {
		patchSet.Replace("/passthroughBehavior", desiredSpec.PassthroughBehavior)
	}
	if delta.DifferentAt("Spec.RequestParameters") {
		patchSet.ForMap("/requestParameters", latestSpec.RequestParameters, desiredSpec.RequestParameters, true)
	}
	if delta.DifferentAt("Spec.RequestTemplates") {
		patchSet.ForMap("/requestTemplates", latestSpec.RequestTemplates, desiredSpec.RequestTemplates, true)
	}
	if delta.DifferentAt("Spec.TimeoutInMillis") {
		var val *string
		if desiredSpec.TimeoutInMillis != nil {
			val = aws.String(strconv.FormatInt(*desiredSpec.TimeoutInMillis, 10))
		}
		patchSet.Replace("/timeoutInMillis", val)
	}
	if delta.DifferentAt("Spec.TLSConfig.InsecureSkipVerification") {
		var val *string
		if desiredSpec.TLSConfig != nil && desiredSpec.TLSConfig.InsecureSkipVerification != nil {
			val = aws.String(strconv.FormatBool(*desiredSpec.TLSConfig.InsecureSkipVerification))
		}
		patchSet.Replace("/tlsConfig/insecureSkipVerification", val)
	}

	if delta.DifferentAt("Spec.URI") {
		patchSet.Replace("/uri", desiredSpec.URI)
	}
	input.PatchOperations = patchSet.GetPatchOperations()
}

func customPreCompare(a, b *resource) {
	if a.ko.Spec.RequestTemplates == nil && b.ko.Spec.RequestTemplates != nil {
		a.ko.Spec.RequestTemplates = map[string]*string{}
	} else if a.ko.Spec.RequestTemplates != nil && b.ko.Spec.RequestTemplates == nil {
		b.ko.Spec.RequestTemplates = map[string]*string{}
	}
	if a.ko.Spec.RequestParameters == nil && b.ko.Spec.RequestParameters != nil {
		a.ko.Spec.RequestParameters = map[string]*string{}
	} else if a.ko.Spec.RequestParameters != nil && b.ko.Spec.RequestParameters == nil {
		b.ko.Spec.RequestParameters = map[string]*string{}
	}
}
