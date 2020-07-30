// +build !ignore_autogenerated

/*
Copyright 2020 The hostpath provisioner operator Authors.

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

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"kubevirt.io/hostpath-provisioner-operator/pkg/apis/hostpathprovisioner/v1alpha1.HostPathProvisioner":       schema_pkg_apis_hostpathprovisioner_v1alpha1_HostPathProvisioner(ref),
		"kubevirt.io/hostpath-provisioner-operator/pkg/apis/hostpathprovisioner/v1alpha1.HostPathProvisionerSpec":   schema_pkg_apis_hostpathprovisioner_v1alpha1_HostPathProvisionerSpec(ref),
		"kubevirt.io/hostpath-provisioner-operator/pkg/apis/hostpathprovisioner/v1alpha1.HostPathProvisionerStatus": schema_pkg_apis_hostpathprovisioner_v1alpha1_HostPathProvisionerStatus(ref),
		"kubevirt.io/hostpath-provisioner-operator/pkg/apis/hostpathprovisioner/v1alpha1.PathConfig":                schema_pkg_apis_hostpathprovisioner_v1alpha1_PathConfig(ref),
	}
}

func schema_pkg_apis_hostpathprovisioner_v1alpha1_HostPathProvisioner(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "HostPathProvisioner is the Schema for the hostpathprovisioners API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("kubevirt.io/hostpath-provisioner-operator/pkg/apis/hostpathprovisioner/v1alpha1.HostPathProvisionerSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("kubevirt.io/hostpath-provisioner-operator/pkg/apis/hostpathprovisioner/v1alpha1.HostPathProvisionerStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "kubevirt.io/hostpath-provisioner-operator/pkg/apis/hostpathprovisioner/v1alpha1.HostPathProvisionerSpec", "kubevirt.io/hostpath-provisioner-operator/pkg/apis/hostpathprovisioner/v1alpha1.HostPathProvisionerStatus"},
	}
}

func schema_pkg_apis_hostpathprovisioner_v1alpha1_HostPathProvisionerSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "HostPathProvisionerSpec defines the desired state of HostPathProvisioner",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"imagePullPolicy": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"pathConfig": {
						SchemaProps: spec.SchemaProps{
							Description: "PathConfig describes the location and layout of PV storage on nodes",
							Ref:         ref("kubevirt.io/hostpath-provisioner-operator/pkg/apis/hostpathprovisioner/v1alpha1.PathConfig"),
						},
					},
				},
				Required: []string{"pathConfig"},
			},
		},
		Dependencies: []string{
			"kubevirt.io/hostpath-provisioner-operator/pkg/apis/hostpathprovisioner/v1alpha1.PathConfig"},
	}
}

func schema_pkg_apis_hostpathprovisioner_v1alpha1_HostPathProvisionerStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "HostPathProvisionerStatus defines the observed state of HostPathProvisioner",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions contains the current conditions observed by the operator",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/custom-resource-status/conditions/v1.Condition"),
									},
								},
							},
						},
					},
					"operatorVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "OperatorVersion The version of the HostPathProvisioner Operator",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"targetVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "TargetVersion The targeted version of the HostPathProvisioner deployment",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"observedVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "ObservedVersion The observed version of the HostPathProvisioner deployment",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/custom-resource-status/conditions/v1.Condition"},
	}
}

func schema_pkg_apis_hostpathprovisioner_v1alpha1_PathConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PathConfig contains the information needed to build the path where the PVs will be created.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"path": {
						SchemaProps: spec.SchemaProps{
							Description: "Path The path the directories for the PVs are created under",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"useNamingPrefix": {
						SchemaProps: spec.SchemaProps{
							Description: "UseNamingPrefix Use the name of the PVC requesting the PV as part of the directory created",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}
