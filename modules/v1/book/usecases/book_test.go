package usecases

import (
	"bookLib/modules/v1/book/domain"
	m_repo "bookLib/modules/v1/book/interfaces/repositories/mock"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tkuchiki/faketime"
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

func TestBookUseCase_CreateBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	location, err := time.LoadLocation("Asia/Jakarta")
	assert.NoError(t, err)
	now := time.Now().In(location)
	tests := []struct {
		nameTest   string
		input      domain.InsertBook
		result     domain.Book
		wantErr    bool
		repository func(repo *m_repo.MockRepositoryPresenter)
	}{
		{
			nameTest: "Test Case 1 Create Book: Success",
			input: domain.InsertBook{
				Name_book: "The Lord of the Rings",
				Author:    "J. R. R. Tolkien",
			},
			result: domain.Book{
				ID:         1,
				Name_book:  "The Lord of the Rings",
				Author:     "J. R. R. Tolkien",
				Created_at: now,
				Updated_at: now,
			},
			repository: func(repo *m_repo.MockRepositoryPresenter) {
				repo.EXPECT().CreateBook(domain.Book{
					Name_book:  "The Lord of the Rings",
					Author:     "J. R. R. Tolkien",
					Created_at: now,
					Updated_at: now,
				},
				).Return(domain.Book{
					ID:         1,
					Name_book:  "The Lord of the Rings",
					Author:     "J. R. R. Tolkien",
					Created_at: now,
					Updated_at: now,
				}, nil)
			},
		},
		{
			nameTest: "Test Case 2 Create Book: Failed",
			input: domain.InsertBook{
				Name_book: "Harry Potter",
				Author:    "J. K. Rowling",
			},
			result:  domain.Book{},
			wantErr: true,
			repository: func(repo *m_repo.MockRepositoryPresenter) {
				repo.EXPECT().CreateBook(domain.Book{
					Name_book:  "Harry Potter",
					Author:     "J. K. Rowling",
					Created_at: now,
					Updated_at: now,
				}).Return(domain.Book{}, errors.New("failed to update database"))
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

			book, err := s.CreateBook(tt.input)
			if (err != nil) && tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.result, book)
		})
	}
}

func TestBookUseCase_UpdateBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	location, err := time.LoadLocation("Asia/Jakarta")
	assert.NoError(t, err)
	now := time.Now().In(location)
	tests := []struct {
		nameTest   string
		id         string
		input      domain.UpdateBook
		result     domain.Book
		wantErr    bool
		repository func(repo *m_repo.MockRepositoryPresenter)
	}{
		{
			nameTest: "Test Case 1 Update Book: Success",
			id:       "1",
			input: domain.UpdateBook{
				Name_book: "The Lord of the Rings",
				Author:    "J. R. R. Tolkien",
			},
			result: domain.Book{
				ID:         1,
				Name_book:  "The Lord of the Rings",
				Author:     "J. R. R. Tolkien",
				Created_at: now,
				Updated_at: now,
			},
			repository: func(repo *m_repo.MockRepositoryPresenter) {
				repo.EXPECT().GetBookByID("1").Return(domain.Book{
					ID:        1,
					Name_book: "The Lord of the Rings",
				}, nil)
				repo.EXPECT().UpdateBook("1", domain.Book{
					Name_book:  "The Lord of the Rings",
					Author:     "J. R. R. Tolkien",
					Updated_at: now,
				},
				).Return(domain.Book{
					ID:         1,
					Name_book:  "The Lord of the Rings",
					Author:     "J. R. R. Tolkien",
					Created_at: now,
					Updated_at: now,
				}, nil)
			},
		},
		{
			nameTest: "Test Case 2 Update Book: Failed, Book Not Found",
			id:       "2",
			input: domain.UpdateBook{
				Name_book: "Harry Potter",
				Author:    "J. K. Rowling",
			},
			result:  domain.Book{},
			wantErr: true,
			repository: func(repo *m_repo.MockRepositoryPresenter) {
				repo.EXPECT().GetBookByID("2").Return(domain.Book{}, errors.New("record not found"))
			},
		},
		{
			nameTest: "Test Case 2 Update Book: Failed Update Book",
			id:       "3",
			input: domain.UpdateBook{
				Name_book: "Harry Potter",
				Author:    "J. K. Rowling",
			},
			result:  domain.Book{},
			wantErr: true,
			repository: func(repo *m_repo.MockRepositoryPresenter) {
				repo.EXPECT().GetBookByID("3").Return(domain.Book{
					ID:        1,
					Name_book: "Harry Potter",
				}, nil)
				repo.EXPECT().UpdateBook("3", domain.Book{
					Name_book:  "Harry Potter",
					Author:     "J. K. Rowling",
					Updated_at: now,
				}).Return(domain.Book{}, errors.New("failed to update database"))
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

			book, err := s.UpdateBook(tt.id, tt.input)
			if (err != nil) && tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.result, book)
		})
	}
}
