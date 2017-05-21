(function () {
    iiApp.Section = {
        ContactList: function () {
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
                    iiApp.Template.ContactList(iiApp.Elements.$contactList, contacts);
                }
            });
        }
    };
})();
