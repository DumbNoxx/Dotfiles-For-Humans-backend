package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SkillJson struct {
	Background1 string   `json:"background1"`
	Background2 string   `json:"background2"`
	Expertise   []string `json:"expertise"`
	Aproach     string   `json:"aproach"`
}

var Skills = []SkillJson{
	{
		Background1: "Over two years of experience in web and mobile development, I specialize in creating clean, efficient, and user-centered digital experiences.",
		Background2: "My approach comines technical expertise with a strong design sensibility, allowing me to build solutions that are both visually compelling and functionally robust.",
		Expertise: []string{
			"Frontend Development (React, Astro, Vue, Sass, Angular — currently learning)",
			"Backend Development (Express, Flask, FastAPI, Go)",
			"Mobile Development (React Native)",
			"UI/UX Design & Implementation",
			"Performance Optimization",
		},
		Aproach: "I believe in minimalism not just as an aesthetic choice, but as a development philosophy. By focusing on what truly matters and eliminating unnecessary complexity, I create digital products that are intuitive, maintainable, and built to last.",
	},
}

func GetJson(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Skills)
}

func main() {
	router := gin.Default()
	router.GET("/Skills", GetJson)

	router.Run(":8080")
}
