package handler

import (
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

// type User struct {
// 	Name string `json:"name,omitempty"`
// 	Age  int    `json:"age,omitempty"`
// }

// func GetEnv() (*viper.Viper, error) {

// 	config := viper.New()
// 	configPath := os.Getenv("CONFIG")

// 	config.SetConfigFile(configPath)

// 	err := config.ReadInConfig()
// 	if err != nil {
// 		return nil, fmt.Errorf("error reading config gile: %w", err)
// 	}

// 	return config, nil

// }

// func ConvertJson() *User {

// 	env, err := GetEnv()
// 	if err != nil {
// 		panic(err)
// 	}

// 	reader, _ := os.Open(env.GetString("PATH"))
// 	encoder := json.NewDecoder(reader)

// 	user := User{}
// 	encoder.Decode(&user)

// 	return &user
// }

// func Greet(ctx *fiber.Ctx) error {
// 	env, err := GetEnv()
// 	if err != nil {
// 		return err
// 	}
// 	ctx.SendString(fmt.Sprintf("Hello %s \nMy Name Is %s \nMy Age Is %d \n",
// 		env.GetString("NAME"),
// 		ConvertJson().Name,
// 		ConvertJson().Age,
// 	))

// 	return nil
// }

// Handler is the entry point for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {

	//Create fiber
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		name := os.Getenv("NAME")
		if name == "" {
			name = "World"
		}
		return c.SendString("Hello " + name)
	})

	adaptor.FiberApp(app)(w, r)
}
