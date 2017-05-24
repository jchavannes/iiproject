(function () {
    var profileSavedTimeout;
    iiApp.Form = {
        /**
         * @param {jQuery} $messageForm
         */
        SendMessage: function ($messageForm) {
            $messageForm.submit(function (e) {
                e.preventDefault();
                var recipient = $messageForm.find("[name=recipient]").val();
                var message = $messageForm.find("[name=message]").val();
                $.ajax({
                    type: "POST",
                    url: iiApp.BaseUrl.Get() + iiApp.URL.MessagesSendSubmit,
                    data: {
                        recipientEid: recipient,
                        message: message
                    },
                    success: function () {
                        window.location.reload();
                    }
                })
            })
        },
        /**
         * @param {jQuery} $deleteMessageForm
         */
        DeleteMessage: function ($deleteMessageForm) {
            $deleteMessageForm.submit(function (e) {
                e.preventDefault();
                var messageId = $deleteMessageForm.find("[name=id]").val();
                $.ajax({
                    type: "POST",
                    url: iiApp.BaseUrl.Get() + iiApp.URL.MessagesDeleteSubmit,
                    data: {
                        id: messageId
                    },
                    success: function () {
                        window.location.reload();
                    }
                })
            })
        },
        /**
         * @param {jQuery} $form
         * @param {jQuery} $profileSaved
         */
        EditProfile: function ($form, $profileSaved) {
            $form.submit(function (e) {
                e.preventDefault();
                var profile = $form.find("[name=profile]").val();
                $.ajax({
                    type: "POST",
                    url: iiApp.BaseUrl.Get() + iiApp.URL.EditProfileSubmit,
                    data: {
                        profile: profile
                    },
                    success: function () {
                        $profileSaved.show();
                        clearTimeout(profileSavedTimeout);
                        profileSavedTimeout = setTimeout(function () {
                            $profileSaved.hide();
                        }, 1500);
                    },
                    error: function (err) {
                        console.log(err);
                    }
                })
            });
        },
        /**
         * @param {jQuery} $deleteContactForm
         * @param {int} contactId
         */
        DeleteContact: function ($deleteContactForm, contactId) {
            $deleteContactForm.submit(function (e) {
                e.preventDefault();
                if (!confirm("Are you sure you want to delete this contact?")) {
                    return;
                }
                $.ajax({
                    type: "POST",
                    url: iiApp.BaseUrl.Get() + iiApp.URL.ContactDeleteSubmit,
                    data: {
                        contactId: contactId
                    },
                    success: function () {
                        iiApp.Section.ContactList();
                    }
                })
            });
        },
        /**
         * @param {jQuery} $contactForm
         */
        AddContact: function ($contactForm) {
            $contactForm.submit(function (e) {
                e.preventDefault();
                var id = $contactForm.find("[name=id]").val();
                $.ajax({
                    type: "POST",
                    url: iiApp.BaseUrl.Get() + iiApp.URL.ContactAddSubmit,
                    data: {
                        id: id
                    },
                    success: function () {
                        iiApp.Section.ContactList();
                    }
                })
            });
        },
        /**
         * @param {jQuery} $form
         */
        ViewId: function ($form) {
            $form.submit(function (e) {
                e.preventDefault();
                var id = $form.find("[name=id]").val();
                $.ajax({
                    type: "POST",
                    url: iiApp.BaseUrl.Get() + iiApp.URL.ViewIdSubmit,
                    data: {
                        id: id
                    },
                    /**
                     * @param {string} profileString
                     */
                    success: function (profileString) {
                        profileString = profileString.replace("\n", "<br/>");
                        iiApp.Elements.$profileViewArea.html(profileString);
                    }
                })
            });
        },
        /**
         * @param {jQuery} $form
         */
        Signup: function ($form) {
            $form.submit(function (e) {
                e.preventDefault();
                var username = $form.find("[name=username]").val();
                var password = $form.find("[name=password]").val();

                if (username.length === 0) {
                    alert("Must enter a username.");
                    return;
                }

                if (password.length === 0) {
                    alert("Must enter a password.");
                    return;
                }

                $.ajax({
                    type: "POST",
                    url: iiApp.BaseUrl.Get() + iiApp.URL.SignupSubmit,
                    data: {
                        username: username,
                        password: password
                    },
                    success: function () {
                        window.location = iiApp.BaseUrl.Get() + iiApp.URL.Dashboard
                    },
                    /**
                     * @param {XMLHttpRequest} xhr
                     */
                    error: function (xhr) {
                        alert("Error creating account:\n" + xhr.responseText + "\nIf this problem persists, try refreshing the page.");
                    }
                });
                return false;
            });
        },
        /**
         * @param {jQuery} $form
         */
        Login: function ($form) {
            $form.submit(function (e) {
                e.preventDefault();
                var username = $form.find("[name=username]").val();
                var password = $form.find("[name=password]").val();

                if (username.length === 0) {
                    alert("Must enter a username.");
                    return;
                }

                if (password.length === 0) {
                    alert("Must enter a password.");
                    return;
                }

                $.ajax({
                    type: "POST",
                    url: iiApp.BaseUrl.Get() + iiApp.URL.LoginSubmit,
                    data: {
                        username: username,
                        password: password
                    },
                    success: function () {
                        window.location = iiApp.BaseUrl.Get() + iiApp.URL.Dashboard
                    },
                    /**
                     * @param {XMLHttpRequest} xhr
                     */
                    error: function (xhr) {
                        alert("Error creating account:\n" + xhr.responseText + "\nIf this problem persists, try refreshing the page.");
                    }
                });
                return false;
            });
        }
    }
})();
