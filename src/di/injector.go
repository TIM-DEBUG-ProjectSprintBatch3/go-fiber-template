package di

import (
	authJwt "github.com/TIM-DEBUG-ProjectSprintBatch3/go-fiber-template/src/auth/jwt"
	"github.com/TIM-DEBUG-ProjectSprintBatch3/go-fiber-template/src/database/postgre"
	userController "github.com/TIM-DEBUG-ProjectSprintBatch3/go-fiber-template/src/http/controllers/user"
	loggerZap "github.com/TIM-DEBUG-ProjectSprintBatch3/go-fiber-template/src/logger/zap"
	userRepository "github.com/TIM-DEBUG-ProjectSprintBatch3/go-fiber-template/src/repositories/user"
	userService "github.com/TIM-DEBUG-ProjectSprintBatch3/go-fiber-template/src/services/user"
	"github.com/TIM-DEBUG-ProjectSprintBatch3/go-fiber-template/src/validation"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samber/do/v2"
)

var Injector *do.RootScope

func init() {
	Injector = do.New()

	//? Setup Database Connection
	do.Provide[*pgxpool.Pool](Injector, postgre.NewPgxConnectInject)

	//? Setup Validation
	//? Validator
	do.Provide[*validator.Validate](Injector, validation.NewValidatorInject)

	//? Logger
	//? Zap
	do.Provide[loggerZap.LoggerInterface](Injector, loggerZap.NewLogHandlerInject)

	//? Setup Auth
	//? JWT Service
	do.Provide[authJwt.JwtServiceInterface](Injector, authJwt.NewJwtServiceInject)

	//? Setup Repositories
	//? User Repository
	do.Provide[userRepository.UserRepositoryInterface](Injector, userRepository.NewUserRepositoryInject)

	//? Setup Services
	//? User Service
	do.Provide[userService.UserServiceInterface](Injector, userService.NewUserServiceInject)

	//? Setup Controller/Handler
	//? User Controller
	do.Provide[userController.UserControllerInterface](Injector, userController.NewUserControllerInject)
}
