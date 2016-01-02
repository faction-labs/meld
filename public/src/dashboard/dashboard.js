var React = require('react');
var api = require("../api/api.js");

module.exports = React.createClass({
    getInitialState: function() {
        return {
            hostname: "",
            version: "",
            map: "",
            players: 0
        }
    },
    componentDidMount() {
        api("/api/rust/status", function(xhr, status){
            if (xhr.status === 200) {
                var o = xhr.responseJSON;
                this.setState({
                    hostname: o.hostname,
                    version: o.version,
                    map: o.map,
                    players: o.players
                });
            }
        }.bind(this));
    },
    render: function () {
        return (
            <div className="ui section-dashboard">
                <div className="ui row">
		    <div className="ui two statistics">
                        <div className="ui small statistic">
                            <div className="value">
                            {this.state.hostname}
                            </div>
                            <div className="label">
                              Hostname
                            </div>
                        </div>
                        <div className="ui small statistic">
                            <div className="value">
                            {this.state.version}
                            </div>
                            <div className="label">
                              Version
                            </div>
                        </div>
                    </div>

                    <div className="ui divider"></div>

		    <div className="ui two statistics">
                        <div className="ui small statistic">
                            <div className="value">
                            {this.state.map}
                            </div>
                            <div className="label">
                              Map
                            </div>
                        </div>
                        <div className="ui small statistic">
                            <div className="value">
                            {this.state.players}
                            </div>
                            <div className="label">
                              Players
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
});
