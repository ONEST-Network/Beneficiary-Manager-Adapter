package middleware

import (
	"github.com/gin-gonic/gin"
)

// AuthenticationMiddleware is a middleware function for user authentication.
// func AuthenticationMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" {
// 			er := models.LicenseError{
// 				Status:    http.StatusUnauthorized,
// 				Message:   "Please check your credentials and try again",
// 				Error:     "no credentials were passed",
// 				Path:      c.Request.URL.Path,
// 				Timestamp: time.Now().Format(time.RFC3339),
// 			}
// 			c.JSON(http.StatusUnauthorized, er)
// 			c.Abort()
// 			return
// 		}
// 		parts := strings.Split(authHeader, " ")
// 		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
// 			er := models.LicenseError{
// 				Status:    http.StatusUnauthorized,
// 				Message:   "Please check your credentials and try again",
// 				Error:     "no credentials were passed",
// 				Path:      c.Request.URL.Path,
// 				Timestamp: time.Now().Format(time.RFC3339),
// 			}
// 			c.JSON(http.StatusUnauthorized, er)
// 			c.Abort()
// 			return
// 		}
// 		tokenString := parts[1]
// 		unverfiedParsedToken, err := jwt.Parse([]byte(tokenString), jwt.WithVerify(false), jwt.WithValidate(true))
// 		if err != nil {
// 			er := models.LicenseError{
// 				Status:    http.StatusUnauthorized,
// 				Message:   "Please check your credentials and try again",
// 				Error:     "token parsing failed",
// 				Path:      c.Request.URL.Path,
// 				Timestamp: time.Now().Format(time.RFC3339),
// 			}
// 			c.JSON(http.StatusUnauthorized, er)
// 			c.Abort()
// 			return
// 		}
//
// 			var userData map[string]interface{}
// 			if err = unverfiedParsedToken.Get("user", &userData); err != nil {
// 				log.Printf("\033[31mError: %s\033[0m", err.Error())
// 				er := models.LicenseError{
// 					Status:    http.StatusUnauthorized,
// 					Message:   "Please check your credentials and try again",
// 					Error:     "incompatible token format",
// 					Path:      c.Request.URL.Path,
// 					Timestamp: time.Now().Format(time.RFC3339),
// 				}
// 				c.JSON(http.StatusUnauthorized, er)
// 				c.Abort()
// 				return
// 			}
// 			userDataBytes, err := json.Marshal(userData)
// 			if err != nil {
// 				log.Printf("\033[31mError: %s\033[0m", err.Error())
// 				er := models.LicenseError{
// 					Status:    http.StatusUnauthorized,
// 					Message:   "Please check your credentials and try again",
// 					Error:     "failed to marshal user data",
// 					Path:      c.Request.URL.Path,
// 					Timestamp: time.Now().Format(time.RFC3339),
// 				}
// 				c.JSON(http.StatusUnauthorized, er)
// 				c.Abort()
// 				return
// 			}
// 			// Unmarshal the JSON bytes into the models.User struct
// 			var user models.User
// 			err = json.Unmarshal(userDataBytes, &user)
// 			if err != nil {
// 				log.Printf("\033[31mError: %s\033[0m", err.Error())
// 				er := models.LicenseError{
// 					Status:    http.StatusUnauthorized,
// 					Message:   "Please check your credentials and try again",
// 					Error:     "incompatible token format",
// 					Path:      c.Request.URL.Path,
// 					Timestamp: time.Now().Format(time.RFC3339),
// 				}
// 				c.JSON(http.StatusUnauthorized, er)
// 				c.Abort()
// 				return
// 			}
// 			if err := db.DB.Where(models.User{Id: user.Id}).First(&user).Error; err != nil {
// 				log.Printf("\033[31mError: %s\033[0m", err.Error())
// 				er := models.LicenseError{
// 					Status:    http.StatusUnauthorized,
// 					Message:   "User not found. Please check your credentials.",
// 					Error:     err.Error(),
// 					Path:      c.Request.URL.Path,
// 					Timestamp: time.Now().Format(time.RFC3339),
// 				}
// 				c.JSON(http.StatusUnauthorized, er)
// 				c.Abort()
// 				return
// 			}
// 			c.Set("username", *user.Username)
// 			c.Set("role", *user.Userlevel)
// 		} else if iss == os.Getenv("OIDC_ISSUER") {
// 			if auth.Jwks == nil || os.Getenv("OIDC_USERNAME_KEY") == "" {
// 				log.Print("\033[31mError: OIDC environment variables not configured properly\033[0m")
// 				er := models.LicenseError{
// 					Status:    http.StatusInternalServerError,
// 					Message:   "Something went wrong",
// 					Error:     "internal server error",
// 					Path:      c.Request.URL.Path,
// 					Timestamp: time.Now().Format(time.RFC3339),
// 				}
// 				c.JSON(http.StatusInternalServerError, er)
// 				c.Abort()
// 				return
// 			}
//
// 			var user models.User
// 			if err := db.DB.Where(models.User{Username: &username}).First(&user).Error; err != nil {
// 				log.Printf("\033[31mError: %s\033[0m", err.Error())
// 				er := models.LicenseError{
// 					Status:    http.StatusUnauthorized,
// 					Message:   "User not found",
// 					Error:     err.Error(),
// 					Path:      c.Request.URL.Path,
// 					Timestamp: time.Now().Format(time.RFC3339),
// 				}
// 				c.JSON(http.StatusUnauthorized, er)
// 				c.Abort()
// 				return
// 			}
// 			c.Set("username", *user.Username)
// 		} else {
// 			log.Printf("\033[31mError: Issuer '%s' not supported or not configured in .env\033[0m", iss)
// 			er := models.LicenseError{
// 				Status:    http.StatusUnauthorized,
// 				Message:   "Please check your credentials and try again",
// 				Error:     "internal server error",
// 				Path:      c.Request.URL.Path,
// 				Timestamp: time.Now().Format(time.RFC3339),
// 			}
// 			c.JSON(http.StatusUnauthorized, er)
// 			c.Abort()
// 			return
// 		}
// 		c.Next()
// 	}
// }

// CORSMiddleware is a middleware function for CORS.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
