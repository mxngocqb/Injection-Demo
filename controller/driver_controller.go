package controller

import (
	"fmt"
	"net/http"
	"os/exec"

	"html/template"
	"injection_testing/model"
	driver "injection_testing/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type DriverController struct {
	driverService driver.DriverService
}

func NewDriverController(driverService driver.DriverService) *DriverController {
	return &DriverController{driverService: driverService}
}

type DriverServerResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// Createdriver
// @Tags Driver
// @Summary Create a new driver
// @Description Create a new driver
// @Accept json
// @Produce json
// @Param driver body model.Driver true "driver object that needs to be added"
// @Success 200
// @Router /drivers [post]
func (c *DriverController) CreateDriver(ctx *gin.Context) {
	log.Info().Msg("Creating Driver")
	createDriver := model.Driver{}
	err := ctx.ShouldBindJSON(&createDriver)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	driver, er := c.driverService.Create(&createDriver)

	if er != nil {
		ctx.JSON(500, gin.H{"error": er.Error()})
		return
	}

	serverReponse := gin.H{
		"code":   200,
		"status": "Created",
		"data":   driver,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, serverReponse)
}

// ReadAllDriver retrieves all drivers from the database.
// @Summary Get all drivers
// @Description Get all drivers from the database and returns them as JSON.
// @Tags Driver
// @Accept json
// @Produce json
// @Success 200 {object} DriverServerResponse "List of drivers"
// @Failure 500 {object} DriverServerResponse "Internal Server Error"
// @Router /drivers [get]
func (c *DriverController) ReadAllDriver(ctx *gin.Context) {
	log.Info().Msg("Reading All Driver")

	drivers, err := c.driverService.ReadAll()
	log.Info().Msgf("Not cache")

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}

	serverResponse := DriverServerResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   drivers,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, serverResponse)
}

// ReadDriverByID retrieves a driver by its ID.
// @Summary Retrieve a driver by ID
// @Description Retrieve a driver from the database by its ID and return it as JSON.
// @Tags Driver
// @Accept json
// @Produce json
// @Param driverID path string true "Driver ID"
// @Success 200 {object} DriverServerResponse "Driver information"
// @Failure 400 {object} DriverServerResponse "Bad Request"
// @Failure 404 {object} DriverServerResponse "Not Found"
// @Failure 500 {object} DriverServerResponse "Internal Server Error"
// @Router /drivers/{driverID} [get]
func (c *DriverController) ReadDriverByID(ctx *gin.Context) {
	log.Info().Msg("Reading Category By ID")
	driverID := ctx.Param("driverID")

	driver, err := c.driverService.ReadByID(driverID)

	if err != nil {
		ctx.JSON(500, DriverServerResponse{Code: http.StatusInternalServerError, Status: "Internal Server Error", Data: err.Error()})
		return
	}

	serverResponse := DriverServerResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   driver,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, serverResponse)
}

// UpdateDriver updates a driver.
// @Summary Update a driver
// @Description Update information of a driver based on the provided data in JSON format.
// @Tags Driver
// @Accept json
// @Produce json
// @Param driverID path string true "Driver ID"
// @Param driver body model.Driver true "Driver object to be updated"
// @Success 200 {object} DriverServerResponse "Updated driver information"
// @Failure 400 {object} DriverServerResponse "Bad Request"
// @Failure 500 {object} DriverServerResponse "Internal Server Error"
// @Router /drivers/{driverID}  [put]
func (c *DriverController) UpdateDriver(ctx *gin.Context) {
	log.Info().Msg("Updating Driver")
	updateDriver := model.Driver{}
	err := ctx.ShouldBindJSON(&updateDriver)
	if err != nil {
		ctx.JSON(400, DriverServerResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   "Nothing",
		})
		return
	}
	driver, er := c.driverService.Update(&updateDriver)
	if er != nil {
		ctx.JSON(500, DriverServerResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Erro",
			Data:   "Nothing",
		})
		return
	}
	serverResponse := DriverServerResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   driver,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, serverResponse)
}

// DeleteDriver deletes a driver by its ID.
// @Summary Delete a driver
// @Description Delete a driver from the database based on its ID.
// @Tags Driver
// @Accept json
// @Produce json
// @Param driverID path string true "Driver ID"
// @Success 200 {object} DriverServerResponse "Driver deleted successfully"
// @Failure 500 {object} DriverServerResponse "Internal Server Error"
// @Router /drivers/{driverID} [delete]
func (c *DriverController) DeleteDriver(ctx *gin.Context) {
	log.Info().Msg("Deleting Driver")
	DriverID := ctx.Param("driverID")
	err := c.driverService.Delete(DriverID)
	if err != nil {
		ctx.JSON(500, DriverServerResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
		return
	}

	serverResponse := DriverServerResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   "OK",
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, serverResponse)
}

// list godoc
// @Summary list an Path address
// @Description Sends ICMP echo requests (list) to the specified Path address.
// @Tags list
// @Accept  json
// @Produce  text/html
// @Param path query string true "Path address to list"
// @Success 200 {string} string "list result"
// @Failure 500 {object} DriverServerResponse "Internal Server Error"
// @Failure 400 {object} DriverServerResponse "Bad Request"
// @Router /list [get]
func (c *DriverController) List(ctx *gin.Context) {
	path := ctx.Query("path")
	log.Info().Msgf("listing %s", path)

	if path == "" {
		ctx.JSON(400, gin.H{"error": "Missing 'ip' query parameter"})
		return
	}

	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("ls %s", path))
	log.Info().Msgf("Executing command: %s", cmd.String())
	output, err := cmd.Output()

	if err != nil {
		log.Error().Err(err).Msg("Failed to execute ping command")
		ctx.JSON(500, gin.H{"error": "Failed to execute ping command"})
		return
	}

	ctx.Data(200, "text/html; charset=utf-8", output)
}

func (c *DriverController) CodeInjection(ctx *gin.Context) {
	userInput := ctx.Query("input")

	
	// Template HTML với {{.}} để hiển thị userInput một cách an toàn
	htmlTemplate := fmt.Sprintf(`<html>
		<head>
			<title>Injection Demo</title>
		</head>
		<body>
			<h1>User Input</h1>
			<p> %s </p>
		</body>
		</html>
	`, userInput)
	// Tạo một template từ chuỗi HTML
	tmpl := template.Must(template.New("inject").Parse(htmlTemplate))
	err := tmpl.Execute(ctx.Writer, nil)
	if err != nil {
		fmt.Println("Error executing template:", err)
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	
}
