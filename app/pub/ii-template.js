(function () {
    iiApp.Template = {
        /**
         * @param {jQuery} $contactList
         * @param {[Contact]} contacts
         */
        ContactList: function ($contactList, contacts) {
            var contactListHtml = "";
            for (var i = 0; i < contacts.length; i++) {
                var contact = contacts[i];
                contactListHtml +=
                    "<p>" +
                    contact.Eid +
                    "</p>";
            }
            if (contacts.length === 0) {
                contactListHtml =
                    "<p>" +
                    "No contacts" +
                    "</p>";
            }
            contactListHtml =
                "<h3>Contact List</h3>" +
                contactListHtml +
                "<hr/>";
            $contactList.html(contactListHtml);
        }
    };
})();
