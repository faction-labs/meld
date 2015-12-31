var React = require('react');
var navigate = require('react-mini-router').navigate;

module.exports = React.createClass({
    click: function(path) {
        navigate(path);
    },
    render: function () {
        return (
            <div className="ui row setup">
                    <img src="./assets/images/setup.png" />

                    <br />

                    <div className="about-text">
                        <h1 className="ui header">
                        Welcome to Meld, the Rust management application.
                        To get started, please click Setup.  This will install
                        the Meld distribution of Rust.
                        </h1>
                    </div>

                    <br />

                    <div className="ui massive blue button" onClick={this.click.bind(this, "/setup")}>Setup</div>
            </div>
        )
    }
});
