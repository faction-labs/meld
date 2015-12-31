var navigate = require('react-mini-router').navigate;

module.exports = React.createClass({
    getInitialState: function() {
        return {}
    },
    componentDidMount() {},
    click: function(path) {
        navigate(path);
    },
    render: function () {
        return (
            <div className="ui large fixed menu">
                <a onClick={this.click.bind(this, "/")} className="item">Meld</a>

                <div className="right item">
                    <a key="1" onClick={this.click.bind(this, "/help")} className="ui button help">Help</a>
                </div>
            </div>
        )
    }
});
