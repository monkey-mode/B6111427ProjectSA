package main

import (
	"github.com/B6111427/app/ent/role"
	"context"
	"fmt"
	"log"

	"github.com/B6111427/app/controllers"
	_ "github.com/B6111427/app/docs"
	"github.com/B6111427/app/ent"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)
//struct for User Data
type User struct{
	Email string
	Name string
	Role int
}
//struct for ClientEntity Data
type ClientEntity struct{
	Name string
	Status string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())
	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewUserController(v1, client)
	controllers.NewBookingController(v1, client)
	controllers.NewBookingTypeController(v1, client)
	controllers.NewClientEntityController(v1, client)
	controllers.NewRoleController(v1, client)

	
	Role := []string{"Library Member","Seller","Librarian"}
	for _, r := range Role{
		client.Role.
			Create().
			SetROLENAME(r).
			Save(context.Background())
	}

	Bookingtype := []string{"รับชม 1 คน","รับชม 2 คน","รับชม 3 คน","รับชม 4 คน","รับชม 5 คน","รับชม 6 คน","รับชม 7 คน"}
	for _, bt := range Bookingtype{
		client.Bookingtype.
			Create().
			SetBOOKTYPENAME(bt).
			Save(context.Background())
	}


	User := []User{
			User{"B6111427@g.sut.sc.th","Suphachai Phetthamrong",1},
			User{"suphachaiphetthamrong@gmail.com","Phetthamrong Suphachai",1},
	}
	
	for _, u := range User {

		r,err := client.Role.
			Query().
			Where(role.IDEQ(int(u.Role))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		client.User.
			Create().
			SetUSEREMAIL(u.Email).
			SetUSERNAME(u.Name).
			SetRoleplay(r).
			Save(context.Background())
	}

	ClientEntity := []ClientEntity{
		ClientEntity{"Video 0n Demand - 01","Available"},
		ClientEntity{"Video 0n Demand - 02","Available"},
		ClientEntity{"Video 0n Demand - 03","Available"},
		ClientEntity{"Video 0n Demand - 04","Available"},
		ClientEntity{"Video 0n Demand - 05","Available"},
		ClientEntity{"Video 0n Demand - 06","Available"},
		ClientEntity{"Video 0n Demand - 07","Available"},
		ClientEntity{"Video 0n Demand - 08","Available"},
	}

	for _, cl := range ClientEntity {
		client.ClientEntity.
			Create().
			SetCLIENTNAME(cl.Name).
			SetCLIENTSTATUS(cl.Status).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
