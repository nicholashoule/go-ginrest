package routes

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/dchest/validator"
	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
	// validator "github.com/go-playground/validator/v10"
)

// URLValidation ...
// IsValidDomain returns true if the domain is valid.
// It uses a simple regular expression to check the domain validity.
func URLValidation(c *gin.Context, hostname string) error {
	err := validator.ValidateDomainByResolvingIt(hostname)
	if err != nil {
		return err
	}

	// ok, move on
	return nil
}

/*
Router groups
*/

// addGetCert ...
// Path: /cert
func addGetCert(rg *gin.RouterGroup) {
	cert := rg.Group("/cert")

	// Default: Full Certificate
	cert.GET("/:hostname", func(c *gin.Context) {
		host := c.Params.ByName("hostname")
		url := host + ":443"

		err := URLValidation(c, host)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Unable to parse URL or not a valid URL",
				"URL":     host,
			})
		} else {
			conn, err := tls.Dial("tcp", url, nil)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Server doesn't support SSL or certificate err",
					"error":   err.Error(),
				})
			}

			err = conn.VerifyHostname(host)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Hostname doesn't match with certificate",
					"error":   err.Error(),
				})
			}

			// nb := conn.ConnectionState().PeerCertificates[0].NotBefore
			// na := conn.ConnectionState().PeerCertificates[0].NotAfter
			c.JSON(http.StatusOK, gin.H{
				// "IssuerOrganization": conn.ConnectionState().PeerCertificates[0].Issuer.Organization,
				// "NotBefore":          nb.Format(time.RFC850),
				// "NotAfter":           na.Format(time.RFC850),
				// "ServerName":         conn.ConnectionState().ServerName,
				"x509Certificate":    conn.ConnectionState().PeerCertificates[0].Raw,
			})
		}
	})

	// Path: /full
	cert.GET("/chain/:hostname", func(c *gin.Context) {
		host := c.Params.ByName("hostname")
		url := host + ":443"

		err := URLValidation(c, host)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Unable to parse URL or not a valid URL",
				"URL":     host,
			})
		} else {
			conn, err := tls.Dial("tcp", url, nil)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Server doesn't support SSL or certificate err",
					"error":   err.Error(),
				})
			}

			err = conn.VerifyHostname(host)
			if err != nil {
				//panic("Hostname doesn't match with certificate: " + err.Error())
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Hostname doesn't match with certificate",
					"error":   err.Error(),
				})
			}

			c.JSON(http.StatusOK, gin.H{
				"PublicCertificate": conn.ConnectionState().PeerCertificates[0].Raw,
				"IntermediateCertificate": conn.ConnectionState().PeerCertificates[1].Raw,
				"RootCA": conn.ConnectionState().PeerCertificates[2].Raw,
			})
		}
	})

	// Path: /subject
	cert.GET("/subject/:hostname", func(c *gin.Context) {
		host := c.Params.ByName("hostname")
		url := host + ":443"

		err := URLValidation(c, host)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Unable to parse URL or not a valid URL",
				"URL":     host,
			})
		} else {
			conn, err := tls.Dial("tcp", url, nil)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Server doesn't support SSL or certificate err",
					"error":   err.Error(),
				})
			}

			err = conn.VerifyHostname(host)
			if err != nil {
				//panic("Hostname doesn't match with certificate: " + err.Error())
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Hostname doesn't match with certificate",
					"error":   err.Error(),
				})
			}

			c.JSON(http.StatusOK, gin.H{
				"Subject": conn.ConnectionState().PeerCertificates[0].Subject,
			})
		}
	})

	// Path: /subject/cn
	cert.GET("/subject/cn/:hostname", func(c *gin.Context) {
		host := c.Params.ByName("hostname")
		url := host + ":443"

		err := URLValidation(c, host)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Unable to parse URL or not a valid URL",
				"URL":     host,
			})
		} else {
			conn, err := tls.Dial("tcp", url, nil)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Server doesn't support SSL or certificate err",
					"error":   err.Error(),
				})
			}

			err = conn.VerifyHostname(host)
			if err != nil {
				//panic("Hostname doesn't match with certificate: " + err.Error())
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Hostname doesn't match with certificate",
					"error":   err.Error(),
				})
			}

			c.JSON(http.StatusOK, gin.H{
				"CommonName": conn.ConnectionState().PeerCertificates[0].Subject.CommonName,
			})
		}
	})

	// Path: /dns
	cert.GET("/dns/:hostname", func(c *gin.Context) {
		host := c.Params.ByName("hostname")
		url := host + ":443"

		err := URLValidation(c, host)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Unable to parse URL or not a valid URL",
				"URL":     host,
			})
		} else {
			conn, err := tls.Dial("tcp", url, nil)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Server doesn't support SSL or certificate err",
					"error":   err.Error(),
				})
			}

			err = conn.VerifyHostname(host)
			if err != nil {
				//panic("Hostname doesn't match with certificate: " + err.Error())
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Hostname doesn't match with certificate",
					"error":   err.Error(),
				})
			}

			c.JSON(http.StatusOK, gin.H{
				"DNSNames": conn.ConnectionState().PeerCertificates[0].DNSNames,
			})
		}
	})

	// Path: /expiry
	cert.GET("/expiry/:hostname", func(c *gin.Context) {
		host := c.Params.ByName("hostname")
		url := host + ":443"

		err := URLValidation(c, host)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Unable to parse URL or not a valid URL",
				"URL":     host,
			})
		} else {
			conn, err := tls.Dial("tcp", url, nil)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Server doesn't support SSL or certificate err",
					"error":   err.Error(),
				})
			}

			err = conn.VerifyHostname(host)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Hostname doesn't match with certificate",
					"error":   err.Error(),
				})
			}

			nb := conn.ConnectionState().PeerCertificates[0].NotBefore
			na := conn.ConnectionState().PeerCertificates[0].NotAfter
			c.JSON(http.StatusOK, gin.H{
				"NotBefore": nb.Format(time.RFC850),
				"NotAfter":  na.Format(time.RFC850),
			})
		}
	})
}
