package signatures

// TaskArg represents a single argument passed to invocation fo a task
type TaskArg struct {
	Type  string
	Value interface{}
}

// TaskSignature represents a single task invocation
type TaskSignature struct {
	UUID       string
	Name       string
	RoutingKey string
	GroupID    string
	Args       []TaskArg
	Immutable  bool
	OnSuccess  []*TaskSignature
	OnError    []*TaskSignature
}

// AdjustRoutingKey makes sure the routing key is correct.
// If the routing key is an empty string:
// a) set it to binding key for direct exchange type
// b) set it to default queue name
func (taskSignature *TaskSignature) AdjustRoutingKey(exchangeType, bindingKey, queueName string) {
	if taskSignature.RoutingKey != "" {
		return
	}

	if exchangeType == "direct" {
		taskSignature.RoutingKey = bindingKey
		return
	}

	taskSignature.RoutingKey = queueName
}
