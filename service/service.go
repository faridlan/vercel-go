package service

import (
	"encoding/json"
	"log"

	_ "embed"

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

//go:embed profile.json
var profileJson []byte

func ConvertJson() (*model.Profile, error) {
	// // Construct the file path based on the environment
	// basePath, _ := os.Getwd()
	// filePath := filepath.Join(basePath, "public", "profile.json")

	// reader, err := os.Open(filePath)
	// if err != nil {
	// 	log.Printf("Error opening profile.json: %v", err)
	// 	return nil, err
	// }
	// defer reader.Close()

	// decoder := json.NewDecoder(reader)

	// profile := &model.Profile{}
	// err = decoder.Decode(profile)
	// if err != nil {
	// 	log.Printf("Error decoding profile.json: %v", err)
	// 	return nil, err
	// }

	// return profile, nil

	profile := &model.Profile{}
	err := json.Unmarshal(profileJson, profile)
	if err != nil {
		log.Printf("Error decoding profile.json: %v", err)
		return nil, err
	}
	return profile, nil
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
