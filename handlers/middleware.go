package handlers

import (
	"context"
	"net/http"
	"github.com/flexGURU/tutorials/data"


)

// MiddlewareValidateProduct validates the product in the request and calls next if ok
func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")

		prod := &data.Products

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)

			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		// validate the product
		errs := p.Validate(prod)
		if len(errs) != 0 {
			p.l.Println("[ERROR] validating product", errs)

			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			prod.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}
		data

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}