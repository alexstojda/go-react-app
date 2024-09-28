import React, { useEffect, useState } from 'react';
import {ReactComponent as Logo} from './logo.svg';
import './App.css';
import {Api} from "./api";

const api = new Api();

function App() {
  const [helloResponse, setHelloResponse] = useState({});

  useEffect(() => {
    getHello()
  })

  function getHello() {
      api.api().helloGet().then((response) => {
        setHelloResponse(response.data.message)
      })
  }

  return (
    <div className="App">
      <header className="App-header">
        <Logo className="App-logo" />
        <span>
          <strong>/api/hello</strong> returned:
        </span>
        <pre>{JSON.stringify(helloResponse)}</pre>
      </header>
    </div>
  );
}

export default App;
