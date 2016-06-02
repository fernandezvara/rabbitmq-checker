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

package main

// Queue is the representation of a queue returned by the API
type Queue struct {
	Arguments          struct{} `json:"arguments"`
	AutoDelete         bool     `json:"auto_delete"`
	BackingQueueStatus struct {
		AvgAckEgressRate  float64       `json:"avg_ack_egress_rate"`
		AvgAckIngressRate float64       `json:"avg_ack_ingress_rate"`
		AvgEgressRate     float64       `json:"avg_egress_rate"`
		AvgIngressRate    float64       `json:"avg_ingress_rate"`
		Delta             []interface{} `json:"delta"`
		Len               int           `json:"len"`
		MirrorSeen        int           `json:"mirror_seen"`
		MirrorSenders     int           `json:"mirror_senders"`
		NextSeqID         int           `json:"next_seq_id"`
		PendingAcks       int           `json:"pending_acks"`
		PersistentCount   int           `json:"persistent_count"`
		Q1                int           `json:"q1"`
		Q2                int           `json:"q2"`
		Q3                int           `json:"q3"`
		Q4                int           `json:"q4"`
		RAMAckCount       int           `json:"ram_ack_count"`
		RAMMsgCount       int           `json:"ram_msg_count"`
		TargetRAMCount    string        `json:"target_ram_count"`
	} `json:"backing_queue_status"`
	Consumers            int    `json:"consumers"`
	Durable              bool   `json:"durable"`
	ExclusiveConsumerTag string `json:"exclusive_consumer_tag"`
	IdleSince            string `json:"idle_since"`
	Memory               int    `json:"memory"`
	MessageStats         struct {
		Ack        int `json:"ack"`
		AckDetails struct {
			Rate int `json:"rate"`
		} `json:"ack_details"`
		Deliver        int `json:"deliver"`
		DeliverDetails struct {
			Rate int `json:"rate"`
		} `json:"deliver_details"`
		DeliverGet        int `json:"deliver_get"`
		DeliverGetDetails struct {
			Rate int `json:"rate"`
		} `json:"deliver_get_details"`
		Publish        int `json:"publish"`
		PublishDetails struct {
			Rate int `json:"rate"`
		} `json:"publish_details"`
	} `json:"message_stats"`
	Name                   string   `json:"name"`
	Node                   string   `json:"node"`
	Policy                 string   `json:"policy"`
	SlaveNodes             []string `json:"slave_nodes"`
	Status                 string   `json:"status"`
	SynchronisedSlaveNodes []string `json:"synchronised_slave_nodes"`
	Vhost                  string   `json:"vhost"`
}
