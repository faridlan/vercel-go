package service

import (
	"encoding/json"
	"os"

	"github.com/faridlan/vercel-go/model"
)

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

func ConvertJson() *model.Profile {

	reader, _ := os.Open("../profile.json")
	encoder := json.NewDecoder(reader)

	profile := model.Profile{}
	encoder.Decode(&profile)

	return &profile

}

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
