package checks

import (
	"fmt"
	"regexp"
)

// Queues returns monitoring status to return
// 0 = ok
// 1 = warning
// 2 = critical
func (c *Check) Queues(queueNames []string, warning, critical int, ignore bool) int {

	var (
		codeToReturn int
		queues       []Queue
		err          error
	)

	err = c.request("GET", "queues", &queues)
	if err != nil {
		fmt.Println("err:", err)
		fmt.Println("Error connecting to server")
		return 2
	}

	for _, queueName := range queueNames {
		for _, queue := range queues {
			if match(&queue, queueName) {
				codeToReturn = newCodeToReturn(codeToReturn, warning, critical, &queue)
				fmt.Printf("'%s' match '%s' Length: %d\n", queueName, queue.Name, queue.Len)
			} else {
				fmt.Printf("'%s' does not match any queue\n", queueName)
				// if it doesnt match and we won't want to ignore it, it's a critical
				if ignore == false {
					codeToReturn = 2
				}
			}
		}
	}

	return codeToReturn
}

// newCodeToReturn returns the code for this check, maintaining an alert if exists
func newCodeToReturn(current, critical, warning int, queue *Queue) int {
	if queue.Len > critical {
		return 2
	}
	if queue.Len > warning {
		if current < 2 {
			current = 1
		}
	}
	return current
}

// match looks if regular expression match
func match(orig *Queue, rx string) bool {
	var reg = regexp.MustCompile(rx)

	return reg.MatchString(orig.Name)
}

// Queue is the representation of a queue returned by the API
type Queue struct {
	Name                   string   `json:"name"`
	Node                   string   `json:"node"`
	Policy                 string   `json:"policy"`
	SlaveNodes             []string `json:"slave_nodes"`
	Status                 string   `json:"status"`
	SynchronisedSlaveNodes []string `json:"synchronised_slave_nodes"`
	Vhost                  string   `json:"vhost"`
	Len                    int      `json:"backing_queue_status.len"`
}
