package handlers

import (
	dto "backend/dto/result"
	tripdto "backend/dto/trip"
	"backend/models"
	"backend/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerTrip struct {
	TripRepository repositories.TripRepository
}

var path_file = "http://localhost:5000/uploads/"

func HandlerTrip(TripRepository repositories.TripRepository) *handlerTrip {
	return &handlerTrip{TripRepository}
}
func (h *handlerTrip) FindTrips(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	trips, err := h.TripRepository.FindTrips()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
	for i, p := range trips {
		trips[i].Image = path_file + p.Image
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: trips}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTrip) GetTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var trip models.Trip

	trip, err := h.TripRepository.GetTrip(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
	trip.Image = path_file + trip.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: trip}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTrip) CreateTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	country_id, _ := strconv.Atoi(r.FormValue("country_id"))
	price, _ := strconv.Atoi(r.FormValue("price"))
	quota, _ := strconv.Atoi(r.FormValue("quota"))
	day, _ := strconv.Atoi(r.FormValue("day"))
	night, _ := strconv.Atoi(r.FormValue("night"))
	request := tripdto.TripRequest{
		Title:          r.FormValue("title"),
		Acomodation:    r.FormValue("acomodation"),
		Transportation: r.FormValue("transportation"),
		Eat:            r.FormValue("eat"),
		Day:            day,
		Night:          night,
		DateTrip:       r.FormValue("date_trip"),
		Price:          price,
		Quota:          quota,
		Description:    r.FormValue("description"),
		CountryID:      country_id,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	trip := models.Trip{
		ID:             request.ID,
		Title:          request.Title,
		Acomodation:    request.Acomodation,
		Transportation: request.Transportation,
		Eat:            request.Eat,
		Day:            request.Day,
		Night:          request.Night,
		DateTrip:       request.DateTrip,
		Price:          request.Price,
		Quota:          request.Quota,
		Description:    request.Description,
		Image:          filename,
		CountryID:      request.CountryID,
	}
	// fmt.Println(trip)
	trip, err = h.TripRepository.CreateTrip(trip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	trip, _ = h.TripRepository.GetTrip(trip.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrip(trip)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseTrip(u models.Trip) tripdto.TripResponse {
	return tripdto.TripResponse{
		ID:             u.ID,
		Title:          u.Title,
		CountryID:      u.CountryID,
		Acomodation:    u.Acomodation,
		Transportation: u.Transportation,
		Eat:            u.Eat,
		Day:            u.Day,
		Night:          u.Night,
		DateTrip:       u.DateTrip,
		Price:          u.Price,
		Quota:          u.Quota,
		Description:    u.Description,
		Image:          u.Image,
	}
}

func (h *handlerTrip) UpdateTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	country_id, _ := strconv.Atoi(r.FormValue("country_id"))
	price, _ := strconv.Atoi(r.FormValue("price"))
	quota, _ := strconv.Atoi(r.FormValue("quota"))
	day, _ := strconv.Atoi(r.FormValue("day"))
	night, _ := strconv.Atoi(r.FormValue("night"))
	request := tripdto.TripRequest{
		Title:          r.FormValue("title"),
		Acomodation:    r.FormValue("acomodation"),
		Transportation: r.FormValue("transportation"),
		Eat:            r.FormValue("eat"),
		Day:            day,
		Night:          night,
		DateTrip:       r.FormValue("date_trip"),
		Price:          price,
		Quota:          quota,
		Description:    r.FormValue("description"),
		CountryID:      country_id,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	trip, err := h.TripRepository.GetTrip(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Title != "" {
		trip.Title = request.Title
	}
	if request.CountryID != 0 {
		trip.CountryID = request.CountryID
	}
	if request.Acomodation != "" {
		trip.Acomodation = request.Acomodation
	}
	if request.Transportation != "" {
		trip.Transportation = request.Transportation
	}
	if request.Eat != "" {
		trip.Eat = request.Eat
	}
	if request.Day != 0 {
		trip.Day = request.Day
	}
	if request.Night != 0 {
		trip.Night = request.Night
	}
	if request.DateTrip != "" {
		trip.DateTrip = request.DateTrip
	}
	if request.Price != 0 {
		trip.Price = request.Price
	}
	if request.Quota != 0 {
		trip.Quota = request.Quota
	}
	if request.Description != "" {
		trip.Description = request.Description
	}
	if request.Image != "" {
		trip.Image = filename
	}
	data, err := h.TripRepository.UpdateTrip(trip, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrip(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTrip) DeleteTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	trip, err := h.TripRepository.GetTrip(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	data, err := h.TripRepository.DeleteTrip(trip, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTrip(data)}
	json.NewEncoder(w).Encode(response)
}
