package order

// Request : 场景执行部分返回，通常需要更换场景
type Request struct {
	Continue     bool
	NextScenario string
	Terminate    bool
}
