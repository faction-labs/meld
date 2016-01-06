var navigate = require('react-mini-router').navigate;

module.exports = React.createClass({
    click: function(path) {
        navigate(path);
    },
    render: function () {
        return (
            <div className="ui secondary small pointing menu">
                <a className="active item">
                  <i className="dashboard icon"></i>
                  Dashboard
                </a>
                <a className="item">
                  <i className="server icon"></i>
                  Nodes
                </a>
                <a className="item">
                  <i className="options icon"></i>
                  Settings
                </a>
            </div>
        )
    }
});
