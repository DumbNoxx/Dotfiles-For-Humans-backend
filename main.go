package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SkillJson struct {
	Background string   `json:"background"`
	Expertise  []string `json:"expertise"`
	Approach   string   `json:"approach"`
}

var Skills = []SkillJson{
	{
		Background: "Over two years of experience in web and mobile development, I specialize in creating clean, efficient, and user-centered digital experiences.",
		Expertise: []string{
			"Frontend Development (React, Astro, Vue, Sass, Angular)",
			"Backend Development (Express, Flask, FastAPI, Go)",
			"Mobile Development (React Native, Flutter)",
			"UI/UX Design & Implementation",
			"Performance Optimization",
		},
		Approach: "I believe in minimalism not just as an aesthetic choice, but as a development philosophy. By focusing on what truly matters and eliminating unnecessary complexity, I create digital products that are intuitive, maintainable, and built to last.",
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
