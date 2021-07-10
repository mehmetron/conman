import { v4 as uuidv4 } from 'uuid';

export const initialState = {
    focus: 2,
    editors: [{
        id: 1,
        name: "main.js",
        content: "One",
        codemirror: null,
    }, {
        id: 2,
        name: "bob.js",
        content: "Two",
        codemirror: null,
    }],
};

export const reducer = (state, action) => {
    switch (action.type) {
        case "add":
        {
            const newEditor = {
                id: uuidv4(),
                name: "james.js",
                content: action.text,
                codemirror: null,
            };
            return {
                ...state,
                editors: [...state.editors, newEditor],
            };
        }
        case "edit":
        {
            const idx = state.editors.findIndex(t => t.id === action.id);
            const editor = Object.assign({}, state.editors[idx]);
            editor.text = action.text;
            const editors = Object.assign([], state.editors);
            editors.splice(idx, 1, editor);
            return {
                ...state,
                editors: editors,
            };
        }
        case "remove":
        {
            const idx = state.editors.findIndex(t => t.id === action.id);
            const editors = Object.assign([], state.editors);
            editors.splice(idx, 1);
            return {
                ...state,
                editors: editors,
            };
        }
        default:
            return state;
    }
};