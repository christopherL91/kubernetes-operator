package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterPhase string

const (
	ClusterPhaseInitial  ClusterPhase = ""
	ClusterPhaseRunning  ClusterPhase = "Running"
	ClusterPhaseCreating ClusterPhase = "Creating"
	ClusterPhaseCleanUp  ClusterPhase = "CleanUp"
	ClusterPhaseFailed   ClusterPhase = "Failed"
	ClusterPhaseDone     ClusterPhase = "Done"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FooList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Foo `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Foo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   FooSpec   `json:"spec"`
	Status FooStatus `json:"status,omitempty"`
}

type FooSpec struct{}

type FooStatus struct {
	Phase             ClusterPhase `json:"phase"`
	Initialized       bool         `json:"initialized"`
	Reason            string       `json:"reason"`
	AvailableReplicas int32        `json:"availableReplicas"`
}
