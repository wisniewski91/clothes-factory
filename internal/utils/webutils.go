package utils

import "github.com/gofiber/fiber/v2"

type AdminPageData struct {
	PageInfo AdminPageInfo
	Data     any
}

type AdminPageInfo struct {
	Title        string
	ContentTitle string
	Path         string
}

func PreparePageData(title string, contentTitle string, path string, data any) *AdminPageData {
	pageInfo := AdminPageInfo{
		Title:        title,
		ContentTitle: contentTitle,
		Path:         path,
	}
	return &AdminPageData{
		PageInfo: pageInfo,
		Data:     data,
	}
}

func Render(c *fiber.Ctx, template string, data *AdminPageData) error {
	partRender := c.Get("Part-render")
	if partRender == "true" {
		return c.Render(template, data)
	}

	return c.Render(template, data, "layout/main")
}
