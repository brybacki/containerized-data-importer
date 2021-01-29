/*
Copyright 2020 The CDI Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package storageclass

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

var _ = FDescribe("AAA", func() {
	It("Should be false if not set", func() {
		client := createClient(
			createStorageClass("TestClass", "test-plugin"),
			createStorageClass("UnknownClass", "unknown-plugin"))

		Expect(getCdiStorageClassCapabilities(client)).ToNot(BeNil())
	})

	It("Should be false if not set", func() {

		sc := createStorageClass("TestClass", "test-plugin")
		storageClassParams := getStorageClassCapabilities(sc)
		Expect(*storageClassParams.accessMode).To(BeEquivalentTo("TestAccessMode"))
		Expect(*storageClassParams.volumeMode).To(BeEquivalentTo("TestVolumeMode"))
	})

})

func createClient(objs ...runtime.Object) client.Client {
	// Register cdi types with the runtime scheme.
	s := scheme.Scheme
	//cdiv1.AddToScheme(s)
	// Create a fake client to mock API calls.
	return fake.NewFakeClientWithScheme(s, objs...)
}

func createStorageClass(name string, provisioner string) *storagev1.StorageClass {
	return &storagev1.StorageClass{
		Provisioner: provisioner,
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
}
