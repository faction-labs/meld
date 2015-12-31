var React = require('react');
var api = require("../api/api.js");

module.exports = React.createClass({
    getInitialState: function() {
        return {
            result: null,
            error: null
        }
    },
    componentDidMount() {},
    handleInstallRust: function(path) {
        console.log("installing rust");
        this.setState({
            result: null,
            error: null
        });

        // make sure steam is installed
        api("/api/install/steam", function(xhr, status){
            if (xhr.status === 200) {
                // install rust
                api("/api/install/rust", function(xhr, status){
                    if (xhr.status === 200) {
                        this.setState({result: xhr.responseText});
                    } else {
                        this.setState({error: xhr.responseText});
                    }
                }.bind(this));
            } else {
                this.setState({error: xhr.responseText});
            }
        }.bind(this));
    },
    handleInstallOxide: function(path) {
        console.log("installing oxide");
        this.setState({
            result: null,
            error: null
        });

        // install oxide
        api("/api/install/oxide", function(xhr, status){
            if (xhr.status === 200) {
                this.setState({result: xhr.responseText});
            } else {
                this.setState({error: xhr.responseText});
            }
        }.bind(this));
    },
    render: function () {
        return (
            <div className="ui row setup">
                <img src="./assets/images/setup.png" />

                <br />

                <div className="setup-text">
                    { this.state.result &&
	                <div className="ui positive message">
	                    <div className="header">
                                Success
	                    </div>
	                    <p>
                                {this.state.result}
	                    </p>
                        </div>
                    }

                    { this.state.error &&
	                <div className="ui negative message">
	                    <div className="header">
                                Error
	                    </div>
	                    <p>
                                {this.state.error}
	                    </p>
                        </div>
                    }

                    <div className="ui button green massive" onClick={this.handleInstallRust}>Install Rust</div>
                    <div className="ui button green massive" onClick={this.handleInstallOxide}>Install Oxide</div>
                </div>
            </div>
        )
    }
});
