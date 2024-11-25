package clock

// 永続化操作を行う際の時刻を固定化できるようにするのが目的のパッケージ定義
// Goのtime.Time型はナノ秒単位の制度のため永続化したデータを取得して比較すると時刻情報が不一致になる
// また、現在の時刻がテスト結果に与える影響を無視したい

import (
	"time"
)

type Clocker interface {
	Now() time.Time
}

type RealClocker struct{}

func (r RealClocker) Now() time.Time {
	return time.Now()
}

// テスト用に固定時刻を返す
type FixedClocker struct{}

func (fc FixedClocker) Now() time.Time {
	return time.Date(2022, 5, 10, 12, 34, 56, 0, time.UTC)
}
