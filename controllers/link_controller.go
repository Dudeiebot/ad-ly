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
	var req request.GetCode

	rules := govalidator.MapData{
		"url":        []string{"required", "url"},
		"customCode": []string{"alpha_space"},
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
