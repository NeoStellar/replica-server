package server

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/swagger"
	"go.mongodb.org/mongo-driver/bson"
)

type FormT struct {
	ID       string   `json:"id,omitempty" bson:"_id,omitempty"`
	Username string   `json:"username" bson:"username"`
	Password string   `json:"password" bson:"password"`
	Roles    []string `json:"roles" bson:"roles"`
	Token    string   `json:"token" bson:"token"`
}

var (
	App          *fiber.App
	API          fiber.Router
	SessionStore *session.Store
)

func init() {
	SessionStore = session.New(session.Config{
		Expiration: 2 * time.Minute,
	})
	App = fiber.New()
	API = App.Group("/api")
	App.Use(
		cors.New(),
		logger.New(),
		getSession,
	)
	App.Get("/swagger/*", swagger.HandlerDefault)
}

func getSession(c *fiber.Ctx) error {
	sess, err := SessionStore.Get(c)
	if err != nil {
		log.Println(err)
	}
	/* var userDoc FormT */
	token := sess.Get("takim_no")
	if len(strings.Split(c.Path(), "swagger")) > 1 {
		return c.Next()
	}
	if c.Path() == "/api/giris" && token == nil {
		return c.Next()
	} else if c.Path() == "/api/giris" && token != nil {
		return c.JSON(fiber.Map{
			"takim_no": token.(int64),
		})
	}
	if token == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	/* userDocx := sess.Get("userDoc")
	if userDocx == nil {
		docString := Redis.Get(token.(string))
		if len(docString) == 0 {
			if _doc, err := GetAuthDocumentByToken(token.(string)); err != nil {
				log.Println("parseError: | GetAuthDocumentByToken" + err.Error())
			} else {
				sess.Set("userDoc", _doc)
				if err := sess.Save(); err != nil {
					log.Println("parseError: | session" + err.Error())
					log.Println(err)
				}
				out, _ := json.Marshal(_doc)
				if err := Redis.Set(_doc.Token, string(out)); err != nil {
					log.Println("parseError: | Redis" + err.Error())
					log.Println(err)
				}
			}
		} else {
			if err := json.Unmarshal([]byte(docString), &userDoc); err != nil {
				log.Println("parseError: " + err.Error())
			}
			sess.Set("userDoc", docString)
			if err := sess.Save(); err != nil {
				log.Println("error: " + err.Error())
			}
		}
		return c.Next()
	}
	if err := json.Unmarshal([]byte(userDocx.(string)), &userDoc); err != nil {
		log.Println("parseError: " + err.Error())
	} */
	return c.Next()
}

/* func isAuthorized(c *fiber.Ctx) error {
	sess, err := SessionStore.Get(c)
	if err != nil {
		log.Println(err)
	}
	token := sess.Get("token")
	if token == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	return c.Next()
}
*/

func GetAuthDocumentByToken(token string) (FormT, error) {
	var auth FormT
	log.Println("App | GetAuthDocumentByToken | token: " + token + "")
	result := Redis.Get(auth.Token)
	if len(result) > 0 {
		if err := json.Unmarshal([]byte(result), &auth); err != nil {
			log.Println("App | GetAuthDocumentByToken | Unmarshal: " + err.Error())
			return auth, err
		}
		return auth, nil
	}
	if err := Auths.FindOne(context.TODO(), bson.D{
		{Key: "token", Value: token},
	}).Decode(&auth); err != nil {
		log.Println("App | GetAuthDocumentByToken: " + err.Error())
		return auth, err
	}
	return auth, nil
}
