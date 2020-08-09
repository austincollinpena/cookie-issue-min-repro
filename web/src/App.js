import React from 'react';

function App() {
  const test = () => {
    fetch("http://api.localhost").then(res => console.log(res))
  }
  test();
  return (
    <div className="App">
     <p>Hello, thanks for the help!</p>
    </div>
  );
}

export default App;
