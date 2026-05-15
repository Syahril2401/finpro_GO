package service

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

// SRLProfile represents one of the 81 pre-defined student profiles.
type SRLProfile struct {
	CombinationID  string           `json:"combination_id"`
	ProfileTitle   string           `json:"profile_title"`
	DeepWork       int              `json:"deep_work_capacity"`
	CognitiveStyle string           `json:"cognitive_style"`
	Consistency    int              `json:"consistency_score"`
	Retention      int              `json:"retention_score"`
	Strengths      json.RawMessage  `json:"strengths"`
	Weaknesses     json.RawMessage  `json:"weaknesses"`
	AreasForGrowth json.RawMessage  `json:"areas_for_growth"`
	AIStrategy     json.RawMessage  `json:"ai_strategy"`
	Recommendations json.RawMessage `json:"recommendations"`
}

// ProfileService provides lookup of pre-defined SRL profiles.
type ProfileService interface {
	GetProfile(combinationID string) (*SRLProfile, error)
}

type profileService struct {
	profiles map[string]*SRLProfile
	mu       sync.RWMutex
}

// NewProfileService loads the 81-profile JSON file and indexes them by combination_id.
func NewProfileService(jsonPath string) ProfileService {
	svc := &profileService{
		profiles: make(map[string]*SRLProfile),
	}

	data, err := os.ReadFile(jsonPath)
	if err != nil {
		log.Fatalf("[Profile] FATAL: Cannot read profiles JSON at %s: %v", jsonPath, err)
	}

	var profiles []SRLProfile
	if err := json.Unmarshal(data, &profiles); err != nil {
		log.Fatalf("[Profile] FATAL: Cannot parse profiles JSON: %v", err)
	}

	for i := range profiles {
		svc.profiles[profiles[i].CombinationID] = &profiles[i]
	}

	log.Printf("[Profile] Loaded %d SRL profiles successfully", len(svc.profiles))
	return svc
}

func (s *profileService) GetProfile(combinationID string) (*SRLProfile, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	profile, exists := s.profiles[combinationID]
	if !exists {
		return nil, fmt.Errorf("profile not found for combination: %s", combinationID)
	}
	return profile, nil
}
