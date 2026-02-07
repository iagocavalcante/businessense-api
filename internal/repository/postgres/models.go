package postgres

import "time"

// User maps to the accounts table
type User struct {
	ID        uint      `gorm:"primaryKey"`
	Email     string    `gorm:"uniqueIndex;not null"`
	FirstName string    `gorm:"column:firstname"`
	LastName  string    `gorm:"column:lastname"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string { return "accounts" }

// Industry maps to the industries table
type Industry struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Industry) TableName() string { return "industries" }

// Company maps to the companies table
type Company struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"not null"`
	IndustryID uint   `gorm:"column:industry_id"`
	UserID     uint   `gorm:"column:user_id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (Company) TableName() string { return "companies" }

// PainPoint maps to the pain_points table
type PainPoint struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (PainPoint) TableName() string { return "pain_points" }

// Issue maps to the issues table
type Issue struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	Stats       string
	RiskFactors string `gorm:"column:riskfactors"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Issue) TableName() string { return "issues" }

// IssuePainPointMap maps to the issues_pain_points_maps table
type IssuePainPointMap struct {
	ID          uint `gorm:"primaryKey"`
	IssueID     uint `gorm:"column:issue_id"`
	PainPointID uint `gorm:"column:pain_point_id"`
}

func (IssuePainPointMap) TableName() string { return "issues_pain_points_maps" }

// Solution maps to the solutions table
type Solution struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Solution) TableName() string { return "solutions" }

// Constraint maps to the constraints table
type Constraint struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Constraint) TableName() string { return "constraints" }

// ConstraintIssueSolutionMap maps to the constraint_issues_solution_maps table
type ConstraintIssueSolutionMap struct {
	ID           uint `gorm:"primaryKey"`
	ConstraintID uint `gorm:"column:constraint_id"`
	IssueID      uint `gorm:"column:issue_id"`
	SolutionID   uint `gorm:"column:solution_id"`
}

func (ConstraintIssueSolutionMap) TableName() string { return "constraint_issues_solution_maps" }

// Project maps to the projects table
type Project struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	CompanyID   uint `gorm:"column:company_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Project) TableName() string { return "projects" }
