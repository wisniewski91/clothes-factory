package parts

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wisniewski91/clothes-factory/internal/storage"
	"github.com/wisniewski91/clothes-factory/internal/utils"
)

var (
	repo    *MysqlPartsRepository
	service *PartsService
)

func Init() {
	repo = NewMysqlPartsRepository(storage.DB)
	service = NewPartsService(repo)
}

func indexHandler(c *fiber.Ctx) error {
	result, err := service.repository.FetchAll()
	if err != nil {
		c.Response().Header.Add("HX-Retarget", "#notification-container")
		c.Response().Header.Add("HX-Reswap", "beforeend")
		return c.SendString(utils.CreateNotification("Problem loading parts!", utils.ErrorNotification{}))
	}
	pageData := utils.PreparePageData("Parts", "Parts - list", routePath, result)
	return utils.Render(c, renderPath("index"), pageData)
}

func createHandler(c *fiber.Ctx) error {
	pageData := utils.PreparePageData("Parts", "Create Parts", routePath, nil)

	return utils.Render(c, renderPath("modal", "create"), pageData)
}

func createPostHandler(c *fiber.Ctx) error {
	var part PartRecord

	err := c.BodyParser(&part)
	if err != nil {
		return c.SendString(utils.CreateNotification("Problem parsing part data", utils.ErrorNotification{}))
	}
	err = utils.Validate(part)
	if err != nil {
		return c.SendString(utils.CreateNotification(err.Error(), utils.ErrorNotification{}))
	}

	result, err := service.Create(&part)
	if err != nil {
		return c.SendString(utils.CreateNotification(err.Error(), utils.ErrorNotification{}))
	}

	return c.SendString(utils.CreateNotification(fmt.Sprintf("Part %s created", result.Name), utils.SuccessNotification{}))
}

func editHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendString(utils.CreateNotification("Problem parsing url param", utils.ErrorNotification{}))
	}

	part, err := service.repository.Fetch(id)
	if err != nil {
		return c.SendString(utils.CreateNotification("No Part with provided id", utils.ErrorNotification{}))
	}
	pageData := utils.PreparePageData("Parts", "Edit Part", routePath, part)

	return utils.Render(c, renderPath("modal", "edit"), pageData)
}

func editPostHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendString(utils.CreateNotification("Problem parsing url param", utils.ErrorNotification{}))
	}
	part, err := service.repository.Fetch(id)
	if err != nil {
		return c.SendString(utils.CreateNotification("Part not exists", utils.ErrorNotification{}))
	}
	err = utils.Validate(part)
	if err != nil {
		return c.SendString(utils.CreateNotification(err.Error(), utils.ErrorNotification{}))
	}

	var postPart PartRecord

	err = c.BodyParser(&postPart)
	if err != nil {
		return c.SendString(utils.CreateNotification("Problem parsing part form data", utils.ErrorNotification{}))
	}
	err = utils.Validate(postPart)
	if err != nil {
		return c.SendString(utils.CreateNotification(err.Error(), utils.ErrorNotification{}))
	}
	postPart.Id = id

	result, err := service.repository.Update(&postPart)
	if err != nil {
		return c.SendString(utils.CreateNotification(err.Error(), utils.ErrorNotification{}))
	}

	return c.SendString(utils.CreateNotification(fmt.Sprintf("Part %s saved", result.Name), utils.SuccessNotification{}))
}
