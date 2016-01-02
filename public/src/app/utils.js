var api = require("../api/api.js");

module.exports = {
    isSetup: function(callback) {
        api("/api/info", function(xhr, status){
            if (xhr.status === 200) {
                var setup = false;
                var info = xhr.responseJSON;

                if (info.steamCMDPath === "" || info.rustPath === "") {
                    callback(false);
                } else {
                    callback(true);
                }
            } else {
                callback(false);
            }
        }.bind(this));
    }
}
