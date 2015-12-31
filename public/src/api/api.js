module.exports = function(url, callback) {
    $.ajax({
        url: url,
        dataType: 'json',
        cache: false,
        complete: function(xhr, status, err) {
            callback(xhr, status, err);
        }.bind(this)
    });
};
