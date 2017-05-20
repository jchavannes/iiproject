(function () {
    var profileSavedTimeout;
    iiApp.Form = {
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
                    success: function() {
                        $profileSaved.show();
                        clearTimeout(profileSavedTimeout);
                        profileSavedTimeout = setTimeout(function() {
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
         * @param {jQuery} $contactForm
         * @param {jQuery} $contactList
         */
        AddContact: function ($contactForm, $contactList) {
            $contactForm.submit(function (e) {
                e.preventDefault();
                var id = $contactForm.find("[name=id]").val();
                $.ajax({
                    type: "POST",
                    url: iiApp.BaseUrl.Get() + iiApp.URL.ContactAddSubmit,
                    data: {
                        id: id
                    },
                    success: function() {
                        iiApp.Section.ContactList($contactList);
                    }
                })
            });
        },
        /**
         * @param {jQuery} $form
         * @param {jQuery} $profileArea
         */
        ViewId: function ($form, $profileArea) {
            $form.submit(function (e) {
                e.preventDefault();
                var id = $form.find("[name=id]").val();
                $.ajax({
                    type: "POST",
                    url: iiApp.BaseUrl.Get() + iiApp.URL.ViewIdSubmit,
                    data: {
                        id: id
                    },
                    success: function(data) {
                        $profileArea.html(data);
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
