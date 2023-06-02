import React from 'react';
import './App.css';
import AllStatus from './components/all-status/allStatus.jsx'
import SingleStatus from './components/single-status/singleStatus.jsx'

const App = () => {
  return (
    <div className="app">
      <header className="app-header">
        <h2>Status Check</h2>
      </header>

      <main>
        <div className="single-status-container">
          <SingleStatus url="/api/v1/google-status" name="Google" />
          <SingleStatus url="/api/v1/amazon-status" name="Amazon" />
        </div>
        <AllStatus />
      </main>

      <footer className="app-footer">
        <p>Bilal Khundmiri, made for Sun Life</p>
      </footer>
    </div>
  );
}

export default App;


