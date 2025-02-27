package userHandler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/DevJonathanSantos/poc-go-api/internal/dto"
	httpErr "github.com/DevJonathanSantos/poc-go-api/internal/handler/httperr"
	"github.com/DevJonathanSantos/poc-go-api/internal/handler/validation"
)

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserDto

	if r.Body == http.NoBody {
		slog.Error("body is empty", slog.String("package", "userHandler"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httpErr.NewBadRequestError("body is required")
		json.NewEncoder(w).Encode(msg)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error to decode body", "err", err, slog.String("package", "userHandler"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httpErr.NewBadRequestError("error to decode body")
		json.NewEncoder(w).Encode(msg)
		return
	}

	httpErr := validation.ValidateHttpData(req)
	if httpErr != nil {
		slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "userHandler"))
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(httpErr)
		return
	}
}
