var navigate = require('react-mini-router').navigate;
var auth = require('./auth.js');
var UserMenu = require('./usermenu.js');

module.exports = React.createClass({
    getInitialState: function() {
        return {
            username: auth.getUsername(),
            isLoggedIn: auth.isLoggedIn()
        }
    },
    componentDidMount() {
        this.setState({
            username: auth.getUsername(),
            isLoggedIn: auth.isLoggedIn()
        });
    },
    click: function(path) {
        navigate(path);
    },
    render: function () {
        return (
            <div className="ui large fixed menu">
                <a onClick={this.click.bind(this, "/")} className="item">Meld</a>
                {
                    this.state.isLoggedIn ? (
                        [
                            <div key="0" className="right item">
                                <UserMenu />
                                <a key="1" onClick={this.click.bind(this, "/help")} className="ui button help">Help</a>
                            </div>
                        ]
                    ) : (
                        <div className="right item">
                            <a onClick={this.click.bind(this, "/login")} className="ui green button">Login</a>
                        </div>
                    )
                }
            </div>
        )
    }
});
