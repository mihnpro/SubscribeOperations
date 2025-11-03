package operations

import "net/http"

func GetAllSubInfo(mux *http.ServeMux) {
	mux.HandleFunc("/getAllSubs", GetAllSubsHandler)
}

func CreateSubInfo(mux *http.ServeMux) {
	mux.HandleFunc("/createSub", CreateSubHandler)
}

func UpdateSubInfo(mux *http.ServeMux) {
	mux.HandleFunc("/updateSub", UpdateSubHandler)
}

func DeleteSubInfo(mux *http.ServeMux) {
	mux.HandleFunc("/deleteSub", DeleteSubHandler)
}

func GetSubInfo(mux *http.ServeMux) {
	mux.HandleFunc("/getSub/{subscriptionID}", GetSubHandler)
}

func GetFullSubPriceByPeriodInfo(mux *http.ServeMux) {
	mux.HandleFunc("/getFullSubPriceByPeriod", GetFullSubPriceByPeriod)
}
