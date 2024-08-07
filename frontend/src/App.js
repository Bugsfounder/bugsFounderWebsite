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
  return (
    <div>
      <Navbar />
      <div className="container m-auto  mt-[130.5px]">
        <div className="dark:text-gray-300 mt-10">
          <Outlet context={{
            publicAxiosInstance,

          }} />
        </div>
      </div>
      <Footer />
    </div>
  );
}

export default App;
