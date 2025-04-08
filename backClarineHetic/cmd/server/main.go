package main

import (
    "backClarineHetic/internal/adapter/controller"
    "backClarineHetic/internal/adapter/repository"
    "backClarineHetic/internal/adapter/router"
    "backClarineHetic/internal/usecase"
    "database/sql"
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
)

func main() {
    // Charger les variables d'environnement
    if err := godotenv.Load(); err != nil {
        log.Println("Pas de fichier .env trouvé, utilisation des variables d'environnement existantes")
    }

    dbUser := os.Getenv("POSTGRES_USER")
    dbPassword := os.Getenv("POSTGRES_PASSWORD")
    dbName := os.Getenv("POSTGRES_DB")
    dbHost := os.Getenv("POSTGRES_HOST")
    dbPort := os.Getenv("POSTGRES_PORT")

    if dbHost == "" {
        dbHost = "localhost"
    }
    if dbPort == "" {
        dbPort = "5432"
    }

    connStr := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Erreur de connexion à la DB :", err)
    }
    defer func() {
        if err := db.Close(); err != nil {
            log.Printf("Erreur lors de la fermeture de la DB: %v", err)
        }
    }()

    if err = db.Ping(); err != nil {
        log.Fatal("Impossible de se connecter à la DB :", err)
    }
    log.Println("Connexion à la DB réussie.")

    userRepo := repository.NewPostgresUserRepo(db)
    authUC := usecase.NewAuthUsecase(userRepo)
    authController := controller.NewAuthController(authUC)

    profileUC := usecase.NewProfileUseCase(userRepo)
    profileController := controller.NewProfileController(profileUC)

    r := gin.Default()
    router.NewAuthRouter(r, authController)
    router.NewProfileRouter(r, profileController)
    if err := r.Run(":9070"); err != nil {
        log.Fatalf("Erreur lors du démarrage du serveur : %v", err)
    }
}
