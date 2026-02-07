package repository

import "businessense/internal/domain"

type UserRepository interface {
	Create(user *domain.User) error
	FindByID(id uint) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
}

type IndustryRepository interface {
	Create(industry *domain.Industry) error
	FindAll() ([]domain.Industry, error)
	Search(query string) ([]domain.Industry, error)
}

type CompanyRepository interface {
	Create(company *domain.Company) error
	FindByID(id uint) (*domain.Company, error)
	FindByUserID(userID uint) ([]domain.Company, error)
}

type PainPointRepository interface {
	Create(pp *domain.PainPoint) error
	FindAll() ([]domain.PainPoint, error)
	Search(query string) ([]domain.PainPoint, error)
}

type IssueRepository interface {
	FindByID(id uint) (*domain.Issue, error)
	FindByPainPointIDs(ppIDs []uint) ([]domain.IssueRelevance, error)
}

type SolutionRepository interface {
	FindByID(id uint) (*domain.Solution, error)
	FindByIssueID(issueID uint) ([]domain.Solution, error)
}

type ConstraintRepository interface {
	FindByIssueAndSolution(issueID, solutionID uint) ([]domain.Constraint, error)
}

type ProjectRepository interface {
	Create(project *domain.Project) error
	FindByID(id uint) (*domain.Project, error)
	FindByCompanyID(companyID uint) ([]domain.Project, error)
}
