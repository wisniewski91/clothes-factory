package clients

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wisniewski91/clothes-factory/internal/storage"
	"github.com/wisniewski91/clothes-factory/internal/utils"
)

var (
	repo    *MysqlClientsRepository
	service *ClientsService
)

func Init() {
	repo = NewMysqlClientsRepository(storage.DB)
	service = NewClientsService(repo)
}

func indexHandler(c *fiber.Ctx) error {
	clients, err := service.repository.FetchAll()
	if err != nil {
		c.Response().Header.Add("HX-Retarget", "#notification-container")
		c.Response().Header.Add("HX-Reswap", "beforeend")
		return c.SendString(utils.CreateNotification("Problem loading Job list!", utils.ErrorNotification{}))
	}
	pageData := utils.PreparePageData("Clients", "Clients - list", routePath, clients)
	return utils.Render(c, renderPath("index"), pageData)
}

func createHandler(c *fiber.Ctx) error {
	pageData := utils.PreparePageData("Clients", "Create client", routePath, nil)

	return utils.Render(c, renderPath("modal", "create"), pageData)
}

func createPostHandler(c *fiber.Ctx) error {
	var addressData utils.AddressData
	err := c.BodyParser(&addressData)
	if err != nil {
		return c.SendString(utils.CreateNotification("Problem parsing address data", utils.ErrorNotification{}))
	}
	err = utils.Validate(addressData)
	if err != nil {
		return c.SendString(utils.CreateNotification(err.Error(), utils.ErrorNotification{}))
	}

	var client ClientRecord

	err = c.BodyParser(&client)
	if err != nil {
		return c.SendString(utils.CreateNotification("Problem parsing client data", utils.ErrorNotification{}))
	}
	err = utils.Validate(client)
	if err != nil {
		return c.SendString(utils.CreateNotification(err.Error(), utils.ErrorNotification{}))
	}
	client.AddressData = addressData

	result, err := service.Create(&client)
	if err != nil {
		return c.SendString(utils.CreateNotification(err.Error(), utils.ErrorNotification{}))
	}

	return c.SendString(utils.CreateNotification(fmt.Sprintf("Client %s created", result.Name), utils.SuccessNotification{}))
}

func editHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendString(utils.CreateNotification("Problem parsing url param", utils.ErrorNotification{}))
	}

	client, err := service.repository.Fetch(id)
	if err != nil {
		return c.SendString(utils.CreateNotification("No client with provided id", utils.ErrorNotification{}))
	}
	pageData := utils.PreparePageData("Clients", "Edit client", routePath, client)

	return utils.Render(c, renderPath("modal", "edit"), pageData)
}

func editPostHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendString(utils.CreateNotification("Problem parsing url param", utils.ErrorNotification{}))
	}
	fmt.Println(id)

	client, err := service.repository.Fetch(id)
	if err != nil {
		return c.SendString(utils.CreateNotification("Client not exists", utils.ErrorNotification{}))
	}
	client.Validate()
	var addressData utils.AddressData
	err = c.BodyParser(&addressData)
	if err != nil {
		return c.SendString(utils.CreateNotification("Problem parsing address data", utils.ErrorNotification{}))
	}
	err = utils.Validate(addressData)
	if err != nil {
		return c.SendString(utils.CreateNotification(err.Error(), utils.ErrorNotification{}))
	}

	var postClient ClientRecord

	err = c.BodyParser(&postClient)
	if err != nil {
		return c.SendString(utils.CreateNotification("Problem parsing client data", utils.ErrorNotification{}))
	}
	err = utils.Validate(postClient)
	if err != nil {
		return c.SendString(utils.CreateNotification(err.Error(), utils.ErrorNotification{}))
	}
	postClient.AddressData = addressData
	postClient.Id = id

	result, err := service.repository.Update(&postClient)
	if err != nil {
		return c.SendString(utils.CreateNotification(err.Error(), utils.ErrorNotification{}))
	}

	return c.SendString(utils.CreateNotification(fmt.Sprintf("Client %s saved", result.Name), utils.SuccessNotification{}))
}
