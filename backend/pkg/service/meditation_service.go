package service

import (
    "encoding/json"
    "errors"
    "sort"
    "time"

    "github.com/elialombardi/piccis-meditation/backend/pkg/dto"
    "github.com/elialombardi/piccis-meditation/backend/pkg/models"
    "gorm.io/gorm"
)

type MeditationService struct{ db *gorm.DB }

func NewMeditationService(db *gorm.DB) *MeditationService {
    return &MeditationService{db: db}
}

func (s *MeditationService) ListTypes() ([]dto.MeditationTypeDTO, error) {
    var types []models.MeditationType
    if err := s.db.Order("id asc").Find(&types).Error; err != nil {
        return nil, err
    }
    out := make([]dto.MeditationTypeDTO, 0, len(types))
    for _, t := range types {
        var script map[string]any
        _ = json.Unmarshal(t.Script, &script)
        out = append(out, dto.MeditationTypeDTO{
            ID: t.ID, Slug: t.Slug, Name: t.Name,
            Description: t.Description, Icon: t.Icon, Script: script,
        })
    }
    return out, nil
}

func (s *MeditationService) CreateSession(userID uint, req dto.CreateSessionRequest) (*dto.SessionDTO, error) {
    var t models.MeditationType
    if err := s.db.First(&t, req.MeditationTypeID).Error; err != nil {
        return nil, errors.New("meditation type not found")
    }
    sess := models.MeditationSession{
        UserID:           userID,
        MeditationTypeID: req.MeditationTypeID,
        DurationSeconds:  req.DurationSeconds,
        CompletedAt:      time.Now().UTC(),
        Notes:            req.Notes,
    }
    if err := s.db.Create(&sess).Error; err != nil {
        return nil, err
    }
    return &dto.SessionDTO{
        ID: sess.ID, MeditationTypeID: sess.MeditationTypeID, TypeName: t.Name,
        DurationSeconds: sess.DurationSeconds, CompletedAt: sess.CompletedAt, Notes: sess.Notes,
    }, nil
}

func (s *MeditationService) ListSessions(userID uint, limit int) ([]dto.SessionDTO, error) {
    if limit <= 0 || limit > 200 {
        limit = 50
    }
    var rows []struct {
        models.MeditationSession
        TypeName string
    }
    err := s.db.Table("meditation_sessions").
        Select("meditation_sessions.*, meditation_types.name as type_name").
        Joins("left join meditation_types on meditation_types.id = meditation_sessions.meditation_type_id").
        Where("meditation_sessions.user_id = ?", userID).
        Order("meditation_sessions.completed_at desc").
        Limit(limit).
        Scan(&rows).Error
    if err != nil {
        return nil, err
    }
    out := make([]dto.SessionDTO, 0, len(rows))
    for _, r := range rows {
        out = append(out, dto.SessionDTO{
            ID: r.ID, MeditationTypeID: r.MeditationTypeID, TypeName: r.TypeName,
            DurationSeconds: r.DurationSeconds, CompletedAt: r.CompletedAt, Notes: r.Notes,
        })
    }
    return out, nil
}

// Stats computes habit metrics. Streak logic: consecutive UTC days with >=1 session,
// counting backwards from today (or yesterday, if today has none yet).
func (s *MeditationService) Stats(userID uint) (*dto.HabitStatsDTO, error) {
    var sessions []models.MeditationSession
    if err := s.db.Where("user_id = ?", userID).Order("completed_at asc").Find(&sessions).Error; err != nil {
        return nil, err
    }

    stats := &dto.HabitStatsDTO{ByType: map[string]int{}}
    daySet := map[string]bool{}
    countsByDay := map[string]int{}

    for _, sess := range sessions {
        stats.TotalSessions++
        stats.TotalMinutes += sess.DurationSeconds / 60
        day := sess.CompletedAt.UTC().Format("2006-01-02")
        daySet[day] = true
        countsByDay[day]++
    }

    // ByType: resolve slugs
    var types []models.MeditationType
    s.db.Find(&types)
    typeIDToSlug := map[uint]string{}
    for _, t := range types {
        typeIDToSlug[t.ID] = t.Slug
    }
    for _, sess := range sessions {
        stats.ByType[typeIDToSlug[sess.MeditationTypeID]]++
    }

    // Current streak
    today := time.Now().UTC().Format("2006-01-02")
    yesterday := time.Now().UTC().AddDate(0, 0, -1).Format("2006-01-02")
    cursor := today
    if !daySet[today] && daySet[yesterday] {
        cursor = yesterday
    }
    for daySet[cursor] {
        stats.CurrentStreakDays++
        t, _ := time.Parse("2006-01-02", cursor)
        cursor = t.AddDate(0, 0, -1).Format("2006-01-02")
    }

    // Longest streak
    days := make([]string, 0, len(daySet))
    for d := range daySet {
        days = append(days, d)
    }
    sort.Strings(days)
    longest, run := 0, 0
    var prev time.Time
    for _, d := range days {
        t, _ := time.Parse("2006-01-02", d)
        if !prev.IsZero() && t.Sub(prev).Hours() == 24 {
            run++
        } else {
            run = 1
        }
        if run > longest {
            longest = run
        }
        prev = t
    }
    stats.LongestStreakDays = longest

    // Last 30 days for chart
    for i := 29; i >= 0; i-- {
        d := time.Now().UTC().AddDate(0, 0, -i).Format("2006-01-02")
        stats.Last30Days = append(stats.Last30Days, dto.DailyCount{Date: d, Count: countsByDay[d]})
    }
    return stats, nil
}