/*
Copyright 2020 The Flux authors

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

package router

import (
	flaggerv1 "github.com/fluxcd/flagger/pkg/apis/flagger/v1beta1"
)

// KubernetesNoopRouter manages nothing. This is useful when one uses Flagger for progressive delivery of
// services that are not load-balanced by a Kubernetes service
type KubernetesNoopRouter struct {
	primaryWeight int
	canaryWeight int
	mirrored bool
}

func (c *KubernetesNoopRouter) SetRoutes(_ *flaggerv1.Canary, pw int, cw int, m bool) error {
	c.primaryWeight = pw
	c.canaryWeight = cw
	c.mirrored = m

	return nil
}

func (c *KubernetesNoopRouter) GetRoutes(_ *flaggerv1.Canary) (primaryWeight int, canaryWeight int, mirrored bool, err error) {
	return c.primaryWeight, c.canaryWeight, c.mirrored, nil
}

func (c *KubernetesNoopRouter) Initialize(_ *flaggerv1.Canary) error {
	return nil
}

func (c *KubernetesNoopRouter) Reconcile(_ *flaggerv1.Canary) error {
	return nil
}

func (c *KubernetesNoopRouter) Finalize(_ *flaggerv1.Canary) error {
	return nil
}
