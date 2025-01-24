package service

import "sync"

type ChatService struct {
	m         sync.Mutex
	ClientIds []string
}

func NewChatService() *ChatService {
	return &ChatService{
		ClientIds: make([]string, 0),
	}
}

func (cs *ChatService) AddClientId(id string) {
	cs.m.Lock()
	defer cs.m.Unlock()
	cs.ClientIds = append(cs.ClientIds, id)
}

func (cs *ChatService) RemoveClientId(id string) {
	cs.m.Lock()
	defer cs.m.Unlock()
	for i, clientId := range cs.ClientIds {
		if clientId == id {
			cs.ClientIds = append(cs.ClientIds[:i], cs.ClientIds[i+1:]...)
			break
		}
	}
}

func (cs *ChatService) GetClientId() []string {
	return cs.ClientIds
}
