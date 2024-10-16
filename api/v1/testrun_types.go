/*
Copyright 2024.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TestRunSpec defines the desired state of TestRun
type TestRunSpec struct {
	TestRunId  string      `json:"testRunId"`
	K6TestRun  string      `json:"k6TestRun"`
	StartedAt  metav1.Time `json:"startedAt"`
	FinishedAt metav1.Time `json:"finishedAt"`
	Succeeded  bool        `json:"succeeded"`
}

// TestRunStatus defines the observed state of TestRun
type TestRunStatus struct {
	Id        string       `json:"id,omitempty"`
	K6TestRun string       `json:"k6TestRun,omitempty"`
	Stage     TestRunStage `json:"stage,omitempty"`

	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:validation:Enum=initialization;initialized;created;started;stopped;finished;error
type TestRunStage string

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// TestRun is the Schema for the testruns API
type TestRun struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TestRunSpec   `json:"spec,omitempty"`
	Status TestRunStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TestRunList contains a list of TestRun
type TestRunList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TestRun `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TestRun{}, &TestRunList{})
}
