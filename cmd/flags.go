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

import "github.com/spf13/cobra"

var (
	server   string
	port     int
	vhost    string
	ssl      bool
	user     string
	password string
)

func initCommonFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&server, "server", "s", "localhost", "server to connect to")
	cmd.PersistentFlags().IntVarP(&port, "port", "p", 15672, "port to connect to")
	cmd.PersistentFlags().StringVarP(&vhost, "vhost", "v", "", "virtual host")
	cmd.PersistentFlags().StringVarP(&user, "user", "", "guest", "user")
	cmd.PersistentFlags().StringVarP(&password, "password", "", "guest", "password")
	cmd.PersistentFlags().BoolVar(&ssl, "ssl", false, "HTTPS server connection")
}
