package handlers

import (
	dto "backend/dto/result"
	transactiondto "backend/dto/transaction"
	"backend/models"
	"backend/repositories"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
	"gopkg.in/gomail.v2"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

// var path_files = "http://localhost:5000/uploads/"

var c = coreapi.Client{
	ServerKey: os.Getenv("SERVER_KEY"), ClientKey: os.Getenv("CLIENT_KEY"),
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}
func (h *handlerTransaction) FindTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transactions, err := h.TransactionRepository.FindTransactions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
	// for i, p := range transactions {
	// 	transactions[i].Attachment = path_files + p.Attachment
	// }
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: transactions}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerTransaction) GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var transaction models.Transaction

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.SuccessResult{Code: http.StatusOK, Data: transaction}
		json.NewEncoder(w).Encode(response)
	}
	// transaction.Attachment = path_file + transaction.Attachment
	// fmt.Println(id)
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: transaction}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userID := int(userInfo["id"].(float64))
	status := userInfo["role"]
	if status == "admin" {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Sorry, you can't make a transaction !!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// dataContex := r.Context().Value("dataFile")
	// filepath := dataContex.(string)

	qty, _ := strconv.Atoi(r.FormValue("qty"))
	trip_id, _ := strconv.Atoi(r.FormValue("trip_id"))
	// user_id, _ := strconv.Atoi(r.FormValue("user_id"))
	total, _ := strconv.Atoi(r.FormValue("total"))

	//create unique idtransaction
	var transIDisMatch = false
	var transactionID int

	for !transIDisMatch {
		transactionID = int(time.Now().Unix())
		request, _ := h.TransactionRepository.GetTransaction(transactionID)
		if request.ID == 0 {
			transIDisMatch = true
		}

	}
	// fmt.Println("ini transUnix", transactionID)

	request := transactiondto.TransactionRequest{
		Qty:    qty,
		Status: r.FormValue("status"),
		Total:  total,
		TripID: trip_id,
		// UserID: userID,
	}
	// log.Print(request)

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// var ctx = context.Background()
	// var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	// var API_KEY = os.Getenv("API_KEY")
	// var API_SECRET = os.Getenv("API_SECRET")

	// // Add Cloudinary credentials
	// cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// // Upload file to Cloudinary
	// resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "dewetour/transaction"})

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	transaction := models.Transaction{
		ID:     transactionID,
		Qty:    request.Qty,
		Status: "Waiting Payment",
		// Attachment: resp.SecureURL,
		Total:  request.Total,
		TripID: request.TripID,
		UserID: userID,
	}
	transaction, err = h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	transaction, _ = h.TransactionRepository.GetTransaction(transaction.ID)
	// fmt.Println(transaction)

	// request token from midtranss
	// initial snap client
	// var s = snap.Client{}
	// s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	// req := &snap.Request{
	// 	TransactionDetails: midtrans.TransactionDetails{
	// 		OrderID:  strconv.Itoa(transaction.ID),
	// 		GrossAmt: int64(transaction.Total),
	// 	},
	// 	CreditCard: &snap.CreditCardDetails{
	// 		Secure: true,
	// 	},
	// 	CustomerDetail: &midtrans.CustomerDetails{
	// 		FName: transaction.User.Fullname,
	// 		Email: transaction.User.Email,
	// 	},
	// }
	// snapResp, _ := s.CreateTransaction(req)
	// fmt.Println(req)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: transaction}
	json.NewEncoder(w).Encode(response)
}
func convertResponseTransaction(u models.Transaction) transactiondto.TransactionResponse {
	return transactiondto.TransactionResponse{
		ID:     u.ID,
		Qty:    u.Qty,
		Status: u.Status,
		// Attachment: u.Attachment,
		Total:  u.Total,
		TripID: u.TripID,
		Trip:   u.Trip,
		UserID: u.UserID,
	}
}

func (h *handlerTransaction) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userID := int(userInfo["id"].(float64))
	status := userInfo["email"]
	if status == "admin@mail.com" {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Sorry, you can't make a transaction"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// dataContex := r.Context().Value("dataFile")
	// filepath := dataContex.(string)

	qty, _ := strconv.Atoi(r.FormValue("qty"))
	trip_id, _ := strconv.Atoi(r.FormValue("trip_id"))
	total, _ := strconv.Atoi(r.FormValue("total"))
	request := transactiondto.TransactionRequest{
		Qty:    qty,
		Status: r.FormValue("status"),
		Total:  total,
		TripID: trip_id,
		UserID: userID,
		// Attachment: filepath,
	}

	// var ctx = context.Background()
	// var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	// var API_KEY = os.Getenv("API_KEY")
	// var API_SECRET = os.Getenv("API_SECRET")

	// // Add Cloudinary credentials
	// cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// // Upload file to Cloudinary
	// resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "dewetour/transaction"})

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	transaction, err := h.TransactionRepository.GetTransaction(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// transaction.Status = "Waiting Approve"

	if request.Qty != 0 {
		transaction.Qty = request.Qty
	}
	if request.Status != "" {
		transaction.Status = request.Status
	}
	// if request.Attachment != "" {
	// 	transaction.Attachment = resp.SecureURL
	// }
	if request.Total != 0 {
		transaction.Total = request.Total
	}
	if request.TripID != 0 {
		transaction.TripID = request.TripID
	}
	// fmt.Println(transaction)

	data, err := h.TransactionRepository.UpdateTransaction(transaction, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	trans, _ := h.TransactionRepository.GetTransaction(data.ID)
	// fmt.Println("ini trans", trans)
	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)
	// Use to midtrans.Production if you want Production Environment (accept real transaction).

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(trans.ID),
			GrossAmt: int64(trans.Total),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: trans.User.Fullname,
			Email: trans.User.Email,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)

	fmt.Println(snapResp)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: snapResp}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerTransaction) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	data, err := h.TransactionRepository.DeleteTransaction(transaction, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) ApproveTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Check ID Transaction"}
		json.NewEncoder(w).Encode(response)
		return
	}

	fmt.Println(transaction.ID)
	fmt.Println(transaction.Status)

	transaction.Status = "Approve"
	data, err := h.TransactionRepository.UpdateTransaction(transaction, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) CancelTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction.Status = "Cancel"
	data, err := h.TransactionRepository.UpdateTransaction(transaction, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}
func SendMail(status string, transaction models.Transaction) {

	var CONFIG_SMTP_HOST = "smtp.gmail.com"
	var CONFIG_SMTP_PORT = 587
	var CONFIG_SENDER_NAME = "Dewetour <demo.dumbways@gmail.com>"
	var CONFIG_AUTH_EMAIL = os.Getenv("EMAIL_SYSTEM")
	var CONFIG_AUTH_PASSWORD = os.Getenv("PASSWORD_SYSTEM")

	var TripName = transaction.Trip.Title
	var price = strconv.Itoa(transaction.Trip.Price)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", transaction.User.Email)
	mailer.SetHeader("Subject", "Transaction Status")
	mailer.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
	  <html lang="en">
		<head>
		<meta charset="UTF-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>Document</title>
		<style>
		  h1 {
		  color: brown;
		  }
		</style>
		</head>
		<body>
		<h2>Transaction :</h2>
		<ul style="list-style-type:none;">
		  <li>Name : %s</li>
		  <li>Total payment: Rp.%s</li>
		  <li>Status : <b>%s</b></li>
		</ul>
		</body>
	  </html>`, TripName, price, status))

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent! to " + transaction.User.Email)
}

func (h *handlerTransaction) Notification(w http.ResponseWriter, r *http.Request) {
	var notificationPayload map[string]interface{}

	// fmt.Println("iniiii", r.Body)

	err := json.NewDecoder(r.Body).Decode(&notificationPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "error di sini"}
		json.NewEncoder(w).Encode(response)
		return
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)
	// fmt.Println(orderId)
	IDtrans, _ := strconv.Atoi(orderId)

	transaction, _ := h.TransactionRepository.GetTransaction(IDtrans)

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			// TODO set transaction status on your database to 'challenge'
			// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
			SendMail("challenge", transaction)
			h.TransactionRepository.UpdateStatus("pending", transaction.ID)

		} else if fraudStatus == "accept" {
			// TODO set transaction status on your database to 'success'
			SendMail("approve", transaction)
			h.TransactionRepository.UpdateStatus("success", transaction.ID)
		}
	} else if transactionStatus == "settlement" {
		// TODO set transaction status on your database to 'success'
		SendMail("success", transaction)
		h.TransactionRepository.UpdateStatus("success", transaction.ID)

	} else if transactionStatus == "deny" {
		// TODO you can ignore 'deny', because most of the time it allows payment retries
		// and later can become success
		SendMail("failed", transaction)
		h.TransactionRepository.UpdateStatus("failed", transaction.ID)

	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		// TODO set transaction status on your database to 'failure'
		SendMail("cancel", transaction)
		h.TransactionRepository.UpdateStatus("failed", transaction.ID)

	} else if transactionStatus == "pending" {
		// TODO set transaction status on your database to 'pending' / waiting payment
		SendMail("pending", transaction)
		h.TransactionRepository.UpdateStatus("pending", transaction.ID)
	}

	w.WriteHeader(http.StatusOK)
}
