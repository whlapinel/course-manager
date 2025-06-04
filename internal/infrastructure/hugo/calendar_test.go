package hugo

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/features/term"
	"log"
	"testing"
	"time"
)

func TestCalendar(t *testing.T) {
	term := dto.Term{
		Term: term.Term{
			Start: time.Date(2025, time.August, 12, 0, 0, 0, 0, time.Local),
			End:   time.Date(2025, time.December, 18, 0, 0, 0, 0, time.Local),
		},
	}
	calendar := NewCalendar(term)
	for _, month := range calendar.Months {
		log.Println(month)
		for _, week := range month.Weeks {
			log.Println(week.Dates)
			for _, date := range week.Dates {
				log.Println(date.Date)
			}
		}
	}

}
