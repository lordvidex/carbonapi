package averageOutsidePercentile

import (
	"testing"
	"time"

	"github.com/go-graphite/carbonapi/expr/helper"
	"github.com/go-graphite/carbonapi/expr/metadata"
	"github.com/go-graphite/carbonapi/expr/types"
	"github.com/go-graphite/carbonapi/pkg/parser"
	th "github.com/go-graphite/carbonapi/tests"
)

func init() {
	md := New("")
	evaluator := th.EvaluatorFromFunc(md[0].F)
	metadata.SetEvaluator(evaluator)
	helper.SetEvaluator(evaluator)
	for _, m := range md {
		metadata.RegisterFunction(m.Name, m.F)
	}
}

// This return is multireturn
func TestAverageOutsidePercentile(t *testing.T) {
	now32 := int64(time.Now().Unix())

	tests := []th.EvalTestItem{
		{
			`averageOutsidePercentile(metric[1234], 30)`,
			map[parser.MetricRequest][]*types.MetricData{
				{Metric: "metric[1234]", From: 0, Until: 1}: {
					types.MakeMetricData("metric1", []float64{7, 7, 7, 7, 7, 7}, 1, now32),
					types.MakeMetricData("metric2", []float64{5, 5, 5, 5, 5, 5}, 1, now32),
					types.MakeMetricData("metric3", []float64{10, 10, 10, 10, 10, 10}, 1, now32),
					types.MakeMetricData("metric4", []float64{1, 1, 1, 1, 1, 1}, 1, now32),
				},
			},
			[]*types.MetricData{
				types.MakeMetricData("metric2", []float64{5, 5, 5, 5, 5, 5}, 1, now32),
				types.MakeMetricData("metric3", []float64{10, 10, 10, 10, 10, 10}, 1, now32),
				types.MakeMetricData("metric4", []float64{1, 1, 1, 1, 1, 1}, 1, now32),
			},
		},
	}

	for _, tt := range tests {
		testName := tt.Target
		t.Run(testName, func(t *testing.T) {
			th.TestEvalExpr(t, &tt)
		})
	}
}
