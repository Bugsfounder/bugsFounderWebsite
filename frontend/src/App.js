import { Outlet } from 'react-router';
import { useEffect, useState } from 'react';
import axios from 'axios';
import Navbar from './Components/utils/Navbar';
import Footer from './Components/utils/Footer';
function App() {

  // const [response, setResponse] = useState("")

  // const fetchData = async () => {
  //   try {
  //     const response = await axios.get('http://localhost:8080/');
  //     console.log(response.data.message);
  //     setResponse(response.data.message);
  //   } catch (error) {
  //     console.error('Error fetching data', error);
  //   }
  // };

  return (
    <div>
      <Navbar />
      <div className="container w-10/12  m-auto">
        <Outlet context={{}} />
      </div>
      <Footer />
    </div>
  );
}

export default App;
