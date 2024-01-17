package api

import (
	"fmt"
	"log"

	"filelineserve/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type HTTPHandler interface {
	Routes(rg *gin.RouterGroup)
	Group() *string
}

type Server interface {
	PushHandlerWithGroup(h HTTPHandler, rg *gin.RouterGroup)
}

type BaseGinServer struct{}

func (*BaseGinServer) PushHandlerWithGroup(h HTTPHandler, rg *gin.RouterGroup) {
	if gs := *h.Group(); gs != "" {
		h.Routes(rg.Group(gs))

		return
	}

	h.Routes(rg)
}

type FileServeAPI struct {
	BaseGinServer
	fileSvc *service.FileService
	Port    int
}

func NewFileServeAPI(port int, fileSvc *service.FileService) *FileServeAPI {
	return &FileServeAPI{
		fileSvc: fileSvc,
		Port:    port,
	}
}

func (s *FileServeAPI) Run() {
	gin.SetMode(gin.ReleaseMode)

	g := gin.New()

	routeGroup := g.Group("/")
	s.PushHandlerWithGroup(NewFileHandler(s.fileSvc), routeGroup)

	// add swagger endpoint for documentarion
	routeGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("running service with Port %d", s.Port)
	if err := g.Run(fmt.Sprintf(":%d", s.Port)); err != nil {
		log.Fatal(err.Error())
	}
}
