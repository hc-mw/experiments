import { useState } from "react";

interface Todo {
  id: number;
  name: string;
}

let count = 0;

function App() {
  const [input, setInput] = useState("");
  const [todos, setTodos] = useState<Todo[]>([]);

  return (
    <>
      <h1>Todos</h1>
      <div className="form">
        <div className="input">
          <input
            type="text"
            value={input}
            onChange={(e) => {
              setInput(e.target.value);
            }}
          />
        </div>
        <div>
          <button
            onClick={() => {
              const obj: Todo = {
                id: count,
                name: input,
              };
              ++count;
              setInput("");
              setTodos([...todos, obj]);
            }}
          >
            Click here
          </button>
        </div>
      </div>

      <ul>
        {todos.map((obj) => {
          return <li>{obj.name}</li>;
        })}
      </ul>
    </>
  );
}

export default App;
