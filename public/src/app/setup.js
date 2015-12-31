var Install = require("./install.js")
var Footer = require("./footer.js")

module.exports = React.createClass({
    render: function () {
        return (
            <div className="pusher">
                <Install />

                <Footer />
            </div>
        )
    }
});
