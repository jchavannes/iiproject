(function () {
    iiApp.Template = {
        /**
         * @param {jQuery} $contactList
         * @param {[Contact]} contacts
         */
        ContactList: function ($contactList, contacts) {
            var contactListHtml = "";
            var i, contact;
            for (i = 0; i < contacts.length; i++) {
                contact = contacts[i];
                contactListHtml +=
                    "<div class='row'>" +
                    "<div class='col-sm-6'>" +
                    "<p>" +
                    contact.Eid +
                    "</p>" +
                    "</div>" +
                    "<div class='col-sm-6'>" +
                    "<form style='display:inline;' id='view-profile-" + contact.Id + "'>" +
                    "<input type='hidden' name='id' value='" + contact.Eid + "' />" +
                    "<input type='submit' class='btn btn-primary' value='View' />" +
                    "</form>" +
                    " " +
                    "<form style='display:inline;' id='delete-contact-" + contact.Id + "'>" +
                    "<input type='submit' class='btn btn-danger' value='Delete' />" +
                    "</form>" +
                    "</div>" +
                    "</div>";
            }
            if (contacts.length === 0) {
                contactListHtml =
                    "<p>" +
                    "No contacts" +
                    "</p>";
            } else {
                contactListHtml =
                    "<div class='row'>" +
                    "<div class='col-sm-6'>" +
                    "<h4>Name</h4>" +
                    "</div>" +
                    "<div class='col-sm-6'>" +
                    "<h4>Actions</h4>" +
                    "</div>" +
                    "</div>" +
                    contactListHtml;
            }
            contactListHtml =
                "<h3>Contact List</h3>" +
                contactListHtml +
                "<hr/>";
            $contactList.html(contactListHtml);
            for (i = 0; i < contacts.length; i++) {
                contact = contacts[i];
                iiApp.Form.DeleteContact($("#delete-contact-" + contact.Id), contact.Id, $contactList);
                iiApp.Form.ViewId($("#view-profile-" + contact.Id));
            }
        }
    };
})();
