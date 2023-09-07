package tenant

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oasdiff/go-common/ds"
	"github.com/sirupsen/logrus"
)

const (
	PathParamTenantId = "tenant-id"
)

type Validator struct {
	dsc ds.Client
}

func NewValidator(dsc ds.Client) *Validator { return &Validator{dsc: dsc} }

func (v *Validator) Validate(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)[PathParamTenantId]
		var t ds.Tenant
		if err := v.dsc.Get(ds.KindTenant, id, &t); err != nil {
			logrus.Infof("tenant not found for request '%s'", r.URL.String())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
