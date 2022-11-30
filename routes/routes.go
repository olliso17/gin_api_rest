package routes

import (
	"gin_api_rest/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.GET("/:nome", controllers.Saudacao)
	r.DELETE("/alunos/:id", controllers.DeletarAluno)
	r.POST("/alunos", controllers.CriarAlunos)
	r.PUT("/alunos/:id", controllers.EditaAluno)
	//se vc quiser usar uma porta diferent r.Run(":5000") defina ela dentro dos parenteses
	r.Run()

}
