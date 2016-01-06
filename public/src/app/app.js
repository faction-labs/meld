var Splash = require("./splash.js");
var Login = require("./login.js");
var Signup = require("./signup.js");
var Dashboard = require("../dashboard/dashboard.js");
var NotFound = require("./notfound.js");

var RouterMixin = require('react-mini-router').RouterMixin;

var App = React.createClass({
    mixins: [RouterMixin],
    routes: {
        '/': 'home',
        '/login': 'login',
        '/signup': 'signup',
        '/dashboard': 'dashboard'
    },
    render: function() {
        return this.renderCurrentRoute();
    },
    home: function() {
        return (
            <Splash />
        )
    },
    login: function() {
        return (
            <Login />
        )
    },
    signup: function() {
        return (
            <Signup />
        )
    },
    dashboard: function() {
        return (
            <Dashboard />
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
