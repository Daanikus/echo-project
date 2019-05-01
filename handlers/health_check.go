package handlers

import "github.com/labstack/echo"

func HealthCheck(c echo.Context) error {
  resp:=renderings.HealthCheckResponse{
    message:"Everything is hood"
  }

  return c.JSON(http.StatusOK, resp)

}
