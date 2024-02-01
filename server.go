package main

import (
	//"fmt"
    "fmt"
	"log"
	"os"

	"os/exec"
	"time"

	"breks/pkg/eks"

	// "github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/session"
	// "github.com/aws/aws-sdk-go/service/eks"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("your_secret_key")

func listClusters2(c *gin.Context){
    region 				:= "us-east-1"
	sessionToken 		:= "" //  optional

	// Create a new EKS client
	eksClient, err := eks.NewEKSClient(
                                os.Getenv("ENV_accessKeyID"),       // *resp.Credentials.AccessKeyId, //
                                os.Getenv("ENV_secretAccessKey"),   // *resp.Credentials.SecretAccessKey, //
                                sessionToken,                       // *resp.Credentials.SessionToken, //
                                region)                             // AWS region

	if err != nil {
		log.Fatalf("Failed to create EKS client: %v", err)
	}

	// List EKS clusters
	clusters, err := eksClient.ListClusters()
	if err != nil {
		log.Fatalf("Failed to list EKS clusters: %v", err)
	}

	fmt.Println("Clusters:", clusters)
    c.JSON(200, gin.H{"clusters": clusters})
}

// func listClusters(c *gin.Context) {
// 	sess := session.Must(session.NewSession(&aws.Config{
// 		Region: aws.String("us-east-1"),
// 	}))

// 	svc := eks.New(sess)
// 	result, err := svc.ListClusters(nil)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(200, gin.H{"clusters": result.Clusters})
// }

func kubectlCommand(c *gin.Context) {
	cmd := exec.Command("kubectl", "get", "pods", "--all-namespaces")
	out, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.String(200, string(out))
}

func login(c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(200, gin.H{"token": tokenString})
}

func authenticate(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Next()
}

func main() {
	r := gin.Default()

    // r.Use(authenticate)
	// r.POST("/login", login)
	// r.GET("/kubectl", kubectlCommand)

    r.GET("/list-clusters", listClusters2)
	r.Run(":8080")
}
