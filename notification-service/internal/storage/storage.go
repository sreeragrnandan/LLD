package storage

type UserPreferences struct {
	Channels []string
}

type InMemoryStorage struct {
	preferences map[string]UserPreferences
}

type Storage interface {
	GetUserPreferences(userID string) (UserPreferences, error)
}

func (s *InMemoryStorage) GetUserPreferences(userID string) (up UserPreferences, err error) {
	if up, found := s.preferences[userID]; found {
		return up, nil
	}
	return
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		preferences: map[string]UserPreferences{
			"user1": {Channels: []string{"email", "sms"}},
			"user2": {Channels: []string{"email"}},
		},
	}
}
