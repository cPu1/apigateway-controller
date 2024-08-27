package util

import (
	"encoding/json"
	"fmt"

	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"

	svcapitypes "github.com/aws-controllers-k8s/apigateway-controller/apis/v1alpha1"
)

// SetResourceIDAnnotation sets an annotation on obj with resource-identifying fields.
func SetResourceIDAnnotation(obj runtime.Object, val interface{}) error {
	if err := setResourceIDAnnotation(obj, val); err != nil {
		return ackerr.NewTerminalError(fmt.Errorf("error setting resource ID annotation: %w", err))
	}
	return nil
}

func setResourceIDAnnotation(obj runtime.Object, val interface{}) error {
	metaAccessor := meta.NewAccessor()
	annotations, err := metaAccessor.Annotations(obj)
	if err != nil {
		return fmt.Errorf("accessing annotations: %w", err)
	}
	if annotations == nil {
		annotations = map[string]string{}
	}
	data, err := json.Marshal(val)
	if err != nil {
		return fmt.Errorf("marshaling annotations: %w", err)
	}
	annotations[svcapitypes.ResourceIDAnnotation] = string(data)
	if err := metaAccessor.SetAnnotations(obj, annotations); err != nil {
		return fmt.Errorf("setting annotations: %w", err)
	}
	return nil
}

func parseResourceIDAnnotation(obj runtime.Object, val interface{}) (bool, error) {
	metaAccessor := meta.NewAccessor()
	annotations, err := metaAccessor.Annotations(obj)
	if err != nil {
		return false, fmt.Errorf("accessing annotations: %w", err)
	}
	data, ok := annotations[svcapitypes.ResourceIDAnnotation]
	if !ok {
		return false, nil
	}
	if err := json.Unmarshal([]byte(data), val); err != nil {
		return false, fmt.Errorf("unmarshaling resource ID annotation: %w", err)
	}
	return true, nil
}

// Resource represents an ACK resource.
type Resource[T any] interface {
	runtime.Object
	DeepCopy() T
}

// UpdateResourceFromAnnotation updates an API Gateway resource by parsing resource-identifying fields from annotations.
func UpdateResourceFromAnnotation[T any, K Resource[K]](obj K, updateResource func(T, K)) error {
	var id T
	found, err := parseResourceIDAnnotation(obj, &id)
	if err != nil {
		return ackerr.NewTerminalError(err)
	}
	if !found {
		return nil
	}
	updateResource(id, obj.DeepCopy())
	return nil
}
