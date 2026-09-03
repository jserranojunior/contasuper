package handlers

import (
    "github.com/gin-gonic/gin"
    "github.com/jserranojunior/sefro/backgo/http/middlewares"
    "github.com/jserranojunior/sefro/backgo/models"
    "golang.org/x/crypto/bcrypt"
)

// AuthLogin a user controllers
func AuthLogin(c *gin.Context) {
    var user models.User
    c.Bind(&user)
    email := user.Email
    password := user.Password

    if email == "" || password == "" {
        c.JSON(401, gin.H{
            "message": "Erro ao tentar fazer login email ou senha em branco",
        })
    } else {
        var user models.User
        // ADICIONADO "name", "cellphone" e "role_id" NO SELECT
        result := DB.Select("id", "password", "name", "cellphone", "role_id").Where("email = ?", email).First(&user)
        
        if result.RowsAffected == 0 {
            c.JSON(401, gin.H{
                "message": "Email ou senha incorretos",
            })
        } else {
            if compareBcrypt(user.Password, password) {
                user.Password = ""
                token := middlewares.GenerateJwt(user.ID)
                
                // RETORNANDO O TOKEN E OS DADOS DO USUÁRIO JUNTO NO JSON
                c.JSON(200, gin.H{
                    "token": &token,
                    "user": gin.H{
                        "name":      user.Name,
                        "email":     user.Email,
                        "cellphone": user.Cellphone,
                        "role_id":   user.RoleID,
                    },
                })
            } else {
                c.JSON(401, gin.H{
                    "message": "Email ou senha incorretos",
                })
            }
        }
    }
}

func compareBcrypt(hashedPassword string, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    if err == nil {
        return true
    } else {
        return false
    }
}