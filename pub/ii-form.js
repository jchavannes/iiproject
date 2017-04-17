(function() {
    iiApp.Form = {
        /**
         * @param {jQuery} $ele
         */
        Signup: function ($ele) {
            $ele.submit(function (e) {
                e.preventDefault();
                var username = $ele.find("[name=username]").val();
                var password = $ele.find("[name=password]").val();

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
        }
    }
})();
