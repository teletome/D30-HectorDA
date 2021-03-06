package servers

import (
	"errors"
	"github.com/dminGod/D30-HectorDA/application/alltrade"
	"github.com/dminGod/D30-HectorDA/constant"
	"github.com/dminGod/D30-HectorDA/logger"
	"github.com/dminGod/D30-HectorDA/model"
	"strings"
)

// Routes store the mapping of routes to the underlying application logic
var Routes map[string]func(model.RequestAbstract) model.ResponseAbstract

func init() {

	Routes = map[string]func(model.RequestAbstract) model.ResponseAbstract{

		// All trade version 1
		"alltrade_stock_adjustment_post": alltrade.StockAdjustmentPost,
		"alltrade_stock_adjustment_get":  alltrade.StockAdjustmentGet,

		"alltrade_obtain_detail_post": alltrade.ObtainDetailPost,
		"alltrade_obtain_detail_get":  alltrade.ObtainDetailGet,

		"alltrade_substock_detail_transfer_post": alltrade.SubStockDetailTransferPost,
		"alltrade_substock_detail_transfer_get":  alltrade.SubStockDetailTransferGet,

		"alltrade_substock_daily_detail_post": alltrade.SubStockDailyDetailPost,
		"alltrade_substock_daily_detail_get":  alltrade.SubStockDailyDetailGet,

		"alltrade_transferout_mismatch_post": alltrade.TransferOutMismatchPost,
		"alltrade_transferout_mismatch_get":  alltrade.TransferOutMismatchGet,

		"alltrade_requestgoods_post": alltrade.RequestGoodsPost,
		"alltrade_requestgoods_get":  alltrade.RequestGoodsGet,

		"alltrade_ordertransfer_post": alltrade.OrderTransferPost,
		"alltrade_ordertransfer_get":  alltrade.OrderTransferGet,

		"alltrade_saleout_detail_post": alltrade.SaleOutDetailPost,
		"alltrade_saleout_detail_get":  alltrade.SaleOutDetailGet,

		"alltrade_checkstock_detail_post": alltrade.CheckStockDetailPost,
		"alltrade_checkstock_detail_get":  alltrade.CheckStockDetailGet,

		// Reports
		"alltrade_reports_requestgoods_get": alltrade.ReportsRequestGoodGet,
		"alltrade_reports_adjuststock_get" : alltrade.ReportsAdjustStockGet,
		"alltrade_reports_checkstockdetail_get"  : alltrade.ReportsCheckStockDetailGet,
		"alltrade_reports_directsaledetail_get" : alltrade.ReportsDirectSaleDetailGet,
		"alltrade_reports_directsalesummary_get" : alltrade.ReportsDirectSaleSummaryGet,
		"alltrade_reports_requestgoodssummary_get" : alltrade.ReportsRequestGoodsSummaryGet,
		"alltrade_reports_transferdetail_get" : alltrade.ReportsTransferDetailGet,
		"alltrade_reports_transfersummary_get" : alltrade.ReportsTransferSummaryGet}
}

// HandleRoutes is used resolve incoming routes and execute the corresponding application logic
func HandleRoutes(reqAbs model.RequestAbstract) (model.ResponseAbstract, error) {

	route := GetRouteName(reqAbs)

	// check if the route exists
	if !RouteExists(route) {
		logger.Write("ERROR", "Route for Application: "+reqAbs.Application+", Action: "+reqAbs.Action+", RequestType: "+reqAbs.HTTPRequestType+" not found")
		return model.ResponseAbstract{}, errors.New("Route not found")
	}

	return Routes[route](reqAbs), nil

}

// RouteExists is used to check if a given route exists
// For example:
//  RoutesExists("alltrade_stock_adjustment_post")
// Output:
//  true
func RouteExists(route string) bool {

	// iterate over each route
	for k := range Routes {

		if route == k {
			return true
		}
	}

	return false
}

// GetRouteName is used to return the route mapping as per the naming convention of Hector
func GetRouteName(reqAbs model.RequestAbstract) string {

	route := strings.ToLower(reqAbs.Application + constant.HectorRouteDelimiter + reqAbs.Action + constant.HectorRouteDelimiter + reqAbs.HTTPRequestType)

	return route
}
