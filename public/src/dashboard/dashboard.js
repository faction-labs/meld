var Menu = require("../app/menu.js");
var AppMenu = require("../app/appmenu.js");

module.exports = React.createClass({
    render: function () {
        return (
            <div className="pusher">
                <Menu />

                <AppMenu />
            </div>
        )
    }
});
