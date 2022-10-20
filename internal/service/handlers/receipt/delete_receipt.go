package handlers

import (
	"Menu-Service/internal/service/helpers"
	requests "Menu-Service/internal/service/requests/receipt"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func DeleteReceipt(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewDeleteReceiptRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	Receipt, err := helpers.ReceiptsQ(r).FilterById(request.ReceiptId).Get()
	if Receipt == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	err = helpers.ReceiptsQ(r).Delete(request.ReceiptId)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to delete receipt")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
