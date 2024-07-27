package handler

import "github.com/bugsfounder/bugsfounderweb/db"

type DB_Handler struct {
	Client db.Client
}

func (h *DB_Handler) DemoFuncHandler() {
	h.Client.DemoFunc()
}
