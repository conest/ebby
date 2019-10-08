package scene

import "time"

// DeltaTime : Delta Time，每次屏幕刷新之间的时间差
type DeltaTime struct {
	Last time.Time
	Dt   float64
}

// NewDT : 生成新的 Delta 实例
func NewDT() DeltaTime {
	return DeltaTime{Last: time.Now()}
}

// Update : 刷新并返回Delta Time
func (d *DeltaTime) Update() float64 {
	dt := time.Since(d.Last).Seconds()
	d.Dt = dt
	d.Last = time.Now()
	return dt
}
