(function () {
    iiApp.Section = {
        /**
         * @param {jQuery} $contactList
         */
        ContactList: function ($contactList) {
            $.ajax({
                url: iiApp.URL.ContactList,
                success: function (data) {
                    /** @type {[Contact]} contacts */
                    var contacts;
                    try {
                        contacts = JSON.parse(data);
                    } catch (e) {
                        console.log(e);
                        return;
                    }
                    iiApp.Template.ContactList($contactList, contacts);
                }
            });
        }
    };
})();
