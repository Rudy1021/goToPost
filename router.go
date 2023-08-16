//go:build debug

package router

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"pds/util"

	bomRoute "pds/controller/routes/ADM/BDM/Bom"
	oemRoute "pds/controller/routes/ADM/BDM/OEM"
	barcodeRoute "pds/controller/routes/ADM/BDM/Product/Barcode"
	docRoute "pds/controller/routes/ADM/BDM/Product/Doc"
	moldRoute "pds/controller/routes/ADM/EDM/Master/Mold"

	biRouteGetPN "pds/controller/routes/BI"
	biRouteInstockCurrent "pds/controller/routes/BI/InstockOkRate/current"
	biRouteInstockMOM "pds/controller/routes/BI/InstockOkRate/monthOnMonth"
	biRoutePdtwhmCurrent "pds/controller/routes/BI/pdtwhm/current"

	costRoute "pds/controller/routes/ADM/BDM/Product/Cost"
	dealRoute "pds/controller/routes/ADM/BDM/Product/Deal"
	extraRoute "pds/controller/routes/ADM/BDM/Product/Extra"
	priceRoute "pds/controller/routes/ADM/BDM/Product/Price"
	stockRoute "pds/controller/routes/ADM/BDM/Product/Stock"

	productRoute "pds/controller/routes/ADM/BDM/Product/Main"
	packRoute "pds/controller/routes/ADM/BDM/Product/Pack"

	sipRouteHD "pds/controller/routes/ADM/BDM/SIP/HD"
	sipRouteTB "pds/controller/routes/ADM/BDM/SIP/TB"
	sopRoute "pds/controller/routes/ADM/BDM/SOP"
	sopRouteHD "pds/controller/routes/ADM/BDM/SOP/HD"
	sopRouteREC "pds/controller/routes/ADM/BDM/SOP/Rec"
	sopRouteTB "pds/controller/routes/ADM/BDM/SOP/TB"
	wpcRouteHD "pds/controller/routes/ADM/BDM/WPC/HD"
	wpcRouteTB "pds/controller/routes/ADM/BDM/WPC/TB"
	addRoute "pds/controller/routes/ADM/SDM/Add"
	contactRoute "pds/controller/routes/ADM/SDM/Contact"
	currencyRoute "pds/controller/routes/ADM/SDM/Currency"
	workSegmentRoute "pds/controller/routes/ADM/SDM/WorkSegment"

	custRouteHD "pds/controller/routes/ADM/SDM/Customer/HD"
	custRouteTB "pds/controller/routes/ADM/SDM/Customer/TB"
	driRoute "pds/controller/routes/ADM/SDM/Dri"
	mgmtcatRoute "pds/controller/routes/ADM/SDM/Mgmtcat"
	mgmtcatCustRoute "pds/controller/routes/ADM/SDM/MgmtcatCust"
	ngRoute "pds/controller/routes/ADM/SDM/Ng"
	oemItemRoute "pds/controller/routes/ADM/SDM/OEM"
	offlineRoute "pds/controller/routes/ADM/SDM/Offline"
	okRoute "pds/controller/routes/ADM/SDM/Ok"
	omsPermissionRoute "pds/controller/routes/ADM/SDM/OmsPermission"
	pauseItemRoute "pds/controller/routes/ADM/SDM/Pause"
	permissionRoute "pds/controller/routes/ADM/SDM/Permission"
	pmsPermissionRoute "pds/controller/routes/ADM/SDM/PmsPermission"
	proccessFlowRoute "pds/controller/routes/ADM/SDM/ProccessFlow"
	productCostCenterRoute "pds/controller/routes/ADM/SDM/ProductCostCenter"
	productWorkCenterRoute "pds/controller/routes/ADM/SDM/ProductWorkCenter"
	productionShopRoute "pds/controller/routes/ADM/SDM/ProductionShop"
	roleRoute "pds/controller/routes/ADM/SDM/Role"
	staffDataRoute "pds/controller/routes/ADM/SDM/StaffData"
	suppRouteHD "pds/controller/routes/ADM/SDM/Supplier/HD"
	suppRouteTB "pds/controller/routes/ADM/SDM/Supplier/TB"
	wareHouseRoute "pds/controller/routes/ADM/SDM/WareHouse"
	wmsPermissionRoute "pds/controller/routes/ADM/SDM/WmsPermission"
	workInspectRoute "pds/controller/routes/ADM/SDM/WorkInspection"
	workLineRoute "pds/controller/routes/ADM/SDM/WorkLine"
	workStationRoute "pds/controller/routes/ADM/SDM/WorkStation"
	wosPermissionRoute "pds/controller/routes/ADM/SDM/WosPermission"
	corpdataRoute "pds/controller/routes/ADM/TDM/CorpData"
	currencyDataRoute "pds/controller/routes/ADM/TDM/CurrencyData"
	deptDataRoute "pds/controller/routes/ADM/TDM/DeptData"
	funcDataRoute "pds/controller/routes/ADM/TDM/FuncData"
	notifyRoute "pds/controller/routes/ADM/TDM/Notify"
	opRouteHD "pds/controller/routes/ADM/TDM/OpData/HD"
	opRouteTB "pds/controller/routes/ADM/TDM/OpData/TB"
	opUnitRoute "pds/controller/routes/ADM/TDM/OpUnit"
	postDataRoute "pds/controller/routes/ADM/TDM/PostData"
	pushDataRoute "pds/controller/routes/ADM/TDM/PushData"
	unitRoute "pds/controller/routes/ADM/TDM/Unit"
	authRoute "pds/controller/routes/Auth"
	initPageRoute "pds/controller/routes/InitPage"
	abcRoute "pds/controller/routes/PDS/ABC"
	apsRoute "pds/controller/routes/PDS/APS"
	dashboardRoute "pds/controller/routes/PDS/Dashboard"
	inpdtRoute "pds/controller/routes/PDS/ProcessData/Inpdt"
	wonSearchRoute "pds/controller/routes/PDS/ProcessData/WoNSearch"
	wpcSearchRoute "pds/controller/routes/PDS/ProcessData/WpcSearch"
	actRoute "pds/controller/routes/PDS/Table/ACT"
	efficiencyRoute "pds/controller/routes/PDS/Table/Efficiency"
	tableHeadline "pds/controller/routes/PDS/Table/Headline"
	spcRoute "pds/controller/routes/SPC"
	selectorRoute "pds/controller/routes/Selector"

	apsOutHd "pds/controller/routes/APS/Mo/Hd"
	apsOutTb "pds/controller/routes/APS/Mo/Tb"

	rcpHdRoute "pds/controller/routes/QMS/Rcp/Hd"
	rcpTbRoute "pds/controller/routes/QMS/Rcp/Tb"
	QmsSipHd "pds/controller/routes/QMS/Sip/Hd"
	QmsSipHdImg "pds/controller/routes/QMS/Sip/Hd/Docs"

	QmsSipTb "pds/controller/routes/QMS/Sip/Tb"
	QmsSipTbImg "pds/controller/routes/QMS/Sip/Tb/Docs"
	QmsSopHd "pds/controller/routes/QMS/Sop/Hd"
	QmsSopRec "pds/controller/routes/QMS/Sop/Rec"
	QmsSopTb "pds/controller/routes/QMS/Sop/Tb"

	apsOutWo "pds/controller/routes/APS/Wo"
	wmsRoute "pds/controller/routes/WMS"
	inventoryRoute "pds/controller/routes/WMS/Inventory"
	invBatchRoute "pds/controller/routes/WMS/Inventory/batch"
	invMonthRoute "pds/controller/routes/WMS/Inventory/month"
	invPriceRoute "pds/controller/routes/WMS/Inventory/price"
	invRouteHD "pds/controller/routes/WMS/Invproject/Hd"
	invProjectRoute "pds/controller/routes/WMS/Invproject/Project"
	invRouteTB "pds/controller/routes/WMS/Invproject/Tb"
	openRoute "pds/controller/routes/WMS/Open"
	pmsOutRoute "pds/controller/routes/WMS/PMS"
	pdtUseRoute "pds/controller/routes/WMS/PdtUse"
	reportRoute "pds/controller/routes/WMS/Report"
	run8090HdRoute "pds/controller/routes/WMS/run80run90/HD"
	run8090TbRoute "pds/controller/routes/WMS/run80run90/TB"

	lmsHistoryRoute "pds/controller/routes/LMS/History"
	lmsLabelRoute "pds/controller/routes/LMS/Label"
	lmsModelHdRoute "pds/controller/routes/LMS/Model/Hd"
	lmsModelPdtRoute "pds/controller/routes/LMS/Model/Pdt"
	lmsModelTbRoute "pds/controller/routes/LMS/Model/Tb"
	lmsPrintRoute "pds/controller/routes/LMS/Print"
	lmsPrinterRoute "pds/controller/routes/LMS/Printer"
	lmsTimecodeHdRoute "pds/controller/routes/LMS/Timecode/Hd"
	lmsTimecodeTbRoute "pds/controller/routes/LMS/Timecode/Tb"

	middleware "pds/middleware"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func Setup_Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(middleware.Cors())
	fmt.Println("AAAAA")
	g1 := router.Group("api")
	{
		// 下拉選單
		g1.GET("/getFunctions", selectorRoute.GetFunctions)
		g1.GET("/getPushlvl", selectorRoute.GetPushlvl)
		g1.GET("/getAllWsID", selectorRoute.GetAllWsID)
		g1.GET("/getWsWl", selectorRoute.GetWsWlAPI)

		// 登入/登出
		g1.POST("/login", authRoute.Login)
		g1.POST("/logout", authRoute.Logout)

		// initPage
		g1.GET("/getUserPerm", initPageRoute.GetUserPerm)

		// -------------------------------------系統資料維護-------------------------------------
		// 員工資料維護
		g1.GET("/getStaff", staffDataRoute.GetStaff)
		g1.POST("/addStaff", staffDataRoute.AddStaff)
		g1.POST("/freezeStaff", staffDataRoute.FreezeStaff)
		g1.POST("/modifyStaff", staffDataRoute.ModifyStaff)

		// 系統登入權限
		g1.GET("/getStaffPerm", permissionRoute.Get)
		g1.POST("/addStaffPerm", permissionRoute.Add)
		g1.POST("/deleteStaffPerm", permissionRoute.Delete)
		g1.POST("/modifyStaffPerm", permissionRoute.Update)
		g1.POST("/changeStaffPerm", permissionRoute.Change)

		// 報工作業權限
		g1.GET("/getWosID", wosPermissionRoute.GetStaffID)
		g1.POST("/addWosPerm", wosPermissionRoute.AddWosPerm)
		g1.POST("/deleteWosPerm", wosPermissionRoute.DeleteWosPerm)
		g1.POST("/modifyWosPerm", wosPermissionRoute.ModifyWosPerm)

		// 倉儲作業權限
		g1.GET("/getWmsPerm", wmsPermissionRoute.GetWmsPerm)
		g1.POST("/addWmsPerm", wmsPermissionRoute.AddWmsPerm)
		g1.POST("/deleteWmsPerm", wmsPermissionRoute.DeleteWmsPerm)
		g1.POST("/modifyWmsPerm", wmsPermissionRoute.ModifyWmsPerm)

		// 採購作業權限
		g1.GET("/getPmsID", pmsPermissionRoute.GetPmsID)
		g1.POST("/addPmsPerm", pmsPermissionRoute.AddPmsPerm)
		g1.POST("/deletePmsPerm", pmsPermissionRoute.DeletePmsPerm)
		g1.POST("/modifyPmsPerm", pmsPermissionRoute.ModifyPmsPerm)

		// 訂購作業權限
		g1.GET("/getOmsID", omsPermissionRoute.GetOmsID)
		g1.POST("/addOmsPerm", omsPermissionRoute.AddOmsPerm)
		g1.POST("/deleteOmsPerm", omsPermissionRoute.DeleteOmsPerm)
		g1.POST("/modifyOmsPerm", omsPermissionRoute.ModifyOmsPerm)

		// 管理分類設定
		g1.GET("/getMgmtcat", mgmtcatRoute.GetMgmtcat)
		g1.POST("/addMgmtcat", mgmtcatRoute.AddMgmtcat)
		g1.POST("/deleteMgmtcat", mgmtcatRoute.DeleteMgmtcat)
		g1.POST("/modifyMgmtcat", mgmtcatRoute.ModifyMgmtcat)

		// 管理分類設定
		g1.GET("/getMgmtcatCust", mgmtcatCustRoute.GetMgmtcatCust)
		g1.POST("/addMgmtcatCust", mgmtcatCustRoute.AddMgmtcatCust)
		g1.POST("/deleteMgmtcatCust", mgmtcatCustRoute.DeleteMgmtcatCust)
		g1.POST("/modifyMgmtcatCust", mgmtcatCustRoute.ModifyMgmtcatCust)

		// 聯絡資料管理
		g1.GET("/getContact", contactRoute.GetContact)
		g1.POST("/addContact", contactRoute.AddContact)
		g1.POST("/deleteContact", contactRoute.DeleteContact)
		g1.POST("/modifyContact", contactRoute.ModifyContact)

		// 檢測項目設定
		g1.GET("/getWi", workInspectRoute.Get)
		g1.POST("/addWi", workInspectRoute.Add)
		g1.POST("/deleteWi", workInspectRoute.Delete)
		g1.POST("/modifyWi", workInspectRoute.Modify)

		// 角色設定
		g1.GET("/getRole", roleRoute.Get)
		g1.POST("/addRole", roleRoute.Add)
		g1.POST("/deleteRole", roleRoute.Delete)
		g1.POST("/updateRole", roleRoute.Update)
		// -------------------------------------基本資料維護-------------------------------------
		// 產品
		g1.GET("/getProductSelect", productRoute.GetProductSelect)
		g1.GET("/getProduct", productRoute.GetProduct)
		g1.POST("/addProduct", productRoute.AddProduct)
		g1.POST("/deleteProduct", productRoute.DeleteProduct)
		g1.POST("/modifyProduct", productRoute.ModifyProduct)
		// 交易料號查詢
		g1.GET("/searchDeal", productRoute.GetDeal)

		// 識別
		g1.GET("/getIdkw", barcodeRoute.GetIdkw)
		g1.POST("/addIdkw", barcodeRoute.AddIdkw)
		g1.POST("/deleteIdkw", barcodeRoute.DeleteIdkw)
		g1.POST("/modifyIdkw", barcodeRoute.ModifyIdkw)

		// 包裝
		g1.GET("/getPack", packRoute.GetPack)
		g1.POST("/addPack", packRoute.AddPack)
		g1.POST("/deletePack", packRoute.DeletePack)
		g1.POST("/modifyPack", packRoute.ModifyPack)

		// 標準成本
		g1.GET("/getCost", costRoute.Get)
		g1.POST("/addCost", costRoute.Add)
		g1.POST("/deleteCost", costRoute.Delete)
		g1.POST("/updateCost", costRoute.Update)

		// 庫存條件
		g1.GET("/getStock", stockRoute.Get)
		g1.POST("/addStock", stockRoute.Add)
		g1.POST("/deleteStock", stockRoute.Delete)
		g1.POST("/updateStock", stockRoute.Update)

		// 庫存料號
		g1.GET("/getDeal", dealRoute.Get)
		g1.POST("/addDeal", dealRoute.Add)
		g1.POST("/deleteDeal", dealRoute.Delete)
		g1.POST("/updateDeal", dealRoute.Update)

		// 標準價格
		g1.GET("/getPrice", priceRoute.Get)
		g1.POST("/addPrice", priceRoute.Add)
		g1.POST("/deletePrice", priceRoute.Delete)
		g1.POST("/updatePrice", priceRoute.Update)

		// 標準尺寸
		g1.GET("/getExtra", extraRoute.Get)
		g1.POST("/addExtra", extraRoute.Add)
		g1.POST("/deleteExtra", extraRoute.Delete)
		g1.POST("/updateExtra", extraRoute.Update)

		// 途程表頭
		g1.POST("/addWpcHd", wpcRouteHD.Add)
		g1.GET("/getWpcHd", wpcRouteHD.Get)
		g1.POST("/updateWpcHd", wpcRouteHD.Update)
		g1.POST("/deleteWpcHd", wpcRouteHD.Delete)

		// 途程表身
		g1.POST("/addWpcTb", wpcRouteTB.Add)
		g1.GET("/getWpcTb", wpcRouteTB.Get)
		g1.POST("/updateWpcTb", wpcRouteTB.Update)
		g1.POST("/deleteWpcTb", wpcRouteTB.Delete)
		g1.POST("/editWpcTb", wpcRouteTB.Edit)

		// 途程下拉選單
		g1.GET("/getWpcIe", selectorRoute.GetWpcIe)

		// 工程資料管理
		g1.POST("/addDoc", docRoute.Add)
		g1.GET("/getDoc", docRoute.Get)
		g1.POST("/updateDoc", docRoute.Update)

		// 製成(工序)主檔
		g1.GET("/getPF", proccessFlowRoute.Get)
		g1.POST("/addPF", proccessFlowRoute.Add)
		g1.POST("/updatePF", proccessFlowRoute.Update)
		g1.POST("/deletePF", proccessFlowRoute.Delete)

		// 生產成本中心
		g1.GET("/getPCC", productCostCenterRoute.Get)
		g1.POST("/addPCC", productCostCenterRoute.Add)
		g1.POST("/updatePCC", productCostCenterRoute.Update)
		g1.POST("/deletePCC", productCostCenterRoute.Delete)

		// 生產工作中心
		g1.GET("/getPWC", productWorkCenterRoute.Get)
		g1.POST("/addPWC", productWorkCenterRoute.Add)
		g1.POST("/updatePWC", productWorkCenterRoute.Update)
		g1.POST("/deletePWC", productWorkCenterRoute.Delete)

		// BOM
		g1.POST("/getTreeStruct", bomRoute.GetTreeStruct)
		g1.GET("/getArrayStruct", bomRoute.GetArrayStruct)
		g1.POST("/getDeleteArrayStruct", bomRoute.GetDeleteArrayStruct)
		g1.GET("/getBom", bomRoute.GetBom)
		g1.POST("/addBom", bomRoute.AddBom)
		g1.POST("/modifyBom", bomRoute.ModifyBom)
		g1.POST("/deleteBom", bomRoute.DeleteBom)

		// 車間
		g1.GET("/getPS", productionShopRoute.GetPS)
		g1.POST("/addPS", productionShopRoute.AddPS)
		g1.POST("/deletePS", productionShopRoute.DeletePS)
		g1.POST("/modifyPS", productionShopRoute.ModifyPS)

		// 工站
		g1.GET("/getStation", workStationRoute.Get)
		g1.POST("/addStation", workStationRoute.Add)
		g1.POST("/deleteStation", workStationRoute.Delete)
		g1.POST("/updateStation", workStationRoute.Update)

		// 工線
		g1.GET("/getLine", workLineRoute.GetLine)
		g1.POST("/addLine", workLineRoute.AddLine)
		g1.POST("/deleteLine", workLineRoute.DeleteLine)
		g1.POST("/modifyLine", workLineRoute.ModifyLine)

		// 工段
		g1.GET("/getWg", workSegmentRoute.Get)
		g1.POST("/addWg", workSegmentRoute.Add)
		g1.POST("/deleteWg", workSegmentRoute.Delete)
		g1.POST("/updateWg", workSegmentRoute.Update)

		// 倉別
		g1.GET("/getWh", wareHouseRoute.GetWh)
		g1.POST("/addWh", wareHouseRoute.AddWh)
		g1.POST("/deleteWh", wareHouseRoute.DeleteWh)
		g1.POST("/modifyWh", wareHouseRoute.ModifyWh)
		g1.POST("/closeWh", wareHouseRoute.Close)

		// 委外基本資料
		g1.GET("/getOEM", oemRoute.Get)
		g1.POST("/deleteOEM", oemRoute.Delete)
		g1.POST("/updateOEM", oemRoute.Update)
		g1.POST("/addOEM", oemRoute.Add)

		// 包裝單位設定 - Ariean
		g1.GET("/getOpUnit", opUnitRoute.Get)
		g1.POST("/addOpUnit", opUnitRoute.Add)
		g1.POST("/modifyOpUnit", opUnitRoute.Update)
		g1.POST("/deleteOpUnit", opUnitRoute.Delete)

		// 計量單位設定
		g1.GET("/getUnit", unitRoute.Get)
		g1.POST("/addUnit", unitRoute.Add)
		g1.POST("/modifyUnit", unitRoute.Update)
		g1.POST("/deleteUnit", unitRoute.Delete)

		// 推播資料設定
		g1.GET("/getNotify", notifyRoute.Get)
		g1.POST("/addNotify", notifyRoute.Add)
		g1.POST("/modifyNotify", notifyRoute.Update)
		g1.POST("/deleteNotify", notifyRoute.Delete)

		// -------------------------------------功能字串管理-------------------------------------
		// 良品項目
		g1.GET("/getOkItem", okRoute.Get)
		g1.POST("/deleteOkItem", okRoute.Delete)
		g1.POST("/updateOkItem", okRoute.Update)
		g1.POST("/addOkItem", okRoute.Add)

		// 不良品項目
		g1.GET("/getNgItem", ngRoute.Get)
		g1.POST("/deleteNgItem", ngRoute.Delete)
		g1.POST("/updateNgItem", ngRoute.Update)
		g1.POST("/addNgItem", ngRoute.Add)

		// 停工復工項目設定
		g1.GET("/getPause", pauseItemRoute.Get)
		g1.POST("/deletePause", pauseItemRoute.Delete)
		g1.POST("/updatePause", pauseItemRoute.Update)
		g1.POST("/addPause", pauseItemRoute.Add)

		// 責任單位設定
		g1.GET("/getDri", driRoute.Get)
		g1.POST("/deleteDri", driRoute.Delete)
		g1.POST("/updateDri", driRoute.Update)
		g1.POST("/addDri", driRoute.Add)

		// 客戶資料
		g1.GET("/getCustHd", custRouteHD.GetCustHd)
		g1.POST("/addCustHd", custRouteHD.AddCustHd)
		g1.POST("/modifyCustHd", custRouteHD.ModifyCustHd)
		g1.POST("/deleteCustHd", custRouteHD.DeleteCustHd)
		g1.GET("/getCustTb", custRouteTB.GetCustTb)
		g1.POST("/addCustTb", custRouteTB.AddCustTb)
		g1.POST("/modifyCustTb", custRouteTB.ModifyCustTb)
		g1.POST("/deleteCustTb", custRouteTB.DeleteCustTb)

		// 廠商資料
		g1.GET("/getSuppHd", suppRouteHD.GetSuppHd)
		g1.POST("/addSuppHd", suppRouteHD.AddSuppHd)
		g1.POST("/modifySuppHd", suppRouteHD.ModifySuppHd)
		g1.POST("/deleteSuppHd", suppRouteHD.DeleteSuppHd)
		g1.GET("/getSuppTb", suppRouteTB.GetSuppTb)
		g1.POST("/addSuppTb", suppRouteTB.AddSuppTb)
		g1.POST("/modifySuppTb", suppRouteTB.ModifySuppTb)
		g1.POST("/deleteSuppTb", suppRouteTB.DeleteSuppTb)

		// 委外項目資料
		g1.GET("/getOemitem", oemItemRoute.Get)
		g1.POST("/addOemitem", oemItemRoute.Add)
		g1.POST("/updateOemitem", oemItemRoute.Update)
		g1.POST("/deleteOemitem", oemItemRoute.Delete)

		// 幣別項目資料
		g1.GET("/getCurrency", currencyRoute.Get)
		g1.POST("/deleteCurrency", currencyRoute.Delete)
		g1.POST("/modifyCurrency", currencyRoute.Update)
		g1.POST("/addCurrency", currencyRoute.Add)
		g1.POST("/sortCurrency", currencyRoute.Sort)

		// 附加項目資料
		g1.GET("/getAdd", addRoute.GetAdd)
		g1.POST("/deleteAdd", addRoute.DeleteAdd)
		g1.POST("/modifyAdd", addRoute.ModifyAdd)
		g1.POST("/addAdd", addRoute.AddAdd)

		// 下線項目設定
		g1.GET("/getOffline", offlineRoute.Get)
		g1.POST("/deleteOffline", offlineRoute.Delete)
		g1.POST("/updateOffline", offlineRoute.Update)
		g1.POST("/addOffline", offlineRoute.Add)

		// -------------------------------------生產資訊查詢-------------------------------------
		// 工站途單查詢
		g1.GET("/getWpcSearch", wpcSearchRoute.Get)

		// 工單履歷查詢
		g1.GET("/getWoNSearch", wonSearchRoute.Get)
		g1.GET("/getWoNDetail", wonSearchRoute.Detail)
		g1.GET("getAllResumeApp", wonSearchRoute.GetAllResumeApp)

		// 工站在製查詢
		g1.GET("/getInpdt", inpdtRoute.Get)

		// -------------------------------------生管派工-------------------------------------
		g1.GET("/getAps", apsRoute.GetAllWpc)
		g1.POST("/modifyAps", apsRoute.ModifyAps)
		g1.POST("/toWos", apsRoute.ToWos)
		g1.POST("/apsClose", apsRoute.Close)

		// -----------------------------------技術資料管理-----------------------------------
		// 部門資料管理
		g1.POST("/addDept", deptDataRoute.Add)
		g1.POST("/deleteDept", deptDataRoute.Delete)
		g1.GET("/getDept", deptDataRoute.Get)
		g1.POST("/modifyDept", deptDataRoute.Update)
		// 推播層級管理
		g1.POST("/addPush", pushDataRoute.Add)
		g1.POST("/deletePush", pushDataRoute.Delete)
		g1.GET("/getPush", pushDataRoute.Get)
		g1.POST("/modifyPush", pushDataRoute.Update)
		// 職稱資料管理
		g1.POST("/addPost", postDataRoute.Add)
		g1.POST("/deletePost", postDataRoute.Delete)
		g1.GET("/getPost", postDataRoute.Get)
		g1.POST("/modifyPost", postDataRoute.Update)
		// 組織功能管理
		g1.POST("/addFunc", funcDataRoute.Add)
		g1.POST("/deleteFunc", funcDataRoute.Delete)
		g1.GET("/getFunc", funcDataRoute.Get)
		g1.POST("/modifyFunc", funcDataRoute.Update)

		// 作業科目管理(表頭)
		g1.POST("/addOpHd", opRouteHD.Add)
		g1.GET("/getOpHd", opRouteHD.Get)
		g1.POST("/modifyOpHd", opRouteHD.Update)
		g1.POST("/deleteOpHd", opRouteHD.Delete)

		// 作業科目管理(表身)
		g1.POST("/addOpTb", opRouteTB.Add)
		g1.GET("/getOpTb", opRouteTB.Get)
		g1.POST("/modifyOpTb", opRouteTB.Update)
		g1.POST("/deleteOpTb", opRouteTB.Delete)

		// 廠別設定
		g1.POST("/addCorp", corpdataRoute.Add)
		g1.GET("/getCorp", corpdataRoute.Get)
		g1.POST("/modifyCorp", corpdataRoute.Update)
		g1.POST("/deleteCorp", corpdataRoute.Delete)

		// 公司幣別資料
		g1.GET("/getCurrencyData", currencyDataRoute.Get)
		g1.POST("/addCurrencyData", currencyDataRoute.Add)
		g1.POST("/modifyCurrencyData", currencyDataRoute.Update)
		g1.POST("/deleteCurrencyData", currencyDataRoute.Delete)

		// SOP表頭
		g1.GET("/getAllSophd", sopRouteHD.GetAllSophd)
		g1.POST("/addSophd", sopRouteHD.AddSophd)
		g1.POST("/deleteSophd", sopRouteHD.DeleteSophd)
		g1.POST("/modifySophd", sopRouteHD.ModifySophd)

		// SOP表身
		g1.GET("/getAllSoptb", sopRouteTB.GetAllSoptb)
		g1.POST("/modifySoptb", sopRouteTB.ModifySoptb)

		// SOP表尾
		g1.GET("/getAllSoprec", sopRouteREC.GetAllSoprec)
		g1.POST("/modifySoprec", sopRouteREC.ModifySoprec)

		// SIP表頭
		g1.GET("/getAllSiphd", sipRouteHD.GetAllSiphd)
		g1.POST("/addSiphd", sipRouteHD.AddSiphd)
		g1.POST("/deleteSiphd", sipRouteHD.DeleteSiphd)
		g1.POST("/modifySiphd", sipRouteHD.ModifySiphd)

		// SIP表身
		g1.GET("/getAllSiptb", sipRouteTB.GetAllSiptb)
		g1.POST("/addSiptb", sipRouteTB.AddSiptb)
		g1.POST("/deleteSiptb", sipRouteTB.DeleteSiptb)
		g1.POST("/modifySiptb", sipRouteTB.ModifySiptb)

		// RCP表頭
		g1.GET("/getRcphd", rcpHdRoute.Get)
		g1.POST("/addRcphd", rcpHdRoute.Add)
		g1.POST("/deleteRcphd", rcpHdRoute.Delete)
		g1.POST("/modifyRcphd", rcpHdRoute.Update)

		// RCP表身
		g1.GET("/getRcptb", rcpTbRoute.Get)
		g1.POST("/addRcptb", rcpTbRoute.Add)
		g1.POST("/deleteRcptb", rcpTbRoute.Delete)
		g1.POST("/modifyRcptb", rcpTbRoute.Update)

		// SOP列印
		g1.GET("/printPosition", sopRoute.PrintPosition)

		// 據此條目創建
		g1.POST("/copySop", sopRoute.CopySop)

		// -------------------------------------資產管理系統-------------------------------------
		// 模具資料管理
		g1.GET("/getMold", moldRoute.Get)
		g1.POST("/addMold", moldRoute.Add)
		g1.POST("/deleteMold", moldRoute.Delete)
		g1.POST("/updateMold", moldRoute.Update)
		g1.POST("/uploadMold", moldRoute.Upload)
		g1.POST("/deleteUploadMold", moldRoute.DeleteUpload)

		// -------------------------------------APS-------------------------------------
		// 工單
		g1.GET("/getApsOutWo", apsOutWo.Get)
		g1.GET("/getApsOutWoSelect", apsOutWo.GetSelect)
		g1.POST("/addApsOutWo", apsOutWo.Add)
		g1.POST("/updateApsOutWo", apsOutWo.Update)
		g1.POST("/deleteApsOutWo", apsOutWo.Delete)

		// 料單表頭
		g1.GET("/getApsOutMoHd", apsOutHd.Get)
		g1.POST("/addApsOutMoHd", apsOutHd.Add)
		g1.POST("/updateApsOutMoHd", apsOutHd.Update)
		g1.POST("/deleteApsOutMoHd", apsOutHd.Delete)

		// 料單表身
		g1.GET("/getApsOutMoTb", apsOutTb.Get)
		g1.POST("/addApsOutMoTb", apsOutTb.Add)
		g1.POST("/updateApsOutMoTb", apsOutTb.Update)
		g1.POST("/deleteApsOutMoTb", apsOutTb.Delete)

		g1.GET("/getQmsSopHd", QmsSopHd.Get)
		g1.POST("/addQmsSopHd", QmsSopHd.Add)
		g1.POST("/updateQmsSopHd", QmsSopHd.Update)
		g1.POST("/deleteQmsSopHd", QmsSopHd.Delete)
		g1.POST("/copyQmsSop", QmsSopHd.CopySop)

		g1.GET("/getQmsSopTb", QmsSopTb.Get)
		// g1.POST("/addQmsSopTb", QmsSopTb.Add)
		g1.POST("/updateQmsSopTb", QmsSopTb.Update)
		// g1.POST("/deleteQmsSopTb", QmsSopTb.Delete)

		g1.GET("/getQmsSopRec", QmsSopRec.Get)
		g1.POST("/updateQmsSopRec", QmsSopRec.Update)

		g1.GET("/getQmsSipHd", QmsSipHd.Get)
		g1.POST("/addQmsSipHd", QmsSipHd.Add)
		g1.POST("/addQmsSipHdImg", QmsSipHdImg.Add)
		g1.POST("/updateQmsSipHd", QmsSipHd.Update)
		g1.POST("/deleteQmsSipHd", QmsSipHd.Delete)

		g1.GET("/getQmsSipTb", QmsSipTb.Get)
		g1.POST("/addQmsSipTb", QmsSipTb.Add)
		g1.POST("/updateQmsSipTb", QmsSipTb.Update)
		g1.POST("/updateQmsSipTbItemno", QmsSipTb.UpdateItemno)
		g1.POST("/deleteQmsSipTb", QmsSipTb.Delete)
		g1.POST("/addQmsSipTbImg", QmsSipTbImg.Add)

		// 標籤機設定
		g1.GET("getPrinter", lmsPrinterRoute.Get)
		g1.POST("addPrinter", lmsPrinterRoute.Add)
		g1.POST("updatePrinter", lmsPrinterRoute.Update)
		g1.POST("deletePrinter", lmsPrinterRoute.Delete)

		// 標籤設定
		g1.GET("getLabel", lmsLabelRoute.Get)
		g1.POST("addLabel", lmsLabelRoute.Add)
		g1.POST("updateLabel", lmsLabelRoute.Update)
		g1.POST("deleteLabel", lmsLabelRoute.Delete)

		// 標籤模板表頭
		g1.GET("getModelHd", lmsModelHdRoute.Get)
		g1.POST("addModelHd", lmsModelHdRoute.Add)
		g1.POST("updateModelHd", lmsModelHdRoute.Update)
		g1.POST("deleteModelHd", lmsModelHdRoute.Delete)

		// 標籤模板表身
		g1.GET("getModelTb", lmsModelTbRoute.Get)

		// 標籤模板料號綁定
		g1.GET("getModelPdt", lmsModelPdtRoute.Get)
		g1.POST("addModelPdt", lmsModelPdtRoute.Add)
		g1.POST("deleteModelPdt", lmsModelPdtRoute.Delete)
		g1.POST("updateModelPdt", lmsModelPdtRoute.Update)

		// 獲取條碼列印序碼和判斷重碼
		g1.POST("print", lmsPrintRoute.Print)
		g1.GET("printWoN", lmsPrintRoute.GetWoN)

		// 列印標籤查詢
		g1.POST("reprint", lmsHistoryRoute.Reprint)
		g1.GET("getModelHistory", lmsHistoryRoute.Get)
		g1.GET("getModelHistoryExcel", lmsHistoryRoute.GetExcel)

		// 條碼日期管理
		g1.GET("getDaycodeHd", lmsTimecodeHdRoute.Get)
		g1.POST("addDaycodeHd", lmsTimecodeHdRoute.Add)
		g1.POST("updateDaycodeHd", lmsTimecodeHdRoute.Update)
		g1.POST("deleteDaycodeHd", lmsTimecodeHdRoute.Delete)
		g1.GET("getDaycodeTb", lmsTimecodeTbRoute.Get)
		g1.POST("updateDaycodeTb", lmsTimecodeTbRoute.Update)

		// -------------------------------------BI-------------------------------------

		g1.GET("getBiQuerycatPN", biRouteGetPN.GetPN)
		g1.GET("getBiParent", biRouteGetPN.GetParent)

		// -------------------------------------BI-Current-------------------------------------
		g1.GET("getBiAllInstockOkRate", biRouteInstockCurrent.GetAll)
		g1.GET("getBiInstockOkRateQuerycat", biRouteInstockCurrent.GetQuerycat)
		g1.GET("getBiInstockOkRatePN", biRouteInstockCurrent.GetPN)

		// -------------------------------------BI-MOM-------------------------------------
		g1.GET("getBiMOMAllQuerycat", biRouteInstockMOM.GetAllQuerycat)
		g1.GET("getBiMOMQuerycat", biRouteInstockMOM.GetQuerycat)
		g1.GET("getBiMOMPN", biRouteInstockMOM.GetPN)

		g1.GET("getBiPdtwhmQuerycat", biRoutePdtwhmCurrent.GetQuerycat)
		g1.GET("getBiPdtwhmPN", biRoutePdtwhmCurrent.GetPN)
	}

	// 異常管理看板
	g2 := router.Group("api/abc")
	{
		g2.GET("/getAbcSelect", selectorRoute.GetAbcSelect)
		g2.GET("/getAllAbc", abcRoute.GetAllAbc)
		g2.GET("/getAllAbcApp", abcRoute.GetAllAbcApp)
		g2.POST("/judge", abcRoute.Judge)
	}
	// 數位生產看板
	g3 := router.Group("api/dashboard")
	{
		g3.GET("/socket", dashboardRoute.DashboardSocket)
	}

	// 綜合效能看板
	g4 := router.Group("api/efficiency")
	{
		g4.GET("/socket", efficiencyRoute.EfficiencySocket)
	}

	// 稼動設備看板
	g5 := router.Group("api/act")
	{
		g5.GET("/socket", actRoute.ActSocket)
	}

	// 設備生產狀況
	g6 := router.Group("api/device")
	{
		g6.GET("/getData", tableHeadline.GetBoardData)
	}

	g7 := router.Group("api")
	{
		g7.GET("/getAllWhID", selectorRoute.GetAllWhID)

		// ------------------------------------WMS共用API------------------------------------
		g7.GET("/getTxnHd", wmsRoute.GetWmsHd)
		g7.GET("/getTxnTb", wmsRoute.GetWmsTb)
		g7.POST("/addTxnHd", wmsRoute.AddWmsHd)
		g7.POST("/updateTxnTb", wmsRoute.UpdateWmsTb) // 靜態修改run21
		g7.POST("/txn", wmsRoute.Txn)                 // 表頭異動
		g7.POST("/confirm", wmsRoute.Confirm)         // 靜態表身提交
		g7.POST("/deleteTxnHd", wmsRoute.Delete)      // 靜態表身提交

		// // -------------------------------------入庫規劃-------------------------------------
		// g7.GET("/getWosOut", planinRoute.WosOut10Select)

		// -------------------------------------庫存查詢-------------------------------------
		g7.GET("getAllInv", inventoryRoute.GetAllInventory)
		g7.GET("getAllWhInv", inventoryRoute.GetAllWhInv)
		g7.GET("getDetail", inventoryRoute.GetDetail)

		// -------------------------------------每月庫存-------------------------------------
		g7.GET("getMonthInv", invMonthRoute.Get)

		// -------------------------------------庫存彙總表-------------------------------------
		g7.GET("getInvPrice", invPriceRoute.Get)

		// -------------------------------------批號查詢-------------------------------------
		g7.GET("getInvBatch", invBatchRoute.Get)
		g7.GET("getInvBatchDetail", invBatchRoute.GetDetail)

		// // -------------------------------------代理入庫-------------------------------------
		// g7.GET("/getAllWoN", hubRoute.GetAllWoN)
		// g7.POST("/addHub", hubRoute.AddHub)
		// g7.POST("/modifyHub", hubRoute.ModifyHub)
		// g7.POST("/deleteHub", hubRoute.DeleteHub)
		// g7.GET("/getAllHub", hubRoute.GetAllHub)

		// --------------------------採購入庫要撈PMS OMS OUT下拉選單--------------------------
		g7.GET("/getPmsOutHd", pmsOutRoute.GetHd)
		g7.GET("/getPmsOutTb", pmsOutRoute.GetTb)

		// --------------------------銷售發貨要撈OMS OUT下拉選單--------------------------
		g7.GET("/getOmsOutHd", pmsOutRoute.GetHd)
		g7.GET("/getOmsOutTb", pmsOutRoute.GetTb)

		// -------------------------------------run80 81 90 91-------------------------------------
		g7.GET("/getrun8090Hd", run8090HdRoute.GetHd)
		g7.GET("/getHdSelect", run8090HdRoute.GetHdSelect)
		g7.GET("/getTb", run8090TbRoute.GetTb)
		g7.POST("/addrun8090Hd", run8090HdRoute.AddHd)
		g7.POST("/updaterun8090Hd", run8090HdRoute.UpdateHd)
		g7.POST("/updaterun8090Tb", run8090TbRoute.UpdateTb)
		g7.POST("/deleterun8090Hd", run8090HdRoute.DeleteHd)
		g7.POST("/deleterun8090Tb", run8090TbRoute.DeleteTb)

		// -------------------------------生產領用-------------------------------
		g7.GET("/getParent", pdtUseRoute.GetParent)
		g7.GET("/getChild", pdtUseRoute.GetChild)

		// -------------------------------------倉庫報表-------------------------------------
		g7.GET("/daily", reportRoute.Daily)
		g7.GET("/dailyDetail", reportRoute.Detail)

		// ---------------------開帳 - 建立run10inv料號與倉庫的連結---------------------
		g7.GET("/getOpenHd", openRoute.GetHd)
		g7.GET("/getOpenTb", openRoute.GetTb)
		g7.POST("/addOpen", openRoute.Add)
		g7.POST("/updateOpen", openRoute.Update)
		g7.POST("/txnOpen", openRoute.Txn)

		// -------------------------------盤點-------------------------------
		// 盤點計畫
		g7.GET("/getInvProject", invProjectRoute.Get)
		g7.POST("/addInvProject", invProjectRoute.Add)
		g7.POST("/deleteInvProject", invProjectRoute.Delete)
		g7.POST("/updateInvProject", invProjectRoute.Update)

		// 盤點單
		g7.GET("/getInvHd", invRouteHD.Get)
		g7.POST("/addInvHd", invRouteHD.Add)
		g7.POST("/updateInvHd", invRouteHD.Update)
		g7.POST("/deleteInvHd", invRouteHD.Delete)
		g7.POST("/txnInvHd", invRouteHD.Txn)

		// 盤點明細
		g7.POST("/updateInvTb", invRouteTB.Update)
		g7.GET("/getInvTb", invRouteTB.Get)
		g7.POST("/updateInvTbQty", invRouteTB.UpdateQty)
	}

	g8 := router.Group("api/spc")
	{
		g8.GET("/getSpcData", spcRoute.GetSpcDataSocket)
		g8.GET("/getSpcDetail", spcRoute.GetSpcDetailSocket)
	}

	// 載入放在web的前端檔案
	router.Use(static.Serve("/", static.LocalFile("./web/build", true)))
	// 如果找不到router回到index.html
	router.NoRoute(func(ctx *gin.Context) {
		file, _ := ioutil.ReadFile("./web/build/index.html")
		// etag := fmt.Sprintf("%x", md5.Sum(file)) //nolint:gosec
		// ctx.Header("ETag", etag)
		// ctx.Header("Cache-Control", "no-cache")

		// if match := ctx.GetHeader("If-None-Match"); match != "" {
		// 	if strings.Contains(match, etag) {
		// 		ctx.Status(http.StatusNotModified)
		// 		// 這裡若沒 return 的話，會執行到 ctx.Data
		// 		return
		// 	}
		// }
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", file)
	})

	util.BuddhaBlessNoBug()

	return router
}
