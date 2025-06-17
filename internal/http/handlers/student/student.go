package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/everestp/go-api/internal/types"
	"github.com/everestp/go-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("creating  a student")
		
		
		var student types.Student 
		err :=  json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err ,io.EOF){
			response.WriteJson(w,http.StatusBadRequest ,response.GeneralError(fmt.Errorf("body is empty")))
			return
		}

		if err!= nil{
			response.WriteJson(w , http.StatusBadRequest ,response.GeneralError(err))
			// response.WriteJson(w ,http.StatusBadRequest)
		}

		// request validation
		 if err := validator.New().Struct(student); err !=nil {
			valigateErrs := err.(validator.ValidationErrors)
 response.WriteJson(w ,http.StatusBadGateway ,response.ValidationError(valigateErrs))
 return
		 }



		 
response.WriteJson(w ,http.StatusCreated ,map[string]string{"success":"OK"})




		// w.Write([]byte("wellcome to Resturant"))
	}
}