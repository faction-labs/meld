var navigate = require('react-mini-router').navigate;

module.exports = React.createClass({
    componentDidMount: function() {
        $(".ui.sidebar").sidebar({
            closeable: false,
            dimPage: false,
            context: $(".section-dashboard")
        })
        .sidebar('attach events', '.section-dashboard .icon .menu .item');
    },
    click: function(path) {
        navigate(path);
    },
    render: function () {
        // for some reason the events below won't fire on the <i> unless
        // onClick is added.  teach me how to React :(
        return (
            <div className="ui labeled icon thin left inline vertical sidebar menu uncover visible">
                <a onClick={this.click.bind(this, "/")}  className="item">
                    <i onClick={this.click.bind(this, "/")}  className="dashboard icon"></i> Dashboard
                </a>
                <a onClick={this.click.bind(this, "/players")}  className="item">
                    <i onClick={this.click.bind(this, "/players")}  className="users icon"></i> Players
                </a>
                <a onClick={this.click.bind(this, "/console")}  className="item">
                    <i onClick={this.click.bind(this, "/console")}  className="terminal icon"></i> Console
                </a>
                <a onClick={this.click.bind(this, "/settings")} className="item">
                    <i onClick={this.click.bind(this, "/settings")}  className="settings icon"></i> Settings
                </a>
            </div>

        )
    }
});
