var React = require('react');
var HelloWorld = require('./HelloWorld.jsx');

// Snag the initial state that was passed from the server side
var initialState = JSON.parse(document.getElementById('initial-state').innerHTML)
var initialState2 = [
         "Vin",
         "Test",
         "Chicken",
         "Duck",
         "Eggs",
         "Fish",
         "Granola",
         "Hash Browns"
       ];

var releaseId = initialState[0].ReleaseId;

React.render(
    <HelloWorld items={initialState2} releaseid={initialState[1].ReleaseId} />,
    document.getElementById('example')
);
