import React, { useEffect, useState } from 'react';
import axios from 'axios';
import './App.css';

function App() {
  const [catFacts, setCatFacts] = useState([]);

  useEffect(() => {
    const fetchCatFacts = async () => {
      try {
        const response = await axios.get('http://localhost:8080/facts');
        setCatFacts(response.data);
      } catch (error) {
        console.error('Error fetching cat facts:', error);
      }
    };
  
    fetchCatFacts();
  }, []);
  
  
  return (
  <div className="App">
  <header className="App-header">
  <h1>Cat Facts</h1>
  </header>
  <main>
  {catFacts.length === 0 && <p>Loading cat facts...</p>}
  <ul>
  {catFacts.map((fact, index) => (
  <li key={index}>{fact.fact || fact.text || JSON.stringify(fact)}</li>
))}

  </ul>
  </main>
  </div>
  );
  }
  
  export default App;