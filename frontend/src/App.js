import { Outlet } from 'react-router';
import { useEffect, useState } from 'react';
import axios from 'axios';
import Navbar from './Components/utils/Navbar';
import Footer from './Components/utils/Footer';
function App() {
  return (
    <div>
      <Navbar />
      <div className="container w-10/12 m-auto">
        <Outlet context={{}} />
      </div>
      <Footer />
    </div>
  );
}

export default App;
