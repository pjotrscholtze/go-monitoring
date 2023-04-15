package controller

import (
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pjotrscholtze/go-monitoring/cmd/go-monitoring/repo"
	"github.com/pjotrscholtze/go-monitoring/models"
	"github.com/pjotrscholtze/go-monitoring/restapi/operations"
)

func ConnectAPI(api *operations.GoMonitoringAPI, tcr repo.TargetCheckRepo) {
	api.ListAllChecksHandler = operations.ListAllChecksHandlerFunc(func(lacp operations.ListAllChecksParams) middleware.Responder {

		res := make([]*models.Check, 0)
		checks := tcr.List()
		for i, check := range checks {
			lcrError := ""
			if check.Result.Error() != nil {
				lcrError = check.Result.Error().Error()
			}
			lcrLC := check.Result.LastCheck().Format(time.RFC3339)
			lcrMessage := check.Result.Message()
			lcrSuccess := check.Result.Success()
			res = append(res, &models.Check{
				CheckName: &checks[i].Check.Name,
				LastCheckResult: &models.CheckResult{
					Error:     &lcrError,
					LastCheck: &lcrLC,
					Message:   &lcrMessage,
					Success:   &lcrSuccess,
				},
				Schedule:   &check.Check.Schedule,
				TargetName: &check.Target.Name,
			})
		}
		return operations.NewListAllChecksOK().WithPayload(res)
	})
	api.GetTargetChecksHandler = operations.GetTargetChecksHandlerFunc(func(gtcp operations.GetTargetChecksParams) middleware.Responder {
		res := make([]*models.CheckResult, 0)
		list, err := tcr.ListHistory(gtcp.TargetName, gtcp.CheckName)
		if err == nil {
			for _, cr := range list {
				lcrError := ""
				if cr.Result.Error() != nil {
					lcrError = cr.Result.Error().Error()
				}
				lcrLC := cr.Result.LastCheck().Format(time.RFC3339)
				lcrMessage := cr.Result.Message()
				lcrSuccess := cr.Result.Success()
				res = append(res, &models.CheckResult{
					Error:     &lcrError,
					LastCheck: &lcrLC,
					Message:   &lcrMessage,
					Success:   &lcrSuccess,
				})
			}
		}
		return operations.NewGetTargetChecksOK().WithPayload(res)
	})
}
