package jobtypes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wisniewski91/clothes-factory/internal/storage"
	"github.com/wisniewski91/clothes-factory/internal/utils"
)

var (
	repo    *MysqlJobtypesRepository
	service *JobtypesService
)

func Init() {
	repo = NewMysqlJobtypesRepository(storage.DB)
	service = NewJobtypesService(repo)
}

func indexHandler(c *fiber.Ctx) error {
	result, err := service.repository.FetchAll()
	if err != nil {
		c.Response().Header.Add("HX-Retarget", "#notification-container")
		c.Response().Header.Add("HX-Reswap", "beforeend")
		return c.SendString(utils.CreateNotification("Problem loading Job list!", utils.ErrorNotification{}))
	}
	pageData := utils.PreparePageData("JobTypes", "JobTypes - list", routePath, result)
	return utils.Render(c, renderPath("index"), pageData)
}

func createHandler(c *fiber.Ctx) error {
	pageData := utils.PreparePageData("JobTypes", "Create job type", routePath, nil)

	return utils.Render(c, renderPath("modal", "create"), pageData)
}

func createPostHandler(c *fiber.Ctx) error {
	var jobType JobtypeRecord

	err := c.BodyParser(&jobType)
	if err != nil {
		return c.SendString(utils.CreateNotification("Problem parsing jobtype data", utils.ErrorNotification{}))
	}
	err = utils.Validate(jobType)
	if err != nil {
		return c.SendString(utils.CreateNotification(err.Error(), utils.ErrorNotification{}))
	}

	result, err := service.Create(&jobType)
	if err != nil {
		return c.SendString(utils.CreateNotification(err.Error(), utils.ErrorNotification{}))
	}

	return c.SendString(utils.CreateNotification(fmt.Sprintf("Job type %s created", result.Name), utils.SuccessNotification{}))
}

func editHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendString(utils.CreateNotification("Problem parsing url param", utils.ErrorNotification{}))
	}

	jobType, err := service.repository.Fetch(id)
	if err != nil {
		return c.SendString(utils.CreateNotification("No job type with provided id", utils.ErrorNotification{}))
	}
	pageData := utils.PreparePageData("JobTypes", "Edit job type", routePath, jobType)

	return utils.Render(c, renderPath("modal", "edit"), pageData)
}

func editPostHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendString(utils.CreateNotification("Problem parsing url param", utils.ErrorNotification{}))
	}
	jobType, err := service.repository.Fetch(id)
	if err != nil {
		return c.SendString(utils.CreateNotification("Job type not exists", utils.ErrorNotification{}))
	}
	err = utils.Validate(jobType)
	if err != nil {
		return c.SendString(utils.CreateNotification(err.Error(), utils.ErrorNotification{}))
	}

	var postJobType JobtypeRecord

	err = c.BodyParser(&postJobType)
	if err != nil {
		return c.SendString(utils.CreateNotification("Problem parsing jobType form data", utils.ErrorNotification{}))
	}
	err = utils.Validate(postJobType)
	if err != nil {
		return c.SendString(utils.CreateNotification(err.Error(), utils.ErrorNotification{}))
	}
	postJobType.Id = id

	result, err := service.repository.Update(&postJobType)
	if err != nil {
		return c.SendString(utils.CreateNotification(err.Error(), utils.ErrorNotification{}))
	}

	return c.SendString(utils.CreateNotification(fmt.Sprintf("Job type %s saved", result.Name), utils.SuccessNotification{}))
}
