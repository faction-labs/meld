var Splash = require("./splash.js");
var Setup = require("./setup.js");
var NotFound = require("./notfound.js");

var RouterMixin = require('react-mini-router').RouterMixin;

var App = React.createClass({
    mixins: [RouterMixin],
    routes: {
        '/': 'home',
        '/setup': 'setup'
    },
    render: function() {
        return this.renderCurrentRoute();
    },
    home: function() {
        return (
            <Splash />
        )
    },
    setup: function() {
        return (
            <Setup />
        )
    },
    notFound: function(path) {
        this.state.path = path;
        return (
            <NotFound path={this.state.path} />
        )
    }
});

React.render(<App />, document.querySelector('body'));
