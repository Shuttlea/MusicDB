/*
 * Music info
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"fmt"
	"gorm.io/gorm"
)

type dbHandler struct {
	DB *gorm.DB
	Addr string
}

func (h dbHandler) ChangePost(w http.ResponseWriter, r *http.Request) {
	group := r.URL.Query().Get("group")
	song := r.URL.Query().Get("song")
	if group == "" || song == "" {
		slog.Debug("Group or song is empty", "Group", group, "Song", song)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Group and Song fields can't be empty"))
		return
	}
	var changes Song
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&changes)
	if err != nil {
		slog.Debug("Request unmarshalling error", "Error", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Request unmarshalling error: " + err.Error()))
		return
	}
	res := h.DB.Model(&changes).Where("\"group\" = ? AND song = ?", group, song).Updates(changes)
	if res.RowsAffected == 0 {
		slog.Debug("This song does't exist", "Group", group, "Song", song)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("This song does't exist"))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h dbHandler) DeleteDelete(w http.ResponseWriter, r *http.Request) {
	group := r.URL.Query().Get("group")
	song := r.URL.Query().Get("song")
	if group == "" || song == "" {
		slog.Debug("Group or song is empty", "Group", group, "Song", song)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Group and Song fields can't be empty"))
		return
	}
	res := h.DB.Where("\"group\" = ? AND song = ?", group, song).Delete(&Song{})
	if res.RowsAffected == 0 {
		slog.Debug("This song does't exist", "Group", group, "Song", song)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("This song does't exist"))
		return
	}
	slog.Debug("Deleted song", "Group", group, "Song", song)
	w.WriteHeader(http.StatusOK)
}

func (h dbHandler) InfoGet(w http.ResponseWriter, r *http.Request) {
	group := r.URL.Query().Get("group")
	song := r.URL.Query().Get("song")
	if group == "" || song == "" {
		slog.Debug("Group or song is empty", "Group", group, "Song", song)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Group and Song fields can't be empty"))
		return
	}
	var s Song
	res := h.DB.Where("\"group\" = ? AND song = ?", group, song).Find(&s)
	if res.RowsAffected == 0 {
		slog.Debug("This song does't exist", "Group", group, "Song", song)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("This song does't exist"))
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	resp, err := json.Marshal(s.SongDetail)
	if err != nil {
		slog.Debug("Response marshalling error", "Error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (h dbHandler) InfoPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var s Song
	err := decoder.Decode(&s)
	if err != nil {
		slog.Debug("Request unmarshalling error", "Error", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Request unmarshalling error: " + err.Error()))
		return
	}
	if s.Group == "" || s.Song == "" {
		slog.Debug("Group or song is empty", "Group", s.Group, "Song", s.Song)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Group and Song fields can't be empty"))
		return
	}
	res := h.DB.Where("\"group\" = ? AND song = ?", s.Group, s.Song).Find(&s)
	if res.RowsAffected != 0 {
		slog.Debug("This song is already exist", "Group", s.Group, "Song", s.Song)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("This song is already exist"))
		return
	}
	h.DB.Create(&s)
	slog.Debug("Created song", "Group", s.Group, "Song", s.Song)
	newRequest,_ := http.NewRequest(http.MethodGet,fmt.Sprintf("http://%s/info?group=%s&song=%s",h.Addr,s.Group,s.Song),nil)
	h.InfoGet(w,newRequest)
}

func (h dbHandler) LyricsGet(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		slog.Debug("Limit conv", "Error", err)
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		slog.Debug("Page conv", "Error", err)
	}
	if limit < 1 {
		limit = 1
	}
	if page < 0 {
		slog.Debug("Page have to be under or equal zero", "Page", page, "Limit", limit)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Page have to be under or equal zero"))
		return
	}
	group := r.URL.Query().Get("group")
	song := r.URL.Query().Get("song")
	if group == "" || song == "" {
		slog.Debug("Group and Song fields can't be empty", "Group", group, "Song", song)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Group and Song fields can't be empty"))
		return
	}
	var s Song
	res := h.DB.Where("\"group\" = ? AND song = ?", group, song).Find(&s)
	if res.RowsAffected == 0 {
		slog.Debug("This song does't exist", "Group", group, "Song", song)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("This song does't exist"))
		return
	}
	couplets := strings.Split(s.Text, "\\n\\n")
	start := page * limit
	stop := page*limit + limit
	if start >= len(couplets) {
		slog.Debug("Page is too big value", "Page", page)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("[]"))
		return
	}
	if stop > len(couplets) {
		stop = len(couplets)
	}
	resp, err := json.Marshal(couplets[start:stop])
	if err != nil {
		slog.Debug("Error marshalling response", "Error", err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (h dbHandler) RootPost(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		slog.Debug("Limit conv", "Error", err)
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		slog.Debug("Page conv", "Error", err)
	}
	if limit < 1 {
		limit = 1
	}
	if page < 0 {
		slog.Debug("Page is under zero")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Page have to be under or equal zero"))
		return
	}
	var s Song
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&s)
	if err != nil && err.Error() != "EOF" {
		slog.Debug("Request unmarshalling error", "Error", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Request unmarshalling error: " + err.Error()))
		return
	}
	var songs []Song
	quer, str := queryBuilder(s)
	h.DB.Limit(limit).Offset(page*limit).Where(quer, str...).Find(&songs)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	resp, err := json.Marshal(songs)
	if err != nil {
		slog.Debug("Response marshalling error", "Error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
