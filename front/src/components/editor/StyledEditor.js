import React, { useState } from "react";
import { Controlled as CodeMirror } from "react-codemirror2";
import {sendMsg} from "../../api";

require("codemirror/mode/xml/xml");
require("codemirror/mode/javascript/javascript");
require("codemirror/mode/go/go");
require("codemirror/mode/python/python");
require("codemirror/mode/css/css");
require("codemirror/mode/jsx/jsx");
require("codemirror/lib/codemirror.css");
require("codemirror/addon/edit/matchbrackets.js");
require("codemirror/addon/edit/closebrackets.js");
require("codemirror/addon/search/match-highlighter.js");

require("codemirror/theme/monokai.css");
// require("codemirror/theme/dracula.css");
// require("codemirror/theme/panda-syntax.css");
// require("codemirror/theme/material.css");
// require("./theme.css");
// require("./darcula.css");
require("../../index.css");
require("./themes/oceanic.css");
require("./themes/rdark.css");
require("./themes/sidewalkchalk.css");
require("./themes/argonaut.css");
require("./themes/friendship-bracelet.css");
require("./themes/vscodedark.css");

const DefaultCode = `package main

import "fmt"


func main() {
    fmt.Println("hello world")
}
`;

const DefaultOptions = {
    theme: "vscode-dark",
    autoCloseBrackets: true,
    matchBrackets: true,
    styleActiveLine: true,
    cursorScrollMargin: 48,
    mode: "text/x-go",
    lineNumbers: true,
    indentUnit: 2,
    tabSize: 2,
    viewportMargin: 99
};

export function StyledEditor(props) {
    const [editor, setEditor] = useState(DefaultCode);

    const options = {
        ...DefaultOptions,
        ...props.options
    };

    const onChange = () => (editor, data, value) => {
        setEditor(value);
    };

    setInterval(console.log(editor), 3000)
    return (
        <React.Fragment>
            <PureEditor
                name="js"
                value={editor}
                options={options}
                onChange={onChange()}
            />
            <Execute code={editor}/>
        </React.Fragment>
    );
}

function PureEditor(props) {
    console.log(`rendering -> ${props.name}`);
    return (
        <CodeMirror
            value={props.value}
            options={props.options}
            onBeforeChange={props.onChange}
        />
    );
}

function Execute({code}) {

    let Command = {
        "operation": 1,
        "entrypoint": "main.go",
        'content': code
    }

    let Work = {
        "command": 1,
        "content": Command,
    }

    function sendRequest() {
        sendMsg(JSON.stringify(Work))
    }

    return (
        <div>
            <button onClick={()=>sendRequest()}>Execute</button>
        </div>
    )
}