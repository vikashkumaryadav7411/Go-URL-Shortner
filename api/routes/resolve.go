package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
	"github.com/vikashkumaryadav7411/Go-URL-Shortner/database"
)

func ResolveURL(c fiber.Ctx) error {

	url := c.Params("url")
	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"error": "short not found in the database",
		})
	} else if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": "cannot connect to DB",
		})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	return c.Redirect().Status(fiber.StatusMovedPermanently).To(value)
}
