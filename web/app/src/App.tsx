import {useEffect, useState} from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import {Api} from "./api";

const api = new Api();

function App() {
  const [helloMessage, setHelloMessage] = useState<string>("");
  const [updateTime, setUpdateTime] = useState<string>("");

  useEffect(() => {
    updateMessage()
  }, [])

  function updateMessage() {
    api.api().helloGet().then((response) => {
      setHelloMessage(response.data.message)
    }).finally(() => {
      setUpdateTime(new Date().toLocaleTimeString())
    })
  }

  return (
    <>
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo"/>
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo"/>
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => updateMessage()}>
          Message is '{helloMessage}'
          <br/>
          Updated at {updateTime}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  )
}

export default App
