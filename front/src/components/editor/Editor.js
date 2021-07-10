import React, {useReducer} from 'react';
import {initialState, reducer} from "./EditorCrud";


export default function Editor() {
    const [state, dispatch] = useReducer(reducer, initialState);
    
    function addTodo(){
       dispatch({type: "add", text: "banana is nice"})
        console.log("10 added todo ")
    }

    return (
        <div>
            <p>Editor penis</p>
            {JSON.stringify(state)}
            <button onClick={()=>addTodo()}>Add todo</button>
        </div>
    )
}



// import React, { useReducer } from 'react'
//
// const reducer = (state, action) => {
//     switch (action.type) {
//         case 'add':
//             return {
//                 ...state,
//                 inputVal: '',
//                 todos: [...state.todos, state.inputVal],
//             }
//         case 'remove':
//             return {
//                 ...state,
//                 todos: state.todos.filter((val, index) => index !== action.payload.id),
//             }
//         case 'updateVal':
//             return { ...state, inputVal: action.payload.value }
//         default:
//             return state
//     }
// }
//
// const initialState = { inputVal: '', todos: [] }
//
// const TodoList = () => {
//     const [{ inputVal, todos }, dispatch] = useReducer(reducer, initialState)
//
//     return (
//         <div>
//             <ul>
//                 {todos.map((todo, index) => (
//                     <li key={index}>
//                         {todo}
//                         <span
//                             onClick={() =>
//                                 dispatch({ type: 'remove', payload: { id: index } })
//                             }
//                         >
//               ❌
//             </span>
//                     </li>
//                 ))}
//             </ul>
//             <input
//                 type="text"
//                 value={inputVal}
//                 onChange={e => dispatch({ type: 'updateVal', payload: e.target })}
//             />
//             <button onClick={() => dispatch({ type: 'add' })}>Add Todo</button>
//         </div>
//     )
// }














//
// import React, { useReducer, useState } from "react";
//
// function reducer(state, action) {
//     switch (action.type) {
//         case "add-todo":
//             return {
//                 todos: [...state.todos, { text: action.text, completed: false }],
//                 todoCount: state.todoCount + 1
//             };
//         case "toggle-todo":
//             return {
//                 todos: state.todos.map((t, idx) =>
//                     idx === action.idx ? { ...t, completed: !t.completed } : t
//                 ),
//                 todoCount: state.todoCount
//             };
//         case "delete-todo":
//             return {
//                 todos: state.todos.filter((t, idx) => idx !== action.idx),
//                 todoCount: state.todoCount - 1
//             };
//         default:
//             return state;
//     }
// }
//
// const App = () => {
//     const [{ todos, todoCount }, dispatch] = useReducer(reducer, {
//         todos: [],
//         todoCount: 0
//     });
//     const [text, setText] = useState();
//
//     return (
//         <div>
//             <form
//                 onSubmit={e => {
//                     e.preventDefault();
//                     dispatch({ type: "add-todo", text });
//                     setText("");
//                 }}
//             >
//                 <input value={text} onChange={e => setText(e.target.value)} />
//             </form>
//             <div>number of todos: {todoCount}</div>
//             {todos.map((t, idx) => (
//                 <div
//                     key={t.idx}
//
//                     onClick={() => dispatch({ type: "toggle-todo", idx })}
//                     onDoubleClick={() => dispatch({type:"delete-todo",idx})}
//                     style={{
//                         textDecoration: t.completed ? "line-through" : ""
//                     }}
//                 >
//                     {t.text}
//                     {idx}
//                 </div>
//             ))}
//             <pre>{JSON.stringify(todos,null,2)}</pre>
//         </div>
//     );
// };
//
// export default App;


