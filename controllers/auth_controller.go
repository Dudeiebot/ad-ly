package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/thedevsaddam/govalidator"

	"github.com/dudeiebot/ad-ly/helpers"
	"github.com/dudeiebot/ad-ly/request"
	"github.com/dudeiebot/ad-ly/services"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var req request.Register

	rules := govalidator.MapData{
		"name":     []string{"required", "alpha_space"},
		"email":    []string{"required", "email"},
		"password": []string{"regex:^.{8,}", "required"},
	}

	opt := govalidator.Options{
		Rules:   rules,
		Request: r,
		Data:    &req,
	}

	validationErrors := helpers.ValidateRequest(opt, "json")

	if len(validationErrors) != 0 {
		helpers.ReturnValidatorErrors(w, validationErrors)
		return
	}

	resp, err, status := services.RegisterUser(req)

	if err != nil {
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(helpers.Message(err.Error()))
		return
	}

	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
	return
}

func VerifyUser(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	data := map[string]interface{}{
		"token": token,
	}

	rules := govalidator.MapData{
		"token": []string{"required", "alpha_num"},
	}

	opts := govalidator.Options{
		Rules:   rules,
		Request: r,
		Data:    &data,
	}

	validationErrors := helpers.ValidateRequest(opts, "query")

	if len(validationErrors) != 0 {
		helpers.ReturnValidatorErrors(w, validationErrors)
		return
	}

	resp, err, status := services.VerifyUser(token)
	if err != nil {
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(helpers.Message(err.Error()))
		return
	}

	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
	return
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var req request.LoginUser

	rules := govalidator.MapData{
		"email":    []string{"required", "email"},
		"password": []string{"regex:^.{8,}", "required"},
	}

	opts := govalidator.Options{
		Rules:   rules,
		Request: r,
		Data:    &req,
	}

	validationErrors := helpers.ValidateRequest(opts, "json")

	if len(validationErrors) != 0 {
		helpers.ReturnValidatorErrors(w, validationErrors)
		return
	}

	resp, err, status := services.LoginUser(req)

	if err != nil {
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(helpers.Message(err.Error()))
		return
	}

	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
	return
}
