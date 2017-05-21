var iiApp;

(function () {
    iiApp = {
        BaseUrl: {
            Set: SetBaseUrl,
            Get: GetBaseUrl
        },
        InitCsrf: InitCsrf,
        URL: {
            Dashboard: "dashboard",
            EditProfileSubmit: "edit-profile-submit",
            ContactAddSubmit: "contact-add-submit",
            ContactDeleteSubmit: "contact-delete-submit",
            ContactList: "contact-list",
            ViewIdSubmit: "view-id-submit",
            SignupSubmit: "signup-submit",
            LoginSubmit: "login-submit"
        }
    };

    var BaseURL = "/";

    /**
     * @param url {string}
     */
    function SetBaseUrl(url) {
        BaseURL = url;
    }

    /**
     * @return {string}
     */
    function GetBaseUrl() {
        return BaseURL;
    }

    /**
     * @param token {string}
     */
    function InitCsrf(token) {
        /**
         * @param method {string}
         * @returns {boolean}
         */
        function csrfSafeMethod(method) {
            // HTTP methods that do not require CSRF protection.
            return (/^(GET|HEAD|OPTIONS|TRACE)$/.test(method));
        }

        $.ajaxSetup({
            crossDomain: false,
            beforeSend: function (xhr, settings) {
                if (!csrfSafeMethod(settings.type)) {
                    xhr.setRequestHeader("X-CSRF-Token", token);
                }
            }
        });
    }
})();

/**
 * @typedef {{
 *   Id: int
 *   Eid: string
 *   PublicKey: string
 * }} Contact
 */
