package rabbitmq

type McoResult struct {
	SenderID    string `json:"senderid" ruby:"senderid"`
	RequestID   string `json:"requestid" ruby:"requestid"`
	SenderAgent string `json:"senderagent" ruby:"senderagent"`
	MsgTime     int64  `json:"msgtime" ruby:"msgtime"`
	Body        string `json:"body" ruby:"body"`
	TTL         int    `json:"ttl" ruby:"ttl"`
	Hash        string `json:"hash" ruby:"hash"`
}
