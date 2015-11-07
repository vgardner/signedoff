var React = require('react');
var HelloWorld = require('./Releases.jsx');

// Snag the initial state that was passed from the server side
var initialState = JSON.parse(document.getElementById('initial-state').innerHTML)

var releases = [];

for (release in initialState) {
    releases.push(initialState[release].ReleaseId);
}

React.render(
    <HelloWorld items={releases} />,
    document.getElementById('releases')
);
