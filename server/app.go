package server

import (
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

type UserData struct {
	ID       string   `json:"id,omitempty" bson:"_id,omitempty"`
	Username string   `json:"kadi" bson:"kadi"`
	Password string   `json:"sifre" bson:"sifre"`
	Roles    []string `json:"roles" bson:"roles"`
	Takim_no int64    `json:"takim_no" bson:"takim_no"`
}

var (
	App          *fiber.App
	API          fiber.Router
	SessionStore *session.Store
	ADMIN        fiber.Router
)

func init() {
	SessionStore = session.New(session.Config{
		Expiration: 2 * time.Minute,
	})
	App = fiber.New()
	API = App.Group("/api")
	ADMIN = App.Group("/admin")
	App.Use(
		cors.New(),
		logger.New(),
		getSession,
	)
	App.Get("/swagger/*", swagger.HandlerDefault)
}

func getSession(c *fiber.Ctx) error {
	var userDoc UserData
	sess, err := SessionStore.Get(c)
	if err != nil {
		log.Println(err)
	}
	/* var userDoc UserData */
	token := sess.Get("takim")
	if len(strings.Split(c.Path(), "swagger")) > 1 {
		return c.Next()
	}
	if c.Path() == "/api/giris" && c.Method() == "POST" {
		return c.Next()
	}
	if c.Path() == "/api/giris" && token == nil {
		return c.Next()
	} else if c.Path() == "/api/giris" && token != nil {
		if err := json.Unmarshal([]byte(token.(string)), &userDoc); err != nil {
			log.Println("parseError: " + err.Error())
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		out, _ := json.Marshal(userDoc)
		sess.Set("takim", string(out))
		if err := sess.Save(); err != nil {
			log.Println(err.Error())
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"takim_no": userDoc.Takim_no,
		})
	}
	if token == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	if err := json.Unmarshal([]byte(token.(string)), &userDoc); err != nil {
		log.Println("parseError: " + err.Error())
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	options := AuthObject{
		Username: userDoc.Username,
		Password: userDoc.Password,
	}
	user, err := GetUserData(c, options)
	if err != nil {
		log.Println("getUserDataError: " + err.Error())
		return c.Status(500).JSON(fiber.Map{
			"message": "no user found with these credentials",
		})
	}
	out, _ := json.Marshal(user)
	sess.Set("takim", string(out))
	if err := sess.Save(); err != nil {
		log.Println(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
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

type AuthObject struct {
	Username string `json:"kadi" bson:"kadi"`
	Password string `json:"sifre" bson:"sifre"`
}

func GetUserData(ctx *fiber.Ctx, options AuthObject) (UserData, error) {
	var userDoc UserData
	if err := UserCollection.FindOne(ctx.Context(), bson.D{
		{Key: "kadi", Value: options.Username},
		{Key: "sifre", Value: options.Password},
	}).Decode(&userDoc); err != nil {
		return userDoc, err
	} else {
		return userDoc, nil
	}
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
/*
func GetAuthDocumentByToken(token string) (UserData, error) {
	var auth UserData
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
} */
