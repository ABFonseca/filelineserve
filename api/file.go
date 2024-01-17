package api

import (
	"filelineserve/data/response"
	"filelineserve/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const ResponseError = "error"

type fileHandler struct {
	svc *service.FileService
}

func NewFileHandler(s *service.FileService) HTTPHandler {
	return &fileHandler{
		svc: s,
	}
}

func (h *fileHandler) Routes(rg *gin.RouterGroup) {
	rg.GET("/:line", h.GetLine)
}

func (h *fileHandler) Group() *string {
	groupName := "lines"
	return &groupName
}

// GetLine godoc
// @Summary Retrieve a single line from a file
// @Schemes
// @Description Get the line with index {line} from the file being served
// @Accept json
// @Param line Index of the line to be read from the file
// @Produce json
// @Success 200 {object} string "line"
// @Router /lines/{id} [get]
func (h *fileHandler) GetLine(c *gin.Context) {
	lineNum := c.Param("line")
	if lineNum == "" {
		c.JSON(http.StatusBadRequest, NewErrorResponse("line index must not be empty"))
		return
	}

	lineNumInt, err := strconv.Atoi(lineNum)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse("line index must be an integer"))
		return
	}

	line, err := h.svc.GetLine(lineNumInt)
	if err != nil {
		c.JSON(http.StatusRequestEntityTooLarge, NewErrorResponse("line index beyond end of file"))
		return
	}

	readLine := &response.ReadLine{
		Line: line,
	}

	resp, err := readLine.Marshal()
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewErrorResponse("error marshaling the response"))
		return
	}
	c.Data(http.StatusOK, "application/json", resp)
}

type ErrorResponse struct {
	Status  string   `json:"status"`
	Message []string `json:"message,omitempty"`
}

func NewErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{Status: ResponseError, Message: []string{message}}
}
