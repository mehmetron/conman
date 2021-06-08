import React, { useState, useEffect } from 'react';

import { connect, sendMsg } from './api';

import {StyledEditor} from "./components/editor/StyledEditor";


function App() {

    const [state, setState] = useState([])

    useEffect(()=> {
        connect((msg) => {
            console.log("New Message")
            setState(prevState => [...prevState, msg])

            console.log(state);
        });
    }, [])

    // function send(event) {
    //     if (event.keyCode === 13) {
    //         sendMsg(event.target.value);
    //         event.target.value = "";
    //     }
    // }

        return (
            <div>
                <StyledEditor />

            </div>
        );

}

export default App;