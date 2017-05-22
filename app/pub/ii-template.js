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
                    "<tr>" +
                    "<td>" +
                    contact.Eid +
                    "</td>" +
                    "<td>" +
                    "<form style='display:inline;' id='view-profile-" + contact.Id + "'>" +
                    "<input type='hidden' name='id' value='" + contact.Eid + "' />" +
                    "<input type='submit' class='btn btn-primary btn-xs' value='View' />" +
                    "</form>" +
                    " " +
                    "<form style='display:inline;' id='delete-contact-" + contact.Id + "'>" +
                    "<input type='submit' class='btn btn-danger btn-xs' value='Remove' />" +
                    "</form>" +
                    "</td>" +
                    "</tr>";
            }
            if (contacts.length === 0) {
                contactListHtml =
                    "<p>" +
                    "No contacts" +
                    "</p>";
            } else {
                contactListHtml =
                    "<div class='table-responsive'>" +
                    "<table class='table table-bordered table-striped'>" +
                    "<thead>" +
                    "<tr>" +
                    "<th>Name</th>" +
                    "<th>Actions</th>" +
                    "</tr>" +
                    "</thead>" +
                    "<tbody>" +
                    contactListHtml +
                    "</tbody>" +
                    "</table>" +
                    "</div>";
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
