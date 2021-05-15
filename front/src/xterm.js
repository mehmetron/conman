import React from "react";
import { Terminal } from "xterm";
import { FitAddon } from "xterm-addon-fit";


// import "./xterm.css";
import "xterm/css/xterm.css";

import { connect, sendMsg } from './api';

import c from "ansi-colors";

let term;
const fitAddon = new FitAddon();

function banana() {
    term.writeln("Type some keys and commands to play around.");
    term.writeln("");
    // term.prompt();
    // term.setOption("cursorBlink", true);
    setTimeout(() => term.clear(), 3000);
}

export default function Xterm() {

    const [command, setCommand] = React.useState([])
    const [chat, setChat] = React.useState([])


    React.useEffect(() => {

        connect((msg) => {
            console.log("New Message", msg)
            setChat((prev)=>[...prev, msg.data])
            console.log("some ", msg.data)
            console.log(chat);
            // term.write("\n")
            term.write(msg.data);

            var shellprompt = "$ ";
            term.write("\r\n" + shellprompt);
        });



        term = new Terminal({
            convertEol: true,
            fontFamily: `'Monaco', monospace`,
            fontSize: 15,
            fontWeight: 900
            // rendererType: "dom" // default is canvas
        });
        const socket = new WebSocket('ws://localhost:8080/ws');

        //Styling
        term.setOption("theme", {
            background: "#1e1e1e",
            foreground: "white"
        });
        term.setOption("cursorBlink", true);

        // Load Fit Addon
        term.loadAddon(fitAddon);


        // Open the terminal in #terminal-container
        term.open(document.getElementById("xterm"));

        //Write text inside the terminal
        term.write(c.magenta("I am ") + c.blue("Blue") + c.red(" and i like it"));

        // Make the terminal's size and geometry fit the size of #terminal-container
        fitAddon.fit();

        term.onKey((key) => {
            const char = key.domEvent.key;
            if (char === "Enter") {
                prompt();
                setCommand([])
            } else if (char === "Backspace") {
                term.write("\b \b");
                setCommand((prev)=>[prev.splice(-1, 1)])
            } else {
                term.write(char);
                setCommand((prev)=>[...prev, char])
            }
        });

        prompt();
    }, [])


    prompt = () => {
        console.log("message is ", command.join(""))
        setTimeout(() => sendMsg(command.join("")), 3000)

        var shellprompt = "$ ";
        term.write("\r\n" + shellprompt);
    };

        const style = {
            backgroundColor: "#252626"
        };
        return (
            <div className="App" style={{ background: "" }}>
                <div style={style}>
                    <button onClick={() => banana()}>dd</button>
                </div>
                <div id="xterm" style={{ height: "100%", width: "100%" }} />
            </div>
        );

}
