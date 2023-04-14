// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ImageLookupPolicyApplyConfiguration represents an declarative configuration of the ImageLookupPolicy type for use
// with apply.
type ImageLookupPolicyApplyConfiguration struct {
	Local *bool `json:"local,omitempty"`
}

// ImageLookupPolicyApplyConfiguration constructs an declarative configuration of the ImageLookupPolicy type for use with
// apply.
func ImageLookupPolicy() *ImageLookupPolicyApplyConfiguration {
	return &ImageLookupPolicyApplyConfiguration{}
}

// WithLocal sets the Local field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Local field is set to the value of the last call.
func (b *ImageLookupPolicyApplyConfiguration) WithLocal(value bool) *ImageLookupPolicyApplyConfiguration {
	b.Local = &value
	return b
}