package web

import (
	"github.com/jchavannes/jgo/web"
	"github.com/jchavannes/iiproject/app/db/auth"
	"errors"
	"net/http"
	"github.com/jchavannes/iiproject/app/db/key"
	"github.com/jchavannes/iiproject/eid/client"
	"github.com/jchavannes/iiproject/app/db/message"
	"fmt"
	"github.com/jchavannes/iiproject/eid"
	"github.com/jchavannes/iiproject/app/db/contact"
	"strconv"
)

var messagesRoute = web.Route{
	Pattern: URL_MESSAGES,
	Handler: func(r *web.Response) {
		if ! auth.IsLoggedIn(r.Session.CookieId) {
			r.SetRedirect(getUrlWithBaseUrl(URL_INDEX, r))
			return
		}

		user, err := auth.GetSessionUser(r.Session.CookieId)
		if err != nil {
			r.Error(fmt.Errorf("Unable to get session user: %s", err), http.StatusUnprocessableEntity)
			return
		}

		contacts, err := contact.GetUserContacts(user.Id)
		if err != nil {
			r.Error(err, http.StatusInternalServerError)
			return
		}

		messages, _ := message.GetMessages(user.Id)
		r.Helper["Messages"] = messages
		r.Helper["Contacts"] = contacts
		r.Render()
	},
}

var messagesSendSubmitRoute = web.Route{
	Pattern: URL_MESSAGES_SEND_SUBMIT,
	CsrfProtect: true,
	Handler: func(r *web.Response) {
		if ! auth.IsLoggedIn(r.Session.CookieId) {
			r.Error(errors.New("Not logged in."), http.StatusUnauthorized)
		}

		recipientEid := r.Request.GetFormValue("recipientEid")
		messageString := r.Request.GetFormValue("message")

		user, err := auth.GetSessionUser(r.Session.CookieId)
		if err != nil {
			r.Error(fmt.Errorf("Unable to get session user: %s", err), http.StatusUnprocessableEntity)
			return
		}

		userKey, err := key.Get(user.Id)
		if err != nil {
			r.Error(fmt.Errorf("Unable to get user key: %s", err), http.StatusUnprocessableEntity)
			return
		}

		idGetResponse, err := client.GetId(recipientEid)
		if err != nil {
			r.Error(fmt.Errorf("Unable to get client id: %s", err), http.StatusUnprocessableEntity)
			return
		}

		userEid := r.Request.HttpRequest.Host + "/id/" + user.Username
		publicKeyPair := eid.KeyPair{
			PublicKey: []byte(idGetResponse.PublicKey),
		}
		sendMessage, err := client.SendMessage(recipientEid, userEid, messageString, publicKeyPair, userKey.GetKeyPair())
		if err != nil {
			r.Error(fmt.Errorf("Unable to send client message: %s", err), http.StatusUnprocessableEntity)
			return
		}

		err = message.Add(user.Id, recipientEid, sendMessage, true)
		if err != nil {
			r.Error(fmt.Errorf("Error adding message to database: %s", err), http.StatusUnprocessableEntity)
			return
		}
	},
}

var messagesDeleteSubmitRoute = web.Route{
	Pattern: URL_MESSAGES_DELETE_SUBMIT,
	CsrfProtect: true,
	Handler: func(r *web.Response) {
		if ! auth.IsLoggedIn(r.Session.CookieId) {
			r.SetRedirect(getUrlWithBaseUrl(URL_INDEX, r))
			return
		}

		messageIdString := r.Request.GetFormValue("id")
		messageId, err := strconv.Atoi(messageIdString)
		if err != nil {
			r.Error(fmt.Errorf("Unable to get session user: %s", err), http.StatusUnprocessableEntity)
			return
		}

		user, err := auth.GetSessionUser(r.Session.CookieId)
		if err != nil {
			r.Error(fmt.Errorf("Unable to get session user: %s", err), http.StatusUnprocessableEntity)
			return
		}

		err = message.Delete(uint(messageId), user.Id)
		if err != nil {
			r.Error(fmt.Errorf("Unable to delete message: %s", err), http.StatusUnprocessableEntity)
			return
		}
	},
}
