package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/thedevsaddam/govalidator"

	"github.com/dudeiebot/ad-ly/helpers"
	"github.com/dudeiebot/ad-ly/request"
	"github.com/dudeiebot/ad-ly/services"
)

func GenerateCode(w http.ResponseWriter, r *http.Request) {
	var req request.GenerateCode

	rules := govalidator.MapData{
		"url":        []string{"required", "url"},
		"customCode": []string{"alpha_num"},
		"expireAt":   []string{"numeric"},
	}

	opt := govalidator.Options{
		Rules:   rules,
		Request: r,
		Data:    &req,
	}

	validationErrrors := helpers.ValidateRequest(opt, "json")

	if len(validationErrrors) != 0 {
		helpers.ReturnValidatorErrors(w, validationErrrors)
		return
	}

	resp, err, status := services.GenerateCode(req, r)
	if err != nil {
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(helpers.Message(err.Error()))
		return
	}

	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
	return
}

func GetUrl(w http.ResponseWriter, r *http.Request) {
	var req request.GetUrl
	req.Code = r.URL.Query().Get("cd")

	rules := govalidator.MapData{
		"cd": []string{"required", "alpha_num"},
	}

	opt := govalidator.Options{
		Rules:   rules,
		Request: r,
		Data:    &req,
	}

	validationErrrors := helpers.ValidateRequest(opt, "query")

	if len(validationErrrors) != 0 {
		helpers.ReturnValidatorErrors(w, validationErrrors)
		return
	}

	resp, err, status := services.GetUrl(req.Code)
	if err != nil {
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}

	http.Redirect(w, r, resp.Url, status)
}
