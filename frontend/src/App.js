import { Outlet } from 'react-router';
import { useEffect, useState } from 'react';
import axios from 'axios';
import Navbar from './Components/utils/Navbar';
import Footer from './Components/utils/Footer';

const publicAxiosInstance = axios.create({
  baseURL: 'http://localhost:8080/api/public',
  withCredentials: true, // If your backend requires credentials
});
function App() {

  const [loggedIn, setLoggedIn] = useState(!!localStorage.getItem("authToken"))
  return (
    <div>
      <Navbar loggedIn={loggedIn} setLoggedIn={setLoggedIn} />
      <div className=" md:mt-[130.5px] mt-[105px]">
        <div className="dark:text-gray-300">
          <Outlet context={{
            publicAxiosInstance,
            loggedIn,
            setLoggedIn
          }} />
        </div>
      </div>
      <Footer />
    </div>
  );
}

export default App;
