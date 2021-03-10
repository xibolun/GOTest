package rabbitmq

//type McoResult struct {
//	SenderID    string `json:"senderid" ruby:"senderid"`
//	RequestID   string `json:"requestid" ruby:"requestid"`
//	SenderAgent string `json:"senderagent" ruby:"senderagent"`
//	MsgTime     int64  `json:"msgtime" ruby:"msgtime"`
//	Body        string `json:"body" ruby:"body"`
//	TTL         int    `json:"ttl" ruby:"ttl"`
//	Hash        string `json:"hash" ruby:"hash"`
//}
type McoResult struct {
	SenderID    string        `json:"senderid" ruby:"senderid"`
	RequestID   string        `json:"requestid" ruby:"requestid"`
	SenderAgent string        `json:"senderagent" ruby:"senderagent"`
	MsgTime     int64         `json:"msgtime" ruby:"msgtime"`
	Body        []byte        `json:"body" ruby:"body"`
	Identity    []interface{} `json:"identity" ruby:"identity"`
	Compound    []interface{} `json:"compound" ruby:"compound"`
	CfClass     []interface{} `json:"cf_class" ruby:"cf_class"`
	Fact        []interface{} `json:"fact" ruby:"fact"`
	Agent       string        `json:"agent" ruby:"agent"`
	CallerID    string        `json:"callerid" ruby:"callerid"`
	Collective  string        `json:"collective" ruby:"collective"`
	TTL         int           `json:"ttl" ruby:"ttl"`
	Hash        string        `json:"hash" ruby:"hash"`
}

type scriptMessage struct {
	BodyStr    string `ruby:"body" json:"-"`
	Body       scriptMessageBody
	SenderID   string              `ruby:"senderid"`
	RequestID  string              `ruby:"requestid"`
	Filter     scriptMessageFilter `ruby:"filter"`
	Collective string              `ruby:"collective"`
	Agent      string              `ruby:"agent"`
	CallerID   string              `ruby:"callerid"`
	TTL        int                 `ruby:"ttl"`
	MsgTime    int64               `ruby:"msgtime"`
	Hash       string              `ruby:"hash"`
}

type scriptMessageFilter struct {
	Agent []string `ruby:"agent;key:string"`
	//Collective string   `ruby:"collective"`
}

type scriptMessageBody struct {
	Agent  string                `ruby:"agent"`
	Action string                `ruby:"action"`
	Caller string                `ruby:"caller"`
	Data   scriptMessageBodyData `ruby:"data"`
}

type scriptMessageBodyData struct {
	Type          string `ruby:"type"`
	User          string `ruby:"user"`
	Command       string `ruby:"command"`
	FileName      string `ruby:"filename"`
	Content       string `ruby:"content"`
	Base64        bool   `ruby:"base64"`
	Params        string `ruby:"params"`
	ScriptType    string `ruby:"scriptType"`
	ProcessResult bool   `ruby:"process_result"`
	Environment   string `ruby:"environment"`
}

const psk = "a36cd839414370e10fd281b8a38a4f48"
