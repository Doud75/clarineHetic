package usecase

import (
    "backClarineHetic/internal/adapter/dto"
    "backClarineHetic/internal/domain"
    "github.com/google/uuid"
)

type ConversationUseCase interface {
    GetOrCreateConv(userUuid uuid.UUID, authUserUuid uuid.UUID) (*dto.ConversationResponse, error)
    SaveMessage(message *domain.Message) error
}

type conversationUseCase struct {
    conversationRepo domain.ConversationRepository
    messageRepo      domain.MessageRepository
}

func NewConversationUseCase(conversationRepo domain.ConversationRepository, messageRepo domain.MessageRepository) ConversationUseCase {
    return &conversationUseCase{
        conversationRepo: conversationRepo,
        messageRepo:      messageRepo,
    }
}

func (c *conversationUseCase) GetOrCreateConv(userUuid uuid.UUID, authUserUUID uuid.UUID) (*dto.ConversationResponse, error) {
    conversation, err := c.conversationRepo.GetByUserIDs(authUserUUID, userUuid)
    if err != nil {
        if err.Error() == "conversation non trouvée" {
            newConv := &domain.Conversation{
                UserIDA: authUserUUID,
                UserIDB: userUuid,
            }
            if err = c.conversationRepo.Create(newConv); err != nil {
                return nil, err
            }
            conversation = newConv
        } else {
            return nil, err
        }
    }
    messages, err := c.messageRepo.GetMessagesByConversationID(conversation.UUID)
    if err != nil {
        return nil, err
    }
    if messages == nil {
        messages = []*domain.Message{}
    }

    response := &dto.ConversationResponse{
        ConversationID: conversation.UUID,
        Messages:       messages,
    }
    return response, nil
}

func (c *conversationUseCase) SaveMessage(message *domain.Message) error {
    return c.messageRepo.Create(message)
}
