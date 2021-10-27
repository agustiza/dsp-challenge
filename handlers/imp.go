package handlers

import (
	"fmt"
	"net/http"
	"wildlife-challenge/util"
)

type ImpHandler struct {

}

func (h *ImpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	impId, bidId := util.ParseQueryImpId(r)

	fmt.Println("callback called: ",impId, bidId) // Todo: settle outstanding balance, count impressions
}

