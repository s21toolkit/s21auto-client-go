package responces

import "encoding/json"

func UnmarshalUserResponce(data []byte) (UserResponce, error) {
	var r UserResponce
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UserResponce) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type UserResponce struct {
	Data Data `json:"data"`
}

type Data struct {
	User    User    `json:"user"`
	Student Student `json:"student"`
}

type Student struct {
	GetExperience GetExperience `json:"getExperience"`
	Typename      string        `json:"__typename"`
}

type GetExperience struct {
	ID               string `json:"id"`
	CookiesCount     int64  `json:"cookiesCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	CoinsCount       int64  `json:"coinsCount"`
	Level            Level  `json:"level"`
	Typename         string `json:"__typename"`
}

type Level struct {
	ID       int64  `json:"id"`
	Range    Range  `json:"range"`
	Typename string `json:"__typename"`
}

type Range struct {
	ID        string `json:"id"`
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}

type User struct {
	GetCurrentUser GetCurrentUser `json:"getCurrentUser"`
	Typename       string         `json:"__typename"`
}

type GetCurrentUser struct {
	ID                     string `json:"id"`
	AvatarURL              string `json:"avatarUrl"`
	Login                  string `json:"login"`
	FirstName              string `json:"firstName"`
	MiddleName             string `json:"middleName"`
	LastName               string `json:"lastName"`
	CurrentSchoolStudentID string `json:"currentSchoolStudentId"`
	Typename               string `json:"__typename"`
}
