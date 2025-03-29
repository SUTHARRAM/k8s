import { useEffect, useState } from 'react';

function App() {
  const [message, setMessage] = useState('');

  useEffect(() => {
    fetch('http://192.168.67.2:30002')  // Kubernetes service name! go-api-service
      .then(res => res.text())
      .then(data => setMessage(data));
  }, []);

  return <h1>{message || "Loading..."}</h1>;
}

export default App;