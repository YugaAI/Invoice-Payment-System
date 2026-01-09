package config

import (
	"fmt"
	model2 "invoice-payment-system/model"
	"log"
	"os"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, fallback to OS env")
	}
}

func BuildDSN(prefix string) string {
	host := os.Getenv(prefix + "_HOST")
	port := os.Getenv(prefix + "_PORT")
	user := os.Getenv(prefix + "_USER")
	pass := os.Getenv(prefix + "_PASSWORD")
	name := os.Getenv(prefix + "_NAME")
	ssl := os.Getenv(prefix + "_SSLMODE")

	if host == "" || port == "" || user == "" || name == "" || pass == "" || ssl == "" {
		log.Fatal("database env is not complete")
	}

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, pass, name, ssl,
	)
}

func InitWriteDB() *gorm.DB {
	LoadEnv()

	dsn := BuildDSN("DB_WRITE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get sql.DB:", err)
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := db.AutoMigrate(
		&model2.Invoices{},
		&model2.Company{},
		&model2.Item{},
		&model2.User{},
	); err != nil {
		log.Fatal("failed to auto migrate:", err)
	}

	log.Println("database connected & migrated")
	return db
}

func InitReadDB() *gorm.DB {
	dsn := BuildDSN("DB_READ")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal("failed to connect READ db:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("READ database ready")
	return db
}

// LOCAL
type Config struct {
	PasetoKey paseto.V4SymmetricKey
}

func LoadPaseto() *Config {
	hexKey := os.Getenv("PASETO_SECRET")
	if hexKey == "" {
		panic("PASETO_SECRET required")
	}

	key, err := paseto.V4SymmetricKeyFromHex(hexKey)
	if err != nil {
		panic("invalid paseto key: " + err.Error())
	}

	return &Config{
		PasetoKey: key,
	}
}

// PUBLIC
type ConfigPublic struct {
	PublicKey  paseto.V4AsymmetricPublicKey
	PrivateKey paseto.V4AsymmetricSecretKey
}

func LoadPasetoPublic() *ConfigPublic {
	privKey := os.Getenv("PRIVATE_KEY_HEX")
	pubKey := os.Getenv("PUBLIC_KEY_HEX")

	if privKey == "" || pubKey == "" {
		panic("PRIVATE_KEY_HEX or PUBLIC_KEY_HEX required")
	}

	secretKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(privKey)
	if err != nil {
		panic("invalid secret key: " + err.Error())
	}

	pubBytes, errPub := paseto.NewV4AsymmetricPublicKeyFromHex(pubKey)
	if errPub != nil {
		panic("invalid public key: " + errPub.Error())
	}

	return &ConfigPublic{
		PublicKey:  pubBytes,
		PrivateKey: secretKey,
	}
}
