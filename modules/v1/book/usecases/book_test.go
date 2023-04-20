package usecases

import (
	"bookLib/modules/v1/book/domain"
	m_repo "bookLib/modules/v1/book/interfaces/repositories/mock"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestBookUseCase_AllBooks(t *testing.T) {
	ctrl := gomock.NewController(t)
	now := time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		nameTest   string
		result     domain.Books
		wantErr    bool
		repository func(repo *m_repo.MockRepositoryPresenter)
	}{
		{
			nameTest: "Test Case 1 Get All Books: Success",
			result: domain.Books{
				{
					ID:         1,
					Name_book:  "The Lord of the Rings",
					Author:     "J. R. R. Tolkien",
					Created_at: now,
					Updated_at: now,
				},
				{
					ID:         2,
					Name_book:  "Harry Potter and the Philosopher's Stone",
					Author:     "J. K. Rowling",
					Created_at: now,
					Updated_at: now,
				},
			},
			repository: func(repo *m_repo.MockRepositoryPresenter) {
				repo.EXPECT().AllBooks().Return(domain.Books{
					{
						ID:         1,
						Name_book:  "The Lord of the Rings",
						Author:     "J. R. R. Tolkien",
						Created_at: now,
						Updated_at: now,
					},
					{
						ID:         2,
						Name_book:  "Harry Potter and the Philosopher's Stone",
						Author:     "J. K. Rowling",
						Created_at: now,
						Updated_at: now,
					},
				}, nil)
			},
		},
		{
			nameTest: "Test Case 2 Get All Books: Failed",
			result:   domain.Books{},
			wantErr:  true,
			repository: func(repo *m_repo.MockRepositoryPresenter) {
				repo.EXPECT().AllBooks().Return(domain.Books{}, errors.New("record not found"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.nameTest, func(t *testing.T) {
			repo := m_repo.NewMockRepositoryPresenter(ctrl)
			if tt.repository != nil {
				tt.repository(repo)
			}

			s := BookUseCase{
				repoBook: repo,
			}

			book, err := s.AllBooks()
			if (err != nil) && tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.result, book)
		})
	}
}

func TestBookUseCase_GetBookByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	now := time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		nameTest   string
		input      string
		result     domain.Book
		wantErr    bool
		repository func(repo *m_repo.MockRepositoryPresenter)
	}{
		{
			nameTest: "Test Case 1 Get Book By ID: Success",
			input:    "1",
			result: domain.Book{
				ID:         1,
				Name_book:  "The Lord of the Rings",
				Author:     "J. R. R. Tolkien",
				Created_at: now,
				Updated_at: now,
			},
			repository: func(repo *m_repo.MockRepositoryPresenter) {
				repo.EXPECT().GetBookByID("1").Return(domain.Book{
					ID:         1,
					Name_book:  "The Lord of the Rings",
					Author:     "J. R. R. Tolkien",
					Created_at: now,
					Updated_at: now,
				}, nil)
			},
		},
		{
			nameTest: "Test Case 2 Get Book By ID: Failed",
			input:    "2",
			result:   domain.Book{},
			wantErr:  true,
			repository: func(repo *m_repo.MockRepositoryPresenter) {
				repo.EXPECT().GetBookByID("2").Return(domain.Book{}, errors.New("record not found"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.nameTest, func(t *testing.T) {
			repo := m_repo.NewMockRepositoryPresenter(ctrl)
			if tt.repository != nil {
				tt.repository(repo)
			}

			s := BookUseCase{
				repoBook: repo,
			}

			book, err := s.GetBookByID(tt.input)
			if (err != nil) && tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.result, book)
		})
	}
}
