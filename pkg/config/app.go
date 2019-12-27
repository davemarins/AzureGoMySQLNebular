package config

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/subosito/gotenv"
	"io/ioutil"
	"log"
	"os"
)

var (
	db *gorm.DB
)

func init() {
	err := gotenv.Load()
	if err != nil {
		fmt.Println("Fatal error: cannot load environment variables")
		return
	}
}

func panicError(err error) {
	if err != nil {
		panic(err)
	}
	return
}

func SSLConnect() {
	rootCertPool := x509.NewCertPool()
	basePath, err := os.Getwd()
	panicError(err)
	pem, _ := ioutil.ReadFile(basePath + "\\BaltimoreCyberTrustRoot.crt.pem")
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatal("Failed to append PEM.")
	}
	err = mysql.RegisterTLSConfig("custom", &tls.Config{RootCAs: rootCertPool})
	panicError(err)
	var connectionString string
	connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true&tls=custom&parseTime=True", os.Getenv("MYSQL_USERNAME"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"))
	d, err := gorm.Open("mysql", connectionString)
	// db, err = sql.Open("mysql", connectionString)
	panicError(err)
	db = d
}

func LocalConnect() {
	// use your local configuration for the testing db (this case is mysql)
	d, err := gorm.Open("mysql", os.Getenv("MYSQL_USERNAME")+":"+os.Getenv("MYSQL_PASSWORD")+"@("+os.Getenv("MYSQL_HOST")+")/"+os.Getenv("MYSQL_DATABASE")+"?charset=utf8&parseTime=True&loc=Local")
	panicError(err)
	db = d
}

func Connect() {
	if os.Getenv("ENVIRONMENT") == "production" {
		SSLConnect()
	} else if os.Getenv("ENVIRONMENT") == "local" {
		LocalConnect()
	} else {
		fmt.Println("Connection environment not available")
		return
	}
}

func GetDB() *gorm.DB {
	return db
}
