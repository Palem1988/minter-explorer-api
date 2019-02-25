package statistics

import (
	"github.com/MinterTeam/minter-explorer-api/chart"
	"github.com/MinterTeam/minter-explorer-api/core"
	"github.com/MinterTeam/minter-explorer-api/errors"
	"github.com/MinterTeam/minter-explorer-api/helpers"
	"github.com/MinterTeam/minter-explorer-api/resource"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type GetTransactionsRequest struct {
	Scale     *string `form:"scale"     binding:"omitempty,eq=minute|eq=hour|eq=day"`
	StartTime *string `form:"startTime" binding:"omitempty,timestamp"`
	EndTime   *string `form:"endTime"   binding:"omitempty,timestamp"`
}

func GetTransactions(c *gin.Context) {
	explorer := c.MustGet("explorer").(*core.Explorer)

	var request GetTransactionsRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		errors.SetValidationErrorResponse(err, c)
		return
	}

	// set default scale
	scale := helpers.DefaultStatisticsScale
	if request.Scale != nil {
		scale = *request.Scale // TODO: validate scale input
	}

	// set default start time
	startTime := time.Now().AddDate(0, 0, helpers.DefaultStatisticsDayDelta).Format("2006-01-02 15:04:05")
	if request.StartTime != nil {
		startTime = *request.StartTime
	}

	// fetch data
	txs := explorer.TransactionRepository.GetTxCountChartDataByFilter(chart.SelectFilter{
		Scale:     scale,
		StartTime: &startTime,
		EndTime:   request.EndTime,
	})

	c.JSON(http.StatusOK, gin.H{
		"data": resource.TransformCollection(txs, chart.TransactionResource{}),
	})
}
