package strdef

// Request : 场景执行部分返回，通常需要更换场景
type Request struct {
	Continue     bool
	NextScenario string
	Terminate    bool
	ResetData    bool
}

// DefaultRequest : 返回默认 Request，执行执行excute
var DefaultRequest = Request{
	Continue:     true,
	NextScenario: "",
	Terminate:    false,
	ResetData:    true,
}
