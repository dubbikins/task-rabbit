package task

import "net/http"

type controller struct {
}

func (c *controller) CreateTask(w http.ResponseWriter, r *http.Request) {

}
func NewController() *controller {
	return &controller{}
}

type Controller interface {
	CreateTask(w http.ResponseWriter, r *http.Request)
}
