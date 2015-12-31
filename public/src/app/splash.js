var Menu = require("./menu.js")
var Setup = require("./setup.js")
var Footer = require("./footer.js")

module.exports = React.createClass({
    render: function () {
        return (
            <div className="pusher">
                <Menu />

                <Setup />

                <Footer />
            </div>
        )
    }
});
