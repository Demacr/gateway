// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

package gatewayapi

import (
	"log"

	egv1a1 "github.com/envoyproxy/gateway/api/v1alpha1"
	"github.com/envoyproxy/gateway/internal/gatewayapi/status"
)

func (t *Translator) ProcessVirtualBackends(virtualBackends []*egv1a1.VirtualBackend) {
	// var res []*egv1a1.VirtualBackend
	log.Println("-------------------------------")
	log.Println(virtualBackends)
	log.Println("-------------------------------")

	for _, virtualBackend := range virtualBackends {
		virtualBackend := virtualBackend.DeepCopy()

		// Ensure VirtualBackends are enabled
		if !t.VirtualBackendEnabled {
			status.UpdateVirtualBackendStatusAcceptedCondition(virtualBackend, false,
				"The VirtualBackend was not accepted since VirtualBackend is not enabled in Envoy Gateway Config")
			log.Println("-------------------------------")
			log.Println("false")
			log.Println("-------------------------------")
		} else {
			status.UpdateVirtualBackendStatusAcceptedCondition(virtualBackend, true, "The VirtualBackend was accepted")
			log.Println("-------------------------------")
			log.Println("true")
			log.Println("-------------------------------")
		}

		// res = append(res, virtualBackend)
	}

	// return res
}
