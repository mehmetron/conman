import React, { useState, useEffect } from 'react';

// import { connect, sendMsg } from './api';

import {StyledEditor} from "./components/editor/StyledEditor";
import Editor from "./components/editor/Editor";



// https://codesandbox.io/s/file-tree-live-forked-ifw6l
// https://codepen.io/Souleste/pen/xxwvVva
function App() {

    var main = `
const myModule = require('./module');
let val = myModule.hello(); // val is "Hello"

console.log("this is the val: ", val)
`

    var module = `
module.exports = {
    hello: function() {
        return "Hello";
    }
}
`
    console.log("main sss: ", JSON.stringify(main))
    console.log("module sss: ", JSON.stringify(module))
    // const [state, setState] = useState([])
    //
    // useEffect(()=> {
    //     connect((msg) => {
    //         console.log("New Message")
    //         setState(prevState => [...prevState, msg])
    //
    //         console.log(state);
    //     });
    // }, [])

    // function send(event) {
    //     if (event.keyCode === 13) {
    //         sendMsg(event.target.value);
    //         event.target.value = "";
    //     }
    // }

        return (
            <div>
                <StyledEditor />

                <Editor />

            </div>
        );

}

export default App;