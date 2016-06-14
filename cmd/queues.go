// Copyright © 2016 Antonio Fernández <antoniofernandezvara@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"

	"github.com/fernandezvara/rabbitmq-checker/checks"
	"github.com/spf13/cobra"
)

var (
	queueWarning           int
	queueCritical          int
	queueNames             []string
	queueIgnoreIfNotExists bool
)

// QueuesCmd represents the base command when called without any subcommands
var QueuesCmd = &cobra.Command{
	Use:   "queues",
	Short: "Alerting for queues message count",
	Long: `RabbitMQ Health Checker

Returns monitor status for the defined queues and conditions to
raise an alert.`,
	Run: queuesCmd,
}

func init() {
	QueuesCmd.Flags().StringSliceVarP(&queueNames, "queue", "q", []string{".*"}, "queue names to watch")
	QueuesCmd.Flags().BoolVarP(&queueIgnoreIfNotExists, "ignore", "i", false, "ignore non-existent queues")
	QueuesCmd.Flags().IntVarP(&queueWarning, "warning", "w", 100, "messages count threshold (WARNING)")
	QueuesCmd.Flags().IntVarP(&queueCritical, "critical", "c", 100, "messages count threshold (CRITICAL)")
}

func queuesCmd(cmd *cobra.Command, args []string) {
	os.Exit(checks.New(ssl, server, port, vhost, user, password).Queues(queueNames, queueWarning, queueCritical, queueIgnoreIfNotExists))
}
