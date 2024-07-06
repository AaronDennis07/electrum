import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.jsx'
import './index.css'
import CreateSessionForm from './CreateSession.jsx'
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { AuthProvider } from './AuthContext.jsx'


ReactDOM.createRoot(document.getElementById('root')).render(
  
  <App/>
   
)
// ReactDOM.render(<App />, document.getElementById('root'));
