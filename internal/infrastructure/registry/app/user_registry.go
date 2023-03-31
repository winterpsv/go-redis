package appRegistry

import (
	ir "task3_3_new/user-management/internal/adapter/db/mongodb"
	au "task3_3_new/user-management/internal/application/service"
	"task3_3_new/user-management/internal/application/usecase"
	ip "task3_3_new/user-management/internal/controller/http/presenter"
	v1 "task3_3_new/user-management/internal/controller/http/v1"
)

func (r *registry) NewUserController() *v1.UserController {
	return v1.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewAuthController() *v1.AuthController {
	return v1.NewAuthController(r.NewAuthInteractor())
}

func (r *registry) NewUserInteractor() *interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter(), r.NewAuth())
}

func (r *registry) NewAuthInteractor() *interactor.AuthInteractor {
	return interactor.NewAuthInteractor(r.NewUserRepository(), r.NewUserPresenter(), r.NewAuth())
}

func (r *registry) NewUserRepository() *ir.UserRepository {
	return ir.NewUserRepository(r.db, r.cfg.MongoCollection)
}

func (r *registry) NewUserPresenter() *ip.UserPresenter {
	return ip.NewUserPresenter()
}

func (r *registry) NewAuth() *au.Auth {
	return au.NewAuth(r.NewUserRepository(), r.cfg)
}
