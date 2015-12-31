var Welcome = require("./welcome.js")
var Footer = require("./footer.js")

module.exports = React.createClass({
    render: function () {
        return (
            <div className="pusher">
                <Welcome />

                <Footer />
            </div>
        )
    }
});
