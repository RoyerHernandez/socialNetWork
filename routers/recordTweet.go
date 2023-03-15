package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/RoyerHernandez/socialNetWork.git/db"
	"github.com/RoyerHernandez/socialNetWork.git/models"
)

func RecordTweet(w http.ResponseWriter, r *http.Request) {
	var (
		message models.Tweet
		status  bool
	)
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "incorrect data !!! ", 400)
		return
	}
	tweet := models.RecordTweet{
		UserID:  UserId,
		Message: message.Message,
		Date:    time.Now(),
	}
	_, status, err = db.InsertTweet(tweet)
	if err != nil {
		http.Error(w, "an error has occurred while intent insert the register, please try again"+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "it has been not inserted the tweet", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
