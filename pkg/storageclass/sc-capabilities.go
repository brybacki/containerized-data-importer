package storageclass

import (
	"context"
	"github.com/pkg/errors"

	storagev1 "k8s.io/api/storage/v1"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// CdiStorageClassCapabilities
// preferred capabilities for selected storage class
type CdiStorageClassCapabilities struct {
	name       string
	provider   string
	accessMode *string
	volumeMode *string
}

type ProviderCapabilities struct {
	// matcher
	accessMode string
	volumeMode string
}

var ProviderDefaults = map[string]ProviderCapabilities{
	"A": ProviderCapabilities{accessMode: "a", volumeMode: "b"},
}

// TODO
func getCdiStorageClassCapabilities(client client.Client) ([]CdiStorageClassCapabilities, error) {
	storageClasses := &storagev1.StorageClassList{}
	if err := client.List(context.TODO(), storageClasses); err != nil {
		return nil, errors.New("unable to retrieve storage classes")
	}

	var scCapabilitiesList []CdiStorageClassCapabilities
	for _, storageClass := range storageClasses.Items {
		scCapabilities := getStorageClassCapabilities(&storageClass)
		scCapabilitiesList = append(scCapabilitiesList, scCapabilities)
	}

	return scCapabilitiesList, nil
}

func getStorageClassCapabilities(sc *storagev1.StorageClass) CdiStorageClassCapabilities {
	provisionerPluginName := sc.Provisioner
	capabilities, found := ProviderDefaults[provisionerPluginName]
	if found {
		return CdiStorageClassCapabilities{
			name:       sc.Name,
			provider:   provisionerPluginName,
			accessMode: &capabilities.accessMode,
			volumeMode: &capabilities.volumeMode,
		}
	} else {
		// do some fancy stuff or return empty defaults
		return CdiStorageClassCapabilities{
			name:       sc.Name,
			accessMode: nil,
			volumeMode: nil,
		}
	}
}
