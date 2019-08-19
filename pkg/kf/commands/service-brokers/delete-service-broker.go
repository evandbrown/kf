// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package servicebrokers

import (
	servicecatalogclient "github.com/google/kf/pkg/client/servicecatalog/clientset/versioned"
	"github.com/google/kf/pkg/kf/commands/config"
	"github.com/google/kf/pkg/kf/commands/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/spf13/cobra"
)

// NewDeleteServiceBrokerCommand adds a namespaced service broker to the service catalog.
func NewDeleteServiceBrokerCommand(p *config.KfParams, client servicecatalogclient.Interface) *cobra.Command {
	var (
		serviceBrokerName string
	)

	deleteCmd := &cobra.Command{
		Use:     "delete-service-broker BROKER_NAME",
		Aliases: []string{"dsb"},
		Short:   "Remove a namespaced service broker from service catalog",
		Example: `  kf delete-service-broker mybroker`,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			serviceBrokerName = args[0]

			cmd.SilenceUsage = true

			if err := utils.ValidateNamespace(p); err != nil {
				return err
			}

			err := client.ServicecatalogV1beta1().ServiceBrokers(p.Namespace).Delete(serviceBrokerName, &metav1.DeleteOptions{})

			if err != nil {
				return err
			}

			return nil
		},
	}

	return deleteCmd
}
