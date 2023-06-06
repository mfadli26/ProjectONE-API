package ws

import (
	"net/http"

	"projectONE/pkg/util/general"

	"github.com/centrifugal/centrifuge"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var NodeCentrifugal *centrifuge.Node

func handleLog(e centrifuge.LogEntry) {
	logrus.Infof("%s: %v", e.Message, e.Fields)
}

func convert(h http.Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		h.ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}
}

func auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		newCtx := centrifuge.SetCredentials(ctx, &centrifuge.Credentials{
			UserID: general.RandSeq(10),
		})

		r = r.WithContext(newCtx)
		h.ServeHTTP(w, r)
	})
}
