package swagger

import (
	"fmt"
	"log/slog"
	"strings"
)

func queryBuilder(s Song) (string, []interface{}) {
	var quer strings.Builder
	str := make([]interface{}, 0, 5)
	if s.Group != "" {
		quer.WriteString("\"group\" LIKE ?")
		str = append(str, fmt.Sprintf("%%%s%%", s.Group))
	}
	if s.Song != "" {
		if quer.Len() > 0 {
			quer.WriteString(" AND ")
		}
		quer.WriteString("song LIKE ?")
		str = append(str, fmt.Sprintf("%%%s%%", s.Song))
	}
	if s.ReleaseDate != "" {
		if quer.Len() > 0 {
			quer.WriteString(" AND ")
		}
		quer.WriteString("release_date LIKE ?")
		str = append(str, fmt.Sprintf("%%%s%%", s.ReleaseDate))
	}
	if s.Text != "" {
		if quer.Len() > 0 {
			quer.WriteString(" AND ")
		}
		quer.WriteString("text LIKE ?")
		str = append(str, fmt.Sprintf("%%%s%%", s.Text))
	}
	if s.Link != "" {
		if quer.Len() > 0 {
			quer.WriteString(" AND ")
		}
		quer.WriteString("link LIKE ?")
		str = append(str, fmt.Sprintf("%%%s%%", s.Link))
	}
	slog.Debug("Builded DB query","Query",quer.String(),"Values",str)
	return quer.String(), str
}
