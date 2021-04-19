package constant

import "time"

// pub/sub
const (
	ClusterID      = "TestCluster"
	DurableId      = "durable-test"
	MaxInflight    = 1000
	AckWait        = time.Second * 3
	DefaultNatsURL = "nats://192.168.1.6:4222"
	DefaultSubject = "defaultSubject"
	DefaultQueue   = "defaultQueue"
	DefaultId      = "client_testId"
)
