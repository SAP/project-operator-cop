/*
SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and project-operator-cop contributors
SPDX-License-Identifier: Apache-2.0
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"

	prometheusv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"

	"github.com/sap/component-operator-runtime/pkg/component"
	componentoperatorruntimetypes "github.com/sap/component-operator-runtime/pkg/types"
)

// ProjectOperatorSpec defines the desired state of ProjectOperator.
type ProjectOperatorSpec struct {
	component.Spec `json:",inline"`
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:default=1
	ReplicaCount int `json:"replicaCount,omitempty"`
	// +optional
	Image                          component.ImageSpec `json:"image"`
	component.KubernetesProperties `json:",inline"`
	NamespacePrefix                string             `json:"namespacePrefix,omitempty"`
	AdminClusterRole               string             `json:"adminClusterRole,omitempty"`
	ViewerClusterRole              string             `json:"viewerClusterRole,omitempty"`
	EnableClusterView              bool               `json:"enableClusterView,omitempty"`
	Metrics                        *MetricsProperties `json:"metrics,omitempty"`
}

// MetricsProperties defines the properties for the metrics server.
type MetricsProperties struct {
	PodMonitor     *PodMonitorProperties     `json:"podMonitor,omitempty"`
	PrometheusRule *PrometheusRuleProperties `json:"prometheusRule,omitempty"`
}

// PodMonitorProperties defines the properties for the pod monitor.
type PodMonitorProperties struct {
	Enabled bool `json:"enabled,omitempty"`
}

// PrometheusRuleProperties defines the properties for the prometheus rule.
type PrometheusRuleProperties struct {
	Enabled bool                `json:"enabled,omitempty"`
	Rules   []prometheusv1.Rule `json:"rules,omitempty"`
}

// ProjectOperatorStatus defines the observed state of ProjectOperator.
type ProjectOperatorStatus struct {
	component.Status `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="State",type=string,JSONPath=`.status.state`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +genclient

// ProjectOperator is the Schema for the projectoperators API.
type ProjectOperator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ProjectOperatorSpec `json:"spec,omitempty"`
	// +kubebuilder:default={"observedGeneration":-1}
	Status ProjectOperatorStatus `json:"status,omitempty"`
}

var _ component.Component = &ProjectOperator{}

// +kubebuilder:object:root=true

// ProjectOperatorList contains a list of ProjectOperator.
type ProjectOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectOperator `json:"items"`
}

func (s *ProjectOperatorSpec) ToUnstructured() map[string]any {
	result, err := runtime.DefaultUnstructuredConverter.ToUnstructured(s)
	if err != nil {
		panic(err)
	}
	return result
}

func (c *ProjectOperator) GetDeploymentNamespace() string {
	if c.Spec.Namespace != "" {
		return c.Spec.Namespace
	}
	return c.Namespace
}

func (c *ProjectOperator) GetDeploymentName() string {
	if c.Spec.Name != "" {
		return c.Spec.Name
	}
	return c.Name
}

func (c *ProjectOperator) GetSpec() componentoperatorruntimetypes.Unstructurable {
	return &c.Spec
}

func (c *ProjectOperator) GetStatus() *component.Status {
	return &c.Status.Status
}

func init() {
	SchemeBuilder.Register(&ProjectOperator{}, &ProjectOperatorList{})
}
