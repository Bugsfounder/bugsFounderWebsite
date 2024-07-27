package handler

import "github.com/bugsfounder/bugsfounderweb/db"

type HandlerForDBHandlers struct {
	Client db.Client
}

func (h_DB *HandlerForDBHandlers) DemoFuncHandler() {
	h_DB.Client.DemoFunc()
}

func (h_DB *HandlerForDBHandlers) CreateOneBlog() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) GetAllBlogs() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) GetOneBlogByID() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) UpdateOneBlogById() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) DeleteOneBlogById() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) CreateOneTutorail() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) GetAllTutorial() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) GetOneTutorialByID() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) UpdateOneTutorialById() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) DeleteOneTutorialById() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) GetAllUsers() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) HandlerGetOneUserByUsernameOrEmail() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) GetOneUserByUsername() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) GetOneUserByEmail() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) CreateOneUser() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) UpdateOneUserByUsernameOrEmail() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
func (h_DB *HandlerForDBHandlers) DeleteOneUserByUsernameOrEmail() {
	// access db functions ex: h.Client.DemoFunc() // in db
}
