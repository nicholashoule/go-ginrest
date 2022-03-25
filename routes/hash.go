package routes

import (
	"encoding/hex"
	"net/http"

	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"

	"github.com/gin-gonic/gin"
)

type Hash struct {
	Text           string `json:"text" binding:"required" validate:"nonzero"`
	HashTypeMD5    string `json:"hash (md5)"`
	HashTypeSHA1   string `json:"hash (sha1)"`
	HashTypeSHA256 string `json:"hash (sha256)"`
	HashTypeSHA512 string `json:"hash (sha512)"`
}

// Group
// Path: /hash
func addHash(rg *gin.RouterGroup) {
	path := rg.Group("/hash")

	// Default: All supported hashes
	// Example:
	// curl -v -X POST http://localhost:5000/v2/hash \
	//   -H 'content-type: application/json' \
	//   -d '{ "text": "testing" }'
	path.POST("/", func(c *gin.Context) {
		msg := &Hash{}

		if err := c.ShouldBindJSON(&msg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		tMD5 := md5.Sum([]byte(msg.Text))
		tSHA1 := sha1.Sum([]byte(msg.Text))
		tSHA256 := sha256.Sum256([]byte(msg.Text))
		tSHA512 := sha512.Sum512([]byte(msg.Text))

		msg.HashTypeMD5 = hex.EncodeToString(tMD5[:])
		msg.HashTypeSHA1 = hex.EncodeToString(tSHA1[:])
		msg.HashTypeSHA256 = hex.EncodeToString(tSHA256[:])
		msg.HashTypeSHA512 = hex.EncodeToString(tSHA512[:])
		c.JSON(http.StatusOK, msg)
	})

	// Path: /md5
	// Example:
	// 	curl -v -X POST http://localhost:5000/v2/hash/md5 \
	//   -H 'content-type: application/json' \
	//   -d '{ "text": "testing" }'
	path.POST("/md5", func(c *gin.Context) {
		msg := &Hash{}

		if err := c.ShouldBindJSON(&msg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

                tMD5 := md5.Sum([]byte(msg.Text))

		msg.HashTypeMD5 = hex.EncodeToString(tMD5[:])
		c.JSON(http.StatusOK, gin.H{"hash (md5)": msg.HashTypeMD5})
	})

	// Path: /sha1
	// Example:
	// 	curl -v -X POST http://localhost:5000/v2/hash/sha1 \
	//   -H 'content-type: application/json' \
	//   -d '{ "text": "testing" }'
	path.POST("/sha1", func(c *gin.Context) {
		msg := &Hash{}

		if err := c.ShouldBindJSON(&msg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

                tSHA1 := sha1.Sum([]byte(msg.Text))

		msg.HashTypeSHA1 = hex.EncodeToString(tSHA1[:])
		c.JSON(http.StatusOK, gin.H{"hash (sha1)": msg.HashTypeSHA1})
	})

	// Path: /sha256
	// Example:
	// 	curl -v -X POST http://localhost:5000/v2/hash/sha256 \
	//   -H 'content-type: application/json' \
	//   -d '{ "text": "testing" }'
	path.POST("/sha256", func(c *gin.Context) {
		msg := &Hash{}

		if err := c.ShouldBindJSON(&msg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

                tSHA256 := sha256.Sum256([]byte(msg.Text))

		msg.HashTypeSHA256 = hex.EncodeToString(tSHA256[:])
		c.JSON(http.StatusOK, gin.H{"hash (sha256)": msg.HashTypeSHA256})
	})

	// Path: /sha512
	// Example:
	// 	curl -v -X POST http://localhost:5000/v2/hash/sha512 \
	//   -H 'content-type: application/json' \
	//   -d '{ "text": "testing" }'
	path.POST("/sha512", func(c *gin.Context) {
		msg := &Hash{}

		if err := c.ShouldBindJSON(&msg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

                tSHA512 := sha512.Sum512([]byte(msg.Text))

		msg.HashTypeSHA512 = hex.EncodeToString(tSHA512[:])
		c.JSON(http.StatusOK, gin.H{"hash (sha256)": msg.HashTypeSHA512})
	})
}
