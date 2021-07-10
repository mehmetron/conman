import React, {useState} from "react";
import {Controlled as CodeMirror} from "react-codemirror2";
// import {sendMsg} from "../../api";

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

// const DefaultCode = `
// package main
//
// import (
// 	"fmt"
// 	"net/http"
// )
//
// func main() {
// 	api()
// }
//
// func api() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "second, %s!", r.URL.Path[1:])
// 	})
// 	http.ListenAndServe(":8090", mux)
// }
// `;
const DefaultCode = `
package main

import (
	"fmt"
)

func main() {
	fmt.Println("please work maaaan")
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
        <>
            <PureEditor
                name="js"
                value={editor}
                options={options}
                onChange={onChange()}
            />
            <Execute code={editor}/>
        </>
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

    // let Command = {
    //     "operation": 1,
    //     "entrypoint": "main.go",
    //     'content': code
    // }
    //
    // let Work = {
    //     "command": 1,
    //     "content": Command,
    // }

    function sendRequest() {
        console.log("code to execute", code)

        // var sample = {
        //     "operation": "run",
        //     "entrypoint": "main.go",
        //     "content": code
        //     // "content": "package main\nimport \"fmt\"\nfunc main() {\nfmt.Println(\"something here working\")\n}"
        // }
        var sample = {
            language: "go",
            version: "1.16.2",
            files: [
                {
                    name: "main.go",
                    content: code
                }
            ],
            stdin: "",
            args: ["1", "2", "3"],
            compile_timeout: 500000,
            run_timeout: 30000,
            compile_memory_limit: -1,
            run_memory_limit: -1
        }

        // postData("http://localhost:2000/api/v2/execute", sample).then(data => {
        //     console.log("141 response ", data); // JSON data parsed by `data.json()` call
        // });

        fetch('http://localhost:2000/api/v2/execute', {
            method: 'POST', // or 'PUT'
            // mode: 'no-cors',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(sample),
        })
            .then(response => response.json())
            .then(data => {
                console.log('Success:', data);
            })
            .catch((error) => {
                console.error('Error:', error);
            });

        // sendMsg(JSON.stringify(Work))
    }


    async function postData(url = '', sample = {}) {
        console.log("stringified ", JSON.stringify(sample))
        // Default options are marked with *
        const response = await fetch(url, {
            method: 'POST', // *GET, POST, PUT, DELETE, etc.
            mode: 'no-cors', // no-cors, *cors, same-origin
            // cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
            // credentials: 'same-origin', // include, *same-origin, omit
            headers: {
                'Content-Type': 'application/json'
                // "accept": "/"
                // 'content-type': 'application/json; charset=utf-8'
                // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            // redirect: 'follow', // manual, *follow, error
            // referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
            body: JSON.stringify(sample) // body data type must match "Content-Type" header
        });
        return response.json(); // parses JSON response into native JavaScript objects
    }

    return (
        <div>
            <button onClick={() => sendRequest()}>Execute</button>
        </div>
    )
}