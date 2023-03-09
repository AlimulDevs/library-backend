package routes

import (
	middlewares "golang/app/middlewares"
	middlewareInstructor "golang/app/middlewares/instructor"
	instructorController "golang/controllers/instructorController"
	mediamoduleController "golang/controllers/mediaModuleController"
	"golang/helper"
	instructorrepository "golang/repository/instructorRepository"
	mediamodulerepository "golang/repository/mediaModuleRepository"
	instructorservice "golang/service/instructorService"
	mediamoduleservice "golang/service/mediaModuleService"
	"golang/util"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *echo.Echo {
	app := echo.New()
	// repository

	app.Validator = &helper.CustomValidator{
		Validator: validator.New(),
	}
	configLogger := middlewares.ConfigLogger{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	configInstructor := middleware.JWTConfig{
		Claims:     &middlewareInstructor.JwtInstructorClaims{},
		SigningKey: []byte(util.GetConfig("TOKEN_SECRET")),
	}

	app.Use(configLogger.Init())
	app.Use(middleware.CORS())

	instructorRp := instructorrepository.Newinstructorrepository(db)
	mediaModuleRp := mediamodulerepository.NewMediaModuleRepository(db)

	// end repository

	// service

	instructorSv := instructorservice.NewinstructorService(instructorRp)
	mediaModuleSv := mediamoduleservice.NewMediaModuleService(mediaModuleRp)

	// end service

	// controller

	instructorsCtr := instructorController.InstructorController{
		InstructorService: instructorSv,
	}

	mediaModuleCtr := mediamoduleController.MediaModuleController{
		MediaModuleService: mediaModuleSv,
	}

	// end controller

	// routes
	instructor := app.Group("/instructor")
	instructor.POST("/register", instructorsCtr.Register)
	instructor.POST("/login", instructorsCtr.Login)

	privateInstructor := app.Group("/instructor", middleware.JWTWithConfig(configInstructor))
	privateInstructor.Use(middlewares.CheckTokenMiddlewareInstructor)
	// private instructor access
	privateInstructor.GET("/media-module", mediaModuleCtr.GetAllMediaModule)
	privateInstructor.GET("/media-module/id=:id", mediaModuleCtr.GetMediaModuleByID)
	privateInstructor.POST("/media_module/create", mediaModuleCtr.CreateMediaModule)
	privateInstructor.PUT("/media_module/update/id=:id", mediaModuleCtr.UpdateMediaModule)
	privateInstructor.DELETE("/media_module/delete/id=:id", mediaModuleCtr.DeleteMediaModule)
	return app
}
