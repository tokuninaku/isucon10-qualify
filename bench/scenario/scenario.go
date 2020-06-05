package scenario

import (
	"context"
	"time"

	"github.com/isucon10-qualify/isucon10-qualify/bench/fails"
)

var (
	ExecutionSeconds = 120
)

func Initialize(ctx context.Context) {
	// Initializeにはタイムアウトを設定
	// レギュレーションにある時間を設定する
	// timeoutSeconds := 180

	ctx, cancel := context.WithTimeout(ctx, 180*time.Second)
	defer cancel()

	err := initialize(ctx)
	if err != nil {
		fails.ErrorsForCheck.Add(err)
	}

}

func Validation(ctx context.Context) {
}