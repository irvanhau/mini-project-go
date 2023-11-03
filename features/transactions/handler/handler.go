package handler

import (
	"MiniProject/features/transactions"
	"MiniProject/helper"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type transactionHandler struct {
	s   transactions.TransactionServiceInterface
	jwt helper.JWTInterface
}

func NewHandler(service transactions.TransactionServiceInterface, jwt helper.JWTInterface) transactions.TransactionHandlerInterface {
	return &transactionHandler{
		s:   service,
		jwt: jwt,
	}
}

// CreateTransaction implements transactions.TransactionHandlerInterface.
func (t *transactionHandler) CreateTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {

		var paramMCU = c.Param("idmcu")
		idMcu, err := strconv.Atoi(paramMCU)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid param id medical checkup", nil))
		}

		var paramUser = c.Param("iduser")
		idUser, err := strconv.Atoi(paramUser)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid param id user", nil))
		}

		jwtIdUser := t.jwt.GetID(c)
		jwtIdInt := jwtIdUser.(float64)
		if idUser != int(jwtIdInt) {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Cannot create other user transaction", nil))
		}

		result, err := t.s.CreateTransaction(idUser, idMcu)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success create transaction", result))
	}
}

// PaymentTransaction implements transactions.TransactionHandlerInterface.
func (t *transactionHandler) PaymentTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramIDTrans = c.Param("idtrans")
		idTrans, err := strconv.Atoi(paramIDTrans)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid param id transaction", nil))
		}

		var paramUser = c.Param("iduser")
		idUser, err := strconv.Atoi(paramUser)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid param id user", nil))
		}

		jwtIdUser := t.jwt.GetID(c)
		jwtIdInt := jwtIdUser.(float64)
		if idUser != int(jwtIdInt) {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Cannot create other user transaction", nil))
		}

		result, err := t.s.PaymentTransaction(idTrans, idUser)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success payment", result))
	}
}

func (t *transactionHandler) NotificationPayment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramIDTrans = c.Param("idtrans")
		idTrans, err := strconv.Atoi(paramIDTrans)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid param id transaction", nil))
		}

		var paramUser = c.Param("iduser")
		idUser, err := strconv.Atoi(paramUser)

		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Invalid param id user", nil))
		}

		jwtIdUser := t.jwt.GetID(c)
		jwtIdInt := jwtIdUser.(float64)
		if idUser != int(jwtIdInt) {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("Cannot create other user transaction", nil))
		}

		result, err := t.s.NotificationPayment(idTrans, idUser)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot process data", nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("Success get notification payment", result))
	}
}

func (t transactionHandler) GetTransactions() echo.HandlerFunc {
	return func(c echo.Context) error {
		role := t.jwt.CheckRole(c)
		if role == "Patient" {
			userId := t.jwt.GetID(c)
			userIdInt := int(userId.(float64))

			resTrans, resDetail, err := t.s.GetUserTransactions(userIdInt)

			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot Process Data", nil))
			}

			response := TransactionResponseDetail{
				UserName:       resTrans.UserName,
				IdentityNumber: resTrans.IdentityNumber,
				DetailInfo: make([]struct {
					TransactionID string "json:\"transaction_id\""
					Status        string "json:\"status\""
					Amount        int    "json:\"amount\""
				}, len(resDetail)),
			}

			for i, trans := range resDetail {
				response.DetailInfo[i].Amount = trans.Amount
				response.DetailInfo[i].TransactionID = trans.TransactionID
				response.DetailInfo[i].Status = trans.Status
			}

			return c.JSON(http.StatusOK, helper.FormatResponse("Success get transactions", response))
		}

		res, err := t.s.GetTransactions()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("Cannot Process Data", nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse("Success get transactions", res))
	}
}
