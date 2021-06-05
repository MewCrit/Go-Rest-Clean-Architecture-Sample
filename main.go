package main

import (
	"fmt"
	"net/http"

	controller "go-clean-architecture/app/controllers"
	router "go-clean-architecture/app/http"
	"go-clean-architecture/app/repository"
	"go-clean-architecture/app/service"
)

var (
	identityRepository repository.RepoInterface       = repository.RepoImpl()
	identityService    service.ServiceInterce         = service.ServiceImpl(identityRepository)
	identityController controller.ControllerInterface = controller.Controller(identityService)
	httpRouter         router.Router                  = router.NewMuxRouter()
)

func main() {

	const port string = ":8000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/users", identityController.GetUsers)

	httpRouter.SERVE(port)

}
