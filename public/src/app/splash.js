var Menu = require("./menu.js")
var Welcome = require("./welcome.js")

module.exports = React.createClass({
    render: function () {
        return (
            <div className="pusher">
                <Menu />

                <Welcome />
            </div>
        )
    }
});
