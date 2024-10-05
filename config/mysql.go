package config

import (
    "os"
    "github.com/DaniloFaraum/studere-backend/models"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
)

func InitMySQL() (*gorm.DB, error) {
    logger := GetLogger("mysql")

    // Load .env file
    err := godotenv.Load("../DB.env")
    if err != nil {
        logger.Errorf("Error loading .env file: %v", err)
        return nil, err
    }

    // Get DSN from environment variable
    dsn := os.Getenv("DB_DSN")
    if dsn == "" {
        logger.Errorf("DB_DSN not found in environment variables")
        return nil, err
    }

    // Connect to the MySQL database
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        logger.Errorf("MySQL could not be initialized: %v", err)
        return nil, err
    }

 // Auto-migrate all schemas
 schemasToMigrate := []interface{}{
    &models.Answer{},
    &models.AnswerTag{},
    &models.Comment{},
    &models.Course{},
    &models.CourseTag{},
    &models.Image{},
    &models.Question{},
    &models.Questionnaire{},
    &models.QuestionnaireQuestion{},
    &models.QuestionTag{},
    &models.Rating{},
    &models.Role{},
    &models.Tag{},
    &models.User{},
}

for _, schema := range schemasToMigrate {
    err = db.AutoMigrate(schema)
    if err != nil {
        logger.Errorf("MySQL AutoMigrate failed for %T: %v", schema, err)
        return nil, err
    }
}

    logger.Info("MySQL successfully initialized and migrated.")
    return db, nil
}
