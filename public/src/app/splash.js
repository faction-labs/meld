var Welcome = require("./welcome.js");
var Dashboard = require("../dashboard/dashboard.js");
var Menu = require("./menu.js");
var Sidebar = require("./sidebar.js");
var api = require("../api/api.js");
var utils = require("./utils.js");

module.exports = React.createClass({
    getInitialState: function() {
        return {
            isSetup: true
        }
    },
    componentDidMount: function() {
        utils.isSetup(function(val) {
            this.setState({
                isSetup: val
            });
        }.bind(this));
    },
    render: function () {
        return (
            <div className="ui bottom attached segment pushable">
                { this.state.isSetup &&
                    <Sidebar />
                }

                    { ! this.state.isSetup &&
                        <Welcome />
                    }

                <div className="pusher">
                    { this.state.isSetup &&
                        <Menu />
                    }

                    { this.state.isSetup &&
                        <Dashboard />
                    }
                </div>
            </div>
        )
    }
});
