package handler

import (
	"context"
	"developer/api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// StartSaleBarcode godoc
// @Router       /start_sale_barcode [POST]
// @Summary      Starting sale with barcode
// @Description  starting sale with barcode
// @Tags         start_sale_barcode
// @Accept       json
// @Produce      json
// @Param 		 start_sale_barcode body models.StartSaleBarcodeReq false "start_sale_barcode"
// @Success      200  {object} models.EndSaleResp
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) StartSaleBarcode(c *gin.Context) {
	startSaleReq := models.StartSaleBarcodeReq{}
	endSaleResp := models.EndSaleResp{}

	err := c.ShouldBindJSON(&startSaleReq)
	if err != nil {
		handleResponse(c, "Error in handlers, shile reading body json from client!", http.StatusBadRequest, err)
		return
	}

	getProduct, err := h.Store.Product().GetByBarcode(context.Background(), startSaleReq.Barcode)
	if err != nil {
		handleResponse(c, "Error in handlers, while getting product by barcode!", http.StatusInternalServerError, err)
		return
	}

	createBasket := models.CreateBasket{
		SaleID:    startSaleReq.SaleID,
		ProductID: getProduct.ID,
		Quantity:  startSaleReq.Quantity,
		Price:     uint(startSaleReq.Quantity) * uint(getProduct.Price),
	}

	_, err = h.Store.Basket().Create(context.Background(), createBasket)
	if err != nil {
		handleResponse(c, "Error in handlers, while creating basket!", http.StatusInternalServerError, err)
		return
	}

	storage, err := h.Store.Storage().GetByProductID(context.Background(), models.PrimaryKey{ID: getProduct.ID})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting storage by product id!", http.StatusInternalServerError, err)
		return
	}

	if storage.Count < startSaleReq.Quantity {
		handleResponse(c, "Don't have enough product!", http.StatusInternalServerError,
			fmt.Sprintf("Sorry, we don't have enough %s product, left %d", getProduct.Name, storage.Count))
		return
	}

	err = h.Store.Storage().UpdateCount(context.Background(), storage.ID, startSaleReq.Quantity)
	if err != nil {
		handleResponse(c, "Error in handlers, while updating count, storage", http.StatusInternalServerError, err)
		return
	}

	endSaleResp.ProductName = getProduct.Name
	endSaleResp.ProductPrice = getProduct.Price
	endSaleResp.ProductQuantity = startSaleReq.Quantity
	endSaleResp.TotalSum = int(createBasket.Price)

	sale, err := h.Store.Sale().GetByID(context.Background(), models.PrimaryKey{ID: startSaleReq.SaleID})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting sale!", http.StatusInternalServerError, err)
		return
	}

	_, err = h.Store.StorageTransaction().Create(context.Background(), models.CreateStorageTransaction{
		StaffID:                sale.CashierID,
		ProductID:              getProduct.ID,
		StorageTransactionType: "minus",
		Price:                  float64(endSaleResp.TotalSum),
		Quantity:               float64(startSaleReq.Quantity),
	})
 
	if err != nil {
		handleResponse(c, "Error in handlers, while creating storage transaction!", http.StatusInternalServerError, err)
		return
	}

	err = h.Store.Sale().UpdateSalePrice(context.Background(), models.UpdateSaleForPrice{
		ID:     sale.ID,
		Price:  endSaleResp.TotalSum,
		Status: "success",
	})
	if err != nil {
		handleResponse(c, "Error in handlers, while updating sale price and status!", http.StatusInternalServerError, err)
		return
	}

	cashier, err := h.Store.Staff().GetByID(context.Background(), models.PrimaryKey{ID: sale.CashierID})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting staff by id!", http.StatusInternalServerError, err)
		return
	}

	cashierTarif, err := h.Store.Tarif().GetByID(context.Background(), models.PrimaryKey{ID: cashier.TarifID})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting tarif by id!", http.StatusInternalServerError, err)
		return
	}

	balance := 0.0

	if cashierTarif.TarifType == "percent" {
		if sale.PaymentType == "cash" {
			balance = float64(endSaleResp.TotalSum) * (cashierTarif.AmountForCash)
			fmt.Println(endSaleResp.TotalSum, cashierTarif.AmountForCash, balance)
		} else {
			balance = float64(endSaleResp.TotalSum) * (cashierTarif.AmountForCard)
		}
	} else {
		if sale.PaymentType == "cash" {
			balance = (cashierTarif.AmountForCash)
		} else {
			balance = (cashierTarif.AmountForCard)
		}
	}

	reqToUpdate := models.UpdateStaffBalanceAndCreateTransaction{
		Cashier: models.StaffInformation{
			StaffID: sale.CashierID,
			Amount:  balance,
		},
		SaleID:          sale.ID,
		TransactionType: "topup",
		SourceType:      "sales",
		Amount:          float64(endSaleResp.TotalSum),
		Description:     "We will be glad to see you again!",
	}

	if sale.ShopAssistantID != "" {
		ShopAssistant, err := h.Store.Staff().GetByID(context.Background(), models.PrimaryKey{ID: sale.ShopAssistantID})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting shop assistant by id!", http.StatusInternalServerError, err)
		return
	}

	ShopAssistantTarif, err := h.Store.Tarif().GetByID(context.Background(), models.PrimaryKey{ID: ShopAssistant.TarifID})
	if err != nil {
		handleResponse(c, "Error in handlers, while getting tarif by id!", http.StatusInternalServerError, err)
		return
	}

	if ShopAssistantTarif.TarifType == "percent" {
		if sale.PaymentType == "cash" {
			balance = float64(endSaleResp.TotalSum) * (ShopAssistantTarif.AmountForCash)
		} else {
			balance = float64(endSaleResp.TotalSum) * (ShopAssistantTarif.AmountForCard)
		}
	} else {
		if sale.PaymentType == "cash" {
			balance = ShopAssistantTarif.AmountForCash
		} else {
			balance = ShopAssistantTarif.AmountForCard
		}
	}
	reqToUpdate.ShopAssistant.StaffID = sale.ShopAssistantID
	reqToUpdate.ShopAssistant.Amount  = balance
	fmt.Println(reqToUpdate.ShopAssistant.Amount)
	}

	err = h.Store.Transaction().UpdateStaffBalanceAndCreateTransaction(context.Background(), reqToUpdate)
	if err != nil{
		handleResponse(c, "Error in handlers, while updating cashier balance and creating transaction!", http.StatusInternalServerError,err)
		return
	}

	handleResponse(c, "success", http.StatusOK, endSaleResp)
}
