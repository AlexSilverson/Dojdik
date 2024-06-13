package controllers

import (
	"fmt"

	"github.com/AlexSilverson/Dojdik/src/entity"
	"github.com/AlexSilverson/Dojdik/src/entity/DTO"
	"github.com/AlexSilverson/Dojdik/src/solver"
	"github.com/gofiber/fiber/v2"
)

// PutCityById Getting City by json
//
//	@Summary		Getting City by Id
//	@Description	Getting City by Id in detail
//	@Tags			Citis
//	@Accept			json
//	@Produce		json
//	@Param			request			body		entity.Input	true	"Request of Creating City Object"
//	@Failure		400				{string}	string
//	@Failure		404				{string}	string
//	@Success		200				{string}	string
//	@Router			/input [post]
func CalcByGivenData(app *fiber.App) fiber.Router {
	return app.Post("/input", func(c *fiber.Ctx) error {
		fmt.Println("here")
		//var inp entity.Input
		var inpDto DTO.InputDTO
		fmt.Println(string(c.Body()))
		err := c.BodyParser(&inpDto)
		inp := inpDto.MapToInput()
		fmt.Println(inpDto)
		fmt.Println(inp)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("City format is not valid")
		}
		ans := solver.Solve(inp)
		fmt.Println(ans)
		if err == nil {
			var out entity.Output
			out.Result = ans
			return c.Status(fiber.StatusOK).JSON(out)
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}
	})
}
