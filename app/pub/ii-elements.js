(function () {
    iiApp.Elements = {
        $contactList: null,
        $profileViewArea: null,
        /**
         * @param {jQuery} $contactList
         */
        SetContactList: function($contactList) {
            iiApp.Elements.$contactList = $contactList;
        },
        /**
         * @param {jQuery} $profileViewArea
         */
        SetProfileViewArea: function($profileViewArea) {
            iiApp.Elements.$profileViewArea = $profileViewArea;
        }
    };
})();
