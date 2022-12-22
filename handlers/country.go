package handlers

import (
	countrydto "backend/dto/country"
	dto "backend/dto/result"
	"backend/models"
	"backend/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerCountry struct {
	CountryRepository repositories.CountryRepository
}

func HandlerCountry(CountryRepository repositories.CountryRepository) *handlerCountry {
	return &handlerCountry{CountryRepository}
}
func (h *handlerCountry) FindCountries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	countries, err := h.CountryRepository.FindCountries()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Countries Not Found"}
		json.NewEncoder(w).Encode(response)
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: countries}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCountry) GetCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var country models.Country
	country, err := h.CountryRepository.GetCountry(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Countries Not Found"}
		json.NewEncoder(w).Encode(response)
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCountry(country)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCountry) CreateCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(countrydto.CountryRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	country := models.Country{
		ID:   request.ID,
		Name: request.Name,
	}
	country, err = h.CountryRepository.CreateCountry(country)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	country, _ = h.CountryRepository.GetCountry(country.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: country}
	json.NewEncoder(w).Encode(response)
}
func convertResponseCountry(u models.Country) countrydto.CountryResponse {
	return countrydto.CountryResponse{
		ID:   u.ID,
		Name: u.Name,
	}
}

func (h *handlerCountry) UpdateCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	request := new(countrydto.CountryRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	country, err := h.CountryRepository.GetCountry(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	if request.Name != "" {
		country.Name = request.Name
	}
	data, err := h.CountryRepository.UpdateCountry(country, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCountry(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCountry) DeleteCountry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	country, err := h.CountryRepository.GetCountry(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	data, err := h.CountryRepository.DeleteCountry(country, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCountry(data)}
	json.NewEncoder(w).Encode(response)
}
