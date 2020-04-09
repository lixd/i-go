package constant

import "time"

// pub/sub
const (
	ClusterID      = "lixd"
	DurableId      = "durable-lixd"
	MaxInflight    = 1000
	AckWait        = time.Second * 3
	DefaultNatsURL = "nats://127.0.0.1:4222"
	DefaultSubject = "defaultSubject"
	DefaultQueue   = "defaultQueue"
	DefaultId      = "client_testId"
)
