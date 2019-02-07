package apis

import (
	"github.com/jam2in/arcus-operator/pkg/apis/jam2in/v1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1.SchemeBuilder.AddToScheme)
}
