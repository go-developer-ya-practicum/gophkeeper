package storage

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/go-developer-ya-practicum/gophkeeper/internal/server/models"
)

// Возможные ошибки при работе с хранилищем UserStorage
var (
	ErrEmailIsAlreadyInUse = errors.New("email is already in use")
	ErrInvalidCredentials  = errors.New("invalid request credentials")
)

// UserStorage определяет интерфейс для хранения учетных данных пользователей
type UserStorage interface {
	// PutUser сохраняет учетные данные пользователя
	PutUser(ctx context.Context, user *models.User) (userID int, err error)
	// GetUser возвращает ID пользователя с указанными учетными данными
	GetUser(ctx context.Context, user *models.User) (userID int, err error)
}

// Возможные ошибки при работе с хранилищем SecretStorage
var (
	ErrSecretNotFound        = errors.New("secret with given name not found")
	ErrSecretNameConflict    = errors.New("secret with given name already exists")
	ErrSecretVersionConflict = errors.New("secret version conflict")
)

// SecretStorage определяет интерфейс для хранения приватных данных пользователей
type SecretStorage interface {
	// GetSecret возвращает секрет с указанным именем name для пользователя c идентификатором userID
	GetSecret(ctx context.Context, name string, userID int) (*models.Secret, error)
	// PutSecret создает или обновляет содержимое секрета secret
	// возвращает новую версию секрета
	PutSecret(ctx context.Context, secret *models.Secret) (uuid.UUID, error)
	// ListSecrets возвращает список всех секретов пользователя с указанным идентификатором
	ListSecrets(ctx context.Context, userID int) ([]*models.Secret, error)
}