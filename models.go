package tcgcollector

import (
	"time"
)

// Common types
type ListResponse[T any] struct {
	Items          []T `json:"items"`
	ItemCount      int `json:"itemCount"`
	TotalItemCount int `json:"totalItemCount"`
	Page           int `json:"page"`
	PageCount      int `json:"pageCount"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

// Audit Log types
type AuditLogEventType struct {
	ID       int    `json:"id"`
	CodeName string `json:"codeName"`
	Name     string `json:"name"`
}

type AuditLogEntry struct {
	ID          int       `json:"id"`
	EventTypeID int       `json:"eventTypeId"`
	UserID      int       `json:"userId"`
	IPAddress   string    `json:"ipAddress"`
	CreatedAt   time.Time `json:"createdAt"`
	Details     string    `json:"details"`
}

// Card types
type Card struct {
	ID          int       `json:"id"`
	SetID       int       `json:"setId"`
	Name        string    `json:"name"`
	Number      string    `json:"number"`
	Rarity      string    `json:"rarity"`
	ImageURL    string    `json:"imageUrl"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Description string    `json:"description"`
}

type CardPrice struct {
	ID        int       `json:"id"`
	CardID    int       `json:"cardId"`
	Price     float64   `json:"price"`
	Currency  string    `json:"currency"`
	Source    string    `json:"source"`
	CreatedAt time.Time `json:"createdAt"`
}

// Card Variant types
type CardVariant struct {
	ID          int       `json:"id"`
	CardID      int       `json:"cardId"`
	TypeID      int       `json:"typeId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageURL    string    `json:"imageUrl"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CardVariantPrice struct {
	ID        int       `json:"id"`
	VariantID int       `json:"variantId"`
	Price     float64   `json:"price"`
	Currency  string    `json:"currency"`
	Source    string    `json:"source"`
	CreatedAt time.Time `json:"createdAt"`
}

// Set types
type Set struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	ReleaseDate string    `json:"releaseDate"`
	TotalCards  int       `json:"totalCards"`
	ImageURL    string    `json:"imageUrl"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// Collection types
type Collection struct {
	ID          int       `json:"id"`
	UserID      int       `json:"userId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsPublic    bool      `json:"isPublic"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CollectionCard struct {
	ID           int       `json:"id"`
	CollectionID int       `json:"collectionId"`
	CardID       int       `json:"cardId"`
	Quantity     int       `json:"quantity"`
	Condition    string    `json:"condition"`
	Notes        string    `json:"notes"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// User types
type User struct {
	ID                                  int        `json:"id"`
	DisplayName                         string     `json:"displayName"`
	EmailAddress                        string     `json:"emailAddress"`
	IsEmailAddressVerified              bool       `json:"isEmailAddressVerified"`
	HasApiAccessToken                   bool       `json:"hasApiAccessToken"`
	IsAdmin                             bool       `json:"isAdmin"`
	CanWriteApiExpansions               bool       `json:"canWriteApiExpansions"`
	CanReadApiCards                     bool       `json:"canReadApiCards"`
	CanReadApiCardsMinimal              bool       `json:"canReadApiCardsMinimal"`
	CanWriteApiCards                    bool       `json:"canWriteApiCards"`
	CanReadApiCardVariants              bool       `json:"canReadApiCardVariants"`
	CanReadApiCardVariantsMinimal       bool       `json:"canReadApiCardVariantsMinimal"`
	CanWriteApiCardVariants             bool       `json:"canWriteApiCardVariants"`
	CanWriteApiCardVariantTypes         bool       `json:"canWriteApiCardVariantTypes"`
	CanWriteApiCardIllustrators         bool       `json:"canWriteApiCardIllustrators"`
	CanWriteApiCardLists                bool       `json:"canWriteApiCardLists"`
	CanReadApiStatistics                bool       `json:"canReadApiStatistics"`
	CanWriteApiTcgPrices                bool       `json:"canWriteApiTcgPrices"`
	CanReadApiUsers                     bool       `json:"canReadApiUsers"`
	IsPremiumEnabled                    bool       `json:"isPremiumEnabled"`
	IsPremiumWithoutSubscriptionEnabled bool       `json:"isPremiumWithoutSubscriptionEnabled"`
	IsPremiumWithSubscriptionEnabled    bool       `json:"isPremiumWithSubscriptionEnabled"`
	PremiumStartDateTime                *time.Time `json:"premiumStartDateTime,omitempty"`
	PreviousPremiumStartDateTime        *time.Time `json:"previousPremiumStartDateTime,omitempty"`
	LastVisitDateTime                   time.Time  `json:"lastVisitDateTime"`
}

type UserPreferences struct {
	ID              int    `json:"id"`
	UserID          int    `json:"userId"`
	DefaultCurrency string `json:"defaultCurrency"`
	Language        string `json:"language"`
}

// Authentication types
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
	User      User      `json:"user"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	User User `json:"user"`
}

// NewsPost represents a news post in the system
type NewsPost struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Summary         string    `json:"summary"`
	TextMd          string    `json:"textMd"`
	Slug            string    `json:"slug"`
	CreatedDateTime time.Time `json:"createdDateTime"`
}

// Image types
type ImageType string

const (
	ImageTypeCardImage        ImageType = "CardImage"
	ImageTypeCardRaritySymbol ImageType = "CardRaritySymbol"
	ImageTypeEnergyTypeSymbol ImageType = "EnergyTypeSymbol"
	ImageTypeSetLogo          ImageType = "SetLogo"
	ImageTypeSetSymbol        ImageType = "SetSymbol"
)

type ImageSize struct {
	URL         string  `json:"url"`
	Width       int     `json:"width"`
	Height      int     `json:"height"`
	AspectRatio float64 `json:"aspectRatio"`
}

type Image struct {
	ID          int         `json:"id"`
	URL         string      `json:"url"`
	ContentType string      `json:"contentType"`
	Size        int64       `json:"size"`
	Width       int         `json:"width"`
	Height      int         `json:"height"`
	Sizes       []ImageSize `json:"sizes"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

// Statistics-related DTOs
type UserStatistics struct {
	UserCount              int               `json:"userCount"`
	MonthlyActiveUserCount int               `json:"monthlyActiveUserCount"`
	Premium                PremiumStatistics `json:"premium"`
}

type PremiumStatistics struct {
	UserCount                          int `json:"userCount"`
	UserWithoutSubscriptionCount       int `json:"userWithoutSubscriptionCount"`
	UserWithSubscriptionCount          int `json:"userWithSubscriptionCount"`
	ActiveSubscriptionCount            int `json:"activeSubscriptionCount"`
	ActiveNonExpiringSubscriptionCount int `json:"activeNonExpiringSubscriptionCount"`
	ActiveExpiringSubscriptionCount    int `json:"activeExpiringSubscriptionCount"`
	SuspendedSubscriptionCount         int `json:"suspendedSubscriptionCount"`
	ExpiredSubscriptionCount           int `json:"expiredSubscriptionCount"`
	CanceledSubscriptionCount          int `json:"canceledSubscriptionCount"`
}

// Card-related DTOs
type CardRule struct {
	ID           int     `json:"id"`
	Name         *string `json:"name,omitempty"`
	Description  string  `json:"description"`
	SortingOrder int     `json:"sortingOrder"`
}

type CardEffect struct {
	ID           int            `json:"id"`
	Type         CardEffectType `json:"type"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	SortingOrder int            `json:"sortingOrder"`
}

type CardAttack struct {
	ID               int                `json:"id"`
	Name             string             `json:"name"`
	Energies         []CardAttackEnergy `json:"energies"`
	HasExtraEnergies bool               `json:"hasExtraEnergies"`
	Damage           *string            `json:"damage,omitempty"`
	Description      *string            `json:"description,omitempty"`
	SortingOrder     int                `json:"sortingOrder"`
}

type CardAttackEnergy struct {
	ID           int        `json:"id"`
	Type         EnergyType `json:"type"`
	Quantity     int        `json:"quantity"`
	SortingOrder int        `json:"sortingOrder"`
}

type CardWeakness struct {
	ID           int        `json:"id"`
	Type         EnergyType `json:"type"`
	Value        string     `json:"value"`
	SortingOrder int        `json:"sortingOrder"`
}

type CardResistance struct {
	ID           int        `json:"id"`
	Type         EnergyType `json:"type"`
	Value        string     `json:"value"`
	SortingOrder int        `json:"sortingOrder"`
}

// Validation error DTO
type ValidationError struct {
	Message      string `json:"message"`
	PropertyPath string `json:"propertyPath"`
}
