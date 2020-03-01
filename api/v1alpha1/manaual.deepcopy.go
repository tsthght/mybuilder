package v1alpha1

import "k8s.io/apimachinery/pkg/runtime"

func (in *SidecarSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
