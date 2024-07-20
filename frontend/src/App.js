import { Outlet } from 'react-router';
import { useEffect, useState } from 'react';
import axios from 'axios';
function App() {

  const [response, setResponse] = useState("")

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get('http://localhost:8080/');
        console.log(response.data.message);
        setResponse(response.data.message);
      } catch (error) {
        console.error('Error fetching data', error);
      }
    };
    fetchData();
  }, []);


  return (
    <div>
      <Outlet context={{
        response,
      }} />
    </div>
  );
}

export default App;
