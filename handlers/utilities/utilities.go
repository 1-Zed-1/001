package utilities

import (
	"fmt"
	"net/url"
	"os"
	"path"

	"github.com/gofiber/fiber/v2"
)

type Directories struct {
	Folders []string
	Files   []string
}

var location string = "web/static/"

func NavigateDirectories(c *fiber.Ctx, foldername string) error {
	foldername, err := url.QueryUnescape(foldername)
	if err != nil {
		fmt.Println(err)
		return &fiber.Error{
			Code:    418,
			Message: "URL Parsing Error",
		}
	}
	structure := Directories{}
	if foldername != "" {
		location = location + foldername + "/"
	}
	entries, err := os.ReadDir(location)
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Internal Server Error",
		}
	}
	for _, element := range entries {
		if element.IsDir() == true {
			structure.Folders = append(structure.Folders, element.Name())
		} else {
			structure.Files = append(structure.Files, element.Name())
		}
	}
	if location != "web/" && location != "./" {
		return c.Render("snippets/directory", fiber.Map{
			"files":   structure.Files,
			"folders": structure.Folders,
			"cwd":     location,
		})
	}
	return c.Render("snippets/directory", fiber.Map{
		"files":   [3]string{"nah", "not", "today"},
		"folders": [3]string{"nope", "not", "happening"},
		"cwd":     "ngl let you sniff around in my directories",
	})
}

func DownloadFile(c *fiber.Ctx, filename string) error {
	return c.Download(location+filename, filename)
}

func NavigateBackwards(c *fiber.Ctx) error {
	location = path.Dir(path.Dir(location))
	location = location + "/"
	NavigateDirectories(c, "")
	return nil
}
