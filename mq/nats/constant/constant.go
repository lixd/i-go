package constant

import "time"

// pub/sub
const (
	ClusterID      = "ND4N3DCKP5XUKBIB7O4MTEB2JURIVSSHV55Z6OQICJR4ZJQHQO22EL7K"
	DurableId      = "durable-test"
	MaxInflight    = 1000
	AckWait        = time.Second * 3
	DefaultNatsURL = "nats://127.0.0.1:4222"
	DefaultSubject = "defaultSubject"
	DefaultQueue   = "defaultQueue"
	DefaultId      = "client_testId"
)
