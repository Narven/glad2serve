package cmd

import (
	"fmt"
	"github.com/gobuffalo/packr"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/cobra"
	"log"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Runs local server with goodies",
	Long:  `Runs local server with goodies`,
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetInt32("port")
		if err != nil {
			log.Fatal(err.Error())
		}

		box := packr.NewBox("../templates")
		// http.Handle("/", http.FileServer(box))

		fmt.Println("Serving:")

		for _, i := range box.List() {
			fmt.Println(fmt.Sprintf("  ▸ http://localhost:%d/%s", port, i))
		}

		app := fiber.New(fiber.Config{
			DisableStartupMessage: true,
			GETOnly:               true,
		})

		app.Use(favicon.New())

		app.Use(recover.New())

		app.Use(logger.New())

		app.Use(cors.New(cors.Config{
			AllowCredentials: true,
			AllowHeaders:     "*",
			AllowMethods:     "GET",
			AllowOrigins:     "*",
		}))

		app.Use(filesystem.New(filesystem.Config{
			Root: box,
		}))

		log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
	},
}

func init() {
	serveCmd.Flags().Int32P("port", "p", 4500, "Specify a port")
	rootCmd.AddCommand(serveCmd)
}
