package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"greenlight.bcc/internal/assert"
)

func TestRegisterUser(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routesTest())
	defer ts.Close()

	const (
		validName     = "Giorno"
		validEmail    = "giovanna@gmail.com"
		validPassword = "GoldenExperience123"
	)

	tests := []struct {
		testName string
		Name     string
		Email    string
		Password string
		wantCode int
	}{
		{
			testName: "valid submission",
			Name:     validName,
			Email:    validEmail,
			Password: validPassword,
			wantCode: http.StatusCreated,
		},
		{
			testName: "empty name",
			Name:     "",
			Email:    validEmail,
			Password: validPassword,
			wantCode: http.StatusUnprocessableEntity,
		},
		{
			testName: "email address not valid",
			Name:     validName,
			// Email:    "not@valid",
			Email:    "not@valid.",
			Password: validPassword,
			wantCode: http.StatusUnprocessableEntity,
		},
		{
			testName: "test for wrong input",
			Name:     validName,
			Email:    validEmail,
			Password: validPassword,
			wantCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			inputData := struct {
				Name     string
				Email    string
				Password string
			}{
				Name:     tt.Name,
				Email:    tt.Email,
				Password: tt.Password,
			}

			b, err := json.Marshal(&inputData)
			fmt.Println(inputData)
			if err != nil {
				t.Fatal("wrong input data")
			}
			if tt.testName == "test for wrong input" {
				b = append(b, 'a')
			}

			code, _, _ := ts.postForm(t, "/v1/users", b)

			assert.Equal(t, code, tt.wantCode)
		})
	}
}
