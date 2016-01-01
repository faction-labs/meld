var React = require('react');
var api = require("../api/api.js");

module.exports = React.createClass({
    getInitialState: function() {
        return {
            loading: false,
            loadingText: "",
            result: null,
            error: null
        }
    },
    componentDidMount() {},
    handleInstallRust: function(path) {
        console.log("installing rust");
        this.setState({
            result: null,
            error: null,
            loading: true,
            loadingText: "Installing Rust.  Please wait.  This may take a while."
        });

        // make sure steam is installed
        api("/api/install/steam", function(xhr, status){
            if (xhr.status === 200) {
                // install rust
                api("/api/install/rust", function(xhr, status){
                    if (xhr.status === 200) {
                        this.setState({
                            result: xhr.responseText,
                            loading: false
                        });
                    } else {
                        this.setState({
                            error: xhr.responseText,
                            loading: false
                        });
                    }
                }.bind(this));
            } else {
                this.setState({
                    error: xhr.responseText,
                    loading: false
                });
            }
        }.bind(this));
    },
    handleInstallOxide: function(path) {
        console.log("installing oxide");
        this.setState({
            result: null,
            error: null,
            loading: true,
            loadingText: "Installing Oxide.  Please wait.  This may take a while."
        });

        // install oxide
        api("/api/install/oxide", function(xhr, status){
            if (xhr.status === 200) {
                this.setState({
                    result: xhr.responseText,
                    loading: false
                });
            } else {
                this.setState({
                    error: xhr.responseText,
                    loading: false
                });
            }
        }.bind(this));
    },
    render: function () {
        return (
            <div className="ui row setup">
                { this.state.loading &&
                    <div className="ui active dimmer">
                        <div className="ui indeterminate text loader">{this.state.loadingText}</div>
                    </div>
                }
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
